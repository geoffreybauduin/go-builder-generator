// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

package builders

import (
	"database/sql"
)

// NamedArgBuilder is an alias of NamedArg to build NamedArg with builder-pattern.
type NamedArgBuilder sql.NamedArg

// NewNamedArgBuilder creates a new NamedArgBuilder.
func NewNamedArgBuilder() *NamedArgBuilder {
	return &NamedArgBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *NamedArgBuilder) Copy() *NamedArgBuilder {
	result := *b
	return &result
}

// Build returns built NamedArg.
func (b *NamedArgBuilder) Build() *sql.NamedArg {

	result := (sql.NamedArg)(*b)
	return &result
}

// SetName sets NamedArg's Name.
func (b *NamedArgBuilder) SetName(name string) *NamedArgBuilder {
	b.Name = name
	return b
}

// SetValue sets NamedArg's Value.
func (b *NamedArgBuilder) SetValue(value any) *NamedArgBuilder {
	b.Value = value
	return b
}
