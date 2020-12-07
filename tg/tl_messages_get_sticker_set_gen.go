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

// MessagesGetStickerSetRequest represents TL type `messages.getStickerSet#2619a90e`.
// Get info about a stickerset
//
// See https://core.telegram.org/method/messages.getStickerSet for reference.
type MessagesGetStickerSetRequest struct {
	// Stickerset
	Stickerset InputStickerSetClass
}

// MessagesGetStickerSetRequestTypeID is TL type id of MessagesGetStickerSetRequest.
const MessagesGetStickerSetRequestTypeID = 0x2619a90e

// Encode implements bin.Encoder.
func (g *MessagesGetStickerSetRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getStickerSet#2619a90e as nil")
	}
	b.PutID(MessagesGetStickerSetRequestTypeID)
	if g.Stickerset == nil {
		return fmt.Errorf("unable to encode messages.getStickerSet#2619a90e: field stickerset is nil")
	}
	if err := g.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getStickerSet#2619a90e: field stickerset: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetStickerSetRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getStickerSet#2619a90e to nil")
	}
	if err := b.ConsumeID(MessagesGetStickerSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getStickerSet#2619a90e: %w", err)
	}

	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getStickerSet#2619a90e: field stickerset: %w", err)
		}
		g.Stickerset = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetStickerSetRequest.
var (
	_ bin.Encoder = &MessagesGetStickerSetRequest{}
	_ bin.Decoder = &MessagesGetStickerSetRequest{}
)

// MessagesGetStickerSet invokes method messages.getStickerSet#2619a90e returning error if any.
// Get info about a stickerset
//
// See https://core.telegram.org/method/messages.getStickerSet for reference.
func (c *Client) MessagesGetStickerSet(ctx context.Context, request *MessagesGetStickerSetRequest) (*MessagesStickerSet, error) {
	var result MessagesStickerSet
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
