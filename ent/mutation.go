// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratosent/ent/department"
	"kratosent/ent/predicate"
	"kratosent/ent/role"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeDepartment = "Department"
	TypeRole       = "Role"
)

// DepartmentMutation represents an operation that mutates the Department nodes in the graph.
type DepartmentMutation struct {
	config
	op                      Op
	typ                     string
	id                      *int64
	is_deleted              *int8
	addis_deleted           *int8
	create_at               *time.Time
	update_at               *time.Time
	name                    *string
	parent_department_id    *int64
	addparent_department_id *int64
	clearedFields           map[string]struct{}
	done                    bool
	oldValue                func(context.Context) (*Department, error)
	predicates              []predicate.Department
}

var _ ent.Mutation = (*DepartmentMutation)(nil)

// departmentOption allows management of the mutation configuration using functional options.
type departmentOption func(*DepartmentMutation)

// newDepartmentMutation creates new mutation for the Department entity.
func newDepartmentMutation(c config, op Op, opts ...departmentOption) *DepartmentMutation {
	m := &DepartmentMutation{
		config:        c,
		op:            op,
		typ:           TypeDepartment,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withDepartmentID sets the ID field of the mutation.
func withDepartmentID(id int64) departmentOption {
	return func(m *DepartmentMutation) {
		var (
			err   error
			once  sync.Once
			value *Department
		)
		m.oldValue = func(ctx context.Context) (*Department, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Department.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withDepartment sets the old Department of the mutation.
func withDepartment(node *Department) departmentOption {
	return func(m *DepartmentMutation) {
		m.oldValue = func(context.Context) (*Department, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m DepartmentMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m DepartmentMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Department entities.
func (m *DepartmentMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *DepartmentMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *DepartmentMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Department.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetIsDeleted sets the "is_deleted" field.
func (m *DepartmentMutation) SetIsDeleted(i int8) {
	m.is_deleted = &i
	m.addis_deleted = nil
}

// IsDeleted returns the value of the "is_deleted" field in the mutation.
func (m *DepartmentMutation) IsDeleted() (r int8, exists bool) {
	v := m.is_deleted
	if v == nil {
		return
	}
	return *v, true
}

// OldIsDeleted returns the old "is_deleted" field's value of the Department entity.
// If the Department object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DepartmentMutation) OldIsDeleted(ctx context.Context) (v int8, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsDeleted is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsDeleted requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsDeleted: %w", err)
	}
	return oldValue.IsDeleted, nil
}

// AddIsDeleted adds i to the "is_deleted" field.
func (m *DepartmentMutation) AddIsDeleted(i int8) {
	if m.addis_deleted != nil {
		*m.addis_deleted += i
	} else {
		m.addis_deleted = &i
	}
}

// AddedIsDeleted returns the value that was added to the "is_deleted" field in this mutation.
func (m *DepartmentMutation) AddedIsDeleted() (r int8, exists bool) {
	v := m.addis_deleted
	if v == nil {
		return
	}
	return *v, true
}

// ResetIsDeleted resets all changes to the "is_deleted" field.
func (m *DepartmentMutation) ResetIsDeleted() {
	m.is_deleted = nil
	m.addis_deleted = nil
}

// SetCreateAt sets the "create_at" field.
func (m *DepartmentMutation) SetCreateAt(t time.Time) {
	m.create_at = &t
}

// CreateAt returns the value of the "create_at" field in the mutation.
func (m *DepartmentMutation) CreateAt() (r time.Time, exists bool) {
	v := m.create_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateAt returns the old "create_at" field's value of the Department entity.
// If the Department object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DepartmentMutation) OldCreateAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateAt: %w", err)
	}
	return oldValue.CreateAt, nil
}

// ResetCreateAt resets all changes to the "create_at" field.
func (m *DepartmentMutation) ResetCreateAt() {
	m.create_at = nil
}

// SetUpdateAt sets the "update_at" field.
func (m *DepartmentMutation) SetUpdateAt(t time.Time) {
	m.update_at = &t
}

// UpdateAt returns the value of the "update_at" field in the mutation.
func (m *DepartmentMutation) UpdateAt() (r time.Time, exists bool) {
	v := m.update_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateAt returns the old "update_at" field's value of the Department entity.
// If the Department object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DepartmentMutation) OldUpdateAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdateAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdateAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateAt: %w", err)
	}
	return oldValue.UpdateAt, nil
}

// ResetUpdateAt resets all changes to the "update_at" field.
func (m *DepartmentMutation) ResetUpdateAt() {
	m.update_at = nil
}

// SetName sets the "name" field.
func (m *DepartmentMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *DepartmentMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Department entity.
// If the Department object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DepartmentMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *DepartmentMutation) ResetName() {
	m.name = nil
}

// SetParentDepartmentID sets the "parent_department_id" field.
func (m *DepartmentMutation) SetParentDepartmentID(i int64) {
	m.parent_department_id = &i
	m.addparent_department_id = nil
}

// ParentDepartmentID returns the value of the "parent_department_id" field in the mutation.
func (m *DepartmentMutation) ParentDepartmentID() (r int64, exists bool) {
	v := m.parent_department_id
	if v == nil {
		return
	}
	return *v, true
}

// OldParentDepartmentID returns the old "parent_department_id" field's value of the Department entity.
// If the Department object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DepartmentMutation) OldParentDepartmentID(ctx context.Context) (v *int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldParentDepartmentID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldParentDepartmentID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldParentDepartmentID: %w", err)
	}
	return oldValue.ParentDepartmentID, nil
}

// AddParentDepartmentID adds i to the "parent_department_id" field.
func (m *DepartmentMutation) AddParentDepartmentID(i int64) {
	if m.addparent_department_id != nil {
		*m.addparent_department_id += i
	} else {
		m.addparent_department_id = &i
	}
}

// AddedParentDepartmentID returns the value that was added to the "parent_department_id" field in this mutation.
func (m *DepartmentMutation) AddedParentDepartmentID() (r int64, exists bool) {
	v := m.addparent_department_id
	if v == nil {
		return
	}
	return *v, true
}

// ClearParentDepartmentID clears the value of the "parent_department_id" field.
func (m *DepartmentMutation) ClearParentDepartmentID() {
	m.parent_department_id = nil
	m.addparent_department_id = nil
	m.clearedFields[department.FieldParentDepartmentID] = struct{}{}
}

// ParentDepartmentIDCleared returns if the "parent_department_id" field was cleared in this mutation.
func (m *DepartmentMutation) ParentDepartmentIDCleared() bool {
	_, ok := m.clearedFields[department.FieldParentDepartmentID]
	return ok
}

// ResetParentDepartmentID resets all changes to the "parent_department_id" field.
func (m *DepartmentMutation) ResetParentDepartmentID() {
	m.parent_department_id = nil
	m.addparent_department_id = nil
	delete(m.clearedFields, department.FieldParentDepartmentID)
}

// Where appends a list predicates to the DepartmentMutation builder.
func (m *DepartmentMutation) Where(ps ...predicate.Department) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the DepartmentMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *DepartmentMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Department, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *DepartmentMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *DepartmentMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Department).
func (m *DepartmentMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *DepartmentMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.is_deleted != nil {
		fields = append(fields, department.FieldIsDeleted)
	}
	if m.create_at != nil {
		fields = append(fields, department.FieldCreateAt)
	}
	if m.update_at != nil {
		fields = append(fields, department.FieldUpdateAt)
	}
	if m.name != nil {
		fields = append(fields, department.FieldName)
	}
	if m.parent_department_id != nil {
		fields = append(fields, department.FieldParentDepartmentID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *DepartmentMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case department.FieldIsDeleted:
		return m.IsDeleted()
	case department.FieldCreateAt:
		return m.CreateAt()
	case department.FieldUpdateAt:
		return m.UpdateAt()
	case department.FieldName:
		return m.Name()
	case department.FieldParentDepartmentID:
		return m.ParentDepartmentID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *DepartmentMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case department.FieldIsDeleted:
		return m.OldIsDeleted(ctx)
	case department.FieldCreateAt:
		return m.OldCreateAt(ctx)
	case department.FieldUpdateAt:
		return m.OldUpdateAt(ctx)
	case department.FieldName:
		return m.OldName(ctx)
	case department.FieldParentDepartmentID:
		return m.OldParentDepartmentID(ctx)
	}
	return nil, fmt.Errorf("unknown Department field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DepartmentMutation) SetField(name string, value ent.Value) error {
	switch name {
	case department.FieldIsDeleted:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsDeleted(v)
		return nil
	case department.FieldCreateAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateAt(v)
		return nil
	case department.FieldUpdateAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateAt(v)
		return nil
	case department.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case department.FieldParentDepartmentID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetParentDepartmentID(v)
		return nil
	}
	return fmt.Errorf("unknown Department field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *DepartmentMutation) AddedFields() []string {
	var fields []string
	if m.addis_deleted != nil {
		fields = append(fields, department.FieldIsDeleted)
	}
	if m.addparent_department_id != nil {
		fields = append(fields, department.FieldParentDepartmentID)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *DepartmentMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case department.FieldIsDeleted:
		return m.AddedIsDeleted()
	case department.FieldParentDepartmentID:
		return m.AddedParentDepartmentID()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DepartmentMutation) AddField(name string, value ent.Value) error {
	switch name {
	case department.FieldIsDeleted:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddIsDeleted(v)
		return nil
	case department.FieldParentDepartmentID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddParentDepartmentID(v)
		return nil
	}
	return fmt.Errorf("unknown Department numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *DepartmentMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(department.FieldParentDepartmentID) {
		fields = append(fields, department.FieldParentDepartmentID)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *DepartmentMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *DepartmentMutation) ClearField(name string) error {
	switch name {
	case department.FieldParentDepartmentID:
		m.ClearParentDepartmentID()
		return nil
	}
	return fmt.Errorf("unknown Department nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *DepartmentMutation) ResetField(name string) error {
	switch name {
	case department.FieldIsDeleted:
		m.ResetIsDeleted()
		return nil
	case department.FieldCreateAt:
		m.ResetCreateAt()
		return nil
	case department.FieldUpdateAt:
		m.ResetUpdateAt()
		return nil
	case department.FieldName:
		m.ResetName()
		return nil
	case department.FieldParentDepartmentID:
		m.ResetParentDepartmentID()
		return nil
	}
	return fmt.Errorf("unknown Department field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *DepartmentMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *DepartmentMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *DepartmentMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *DepartmentMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *DepartmentMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *DepartmentMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *DepartmentMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Department unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *DepartmentMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Department edge %s", name)
}

// RoleMutation represents an operation that mutates the Role nodes in the graph.
type RoleMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	is_deleted    *int8
	addis_deleted *int8
	create_at     *time.Time
	update_at     *time.Time
	name          *string
	description   *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Role, error)
	predicates    []predicate.Role
}

var _ ent.Mutation = (*RoleMutation)(nil)

// roleOption allows management of the mutation configuration using functional options.
type roleOption func(*RoleMutation)

// newRoleMutation creates new mutation for the Role entity.
func newRoleMutation(c config, op Op, opts ...roleOption) *RoleMutation {
	m := &RoleMutation{
		config:        c,
		op:            op,
		typ:           TypeRole,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRoleID sets the ID field of the mutation.
func withRoleID(id int64) roleOption {
	return func(m *RoleMutation) {
		var (
			err   error
			once  sync.Once
			value *Role
		)
		m.oldValue = func(ctx context.Context) (*Role, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Role.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRole sets the old Role of the mutation.
func withRole(node *Role) roleOption {
	return func(m *RoleMutation) {
		m.oldValue = func(context.Context) (*Role, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RoleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RoleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Role entities.
func (m *RoleMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *RoleMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *RoleMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Role.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetIsDeleted sets the "is_deleted" field.
func (m *RoleMutation) SetIsDeleted(i int8) {
	m.is_deleted = &i
	m.addis_deleted = nil
}

// IsDeleted returns the value of the "is_deleted" field in the mutation.
func (m *RoleMutation) IsDeleted() (r int8, exists bool) {
	v := m.is_deleted
	if v == nil {
		return
	}
	return *v, true
}

// OldIsDeleted returns the old "is_deleted" field's value of the Role entity.
// If the Role object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RoleMutation) OldIsDeleted(ctx context.Context) (v int8, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsDeleted is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsDeleted requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsDeleted: %w", err)
	}
	return oldValue.IsDeleted, nil
}

// AddIsDeleted adds i to the "is_deleted" field.
func (m *RoleMutation) AddIsDeleted(i int8) {
	if m.addis_deleted != nil {
		*m.addis_deleted += i
	} else {
		m.addis_deleted = &i
	}
}

// AddedIsDeleted returns the value that was added to the "is_deleted" field in this mutation.
func (m *RoleMutation) AddedIsDeleted() (r int8, exists bool) {
	v := m.addis_deleted
	if v == nil {
		return
	}
	return *v, true
}

// ResetIsDeleted resets all changes to the "is_deleted" field.
func (m *RoleMutation) ResetIsDeleted() {
	m.is_deleted = nil
	m.addis_deleted = nil
}

// SetCreateAt sets the "create_at" field.
func (m *RoleMutation) SetCreateAt(t time.Time) {
	m.create_at = &t
}

// CreateAt returns the value of the "create_at" field in the mutation.
func (m *RoleMutation) CreateAt() (r time.Time, exists bool) {
	v := m.create_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateAt returns the old "create_at" field's value of the Role entity.
// If the Role object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RoleMutation) OldCreateAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateAt: %w", err)
	}
	return oldValue.CreateAt, nil
}

// ResetCreateAt resets all changes to the "create_at" field.
func (m *RoleMutation) ResetCreateAt() {
	m.create_at = nil
}

// SetUpdateAt sets the "update_at" field.
func (m *RoleMutation) SetUpdateAt(t time.Time) {
	m.update_at = &t
}

// UpdateAt returns the value of the "update_at" field in the mutation.
func (m *RoleMutation) UpdateAt() (r time.Time, exists bool) {
	v := m.update_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateAt returns the old "update_at" field's value of the Role entity.
// If the Role object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RoleMutation) OldUpdateAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdateAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdateAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateAt: %w", err)
	}
	return oldValue.UpdateAt, nil
}

// ResetUpdateAt resets all changes to the "update_at" field.
func (m *RoleMutation) ResetUpdateAt() {
	m.update_at = nil
}

// SetName sets the "name" field.
func (m *RoleMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *RoleMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Role entity.
// If the Role object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RoleMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *RoleMutation) ResetName() {
	m.name = nil
}

// SetDescription sets the "description" field.
func (m *RoleMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *RoleMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Role entity.
// If the Role object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RoleMutation) OldDescription(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ClearDescription clears the value of the "description" field.
func (m *RoleMutation) ClearDescription() {
	m.description = nil
	m.clearedFields[role.FieldDescription] = struct{}{}
}

// DescriptionCleared returns if the "description" field was cleared in this mutation.
func (m *RoleMutation) DescriptionCleared() bool {
	_, ok := m.clearedFields[role.FieldDescription]
	return ok
}

// ResetDescription resets all changes to the "description" field.
func (m *RoleMutation) ResetDescription() {
	m.description = nil
	delete(m.clearedFields, role.FieldDescription)
}

// Where appends a list predicates to the RoleMutation builder.
func (m *RoleMutation) Where(ps ...predicate.Role) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the RoleMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *RoleMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Role, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *RoleMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *RoleMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Role).
func (m *RoleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *RoleMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.is_deleted != nil {
		fields = append(fields, role.FieldIsDeleted)
	}
	if m.create_at != nil {
		fields = append(fields, role.FieldCreateAt)
	}
	if m.update_at != nil {
		fields = append(fields, role.FieldUpdateAt)
	}
	if m.name != nil {
		fields = append(fields, role.FieldName)
	}
	if m.description != nil {
		fields = append(fields, role.FieldDescription)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *RoleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case role.FieldIsDeleted:
		return m.IsDeleted()
	case role.FieldCreateAt:
		return m.CreateAt()
	case role.FieldUpdateAt:
		return m.UpdateAt()
	case role.FieldName:
		return m.Name()
	case role.FieldDescription:
		return m.Description()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *RoleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case role.FieldIsDeleted:
		return m.OldIsDeleted(ctx)
	case role.FieldCreateAt:
		return m.OldCreateAt(ctx)
	case role.FieldUpdateAt:
		return m.OldUpdateAt(ctx)
	case role.FieldName:
		return m.OldName(ctx)
	case role.FieldDescription:
		return m.OldDescription(ctx)
	}
	return nil, fmt.Errorf("unknown Role field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RoleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case role.FieldIsDeleted:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsDeleted(v)
		return nil
	case role.FieldCreateAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateAt(v)
		return nil
	case role.FieldUpdateAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateAt(v)
		return nil
	case role.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case role.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	}
	return fmt.Errorf("unknown Role field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *RoleMutation) AddedFields() []string {
	var fields []string
	if m.addis_deleted != nil {
		fields = append(fields, role.FieldIsDeleted)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *RoleMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case role.FieldIsDeleted:
		return m.AddedIsDeleted()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RoleMutation) AddField(name string, value ent.Value) error {
	switch name {
	case role.FieldIsDeleted:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddIsDeleted(v)
		return nil
	}
	return fmt.Errorf("unknown Role numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *RoleMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(role.FieldDescription) {
		fields = append(fields, role.FieldDescription)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *RoleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *RoleMutation) ClearField(name string) error {
	switch name {
	case role.FieldDescription:
		m.ClearDescription()
		return nil
	}
	return fmt.Errorf("unknown Role nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *RoleMutation) ResetField(name string) error {
	switch name {
	case role.FieldIsDeleted:
		m.ResetIsDeleted()
		return nil
	case role.FieldCreateAt:
		m.ResetCreateAt()
		return nil
	case role.FieldUpdateAt:
		m.ResetUpdateAt()
		return nil
	case role.FieldName:
		m.ResetName()
		return nil
	case role.FieldDescription:
		m.ResetDescription()
		return nil
	}
	return fmt.Errorf("unknown Role field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *RoleMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RoleMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RoleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RoleMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RoleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RoleMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RoleMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Role unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RoleMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Role edge %s", name)
}
