// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/kindlyfire/longa/ent/link"
	"github.com/kindlyfire/longa/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	linkFields := schema.Link{}.Fields()
	_ = linkFields
	// linkDescCreatedAt is the schema descriptor for createdAt field.
	linkDescCreatedAt := linkFields[2].Descriptor()
	// link.DefaultCreatedAt holds the default value on creation for the createdAt field.
	link.DefaultCreatedAt = linkDescCreatedAt.Default.(func() time.Time)
}
