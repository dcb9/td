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

// MessagesGetUnreadReactionsRequest represents TL type `messages.getUnreadReactions#e85bae1a`.
// Get unread reactions to messages you sent
//
// See https://core.telegram.org/method/messages.getUnreadReactions for reference.
type MessagesGetUnreadReactionsRequest struct {
	// Peer
	Peer InputPeerClass
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetID int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	AddOffset int
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
	// Only return reactions for messages up until this message ID
	MaxID int
	// Only return reactions for messages starting from this message ID
	MinID int
}

// MessagesGetUnreadReactionsRequestTypeID is TL type id of MessagesGetUnreadReactionsRequest.
const MessagesGetUnreadReactionsRequestTypeID = 0xe85bae1a

// Ensuring interfaces in compile-time for MessagesGetUnreadReactionsRequest.
var (
	_ bin.Encoder     = &MessagesGetUnreadReactionsRequest{}
	_ bin.Decoder     = &MessagesGetUnreadReactionsRequest{}
	_ bin.BareEncoder = &MessagesGetUnreadReactionsRequest{}
	_ bin.BareDecoder = &MessagesGetUnreadReactionsRequest{}
)

func (g *MessagesGetUnreadReactionsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.OffsetID == 0) {
		return false
	}
	if !(g.AddOffset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}
	if !(g.MaxID == 0) {
		return false
	}
	if !(g.MinID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetUnreadReactionsRequest) String() string {
	if g == nil {
		return "MessagesGetUnreadReactionsRequest(nil)"
	}
	type Alias MessagesGetUnreadReactionsRequest
	return fmt.Sprintf("MessagesGetUnreadReactionsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetUnreadReactionsRequest from given interface.
func (g *MessagesGetUnreadReactionsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetOffsetID() (value int)
	GetAddOffset() (value int)
	GetLimit() (value int)
	GetMaxID() (value int)
	GetMinID() (value int)
}) {
	g.Peer = from.GetPeer()
	g.OffsetID = from.GetOffsetID()
	g.AddOffset = from.GetAddOffset()
	g.Limit = from.GetLimit()
	g.MaxID = from.GetMaxID()
	g.MinID = from.GetMinID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetUnreadReactionsRequest) TypeID() uint32 {
	return MessagesGetUnreadReactionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetUnreadReactionsRequest) TypeName() string {
	return "messages.getUnreadReactions"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetUnreadReactionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getUnreadReactions",
		ID:   MessagesGetUnreadReactionsRequestTypeID,
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
			Name:       "OffsetID",
			SchemaName: "offset_id",
		},
		{
			Name:       "AddOffset",
			SchemaName: "add_offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "MaxID",
			SchemaName: "max_id",
		},
		{
			Name:       "MinID",
			SchemaName: "min_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetUnreadReactionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getUnreadReactions#e85bae1a as nil")
	}
	b.PutID(MessagesGetUnreadReactionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetUnreadReactionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getUnreadReactions#e85bae1a as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getUnreadReactions#e85bae1a: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getUnreadReactions#e85bae1a: field peer: %w", err)
	}
	b.PutInt(g.OffsetID)
	b.PutInt(g.AddOffset)
	b.PutInt(g.Limit)
	b.PutInt(g.MaxID)
	b.PutInt(g.MinID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetUnreadReactionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getUnreadReactions#e85bae1a to nil")
	}
	if err := b.ConsumeID(MessagesGetUnreadReactionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getUnreadReactions#e85bae1a: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetUnreadReactionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getUnreadReactions#e85bae1a to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadReactions#e85bae1a: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadReactions#e85bae1a: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadReactions#e85bae1a: field add_offset: %w", err)
		}
		g.AddOffset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadReactions#e85bae1a: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadReactions#e85bae1a: field max_id: %w", err)
		}
		g.MaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getUnreadReactions#e85bae1a: field min_id: %w", err)
		}
		g.MinID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetUnreadReactionsRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetOffsetID returns value of OffsetID field.
func (g *MessagesGetUnreadReactionsRequest) GetOffsetID() (value int) {
	if g == nil {
		return
	}
	return g.OffsetID
}

// GetAddOffset returns value of AddOffset field.
func (g *MessagesGetUnreadReactionsRequest) GetAddOffset() (value int) {
	if g == nil {
		return
	}
	return g.AddOffset
}

// GetLimit returns value of Limit field.
func (g *MessagesGetUnreadReactionsRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetMaxID returns value of MaxID field.
func (g *MessagesGetUnreadReactionsRequest) GetMaxID() (value int) {
	if g == nil {
		return
	}
	return g.MaxID
}

// GetMinID returns value of MinID field.
func (g *MessagesGetUnreadReactionsRequest) GetMinID() (value int) {
	if g == nil {
		return
	}
	return g.MinID
}

// MessagesGetUnreadReactions invokes method messages.getUnreadReactions#e85bae1a returning error if any.
// Get unread reactions to messages you sent
//
// See https://core.telegram.org/method/messages.getUnreadReactions for reference.
func (c *Client) MessagesGetUnreadReactions(ctx context.Context, request *MessagesGetUnreadReactionsRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
