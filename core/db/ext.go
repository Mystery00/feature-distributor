package db

import (
	"context"
	"feature-distributor/core/db/model"
	"feature-distributor/core/db/query"
)

func SelectToggleValues(ctx context.Context, projectId int64, toggleKey string) (model.Toggle, []model.ToggleValue, error) {
	t := query.Toggle
	tv := query.ToggleValue
	var toggle model.Toggle
	err := t.WithContext(ctx).Select(t.ALL).Where(t.ProjectID.Eq(projectId), t.Key.Eq(toggleKey)).Scan(&toggle)
	if err != nil {
		return model.Toggle{}, nil, err
	}
	values := make([]model.ToggleValue, 0)
	err = t.WithContext(ctx).Select(tv.ALL).LeftJoin(tv, t.ID.EqCol(tv.ToggleID)).Where(t.ProjectID.Eq(projectId), t.Key.Eq(toggleKey)).Scan(&values)
	return toggle, values, err
}
