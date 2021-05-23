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

// MessagesGetScheduledHistoryRequest represents TL type `messages.getScheduledHistory#e2c2685b`.
// Get scheduled messages
//
// See https://core.telegram.org/method/messages.getScheduledHistory for reference.
type MessagesGetScheduledHistoryRequest struct {
	// Peer
	Peer InputPeerClass
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// MessagesGetScheduledHistoryRequestTypeID is TL type id of MessagesGetScheduledHistoryRequest.
const MessagesGetScheduledHistoryRequestTypeID = 0xe2c2685b

func (g *MessagesGetScheduledHistoryRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetScheduledHistoryRequest) String() string {
	if g == nil {
		return "MessagesGetScheduledHistoryRequest(nil)"
	}
	type Alias MessagesGetScheduledHistoryRequest
	return fmt.Sprintf("MessagesGetScheduledHistoryRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetScheduledHistoryRequest from given interface.
func (g *MessagesGetScheduledHistoryRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetHash() (value int)
}) {
	g.Peer = from.GetPeer()
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetScheduledHistoryRequest) TypeID() uint32 {
	return MessagesGetScheduledHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetScheduledHistoryRequest) TypeName() string {
	return "messages.getScheduledHistory"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetScheduledHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getScheduledHistory",
		ID:   MessagesGetScheduledHistoryRequestTypeID,
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
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetScheduledHistoryRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getScheduledHistory#e2c2685b as nil")
	}
	b.PutID(MessagesGetScheduledHistoryRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetScheduledHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getScheduledHistory#e2c2685b as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getScheduledHistory#e2c2685b: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getScheduledHistory#e2c2685b: field peer: %w", err)
	}
	b.PutInt(g.Hash)
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetScheduledHistoryRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// GetHash returns value of Hash field.
func (g *MessagesGetScheduledHistoryRequest) GetHash() (value int) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *MessagesGetScheduledHistoryRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getScheduledHistory#e2c2685b to nil")
	}
	if err := b.ConsumeID(MessagesGetScheduledHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getScheduledHistory#e2c2685b: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetScheduledHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getScheduledHistory#e2c2685b to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getScheduledHistory#e2c2685b: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getScheduledHistory#e2c2685b: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetScheduledHistoryRequest.
var (
	_ bin.Encoder     = &MessagesGetScheduledHistoryRequest{}
	_ bin.Decoder     = &MessagesGetScheduledHistoryRequest{}
	_ bin.BareEncoder = &MessagesGetScheduledHistoryRequest{}
	_ bin.BareDecoder = &MessagesGetScheduledHistoryRequest{}
)

// MessagesGetScheduledHistory invokes method messages.getScheduledHistory#e2c2685b returning error if any.
// Get scheduled messages
//
// Possible errors:
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.getScheduledHistory for reference.
func (c *Client) MessagesGetScheduledHistory(ctx context.Context, request *MessagesGetScheduledHistoryRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
