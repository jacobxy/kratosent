// Code generated by ent, DO NOT EDIT.

package ent

import (
	"kratosent/ent/department"
	"kratosent/ent/role"
	"kratosent/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	departmentMixin := schema.Department{}.Mixin()
	departmentMixinFields0 := departmentMixin[0].Fields()
	_ = departmentMixinFields0
	departmentFields := schema.Department{}.Fields()
	_ = departmentFields
	// departmentDescIsDeleted is the schema descriptor for is_deleted field.
	departmentDescIsDeleted := departmentMixinFields0[1].Descriptor()
	// department.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	department.DefaultIsDeleted = int8(departmentDescIsDeleted.Default.(int8))
	// departmentDescCreateAt is the schema descriptor for create_at field.
	departmentDescCreateAt := departmentMixinFields0[2].Descriptor()
	// department.DefaultCreateAt holds the default value on creation for the create_at field.
	department.DefaultCreateAt = departmentDescCreateAt.Default.(func() time.Time)
	// departmentDescUpdateAt is the schema descriptor for update_at field.
	departmentDescUpdateAt := departmentMixinFields0[3].Descriptor()
	// department.DefaultUpdateAt holds the default value on creation for the update_at field.
	department.DefaultUpdateAt = departmentDescUpdateAt.Default.(func() time.Time)
	// department.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	department.UpdateDefaultUpdateAt = departmentDescUpdateAt.UpdateDefault.(func() time.Time)
	// departmentDescName is the schema descriptor for name field.
	departmentDescName := departmentFields[0].Descriptor()
	// department.NameValidator is a validator for the "name" field. It is called by the builders before save.
	department.NameValidator = departmentDescName.Validators[0].(func(string) error)
	roleMixin := schema.Role{}.Mixin()
	roleMixinFields0 := roleMixin[0].Fields()
	_ = roleMixinFields0
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescIsDeleted is the schema descriptor for is_deleted field.
	roleDescIsDeleted := roleMixinFields0[1].Descriptor()
	// role.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	role.DefaultIsDeleted = int8(roleDescIsDeleted.Default.(int8))
	// roleDescCreateAt is the schema descriptor for create_at field.
	roleDescCreateAt := roleMixinFields0[2].Descriptor()
	// role.DefaultCreateAt holds the default value on creation for the create_at field.
	role.DefaultCreateAt = roleDescCreateAt.Default.(func() time.Time)
	// roleDescUpdateAt is the schema descriptor for update_at field.
	roleDescUpdateAt := roleMixinFields0[3].Descriptor()
	// role.DefaultUpdateAt holds the default value on creation for the update_at field.
	role.DefaultUpdateAt = roleDescUpdateAt.Default.(func() time.Time)
	// role.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	role.UpdateDefaultUpdateAt = roleDescUpdateAt.UpdateDefault.(func() time.Time)
	// roleDescName is the schema descriptor for name field.
	roleDescName := roleFields[0].Descriptor()
	// role.NameValidator is a validator for the "name" field. It is called by the builders before save.
	role.NameValidator = roleDescName.Validators[0].(func(string) error)
}
