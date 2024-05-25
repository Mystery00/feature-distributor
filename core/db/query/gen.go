// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q           = new(Query)
	Project     *project
	Toggle      *toggle
	ToggleValue *toggleValue
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Project = &Q.Project
	Toggle = &Q.Toggle
	ToggleValue = &Q.ToggleValue
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:          db,
		Project:     newProject(db, opts...),
		Toggle:      newToggle(db, opts...),
		ToggleValue: newToggleValue(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Project     project
	Toggle      toggle
	ToggleValue toggleValue
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:          db,
		Project:     q.Project.clone(db),
		Toggle:      q.Toggle.clone(db),
		ToggleValue: q.ToggleValue.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:          db,
		Project:     q.Project.replaceDB(db),
		Toggle:      q.Toggle.replaceDB(db),
		ToggleValue: q.ToggleValue.replaceDB(db),
	}
}

type queryCtx struct {
	Project     IProjectDo
	Toggle      IToggleDo
	ToggleValue IToggleValueDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Project:     q.Project.WithContext(ctx),
		Toggle:      q.Toggle.WithContext(ctx),
		ToggleValue: q.ToggleValue.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}