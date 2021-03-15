package proto

import (
	"fmt"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
)

// FutureSalt represents TL type `future_salt#949d9dc`.
type FutureSalt struct {
	// ValidSince field of FutureSalt.
	ValidSince int
	// ValidUntil field of FutureSalt.
	ValidUntil int
	// Salt field of FutureSalt.
	Salt int64
}

// FutureSaltTypeID is TL type id of FutureSalt.
const FutureSaltTypeID = 0x949d9dc

// String implements fmt.Stringer.
func (f *FutureSalt) String() string {
	if f == nil {
		return "FutureSalt(nil)"
	}
	type Alias FutureSalt
	return fmt.Sprintf("FutureSalt%+v", Alias(*f))
}

// FillFrom fills FutureSalt from given interface.
func (f *FutureSalt) FillFrom(from interface {
	GetValidSince() (value int)
	GetValidUntil() (value int)
	GetSalt() (value int64)
}) {
	f.ValidSince = from.GetValidSince()
	f.ValidUntil = from.GetValidUntil()
	f.Salt = from.GetSalt()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FutureSalt) TypeID() uint32 {
	return FutureSaltTypeID
}

// TypeName returns name of type in TL schema.
func (*FutureSalt) TypeName() string {
	return "future_salt"
}

// TypeInfo returns info about TL type.
func (f *FutureSalt) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "future_salt",
		ID:   FutureSaltTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ValidSince",
			SchemaName: "valid_since",
		},
		{
			Name:       "ValidUntil",
			SchemaName: "valid_until",
		},
		{
			Name:       "Salt",
			SchemaName: "salt",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *FutureSalt) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode future_salt#949d9dc as nil")
	}
	b.PutID(FutureSaltTypeID)
	b.PutInt(f.ValidSince)
	b.PutInt(f.ValidUntil)
	b.PutLong(f.Salt)
	return nil
}

// GetValidSince returns value of ValidSince field.
func (f *FutureSalt) GetValidSince() (value int) {
	return f.ValidSince
}

// GetValidUntil returns value of ValidUntil field.
func (f *FutureSalt) GetValidUntil() (value int) {
	return f.ValidUntil
}

// GetSalt returns value of Salt field.
func (f *FutureSalt) GetSalt() (value int64) {
	return f.Salt
}

// Decode implements bin.Decoder.
func (f *FutureSalt) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode future_salt#949d9dc to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode future_salt#949d9dc: field valid_since: %w", err)
		}
		f.ValidSince = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode future_salt#949d9dc: field valid_until: %w", err)
		}
		f.ValidUntil = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode future_salt#949d9dc: field salt: %w", err)
		}
		f.Salt = value
	}
	return nil
}

// Ensuring interfaces in compile-time for FutureSalt.
var (
	_ bin.Encoder = &FutureSalt{}
	_ bin.Decoder = &FutureSalt{}
)