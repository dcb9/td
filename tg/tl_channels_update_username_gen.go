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

// ChannelsUpdateUsernameRequest represents TL type `channels.updateUsername#3514b3de`.
// Change the username of a supergroup/channel
//
// See https://core.telegram.org/method/channels.updateUsername for reference.
type ChannelsUpdateUsernameRequest struct {
	// Channel
	Channel InputChannelClass
	// New username
	Username string
}

// ChannelsUpdateUsernameRequestTypeID is TL type id of ChannelsUpdateUsernameRequest.
const ChannelsUpdateUsernameRequestTypeID = 0x3514b3de

// Encode implements bin.Encoder.
func (u *ChannelsUpdateUsernameRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode channels.updateUsername#3514b3de as nil")
	}
	b.PutID(ChannelsUpdateUsernameRequestTypeID)
	if u.Channel == nil {
		return fmt.Errorf("unable to encode channels.updateUsername#3514b3de: field channel is nil")
	}
	if err := u.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.updateUsername#3514b3de: field channel: %w", err)
	}
	b.PutString(u.Username)
	return nil
}

// Decode implements bin.Decoder.
func (u *ChannelsUpdateUsernameRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode channels.updateUsername#3514b3de to nil")
	}
	if err := b.ConsumeID(ChannelsUpdateUsernameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.updateUsername#3514b3de: %w", err)
	}

	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.updateUsername#3514b3de: field channel: %w", err)
		}
		u.Channel = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.updateUsername#3514b3de: field username: %w", err)
		}
		u.Username = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsUpdateUsernameRequest.
var (
	_ bin.Encoder = &ChannelsUpdateUsernameRequest{}
	_ bin.Decoder = &ChannelsUpdateUsernameRequest{}
)

// ChannelsUpdateUsername invokes method channels.updateUsername#3514b3de returning error if any.
// Change the username of a supergroup/channel
//
// See https://core.telegram.org/method/channels.updateUsername for reference.
func (c *Client) ChannelsUpdateUsername(ctx context.Context, request *ChannelsUpdateUsernameRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
