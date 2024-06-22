package provider

import (
	"context"
	"encoding/json"
	"feature-distributor/common/alert"
	"feature-distributor/core/db"
	"feature-distributor/core/db/enum"
	"feature-distributor/core/db/query"
	"fmt"
	"github.com/bombsimon/logrusr/v4"
	"github.com/open-feature/go-sdk/openfeature"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
)

func Init() error {
	log := logrusr.New(logrus.WithField("source", "grpc"))
	openfeature.SetLogger(log)
	return openfeature.SetProvider(CoreProvider{})
}

type CoreProvider struct {
}

func (i CoreProvider) Metadata() openfeature.Metadata {
	return openfeature.Metadata{
		Name: "feature-distributor",
	}
}

func (i CoreProvider) Hooks() []openfeature.Hook {
	return []openfeature.Hook{}
}

// resolutionStatic 没有动态规则时的详情信息
func resolutionStatic() openfeature.ProviderResolutionDetail {
	return openfeature.ProviderResolutionDetail{
		Reason: openfeature.StaticReason,
	}
}

// resolutionDynamic 动态规则计算之后返回的详情信息
func resolutionDynamic() openfeature.ProviderResolutionDetail {
	return openfeature.ProviderResolutionDetail{
		Reason: openfeature.TargetingMatchReason,
	}
}

// resolutionDisable 禁用规则后返回的详情信息
func resolutionDisable() openfeature.ProviderResolutionDetail {
	return openfeature.ProviderResolutionDetail{
		Reason: openfeature.DisabledReason,
	}
}

func genericResolve[T comparable](ctx context.Context, evalCtx openfeature.FlattenedContext, toggleKey string, toggleType enum.ValueType, defaultValue T, m func(string) (T, error)) (T, openfeature.ProviderResolutionDetail) {
	projectKey := ctx.Value("projectKey").(string)
	p := query.Project
	project, err := p.Where(p.Key.Eq(projectKey)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return defaultValue, openfeature.ProviderResolutionDetail{
				ResolutionError: openfeature.NewFlagNotFoundResolutionError(fmt.Sprintf("%d", alert.ProjectNotExist)),
				Reason:          openfeature.ErrorReason,
			}
		}
		logrus.Errorf("failed to get project, %v", err)
		return defaultValue, openfeature.ProviderResolutionDetail{
			ResolutionError: openfeature.NewGeneralResolutionError(fmt.Sprintf("%d", alert.ServerInternalError)),
			Reason:          openfeature.UnknownReason,
		}
	}
	toggle, toggleValues, err := db.SelectToggleValues(ctx, project.ID, toggleKey)
	if err != nil {
		logrus.Errorf("failed to get toggle, %v", err)
		return defaultValue, openfeature.ProviderResolutionDetail{
			ResolutionError: openfeature.NewGeneralResolutionError(fmt.Sprintf("%d", alert.ServerInternalError)),
			Reason:          openfeature.UnknownReason,
		}
	}
	if toggle == nil {
		return defaultValue, openfeature.ProviderResolutionDetail{
			ResolutionError: openfeature.NewFlagNotFoundResolutionError(fmt.Sprintf("%d", alert.ToggleNotExist)),
			Reason:          openfeature.ErrorReason,
		}
	}
	if enum.ValueTypeEnum(toggle.ValueType) != toggleType {
		return defaultValue, openfeature.ProviderResolutionDetail{
			ResolutionError: openfeature.NewTypeMismatchResolutionError(fmt.Sprintf("%d", alert.InvalidToggleType)),
			Reason:          openfeature.ErrorReason,
		}
	}
	valueMap := make(map[int64]T)
	for _, value := range toggleValues {
		vv, err := m(value.Value)
		if err != nil {
			logrus.Errorf("failed to parse value, %v", err)
			return defaultValue, openfeature.ProviderResolutionDetail{
				ResolutionError: openfeature.NewParseErrorResolutionError(fmt.Sprintf("%d", alert.ServerInternalError)),
				Reason:          openfeature.ErrorReason,
			}
		}
		valueMap[value.ID] = vv
	}
	rolloutKey := evalCtx[openfeature.TargetingKey]
	logrus.Infof("rolloutKey: %s", rolloutKey)
	//TODO 灰度规则
	if !toggle.Enable {
		return valueMap[toggle.ReturnValueWhenDisable], resolutionDisable()
	}
	return valueMap[toggle.DefaultValue], resolutionStatic()
}

func (i CoreProvider) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool, evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {
	v, r := genericResolve(ctx, evalCtx, flag, enum.BoolValueType, defaultValue, func(s string) (bool, error) {
		switch s {
		case "true":
			return true, nil
		case "false":
			return false, nil
		default:
			return false, errors.New(fmt.Sprintf("invalid bool value: %s", s))
		}
	})
	return openfeature.BoolResolutionDetail{
		Value:                    v,
		ProviderResolutionDetail: r,
	}
}

func (i CoreProvider) StringEvaluation(ctx context.Context, flag string, defaultValue string, evalCtx openfeature.FlattenedContext) openfeature.StringResolutionDetail {
	v, r := genericResolve(ctx, evalCtx, flag, enum.StringValueType, defaultValue, func(s string) (string, error) {
		return s, nil
	})
	return openfeature.StringResolutionDetail{
		Value:                    v,
		ProviderResolutionDetail: r,
	}
}

func (i CoreProvider) FloatEvaluation(ctx context.Context, flag string, defaultValue float64, evalCtx openfeature.FlattenedContext) openfeature.FloatResolutionDetail {
	v, r := genericResolve(ctx, evalCtx, flag, enum.FloatValueType, defaultValue, func(s string) (float64, error) {
		return strconv.ParseFloat(s, 64)
	})
	return openfeature.FloatResolutionDetail{
		Value:                    v,
		ProviderResolutionDetail: r,
	}
}

func (i CoreProvider) IntEvaluation(ctx context.Context, flag string, defaultValue int64, evalCtx openfeature.FlattenedContext) openfeature.IntResolutionDetail {
	v, r := genericResolve(ctx, evalCtx, flag, enum.IntValueType, defaultValue, func(s string) (int64, error) {
		return strconv.ParseInt(s, 10, 64)
	})
	return openfeature.IntResolutionDetail{
		Value:                    v,
		ProviderResolutionDetail: r,
	}
}

func (i CoreProvider) ObjectEvaluation(ctx context.Context, flag string, defaultValue interface{}, evalCtx openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {
	v, r := genericResolve(ctx, evalCtx, flag, enum.JsonValueType, defaultValue, func(s string) (interface{}, error) {
		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(s), &m)
		return m, err
	})
	return openfeature.InterfaceResolutionDetail{
		Value:                    v,
		ProviderResolutionDetail: r,
	}
}

func (i CoreProvider) Init(evaluationContext openfeature.EvaluationContext) error {
	return nil
}

func (i CoreProvider) Status() openfeature.State {
	return openfeature.ReadyState
}

func (i CoreProvider) Shutdown() {
}
