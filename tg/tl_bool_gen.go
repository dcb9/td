// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// BoolFalse represents TL type `boolFalse#bc799737`.
// Constructor may be interpreted as a booleanfalse value.
//
// See https://core.telegram.org/constructor/boolFalse for reference.
type BoolFalse struct {
}

// BoolFalseTypeID is TL type id of BoolFalse.
const BoolFalseTypeID = 0xbc799737

// Encode implements bin.Encoder.
func (b *BoolFalse) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode boolFalse#bc799737 as nil")
	}
	buf.PutID(BoolFalseTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (b *BoolFalse) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode boolFalse#bc799737 to nil")
	}
	if err := buf.ConsumeID(BoolFalseTypeID); err != nil {
		return fmt.Errorf("unable to decode boolFalse#bc799737: %w", err)
	}

	return nil
}

// construct implements constructor of BoolClass.
func (b BoolFalse) construct() BoolClass { return &b }

// Ensuring interfaces in compile-time for BoolFalse.
var (
	_ bin.Encoder = &BoolFalse{}
	_ bin.Decoder = &BoolFalse{}

	_ BoolClass = &BoolFalse{}
)

// BoolTrue represents TL type `boolTrue#997275b5`.
// The constructor can be interpreted as a booleantrue value.
//
// See https://core.telegram.org/constructor/boolTrue for reference.
type BoolTrue struct {
}

// BoolTrueTypeID is TL type id of BoolTrue.
const BoolTrueTypeID = 0x997275b5

// Encode implements bin.Encoder.
func (b *BoolTrue) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode boolTrue#997275b5 as nil")
	}
	buf.PutID(BoolTrueTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (b *BoolTrue) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode boolTrue#997275b5 to nil")
	}
	if err := buf.ConsumeID(BoolTrueTypeID); err != nil {
		return fmt.Errorf("unable to decode boolTrue#997275b5: %w", err)
	}

	return nil
}

// construct implements constructor of BoolClass.
func (b BoolTrue) construct() BoolClass { return &b }

// Ensuring interfaces in compile-time for BoolTrue.
var (
	_ bin.Encoder = &BoolTrue{}
	_ bin.Decoder = &BoolTrue{}

	_ BoolClass = &BoolTrue{}
)

// BoolClass represents Bool generic type.
//
// See https://core.telegram.org/type/Bool for reference.
//
// Example:
//  g, err := DecodeBool(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *BoolFalse: // boolFalse#bc799737
//  case *BoolTrue: // boolTrue#997275b5
//  default: panic(v)
//  }
type BoolClass interface {
	bin.Encoder
	bin.Decoder
	construct() BoolClass
}

// DecodeBool implements binary de-serialization for BoolClass.
func DecodeBool(buf *bin.Buffer) (BoolClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case BoolFalseTypeID:
		// Decoding boolFalse#bc799737.
		v := BoolFalse{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BoolClass: %w", err)
		}
		return &v, nil
	case BoolTrueTypeID:
		// Decoding boolTrue#997275b5.
		v := BoolTrue{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BoolClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BoolClass: %w", bin.NewUnexpectedID(id))
	}
}

// Bool boxes the BoolClass providing a helper.
type BoolBox struct {
	Bool BoolClass
}

// Decode implements bin.Decoder for BoolBox.
func (b *BoolBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode BoolBox to nil")
	}
	v, err := DecodeBool(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Bool = v
	return nil
}

// Encode implements bin.Encode for BoolBox.
func (b *BoolBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Bool == nil {
		return fmt.Errorf("unable to encode BoolClass as nil")
	}
	return b.Bool.Encode(buf)
}
