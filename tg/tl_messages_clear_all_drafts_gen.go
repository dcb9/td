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

// MessagesClearAllDraftsRequest represents TL type `messages.clearAllDrafts#7e58ee9c`.
// Clear all drafts.
//
// See https://core.telegram.org/method/messages.clearAllDrafts for reference.
type MessagesClearAllDraftsRequest struct {
}

// MessagesClearAllDraftsRequestTypeID is TL type id of MessagesClearAllDraftsRequest.
const MessagesClearAllDraftsRequestTypeID = 0x7e58ee9c

// Encode implements bin.Encoder.
func (c *MessagesClearAllDraftsRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.clearAllDrafts#7e58ee9c as nil")
	}
	b.PutID(MessagesClearAllDraftsRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (c *MessagesClearAllDraftsRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.clearAllDrafts#7e58ee9c to nil")
	}
	if err := b.ConsumeID(MessagesClearAllDraftsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.clearAllDrafts#7e58ee9c: %w", err)
	}

	return nil
}

// Ensuring interfaces in compile-time for MessagesClearAllDraftsRequest.
var (
	_ bin.Encoder = &MessagesClearAllDraftsRequest{}
	_ bin.Decoder = &MessagesClearAllDraftsRequest{}
)

// MessagesClearAllDrafts invokes method messages.clearAllDrafts#7e58ee9c returning error if any.
// Clear all drafts.
//
// See https://core.telegram.org/method/messages.clearAllDrafts for reference.
func (c *Client) MessagesClearAllDrafts(ctx context.Context, request *MessagesClearAllDraftsRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
