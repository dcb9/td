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

// ChannelsGetLeftChannelsRequest represents TL type `channels.getLeftChannels#8341ecc0`.
// Get a list of channels/supergroups we left
//
// See https://core.telegram.org/method/channels.getLeftChannels for reference.
type ChannelsGetLeftChannelsRequest struct {
	// Offset for pagination
	Offset int
}

// ChannelsGetLeftChannelsRequestTypeID is TL type id of ChannelsGetLeftChannelsRequest.
const ChannelsGetLeftChannelsRequestTypeID = 0x8341ecc0

// Encode implements bin.Encoder.
func (g *ChannelsGetLeftChannelsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getLeftChannels#8341ecc0 as nil")
	}
	b.PutID(ChannelsGetLeftChannelsRequestTypeID)
	b.PutInt(g.Offset)
	return nil
}

// Decode implements bin.Decoder.
func (g *ChannelsGetLeftChannelsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getLeftChannels#8341ecc0 to nil")
	}
	if err := b.ConsumeID(ChannelsGetLeftChannelsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getLeftChannels#8341ecc0: %w", err)
	}

	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getLeftChannels#8341ecc0: field offset: %w", err)
		}
		g.Offset = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsGetLeftChannelsRequest.
var (
	_ bin.Encoder = &ChannelsGetLeftChannelsRequest{}
	_ bin.Decoder = &ChannelsGetLeftChannelsRequest{}
)

// ChannelsGetLeftChannels invokes method channels.getLeftChannels#8341ecc0 returning error if any.
// Get a list of channels/supergroups we left
//
// See https://core.telegram.org/method/channels.getLeftChannels for reference.
func (c *Client) ChannelsGetLeftChannels(ctx context.Context, request *ChannelsGetLeftChannelsRequest) (MessagesChatsClass, error) {
	var result MessagesChatsBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Chats, nil
}
