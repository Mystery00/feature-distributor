package grpc

import (
	"context"
	"errors"
	"feature-distributor/common/alert"
	"feature-distributor/common/attribute"
	"feature-distributor/common/operation"
	"feature-distributor/core/db/enum"
	"feature-distributor/core/db/model"
	"feature-distributor/core/db/query"
	"feature-distributor/core/pb"
	"gorm.io/gorm"
	"sort"
)

type ReqGroupServer struct {
	pb.UnimplementedReqGroupServiceServer
}

func (r ReqGroupServer) ListReqGroup(ctx context.Context, in *pb.ListReqGroupRequest) (*pb.ListReqGroupResponse, error) {
	rg := query.ReqGroup
	rgc := rg.WithContext(ctx)
	offset := int((in.GetIndex() - 1) * in.GetSize())
	limit := int(in.GetSize())
	list, total, err := rgc.FindByPage(offset, limit)
	if err != nil {
		return nil, err
	}
	resultList := make([]*pb.ListItemReqGroup, 0, len(list))
	for _, item := range list {
		resultList = append(resultList, &pb.ListItemReqGroup{
			GroupId:     item.GroupID,
			Title:       item.Title,
			Key:         item.Key,
			Description: item.Description,
			CreateTime:  item.CreateTime.UnixMilli(),
			UpdateTime:  item.UpdateTime.UnixMilli(),
		})
	}
	return &pb.ListReqGroupResponse{
		Total:  total,
		Groups: resultList,
	}, nil
}

func (r ReqGroupServer) GetReqGroup(ctx context.Context, in *pb.GetReqGroupRequest) (*pb.ReqGroup, error) {
	return nil, nil
}

func (r ReqGroupServer) CreateReqGroup(ctx context.Context, in *pb.ReqGroup) (*pb.ReqGroupOperationResponse, error) {
	rg := query.ReqGroup
	rgo := query.ReqGroupOption
	rgc := rg.WithContext(ctx)
	rgoc := rgo.WithContext(ctx)

	//查询是否存在相同的key
	reqGroup, err := rg.Where(rg.Key.Eq(in.GetKey())).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if reqGroup != nil {
		return nil, alert.Error(alert.ReqGroupExist)
	}
	//检测数据
	for _, option := range in.GetOptions() {
		operationType := operation.ParseType(option.GetOperationType())
		attributeType := attribute.ParseType(option.GetAttrType())
		if !operationType.ForAttributeType(*attributeType) {
			return nil, alert.Error(alert.InvalidOperationType)
		}
	}
	//对option按照index做一下排序
	indexedOptions := make(map[int][]*pb.ReqGroupOption)
	for _, option := range in.GetOptions() {
		indexedOptions[int(option.GetIndex())] = append(indexedOptions[int(option.GetIndex())], option)
	}
	indexs := make([]int, 0)
	for i := range indexedOptions {
		indexs = append(indexs, i)
	}
	sort.Ints(indexs)
	//保存数据
	reqGroup = &model.ReqGroup{
		Title:       in.GetTitle(),
		Key:         in.GetKey(),
		Description: in.GetDescription(),
	}
	err = rgc.Save(reqGroup)
	if err != nil {
		return nil, err
	}
	//保存选项
	options := make([]*model.ReqGroupOption, 0)
	for _, index := range indexs {
		list := indexedOptions[index]
		for _, option := range list {
			reqGroupOption := &model.ReqGroupOption{
				GroupID:        reqGroup.GroupID,
				ListNum:        int64(index),
				AttributeType:  int8(enum.ParseAttributeType(option.GetAttrType())),
				AttributeName:  option.GetAttrName(),
				OperationType:  int8(enum.ParseOperationType(option.GetOperationType())),
				AttributeValue: option.GetAttrValue(),
			}
			err := rgoc.Save(reqGroupOption)
			if err != nil {
				return nil, err
			}
			options = append(options, reqGroupOption)
		}
	}
	return &pb.ReqGroupOperationResponse{
		GroupId: reqGroup.GroupID,
	}, nil
}

func (r ReqGroupServer) DeleteReqGroup(ctx context.Context, in *pb.GetReqGroupRequest) (*pb.ReqGroupOperationResponse, error) {
	rg := query.ReqGroup
	rgo := query.ReqGroupOption
	rgc := rg.WithContext(ctx)
	rgoc := rgo.WithContext(ctx)

	reqGroup, err := rgc.Where(rg.GroupID.Eq(in.GetGroupId())).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, alert.Error(alert.ReqGroupNotExist)
		}
		return nil, err
	}
	_, err = rgc.Where(rg.GroupID.Eq(in.GetGroupId())).Delete()
	if err != nil {
		return nil, err
	}
	_, err = rgoc.Where(rgo.GroupID.Eq(in.GetGroupId())).Delete()
	if err != nil {
		return nil, err
	}
	return &pb.ReqGroupOperationResponse{
		GroupId: reqGroup.GroupID,
	}, nil
}