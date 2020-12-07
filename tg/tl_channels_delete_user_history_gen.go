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

// ChannelsDeleteUserHistoryRequest represents TL type `channels.deleteUserHistory#d10dd71b`.
// Delete all messages sent by a certain user in a supergroup
//
// See https://core.telegram.org/method/channels.deleteUserHistory for reference.
type ChannelsDeleteUserHistoryRequest struct {
	// Supergroup
	Channel InputChannelClass
	// User whose messages should be deleted
	UserID InputUserClass
}

// ChannelsDeleteUserHistoryRequestTypeID is TL type id of ChannelsDeleteUserHistoryRequest.
const ChannelsDeleteUserHistoryRequestTypeID = 0xd10dd71b

// Encode implements bin.Encoder.
func (d *ChannelsDeleteUserHistoryRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteUserHistory#d10dd71b as nil")
	}
	b.PutID(ChannelsDeleteUserHistoryRequestTypeID)
	if d.Channel == nil {
		return fmt.Errorf("unable to encode channels.deleteUserHistory#d10dd71b: field channel is nil")
	}
	if err := d.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteUserHistory#d10dd71b: field channel: %w", err)
	}
	if d.UserID == nil {
		return fmt.Errorf("unable to encode channels.deleteUserHistory#d10dd71b: field user_id is nil")
	}
	if err := d.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteUserHistory#d10dd71b: field user_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *ChannelsDeleteUserHistoryRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteUserHistory#d10dd71b to nil")
	}
	if err := b.ConsumeID(ChannelsDeleteUserHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.deleteUserHistory#d10dd71b: %w", err)
	}

	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteUserHistory#d10dd71b: field channel: %w", err)
		}
		d.Channel = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteUserHistory#d10dd71b: field user_id: %w", err)
		}
		d.UserID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsDeleteUserHistoryRequest.
var (
	_ bin.Encoder = &ChannelsDeleteUserHistoryRequest{}
	_ bin.Decoder = &ChannelsDeleteUserHistoryRequest{}
)

// ChannelsDeleteUserHistory invokes method channels.deleteUserHistory#d10dd71b returning error if any.
// Delete all messages sent by a certain user in a supergroup
//
// See https://core.telegram.org/method/channels.deleteUserHistory for reference.
func (c *Client) ChannelsDeleteUserHistory(ctx context.Context, request *ChannelsDeleteUserHistoryRequest) (*MessagesAffectedHistory, error) {
	var result MessagesAffectedHistory
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
