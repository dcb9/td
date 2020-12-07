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

// InputPhotoEmpty represents TL type `inputPhotoEmpty#1cd7bf0d`.
// Empty constructor.
//
// See https://core.telegram.org/constructor/inputPhotoEmpty for reference.
type InputPhotoEmpty struct {
}

// InputPhotoEmptyTypeID is TL type id of InputPhotoEmpty.
const InputPhotoEmptyTypeID = 0x1cd7bf0d

// Encode implements bin.Encoder.
func (i *InputPhotoEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPhotoEmpty#1cd7bf0d as nil")
	}
	b.PutID(InputPhotoEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPhotoEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPhotoEmpty#1cd7bf0d to nil")
	}
	if err := b.ConsumeID(InputPhotoEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPhotoEmpty#1cd7bf0d: %w", err)
	}

	return nil
}

// construct implements constructor of InputPhotoClass.
func (i InputPhotoEmpty) construct() InputPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputPhotoEmpty.
var (
	_ bin.Encoder = &InputPhotoEmpty{}
	_ bin.Decoder = &InputPhotoEmpty{}

	_ InputPhotoClass = &InputPhotoEmpty{}
)

// InputPhoto represents TL type `inputPhoto#3bb3b94a`.
// Defines a photo for further interaction.
//
// See https://core.telegram.org/constructor/inputPhoto for reference.
type InputPhoto struct {
	// Photo identifier
	ID int64
	// access_hash value from the photo constructor
	AccessHash int64
	// File reference
	FileReference []byte
}

// InputPhotoTypeID is TL type id of InputPhoto.
const InputPhotoTypeID = 0x3bb3b94a

// Encode implements bin.Encoder.
func (i *InputPhoto) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPhoto#3bb3b94a as nil")
	}
	b.PutID(InputPhotoTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	b.PutBytes(i.FileReference)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPhoto) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPhoto#3bb3b94a to nil")
	}
	if err := b.ConsumeID(InputPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPhoto#3bb3b94a: %w", err)
	}

	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoto#3bb3b94a: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoto#3bb3b94a: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoto#3bb3b94a: field file_reference: %w", err)
		}
		i.FileReference = value
	}
	return nil
}

// construct implements constructor of InputPhotoClass.
func (i InputPhoto) construct() InputPhotoClass { return &i }

// Ensuring interfaces in compile-time for InputPhoto.
var (
	_ bin.Encoder = &InputPhoto{}
	_ bin.Decoder = &InputPhoto{}

	_ InputPhotoClass = &InputPhoto{}
)

// InputPhotoClass represents InputPhoto generic type.
//
// See https://core.telegram.org/type/InputPhoto for reference.
//
// Example:
//  g, err := DecodeInputPhoto(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputPhotoEmpty: // inputPhotoEmpty#1cd7bf0d
//  case *InputPhoto: // inputPhoto#3bb3b94a
//  default: panic(v)
//  }
type InputPhotoClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputPhotoClass
}

// DecodeInputPhoto implements binary de-serialization for InputPhotoClass.
func DecodeInputPhoto(buf *bin.Buffer) (InputPhotoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputPhotoEmptyTypeID:
		// Decoding inputPhotoEmpty#1cd7bf0d.
		v := InputPhotoEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPhotoClass: %w", err)
		}
		return &v, nil
	case InputPhotoTypeID:
		// Decoding inputPhoto#3bb3b94a.
		v := InputPhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPhotoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputPhotoClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputPhoto boxes the InputPhotoClass providing a helper.
type InputPhotoBox struct {
	InputPhoto InputPhotoClass
}

// Decode implements bin.Decoder for InputPhotoBox.
func (b *InputPhotoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputPhotoBox to nil")
	}
	v, err := DecodeInputPhoto(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputPhoto = v
	return nil
}

// Encode implements bin.Encode for InputPhotoBox.
func (b *InputPhotoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputPhoto == nil {
		return fmt.Errorf("unable to encode InputPhotoClass as nil")
	}
	return b.InputPhoto.Encode(buf)
}
