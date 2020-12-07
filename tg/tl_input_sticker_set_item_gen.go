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

// InputStickerSetItem represents TL type `inputStickerSetItem#ffa0a496`.
// Sticker in a stickerset
//
// See https://core.telegram.org/constructor/inputStickerSetItem for reference.
type InputStickerSetItem struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// The sticker
	Document InputDocumentClass
	// Associated emoji
	Emoji string
	// Coordinates for mask sticker
	//
	// Use SetMaskCoords and GetMaskCoords helpers.
	MaskCoords MaskCoords
}

// InputStickerSetItemTypeID is TL type id of InputStickerSetItem.
const InputStickerSetItemTypeID = 0xffa0a496

// Encode implements bin.Encoder.
func (i *InputStickerSetItem) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerSetItem#ffa0a496 as nil")
	}
	b.PutID(InputStickerSetItemTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStickerSetItem#ffa0a496: field flags: %w", err)
	}
	if i.Document == nil {
		return fmt.Errorf("unable to encode inputStickerSetItem#ffa0a496: field document is nil")
	}
	if err := i.Document.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStickerSetItem#ffa0a496: field document: %w", err)
	}
	b.PutString(i.Emoji)
	if i.Flags.Has(0) {
		if err := i.MaskCoords.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputStickerSetItem#ffa0a496: field mask_coords: %w", err)
		}
	}
	return nil
}

// SetMaskCoords sets value of MaskCoords conditional field.
func (i *InputStickerSetItem) SetMaskCoords(value MaskCoords) {
	i.Flags.Set(0)
	i.MaskCoords = value
}

// GetMaskCoords returns value of MaskCoords conditional field and
// boolean which is true if field was set.
func (i *InputStickerSetItem) GetMaskCoords() (value MaskCoords, ok bool) {
	if !i.Flags.Has(0) {
		return value, false
	}
	return i.MaskCoords, true
}

// Decode implements bin.Decoder.
func (i *InputStickerSetItem) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerSetItem#ffa0a496 to nil")
	}
	if err := b.ConsumeID(InputStickerSetItemTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickerSetItem#ffa0a496: %w", err)
	}

	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputStickerSetItem#ffa0a496: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerSetItem#ffa0a496: field document: %w", err)
		}
		i.Document = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerSetItem#ffa0a496: field emoji: %w", err)
		}
		i.Emoji = value
	}
	if i.Flags.Has(0) {
		if err := i.MaskCoords.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputStickerSetItem#ffa0a496: field mask_coords: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for InputStickerSetItem.
var (
	_ bin.Encoder = &InputStickerSetItem{}
	_ bin.Decoder = &InputStickerSetItem{}
)
