// Code generated by go-builder-generator (https://github.com/kilianpaquier/go-builder-generator). DO NOT EDIT.

package builders

import (
	"context"

	"github.com/kilianpaquier/go-builder-generator/testdata/success_channels"
)

// ChanBuilder is an alias of Chan to build Chan with builder-pattern.
type ChanBuilder success_channels.Chan

// NewChanBuilder creates a new ChanBuilder.
func NewChanBuilder() *ChanBuilder {
	return &ChanBuilder{}
}

// Copy reassigns the builder struct (behind pointer) to a new pointer and returns it.
func (b *ChanBuilder) Copy() *ChanBuilder {
	result := *b
	return &result
}

// Build returns built Chan.
func (b *ChanBuilder) Build() *success_channels.Chan {

	result := (success_channels.Chan)(*b)
	return &result
}

// SetChanField sets Chan's ChanField.
func (b *ChanBuilder) SetChanField(chanField chan int64) *ChanBuilder {
	b.ChanField = chanField
	return b
}

// SetChanFieldAlias sets Chan's ChanFieldAlias.
func (b *ChanBuilder) SetChanFieldAlias(chanFieldAlias chan success_channels.Int64Alias) *ChanBuilder {
	b.ChanFieldAlias = chanFieldAlias
	return b
}

// SetChanFieldPtrAlias sets Chan's ChanFieldPtrAlias.
func (b *ChanBuilder) SetChanFieldPtrAlias(chanFieldPtrAlias chan *success_channels.Int64Alias) *ChanBuilder {
	b.ChanFieldPtrAlias = &chanFieldPtrAlias
	return b
}

// SetChanFieldSliceAlias sets Chan's ChanFieldSliceAlias.
func (b *ChanBuilder) SetChanFieldSliceAlias(chanFieldSliceAlias chan []success_channels.FuncAlias) *ChanBuilder {
	b.ChanFieldSliceAlias = chanFieldSliceAlias
	return b
}

// SetChanFieldSliceFunc sets Chan's ChanFieldSliceFunc.
func (b *ChanBuilder) SetChanFieldSliceFunc(chanFieldSliceFunc chan []func(success_channels.Int64Alias) (err error)) *ChanBuilder {
	b.ChanFieldSliceFunc = chanFieldSliceFunc
	return b
}

// SetChanFieldWithPkg sets Chan's ChanFieldWithPkg.
func (b *ChanBuilder) SetChanFieldWithPkg(chanFieldWithPkg chan context.Context) *ChanBuilder {
	b.ChanFieldWithPkg = chanFieldWithPkg
	return b
}

// SetRChanField sets Chan's RChanField.
func (b *ChanBuilder) SetRChanField(rChanField <-chan int64) *ChanBuilder {
	b.RChanField = rChanField
	return b
}

// SetRChanFieldAlias sets Chan's RChanFieldAlias.
func (b *ChanBuilder) SetRChanFieldAlias(rChanFieldAlias <-chan success_channels.Int64Alias) *ChanBuilder {
	b.RChanFieldAlias = rChanFieldAlias
	return b
}

// SetRChanFieldPtrAlias sets Chan's RChanFieldPtrAlias.
func (b *ChanBuilder) SetRChanFieldPtrAlias(rChanFieldPtrAlias <-chan *success_channels.Int64Alias) *ChanBuilder {
	b.RChanFieldPtrAlias = &rChanFieldPtrAlias
	return b
}

// SetRChanFieldWithPkg sets Chan's RChanFieldWithPkg.
func (b *ChanBuilder) SetRChanFieldWithPkg(rChanFieldWithPkg <-chan context.Context) *ChanBuilder {
	b.RChanFieldWithPkg = rChanFieldWithPkg
	return b
}

// SetSChanField sets Chan's SChanField.
func (b *ChanBuilder) SetSChanField(sChanField chan<- int64) *ChanBuilder {
	b.SChanField = sChanField
	return b
}

// SetSChanFieldAlias sets Chan's SChanFieldAlias.
func (b *ChanBuilder) SetSChanFieldAlias(sChanFieldAlias chan<- success_channels.Int64Alias) *ChanBuilder {
	b.SChanFieldAlias = sChanFieldAlias
	return b
}

// SetSChanFieldPtrAlias sets Chan's SChanFieldPtrAlias.
func (b *ChanBuilder) SetSChanFieldPtrAlias(sChanFieldPtrAlias chan<- *success_channels.Int64Alias) *ChanBuilder {
	b.SChanFieldPtrAlias = &sChanFieldPtrAlias
	return b
}

// SetSChanFieldWithPkg sets Chan's SChanFieldWithPkg.
func (b *ChanBuilder) SetSChanFieldWithPkg(sChanFieldWithPkg chan<- context.Context) *ChanBuilder {
	b.SChanFieldWithPkg = sChanFieldWithPkg
	return b
}
