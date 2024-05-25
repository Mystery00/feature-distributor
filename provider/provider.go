package provider

import (
	"context"
	"github.com/open-feature/go-sdk/openfeature"
)

type FeatureDistributorProvider struct{}

func (i FeatureDistributorProvider) Metadata() openfeature.Metadata {
	return openfeature.Metadata{
		Name: "feature-distributor",
	}
}

func (i FeatureDistributorProvider) Hooks() []openfeature.Hook {
	return []openfeature.Hook{}
}

func (i FeatureDistributorProvider) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool, evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {

}

func (i FeatureDistributorProvider) StringEvaluation(ctx context.Context, flag string, defaultValue string, evalCtx openfeature.FlattenedContext) openfeature.StringResolutionDetail {

}

func (i FeatureDistributorProvider) FloatEvaluation(ctx context.Context, flag string, defaultValue float64, evalCtx openfeature.FlattenedContext) openfeature.FloatResolutionDetail {

}

func (i FeatureDistributorProvider) IntEvaluation(ctx context.Context, flag string, defaultValue int64, evalCtx openfeature.FlattenedContext) openfeature.IntResolutionDetail {

}

func (i FeatureDistributorProvider) ObjectEvaluation(ctx context.Context, flag string, defaultValue interface{}, evalCtx openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {

}

func (i FeatureDistributorProvider) Init(evaluationContext openfeature.EvaluationContext) error {

}

func (i FeatureDistributorProvider) Status() openfeature.State {
	return openfeature.ReadyState
}

func (i FeatureDistributorProvider) Shutdown() {

}

func (i FeatureDistributorProvider) EventChannel() <-chan openfeature.Event {

}
