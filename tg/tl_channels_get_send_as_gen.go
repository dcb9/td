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
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// ChannelsGetSendAsRequest represents TL type `channels.getSendAs#dc770ee`.
// Obtains a list of peers that can be used to send messages in a specific group
//
// See https://core.telegram.org/method/channels.getSendAs for reference.
type ChannelsGetSendAsRequest struct {
	// The group where we intend to send messages
	Peer InputPeerClass
}

// ChannelsGetSendAsRequestTypeID is TL type id of ChannelsGetSendAsRequest.
const ChannelsGetSendAsRequestTypeID = 0xdc770ee

// Ensuring interfaces in compile-time for ChannelsGetSendAsRequest.
var (
	_ bin.Encoder     = &ChannelsGetSendAsRequest{}
	_ bin.Decoder     = &ChannelsGetSendAsRequest{}
	_ bin.BareEncoder = &ChannelsGetSendAsRequest{}
	_ bin.BareDecoder = &ChannelsGetSendAsRequest{}
)

func (g *ChannelsGetSendAsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChannelsGetSendAsRequest) String() string {
	if g == nil {
		return "ChannelsGetSendAsRequest(nil)"
	}
	type Alias ChannelsGetSendAsRequest
	return fmt.Sprintf("ChannelsGetSendAsRequest%+v", Alias(*g))
}

// FillFrom fills ChannelsGetSendAsRequest from given interface.
func (g *ChannelsGetSendAsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
}) {
	g.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsGetSendAsRequest) TypeID() uint32 {
	return ChannelsGetSendAsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsGetSendAsRequest) TypeName() string {
	return "channels.getSendAs"
}

// TypeInfo returns info about TL type.
func (g *ChannelsGetSendAsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.getSendAs",
		ID:   ChannelsGetSendAsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *ChannelsGetSendAsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getSendAs#dc770ee as nil")
	}
	b.PutID(ChannelsGetSendAsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ChannelsGetSendAsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getSendAs#dc770ee as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode channels.getSendAs#dc770ee: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getSendAs#dc770ee: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *ChannelsGetSendAsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getSendAs#dc770ee to nil")
	}
	if err := b.ConsumeID(ChannelsGetSendAsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getSendAs#dc770ee: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ChannelsGetSendAsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getSendAs#dc770ee to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.getSendAs#dc770ee: field peer: %w", err)
		}
		g.Peer = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *ChannelsGetSendAsRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// ChannelsGetSendAs invokes method channels.getSendAs#dc770ee returning error if any.
// Obtains a list of peers that can be used to send messages in a specific group
//
// Possible errors:
//
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/channels.getSendAs for reference.
func (c *Client) ChannelsGetSendAs(ctx context.Context, peer InputPeerClass) (*ChannelsSendAsPeers, error) {
	var result ChannelsSendAsPeers

	request := &ChannelsGetSendAsRequest{
		Peer: peer,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
