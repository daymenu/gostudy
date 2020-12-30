// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/daymenu/gostudy/examples/ent/ent/pet"
	"github.com/daymenu/gostudy/examples/ent/ent/schema"
	"github.com/daymenu/gostudy/examples/ent/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	petMixin := schema.Pet{}.Mixin()
	petMixinFields0 := petMixin[0].Fields()
	petMixinFields1 := petMixin[1].Fields()
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescCreateTime is the schema descriptor for create_time field.
	petDescCreateTime := petMixinFields0[0].Descriptor()
	// pet.DefaultCreateTime holds the default value on creation for the create_time field.
	pet.DefaultCreateTime = petDescCreateTime.Default.(func() time.Time)
	// petDescUpdateTime is the schema descriptor for update_time field.
	petDescUpdateTime := petMixinFields0[1].Descriptor()
	// pet.DefaultUpdateTime holds the default value on creation for the update_time field.
	pet.DefaultUpdateTime = petDescUpdateTime.Default.(func() time.Time)
	// pet.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	pet.UpdateDefaultUpdateTime = petDescUpdateTime.UpdateDefault.(func() time.Time)
	// petDescStatus is the schema descriptor for status field.
	petDescStatus := petMixinFields1[0].Descriptor()
	// pet.DefaultStatus holds the default value on creation for the status field.
	pet.DefaultStatus = petDescStatus.Default.(int)
	// petDescName is the schema descriptor for name field.
	petDescName := petFields[0].Descriptor()
	// pet.DefaultName holds the default value on creation for the name field.
	pet.DefaultName = petDescName.Default.(string)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	userMixinFields1 := userMixin[1].Fields()
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userMixinFields1[0].Descriptor()
	// user.DefaultStatus holds the default value on creation for the status field.
	user.DefaultStatus = userDescStatus.Default.(int)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.DefaultEmail holds the default value on creation for the email field.
	user.DefaultEmail = userDescEmail.Default.(string)
}
