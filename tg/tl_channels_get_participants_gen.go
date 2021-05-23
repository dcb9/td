// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// ChannelsGetParticipantsRequest represents TL type `channels.getParticipants#123e05e9`.
// Get the participants of a supergroup/channel¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.getParticipants for reference.
type ChannelsGetParticipantsRequest struct {
	// Channel
	Channel InputChannelClass
	// Which participant types to fetch
	Filter ChannelParticipantsFilterClass
	// Offset¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Offset int
	// Limit¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
	// Hash¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Hash int
}

// ChannelsGetParticipantsRequestTypeID is TL type id of ChannelsGetParticipantsRequest.
const ChannelsGetParticipantsRequestTypeID = 0x123e05e9

func (g *ChannelsGetParticipantsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Channel == nil) {
		return false
	}
	if !(g.Filter == nil) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChannelsGetParticipantsRequest) String() string {
	if g == nil {
		return "ChannelsGetParticipantsRequest(nil)"
	}
	type Alias ChannelsGetParticipantsRequest
	return fmt.Sprintf("ChannelsGetParticipantsRequest%+v", Alias(*g))
}

// FillFrom fills ChannelsGetParticipantsRequest from given interface.
func (g *ChannelsGetParticipantsRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetFilter() (value ChannelParticipantsFilterClass)
	GetOffset() (value int)
	GetLimit() (value int)
	GetHash() (value int)
}) {
	g.Channel = from.GetChannel()
	g.Filter = from.GetFilter()
	g.Offset = from.GetOffset()
	g.Limit = from.GetLimit()
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsGetParticipantsRequest) TypeID() uint32 {
	return ChannelsGetParticipantsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsGetParticipantsRequest) TypeName() string {
	return "channels.getParticipants"
}

// TypeInfo returns info about TL type.
func (g *ChannelsGetParticipantsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.getParticipants",
		ID:   ChannelsGetParticipantsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *ChannelsGetParticipantsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getParticipants#123e05e9 as nil")
	}
	b.PutID(ChannelsGetParticipantsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ChannelsGetParticipantsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getParticipants#123e05e9 as nil")
	}
	if g.Channel == nil {
		return fmt.Errorf("unable to encode channels.getParticipants#123e05e9: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getParticipants#123e05e9: field channel: %w", err)
	}
	if g.Filter == nil {
		return fmt.Errorf("unable to encode channels.getParticipants#123e05e9: field filter is nil")
	}
	if err := g.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getParticipants#123e05e9: field filter: %w", err)
	}
	b.PutInt(g.Offset)
	b.PutInt(g.Limit)
	b.PutInt(g.Hash)
	return nil
}

// GetChannel returns value of Channel field.
func (g *ChannelsGetParticipantsRequest) GetChannel() (value InputChannelClass) {
	return g.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (g *ChannelsGetParticipantsRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return g.Channel.AsNotEmpty()
}

// GetFilter returns value of Filter field.
func (g *ChannelsGetParticipantsRequest) GetFilter() (value ChannelParticipantsFilterClass) {
	return g.Filter
}

// GetOffset returns value of Offset field.
func (g *ChannelsGetParticipantsRequest) GetOffset() (value int) {
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *ChannelsGetParticipantsRequest) GetLimit() (value int) {
	return g.Limit
}

// GetHash returns value of Hash field.
func (g *ChannelsGetParticipantsRequest) GetHash() (value int) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *ChannelsGetParticipantsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getParticipants#123e05e9 to nil")
	}
	if err := b.ConsumeID(ChannelsGetParticipantsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getParticipants#123e05e9: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ChannelsGetParticipantsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getParticipants#123e05e9 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.getParticipants#123e05e9: field channel: %w", err)
		}
		g.Channel = value
	}
	{
		value, err := DecodeChannelParticipantsFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.getParticipants#123e05e9: field filter: %w", err)
		}
		g.Filter = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getParticipants#123e05e9: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getParticipants#123e05e9: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getParticipants#123e05e9: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsGetParticipantsRequest.
var (
	_ bin.Encoder     = &ChannelsGetParticipantsRequest{}
	_ bin.Decoder     = &ChannelsGetParticipantsRequest{}
	_ bin.BareEncoder = &ChannelsGetParticipantsRequest{}
	_ bin.BareDecoder = &ChannelsGetParticipantsRequest{}
)

// ChannelsGetParticipants invokes method channels.getParticipants#123e05e9 returning error if any.
// Get the participants of a supergroup/channel¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 INPUT_CONSTRUCTOR_INVALID: The provided constructor is invalid
//
// See https://core.telegram.org/method/channels.getParticipants for reference.
// Can be used by bots.
func (c *Client) ChannelsGetParticipants(ctx context.Context, request *ChannelsGetParticipantsRequest) (ChannelsChannelParticipantsClass, error) {
	var result ChannelsChannelParticipantsBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ChannelParticipants, nil
}
