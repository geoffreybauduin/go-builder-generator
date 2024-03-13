// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

package builders

import (
	"github.com/kilianpaquier/go-builder-generator/testdata/success_funcs"
)

// FuncBuilder is an alias of Func to build Func with builder-pattern.
type FuncBuilder success_funcs.Func

// NewFuncBuilder creates a new FuncBuilder.
func NewFuncBuilder() *FuncBuilder {
	return &FuncBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *FuncBuilder) Copy() *FuncBuilder {
	result := *b
	return &result
}

// Build returns built Func.
func (b *FuncBuilder) Build() *success_funcs.Func {

	result := (success_funcs.Func)(*b)
	return &result
}

// SetFuncField sets Func's FuncField.
func (b *FuncBuilder) SetFuncField(funcField func(int64) string) *FuncBuilder {
	b.FuncField = funcField
	return b
}

// SetFuncFieldAlias sets Func's FuncFieldAlias.
func (b *FuncBuilder) SetFuncFieldAlias(funcFieldAlias func(success_funcs.Int64Alias) string) *FuncBuilder {
	b.FuncFieldAlias = funcFieldAlias
	return b
}

// SetFuncFieldAliasMultiple sets Func's FuncFieldAliasMultiple.
func (b *FuncBuilder) SetFuncFieldAliasMultiple(funcFieldAliasMultiple func(success_funcs.Int64Alias, success_funcs.FuncAlias) (string, error)) *FuncBuilder {
	b.FuncFieldAliasMultiple = funcFieldAliasMultiple
	return b
}

// SetFuncFieldAliasNamed sets Func's FuncFieldAliasNamed.
func (b *FuncBuilder) SetFuncFieldAliasNamed(funcFieldAliasNamed func(in success_funcs.Int64Alias) (out success_funcs.FuncAlias)) *FuncBuilder {
	b.FuncFieldAliasNamed = funcFieldAliasNamed
	return b
}

// SetFuncFieldChan sets Func's FuncFieldChan.
func (b *FuncBuilder) SetFuncFieldChan(funcFieldChan func(c chan<- success_funcs.Int64Alias) error) *FuncBuilder {
	b.FuncFieldChan = funcFieldChan
	return b
}

// SetFuncFieldCResult sets Func's FuncFieldCResult.
func (b *FuncBuilder) SetFuncFieldCResult(funcFieldCResult func(int64, string) (func(), error)) *FuncBuilder {
	b.FuncFieldCResult = funcFieldCResult
	return b
}

// SetFuncFieldMultiple sets Func's FuncFieldMultiple.
func (b *FuncBuilder) SetFuncFieldMultiple(funcFieldMultiple func(int64, string) (string, error)) *FuncBuilder {
	b.FuncFieldMultiple = funcFieldMultiple
	return b
}

// SetFuncFieldNamed sets Func's FuncFieldNamed.
func (b *FuncBuilder) SetFuncFieldNamed(funcFieldNamed func(in int64) (out string)) *FuncBuilder {
	b.FuncFieldNamed = funcFieldNamed
	return b
}

// SetFuncFieldPtrAlias sets Func's FuncFieldPtrAlias.
func (b *FuncBuilder) SetFuncFieldPtrAlias(funcFieldPtrAlias func(in *success_funcs.Int64Alias) (out *success_funcs.FuncAlias, err error)) *FuncBuilder {
	b.FuncFieldPtrAlias = &funcFieldPtrAlias
	return b
}
