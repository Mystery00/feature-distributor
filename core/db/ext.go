package db

import (
	"context"
	"errors"
	"feature-distributor/core/db/model"
	"feature-distributor/core/db/query"
	"gorm.io/gorm"
)

func SelectToggleValues(ctx context.Context, projectId int64, toggleKey string) (*model.Toggle, []model.ToggleValue, error) {
	t := query.Toggle
	tv := query.ToggleValue
	toggle, err := t.WithContext(ctx).Where(t.Key.Eq(toggleKey), t.ProjectID.Eq(projectId)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, nil
	}
	values := make([]model.ToggleValue, 0)
	err = t.WithContext(ctx).Select(tv.ALL).LeftJoin(tv, t.ID.EqCol(tv.ToggleID)).Where(t.ProjectID.Eq(projectId), t.Key.Eq(toggleKey)).Scan(&values)
	return toggle, values, err
}
