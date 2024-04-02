// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"news-subscriber-bff/internal/data/ent/articlemeta"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ArticleMetaCreate is the builder for creating a ArticleMeta entity.
type ArticleMetaCreate struct {
	config
	mutation *ArticleMetaMutation
	hooks    []Hook
}

// SetUID sets the "uid" field.
func (amc *ArticleMetaCreate) SetUID(i int64) *ArticleMetaCreate {
	amc.mutation.SetUID(i)
	return amc
}

// SetTitle sets the "title" field.
func (amc *ArticleMetaCreate) SetTitle(s string) *ArticleMetaCreate {
	amc.mutation.SetTitle(s)
	return amc
}

// SetTags sets the "tags" field.
func (amc *ArticleMetaCreate) SetTags(s []string) *ArticleMetaCreate {
	amc.mutation.SetTags(s)
	return amc
}

// SetCreateTime sets the "createTime" field.
func (amc *ArticleMetaCreate) SetCreateTime(t time.Time) *ArticleMetaCreate {
	amc.mutation.SetCreateTime(t)
	return amc
}

// SetNillableCreateTime sets the "createTime" field if the given value is not nil.
func (amc *ArticleMetaCreate) SetNillableCreateTime(t *time.Time) *ArticleMetaCreate {
	if t != nil {
		amc.SetCreateTime(*t)
	}
	return amc
}

// SetUpdateTime sets the "updateTime" field.
func (amc *ArticleMetaCreate) SetUpdateTime(t time.Time) *ArticleMetaCreate {
	amc.mutation.SetUpdateTime(t)
	return amc
}

// SetNillableUpdateTime sets the "updateTime" field if the given value is not nil.
func (amc *ArticleMetaCreate) SetNillableUpdateTime(t *time.Time) *ArticleMetaCreate {
	if t != nil {
		amc.SetUpdateTime(*t)
	}
	return amc
}

// Mutation returns the ArticleMetaMutation object of the builder.
func (amc *ArticleMetaCreate) Mutation() *ArticleMetaMutation {
	return amc.mutation
}

// Save creates the ArticleMeta in the database.
func (amc *ArticleMetaCreate) Save(ctx context.Context) (*ArticleMeta, error) {
	amc.defaults()
	return withHooks(ctx, amc.sqlSave, amc.mutation, amc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (amc *ArticleMetaCreate) SaveX(ctx context.Context) *ArticleMeta {
	v, err := amc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (amc *ArticleMetaCreate) Exec(ctx context.Context) error {
	_, err := amc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amc *ArticleMetaCreate) ExecX(ctx context.Context) {
	if err := amc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (amc *ArticleMetaCreate) defaults() {
	if _, ok := amc.mutation.CreateTime(); !ok {
		v := articlemeta.DefaultCreateTime()
		amc.mutation.SetCreateTime(v)
	}
	if _, ok := amc.mutation.UpdateTime(); !ok {
		v := articlemeta.DefaultUpdateTime()
		amc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (amc *ArticleMetaCreate) check() error {
	if _, ok := amc.mutation.UID(); !ok {
		return &ValidationError{Name: "uid", err: errors.New(`ent: missing required field "ArticleMeta.uid"`)}
	}
	if _, ok := amc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "ArticleMeta.title"`)}
	}
	if v, ok := amc.mutation.Title(); ok {
		if err := articlemeta.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "ArticleMeta.title": %w`, err)}
		}
	}
	if _, ok := amc.mutation.Tags(); !ok {
		return &ValidationError{Name: "tags", err: errors.New(`ent: missing required field "ArticleMeta.tags"`)}
	}
	if _, ok := amc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "createTime", err: errors.New(`ent: missing required field "ArticleMeta.createTime"`)}
	}
	if _, ok := amc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "updateTime", err: errors.New(`ent: missing required field "ArticleMeta.updateTime"`)}
	}
	return nil
}

func (amc *ArticleMetaCreate) sqlSave(ctx context.Context) (*ArticleMeta, error) {
	if err := amc.check(); err != nil {
		return nil, err
	}
	_node, _spec := amc.createSpec()
	if err := sqlgraph.CreateNode(ctx, amc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	amc.mutation.id = &_node.ID
	amc.mutation.done = true
	return _node, nil
}

func (amc *ArticleMetaCreate) createSpec() (*ArticleMeta, *sqlgraph.CreateSpec) {
	var (
		_node = &ArticleMeta{config: amc.config}
		_spec = sqlgraph.NewCreateSpec(articlemeta.Table, sqlgraph.NewFieldSpec(articlemeta.FieldID, field.TypeInt))
	)
	if value, ok := amc.mutation.UID(); ok {
		_spec.SetField(articlemeta.FieldUID, field.TypeInt64, value)
		_node.UID = value
	}
	if value, ok := amc.mutation.Title(); ok {
		_spec.SetField(articlemeta.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := amc.mutation.Tags(); ok {
		_spec.SetField(articlemeta.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := amc.mutation.CreateTime(); ok {
		_spec.SetField(articlemeta.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := amc.mutation.UpdateTime(); ok {
		_spec.SetField(articlemeta.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	return _node, _spec
}

// ArticleMetaCreateBulk is the builder for creating many ArticleMeta entities in bulk.
type ArticleMetaCreateBulk struct {
	config
	err      error
	builders []*ArticleMetaCreate
}

// Save creates the ArticleMeta entities in the database.
func (amcb *ArticleMetaCreateBulk) Save(ctx context.Context) ([]*ArticleMeta, error) {
	if amcb.err != nil {
		return nil, amcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(amcb.builders))
	nodes := make([]*ArticleMeta, len(amcb.builders))
	mutators := make([]Mutator, len(amcb.builders))
	for i := range amcb.builders {
		func(i int, root context.Context) {
			builder := amcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArticleMetaMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, amcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, amcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, amcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (amcb *ArticleMetaCreateBulk) SaveX(ctx context.Context) []*ArticleMeta {
	v, err := amcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (amcb *ArticleMetaCreateBulk) Exec(ctx context.Context) error {
	_, err := amcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amcb *ArticleMetaCreateBulk) ExecX(ctx context.Context) {
	if err := amcb.Exec(ctx); err != nil {
		panic(err)
	}
}