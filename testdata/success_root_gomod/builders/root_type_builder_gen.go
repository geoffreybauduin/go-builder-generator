// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

package builders

import (
	"github.com/kilianpaquier/go-builder-generator/root"
)

// RootTypeBuilder is an alias of RootType to build RootType with builder-pattern.
type RootTypeBuilder root.RootType

// NewRootTypeBuilder creates a new RootTypeBuilder.
func NewRootTypeBuilder() *RootTypeBuilder {
	return &RootTypeBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *RootTypeBuilder) Copy() *RootTypeBuilder {
	result := *b
	return &result
}

// Build returns built RootType.
func (b *RootTypeBuilder) Build() *root.RootType {

	result := (root.RootType)(*b)
	return &result
}

// SetField sets RootType's Field.
func (b *RootTypeBuilder) SetField(field int64) *RootTypeBuilder {
	b.Field = field
	return b
}
