// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/matsuokashuhei/landin/ent/class"
	"github.com/matsuokashuhei/landin/ent/instructor"
	"github.com/matsuokashuhei/landin/ent/member"
	"github.com/matsuokashuhei/landin/ent/membersclass"
	"github.com/matsuokashuhei/landin/ent/room"
	"github.com/matsuokashuhei/landin/ent/schedule"
	"github.com/matsuokashuhei/landin/ent/schema"
	"github.com/matsuokashuhei/landin/ent/school"
	"github.com/matsuokashuhei/landin/ent/studio"
	"github.com/matsuokashuhei/landin/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	classMixin := schema.Class{}.Mixin()
	classMixinFields0 := classMixin[0].Fields()
	_ = classMixinFields0
	classFields := schema.Class{}.Fields()
	_ = classFields
	// classDescCreateTime is the schema descriptor for create_time field.
	classDescCreateTime := classMixinFields0[0].Descriptor()
	// class.DefaultCreateTime holds the default value on creation for the create_time field.
	class.DefaultCreateTime = classDescCreateTime.Default.(func() time.Time)
	// classDescUpdateTime is the schema descriptor for update_time field.
	classDescUpdateTime := classMixinFields0[1].Descriptor()
	// class.DefaultUpdateTime holds the default value on creation for the update_time field.
	class.DefaultUpdateTime = classDescUpdateTime.Default.(func() time.Time)
	// class.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	class.UpdateDefaultUpdateTime = classDescUpdateTime.UpdateDefault.(func() time.Time)
	instructorMixin := schema.Instructor{}.Mixin()
	instructorMixinFields0 := instructorMixin[0].Fields()
	_ = instructorMixinFields0
	instructorFields := schema.Instructor{}.Fields()
	_ = instructorFields
	// instructorDescCreateTime is the schema descriptor for create_time field.
	instructorDescCreateTime := instructorMixinFields0[0].Descriptor()
	// instructor.DefaultCreateTime holds the default value on creation for the create_time field.
	instructor.DefaultCreateTime = instructorDescCreateTime.Default.(func() time.Time)
	// instructorDescUpdateTime is the schema descriptor for update_time field.
	instructorDescUpdateTime := instructorMixinFields0[1].Descriptor()
	// instructor.DefaultUpdateTime holds the default value on creation for the update_time field.
	instructor.DefaultUpdateTime = instructorDescUpdateTime.Default.(func() time.Time)
	// instructor.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	instructor.UpdateDefaultUpdateTime = instructorDescUpdateTime.UpdateDefault.(func() time.Time)
	memberMixin := schema.Member{}.Mixin()
	memberMixinFields0 := memberMixin[0].Fields()
	_ = memberMixinFields0
	memberFields := schema.Member{}.Fields()
	_ = memberFields
	// memberDescCreateTime is the schema descriptor for create_time field.
	memberDescCreateTime := memberMixinFields0[0].Descriptor()
	// member.DefaultCreateTime holds the default value on creation for the create_time field.
	member.DefaultCreateTime = memberDescCreateTime.Default.(func() time.Time)
	// memberDescUpdateTime is the schema descriptor for update_time field.
	memberDescUpdateTime := memberMixinFields0[1].Descriptor()
	// member.DefaultUpdateTime holds the default value on creation for the update_time field.
	member.DefaultUpdateTime = memberDescUpdateTime.Default.(func() time.Time)
	// member.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	member.UpdateDefaultUpdateTime = memberDescUpdateTime.UpdateDefault.(func() time.Time)
	membersclassMixin := schema.MembersClass{}.Mixin()
	membersclassMixinFields0 := membersclassMixin[0].Fields()
	_ = membersclassMixinFields0
	membersclassFields := schema.MembersClass{}.Fields()
	_ = membersclassFields
	// membersclassDescCreateTime is the schema descriptor for create_time field.
	membersclassDescCreateTime := membersclassMixinFields0[0].Descriptor()
	// membersclass.DefaultCreateTime holds the default value on creation for the create_time field.
	membersclass.DefaultCreateTime = membersclassDescCreateTime.Default.(func() time.Time)
	// membersclassDescUpdateTime is the schema descriptor for update_time field.
	membersclassDescUpdateTime := membersclassMixinFields0[1].Descriptor()
	// membersclass.DefaultUpdateTime holds the default value on creation for the update_time field.
	membersclass.DefaultUpdateTime = membersclassDescUpdateTime.Default.(func() time.Time)
	// membersclass.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	membersclass.UpdateDefaultUpdateTime = membersclassDescUpdateTime.UpdateDefault.(func() time.Time)
	roomMixin := schema.Room{}.Mixin()
	roomMixinFields0 := roomMixin[0].Fields()
	_ = roomMixinFields0
	roomFields := schema.Room{}.Fields()
	_ = roomFields
	// roomDescCreateTime is the schema descriptor for create_time field.
	roomDescCreateTime := roomMixinFields0[0].Descriptor()
	// room.DefaultCreateTime holds the default value on creation for the create_time field.
	room.DefaultCreateTime = roomDescCreateTime.Default.(func() time.Time)
	// roomDescUpdateTime is the schema descriptor for update_time field.
	roomDescUpdateTime := roomMixinFields0[1].Descriptor()
	// room.DefaultUpdateTime holds the default value on creation for the update_time field.
	room.DefaultUpdateTime = roomDescUpdateTime.Default.(func() time.Time)
	// room.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	room.UpdateDefaultUpdateTime = roomDescUpdateTime.UpdateDefault.(func() time.Time)
	scheduleMixin := schema.Schedule{}.Mixin()
	scheduleMixinFields0 := scheduleMixin[0].Fields()
	_ = scheduleMixinFields0
	scheduleFields := schema.Schedule{}.Fields()
	_ = scheduleFields
	// scheduleDescCreateTime is the schema descriptor for create_time field.
	scheduleDescCreateTime := scheduleMixinFields0[0].Descriptor()
	// schedule.DefaultCreateTime holds the default value on creation for the create_time field.
	schedule.DefaultCreateTime = scheduleDescCreateTime.Default.(func() time.Time)
	// scheduleDescUpdateTime is the schema descriptor for update_time field.
	scheduleDescUpdateTime := scheduleMixinFields0[1].Descriptor()
	// schedule.DefaultUpdateTime holds the default value on creation for the update_time field.
	schedule.DefaultUpdateTime = scheduleDescUpdateTime.Default.(func() time.Time)
	// schedule.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	schedule.UpdateDefaultUpdateTime = scheduleDescUpdateTime.UpdateDefault.(func() time.Time)
	schoolMixin := schema.School{}.Mixin()
	schoolMixinFields0 := schoolMixin[0].Fields()
	_ = schoolMixinFields0
	schoolFields := schema.School{}.Fields()
	_ = schoolFields
	// schoolDescCreateTime is the schema descriptor for create_time field.
	schoolDescCreateTime := schoolMixinFields0[0].Descriptor()
	// school.DefaultCreateTime holds the default value on creation for the create_time field.
	school.DefaultCreateTime = schoolDescCreateTime.Default.(func() time.Time)
	// schoolDescUpdateTime is the schema descriptor for update_time field.
	schoolDescUpdateTime := schoolMixinFields0[1].Descriptor()
	// school.DefaultUpdateTime holds the default value on creation for the update_time field.
	school.DefaultUpdateTime = schoolDescUpdateTime.Default.(func() time.Time)
	// school.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	school.UpdateDefaultUpdateTime = schoolDescUpdateTime.UpdateDefault.(func() time.Time)
	studioMixin := schema.Studio{}.Mixin()
	studioMixinFields0 := studioMixin[0].Fields()
	_ = studioMixinFields0
	studioFields := schema.Studio{}.Fields()
	_ = studioFields
	// studioDescCreateTime is the schema descriptor for create_time field.
	studioDescCreateTime := studioMixinFields0[0].Descriptor()
	// studio.DefaultCreateTime holds the default value on creation for the create_time field.
	studio.DefaultCreateTime = studioDescCreateTime.Default.(func() time.Time)
	// studioDescUpdateTime is the schema descriptor for update_time field.
	studioDescUpdateTime := studioMixinFields0[1].Descriptor()
	// studio.DefaultUpdateTime holds the default value on creation for the update_time field.
	studio.DefaultUpdateTime = studioDescUpdateTime.Default.(func() time.Time)
	// studio.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	studio.UpdateDefaultUpdateTime = studioDescUpdateTime.UpdateDefault.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
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
}
