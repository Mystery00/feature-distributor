package grpc

import (
	"context"
	"errors"
	"feature-distributor/common/alert"
	"feature-distributor/common/value"
	"feature-distributor/core/db/enum"
	"feature-distributor/core/db/model"
	"feature-distributor/core/db/query"
	"feature-distributor/core/pb"
	"feature-distributor/core/provider"
	"github.com/open-feature/go-sdk/openfeature"
	"gorm.io/gorm"
)

var client = openfeature.NewClient("feature-distributor")

type ToggleServer struct {
	pb.UnimplementedToggleServiceServer
}

func (s *ToggleServer) ListToggle(ctx context.Context, in *pb.ListToggleRequest) (*pb.ListToggleResponse, error) {
	t := query.Toggle
	tc := t.WithContext(ctx)

	offset := int((in.GetIndex() - 1) * in.GetSize())
	limit := int(in.GetSize())
	list, total, err := tc.Where(t.ProjectID.Eq(in.GetProjectId())).FindByPage(offset, limit)
	if err != nil {
		return nil, err
	}
	toggleList := make([]*pb.ListItemToggle, 0)
	for _, item := range list {
		toggleList = append(toggleList, &pb.ListItemToggle{
			Id:          item.ID,
			Enabled:     item.Enable,
			Title:       item.Title,
			Key:         item.Key,
			Description: item.Description,
			ValueType:   enum.ValueTypeEnum(item.ValueType).String(),
		})
	}
	return &pb.ListToggleResponse{
		Total:   total,
		Toggles: toggleList,
	}, nil
}

func (s *ToggleServer) GetToggle(ctx context.Context, in *pb.GetToggleRequest) (*pb.Toggle, error) {
	t := query.Toggle
	tv := query.ToggleValue
	tc := t.WithContext(ctx)
	tvc := tv.WithContext(ctx)

	toggle, err := tc.Where(t.ID.Eq(in.GetId())).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, alert.Error(alert.ToggleNotExist)
		}
		return nil, err
	}
	values, err := tvc.Where(tv.ToggleID.Eq(toggle.ID)).Find()
	if err != nil {
		return nil, err
	}
	return convertToToggle(toggle, values), nil
}

func (s *ToggleServer) GetToggleValue(ctx context.Context, in *pb.GetToggleValueRequest) (*pb.GetToggleValueResponse, error) {
	attributes := make(map[string]interface{})
	for k, v := range in.GetReqUser().GetAttributes() {
		attributes[k] = v
	}
	c := openfeature.NewEvaluationContext(in.GetReqUser().GetRolloutKey(), attributes)
	cc := context.WithValue(ctx, "projectKey", in.GetProjectKey())
	toggle, err := getToggle(cc, c, in.GetToggleKey(), in.GetToggleType())
	if err != nil {
		return nil, err
	}
	return &pb.GetToggleValueResponse{
		ToggleKey:   in.GetToggleKey(),
		ResultValue: *toggle,
	}, nil
}

func getToggle(ctx context.Context, evalContext openfeature.EvaluationContext, toggleKey, toggleType string) (*string, error) {
	switch toggleType {
	case "bool":
		v, err := client.BooleanValue(ctx, toggleKey, false, evalContext)
		if err != nil {
			return nil, provider.DealError(err)
		}
		toString, err := value.ToString(toggleType, v)
		if err != nil {
			return nil, err
		}
		return &toString, nil
	case "string":
		v, err := client.StringValue(ctx, toggleKey, "", evalContext)
		if err != nil {
			return nil, provider.DealError(err)
		}
		toString, err := value.ToString(toggleType, v)
		if err != nil {
			return nil, err
		}
		return &toString, nil
	case "float":
		v, err := client.FloatValue(ctx, toggleKey, 0, evalContext)
		if err != nil {
			return nil, provider.DealError(err)
		}
		toString, err := value.ToString(toggleType, v)
		if err != nil {
			return nil, err
		}
		return &toString, nil
	case "int":
		v, err := client.IntValue(ctx, toggleKey, 0, evalContext)
		if err != nil {
			return nil, provider.DealError(err)
		}
		toString, err := value.ToString(toggleType, v)
		if err != nil {
			return nil, err
		}
		return &toString, nil
	case "json":
		v, err := client.ObjectValue(ctx, toggleKey, "{}", evalContext)
		if err != nil {
			return nil, provider.DealError(err)
		}
		toString, err := value.ToString(toggleType, v)
		if err != nil {
			return nil, err
		}
		return &toString, nil
	default:
		return nil, errors.New("invalid toggle type")
	}
}

func (s *ToggleServer) SaveToggle(ctx context.Context, in *pb.SaveToggleRequest) (*pb.Toggle, error) {
	if in.ToggleId != nil {
		return updateToggle(ctx, in)
	} else {
		return insertToggle(ctx, in)
	}
}

func insertToggle(ctx context.Context, in *pb.SaveToggleRequest) (*pb.Toggle, error) {
	t := query.Toggle
	tv := query.ToggleValue
	tc := t.WithContext(ctx)
	tvc := tv.WithContext(ctx)
	//查询是否存在相同的key
	toggle, err := tc.Where(t.Key.Eq(in.GetKey())).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if toggle != nil {
		return nil, alert.Error(alert.ToggleExist)
	}
	//检测数据
	if int(in.GetDisabledValue()) <= 0 || int(in.GetDisabledValue()) > len(in.GetValues()) {
		return nil, alert.Error(alert.InvalidToggleValue)
	}
	//保存数据
	toggle = &model.Toggle{
		ProjectID:   in.GetProjectId(),
		Enable:      in.GetEnabled(),
		Title:       in.GetTitle(),
		Key:         in.GetKey(),
		Description: in.GetDescription(),
		ValueType:   int8(enum.ParseValueType(in.GetValueType())),
		//先填充默认值
		DefaultValue: 0,
		//先填充默认值
		ReturnValueWhenDisable: 0,
	}
	err = tc.Save(toggle)
	if err != nil {
		return nil, err
	}
	//删除所有旧的值
	values := make([]*model.ToggleValue, 0)
	for _, item := range in.GetValues() {
		toggleValue := &model.ToggleValue{
			ToggleID:    toggle.ID,
			Title:       item.GetTitle(),
			Value:       item.GetValue(),
			Description: item.GetDescription(),
		}
		err := tvc.Save(toggleValue)
		if err != nil {
			return nil, err
		}
		values = append(values, toggleValue)
	}
	//更新提前存入的默认值
	disabledValueId := values[int(in.GetDisabledValue())-1].ID
	_, err = tc.Where(t.ID.Eq(toggle.ID)).Updates(model.Toggle{
		ReturnValueWhenDisable: disabledValueId,
	})
	if err != nil {
		return nil, err
	}
	return convertToToggle(toggle, values), nil
}

func updateToggle(ctx context.Context, in *pb.SaveToggleRequest) (*pb.Toggle, error) {
	t := query.Toggle
	tv := query.ToggleValue
	tc := t.WithContext(ctx)
	tvc := tv.WithContext(ctx)
	//检测数据
	if int(in.GetDefaultValue()) <= 0 || int(in.GetDisabledValue()) <= 0 || int(in.GetDefaultValue()) > len(in.GetValues()) || int(in.GetDisabledValue()) > len(in.GetValues()) {
		return nil, alert.Error(alert.InvalidToggleValue)
	}
	//删除所有旧的值
	values := make([]*model.ToggleValue, 0)
	for _, item := range in.GetValues() {
		toggleValue := &model.ToggleValue{
			ToggleID:    in.GetToggleId(),
			Title:       item.GetTitle(),
			Value:       item.GetValue(),
			Description: item.GetDescription(),
		}
		err := tvc.Save(toggleValue)
		if err != nil {
			return nil, err
		}
		values = append(values, toggleValue)
	}
	//保存数据
	defaultValueId := values[int(in.GetDefaultValue())-1].ID
	disabledValueId := values[int(in.GetDisabledValue())-1].ID
	_, err := tc.Where(t.ID.Eq(in.GetToggleId())).Updates(model.Toggle{
		Enable:                 in.GetEnabled(),
		Title:                  in.GetTitle(),
		Description:            in.GetDescription(),
		ValueType:              int8(enum.ParseValueType(in.GetValueType())),
		DefaultValue:           defaultValueId,
		ReturnValueWhenDisable: disabledValueId,
	})
	if err != nil {
		return nil, err
	}
	toggle, _ := tc.Where(t.ID.Eq(in.GetToggleId())).First()
	return convertToToggle(toggle, values), nil
}

func convertToToggle(toggle *model.Toggle, values []*model.ToggleValue) *pb.Toggle {
	if toggle == nil {
		return nil
	}
	valueList := make([]*pb.ToggleValue, 0)
	for _, item := range values {
		valueList = append(valueList, &pb.ToggleValue{
			Id:          item.ID,
			ToggleId:    item.ToggleID,
			Title:       item.Title,
			Value:       item.Value,
			Description: item.Description,
			CreateTime:  item.CreateTime.UnixMilli(),
		})
	}
	return &pb.Toggle{
		Id:            toggle.ID,
		ProjectId:     toggle.ProjectID,
		Enabled:       toggle.Enable,
		Title:         toggle.Title,
		Key:           toggle.Key,
		Description:   toggle.Description,
		ValueType:     enum.ValueTypeEnum(toggle.ValueType).String(),
		DefaultValue:  toggle.DefaultValue,
		DisabledValue: toggle.ReturnValueWhenDisable,
		CreateTime:    toggle.CreateTime.UnixMilli(),
		UpdateTime:    toggle.UpdateTime.UnixMilli(),
		Values:        valueList,
	}
}
