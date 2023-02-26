// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratosent/ent/department"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DepartmentCreate is the builder for creating a Department entity.
type DepartmentCreate struct {
	config
	mutation *DepartmentMutation
	hooks    []Hook
}

// SetIsDeleted sets the "is_deleted" field.
func (dc *DepartmentCreate) SetIsDeleted(i int8) *DepartmentCreate {
	dc.mutation.SetIsDeleted(i)
	return dc
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableIsDeleted(i *int8) *DepartmentCreate {
	if i != nil {
		dc.SetIsDeleted(*i)
	}
	return dc
}

// SetCreateAt sets the "create_at" field.
func (dc *DepartmentCreate) SetCreateAt(t time.Time) *DepartmentCreate {
	dc.mutation.SetCreateAt(t)
	return dc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableCreateAt(t *time.Time) *DepartmentCreate {
	if t != nil {
		dc.SetCreateAt(*t)
	}
	return dc
}

// SetUpdateAt sets the "update_at" field.
func (dc *DepartmentCreate) SetUpdateAt(t time.Time) *DepartmentCreate {
	dc.mutation.SetUpdateAt(t)
	return dc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableUpdateAt(t *time.Time) *DepartmentCreate {
	if t != nil {
		dc.SetUpdateAt(*t)
	}
	return dc
}

// SetName sets the "name" field.
func (dc *DepartmentCreate) SetName(s string) *DepartmentCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetParentDepartmentID sets the "parent_department_id" field.
func (dc *DepartmentCreate) SetParentDepartmentID(i int64) *DepartmentCreate {
	dc.mutation.SetParentDepartmentID(i)
	return dc
}

// SetNillableParentDepartmentID sets the "parent_department_id" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableParentDepartmentID(i *int64) *DepartmentCreate {
	if i != nil {
		dc.SetParentDepartmentID(*i)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DepartmentCreate) SetID(i int64) *DepartmentCreate {
	dc.mutation.SetID(i)
	return dc
}

// Mutation returns the DepartmentMutation object of the builder.
func (dc *DepartmentCreate) Mutation() *DepartmentMutation {
	return dc.mutation
}

// Save creates the Department in the database.
func (dc *DepartmentCreate) Save(ctx context.Context) (*Department, error) {
	dc.defaults()
	return withHooks[*Department, DepartmentMutation](ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DepartmentCreate) SaveX(ctx context.Context) *Department {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DepartmentCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DepartmentCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DepartmentCreate) defaults() {
	if _, ok := dc.mutation.IsDeleted(); !ok {
		v := department.DefaultIsDeleted
		dc.mutation.SetIsDeleted(v)
	}
	if _, ok := dc.mutation.CreateAt(); !ok {
		v := department.DefaultCreateAt()
		dc.mutation.SetCreateAt(v)
	}
	if _, ok := dc.mutation.UpdateAt(); !ok {
		v := department.DefaultUpdateAt()
		dc.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DepartmentCreate) check() error {
	if _, ok := dc.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`ent: missing required field "Department.is_deleted"`)}
	}
	if _, ok := dc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "Department.create_at"`)}
	}
	if _, ok := dc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "Department.update_at"`)}
	}
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Department.name"`)}
	}
	if v, ok := dc.mutation.Name(); ok {
		if err := department.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Department.name": %w`, err)}
		}
	}
	return nil
}

func (dc *DepartmentCreate) sqlSave(ctx context.Context) (*Department, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DepartmentCreate) createSpec() (*Department, *sqlgraph.CreateSpec) {
	var (
		_node = &Department{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(department.Table, sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt64))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.IsDeleted(); ok {
		_spec.SetField(department.FieldIsDeleted, field.TypeInt8, value)
		_node.IsDeleted = value
	}
	if value, ok := dc.mutation.CreateAt(); ok {
		_spec.SetField(department.FieldCreateAt, field.TypeTime, value)
		_node.CreateAt = value
	}
	if value, ok := dc.mutation.UpdateAt(); ok {
		_spec.SetField(department.FieldUpdateAt, field.TypeTime, value)
		_node.UpdateAt = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(department.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.ParentDepartmentID(); ok {
		_spec.SetField(department.FieldParentDepartmentID, field.TypeInt64, value)
		_node.ParentDepartmentID = &value
	}
	return _node, _spec
}

// DepartmentCreateBulk is the builder for creating many Department entities in bulk.
type DepartmentCreateBulk struct {
	config
	builders []*DepartmentCreate
}

// Save creates the Department entities in the database.
func (dcb *DepartmentCreateBulk) Save(ctx context.Context) ([]*Department, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Department, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DepartmentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DepartmentCreateBulk) SaveX(ctx context.Context) []*Department {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DepartmentCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DepartmentCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
