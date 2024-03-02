// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

package builders

import (
	"github.com/kilianpaquier/go-builder-generator/testdata/success_interface"
)

// InterfaceNoFieldsBuilder is an alias of InterfaceNoFields to build InterfaceNoFields with builder-pattern.
type InterfaceNoFieldsBuilder success_interface.InterfaceNoFields

// NewInterfaceNoFieldsBuilder creates a new InterfaceNoFieldsBuilder.
func NewInterfaceNoFieldsBuilder() *InterfaceNoFieldsBuilder {
	return &InterfaceNoFieldsBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *InterfaceNoFieldsBuilder) Copy() *InterfaceNoFieldsBuilder {
	c := *b
	return &c
}

// Build returns built InterfaceNoFields.
func (b *InterfaceNoFieldsBuilder) Build() *success_interface.InterfaceNoFields {
	c := (success_interface.InterfaceNoFields)(*b)
	return &c
}

// SetNoFields sets InterfaceNoFields's NoFields.
func (b *InterfaceNoFieldsBuilder) SetNoFields(noFields interface{}) *InterfaceNoFieldsBuilder {
	b.NoFields = noFields
	return b
}