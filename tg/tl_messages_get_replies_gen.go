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

// MessagesGetRepliesRequest represents TL type `messages.getReplies#24b581ba`.
// Get messages in a reply thread
//
// See https://core.telegram.org/method/messages.getReplies for reference.
type MessagesGetRepliesRequest struct {
	// Peer
	Peer InputPeerClass
	// Message ID
	MsgID int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetID int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetDate int
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
	// If a positive value was transferred, the method will return only messages with ID
	// smaller than max_id
	MaxID int
	// If a positive value was transferred, the method will return only messages with ID
	// bigger than min_id
	MinID int
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// MessagesGetRepliesRequestTypeID is TL type id of MessagesGetRepliesRequest.
const MessagesGetRepliesRequestTypeID = 0x24b581ba

func (g *MessagesGetRepliesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.MsgID == 0) {
		return false
	}
	if !(g.OffsetID == 0) {
		return false
	}
	if !(g.OffsetDate == 0) {
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
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetRepliesRequest) String() string {
	if g == nil {
		return "MessagesGetRepliesRequest(nil)"
	}
	type Alias MessagesGetRepliesRequest
	return fmt.Sprintf("MessagesGetRepliesRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetRepliesRequest from given interface.
func (g *MessagesGetRepliesRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetOffsetID() (value int)
	GetOffsetDate() (value int)
	GetAddOffset() (value int)
	GetLimit() (value int)
	GetMaxID() (value int)
	GetMinID() (value int)
	GetHash() (value int)
}) {
	g.Peer = from.GetPeer()
	g.MsgID = from.GetMsgID()
	g.OffsetID = from.GetOffsetID()
	g.OffsetDate = from.GetOffsetDate()
	g.AddOffset = from.GetAddOffset()
	g.Limit = from.GetLimit()
	g.MaxID = from.GetMaxID()
	g.MinID = from.GetMinID()
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetRepliesRequest) TypeID() uint32 {
	return MessagesGetRepliesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetRepliesRequest) TypeName() string {
	return "messages.getReplies"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetRepliesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getReplies",
		ID:   MessagesGetRepliesRequestTypeID,
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
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "OffsetID",
			SchemaName: "offset_id",
		},
		{
			Name:       "OffsetDate",
			SchemaName: "offset_date",
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
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetRepliesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getReplies#24b581ba as nil")
	}
	b.PutID(MessagesGetRepliesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetRepliesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getReplies#24b581ba as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getReplies#24b581ba: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getReplies#24b581ba: field peer: %w", err)
	}
	b.PutInt(g.MsgID)
	b.PutInt(g.OffsetID)
	b.PutInt(g.OffsetDate)
	b.PutInt(g.AddOffset)
	b.PutInt(g.Limit)
	b.PutInt(g.MaxID)
	b.PutInt(g.MinID)
	b.PutInt(g.Hash)
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetRepliesRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// GetMsgID returns value of MsgID field.
func (g *MessagesGetRepliesRequest) GetMsgID() (value int) {
	return g.MsgID
}

// GetOffsetID returns value of OffsetID field.
func (g *MessagesGetRepliesRequest) GetOffsetID() (value int) {
	return g.OffsetID
}

// GetOffsetDate returns value of OffsetDate field.
func (g *MessagesGetRepliesRequest) GetOffsetDate() (value int) {
	return g.OffsetDate
}

// GetAddOffset returns value of AddOffset field.
func (g *MessagesGetRepliesRequest) GetAddOffset() (value int) {
	return g.AddOffset
}

// GetLimit returns value of Limit field.
func (g *MessagesGetRepliesRequest) GetLimit() (value int) {
	return g.Limit
}

// GetMaxID returns value of MaxID field.
func (g *MessagesGetRepliesRequest) GetMaxID() (value int) {
	return g.MaxID
}

// GetMinID returns value of MinID field.
func (g *MessagesGetRepliesRequest) GetMinID() (value int) {
	return g.MinID
}

// GetHash returns value of Hash field.
func (g *MessagesGetRepliesRequest) GetHash() (value int) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *MessagesGetRepliesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getReplies#24b581ba to nil")
	}
	if err := b.ConsumeID(MessagesGetRepliesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getReplies#24b581ba: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetRepliesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getReplies#24b581ba to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field offset_date: %w", err)
		}
		g.OffsetDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field add_offset: %w", err)
		}
		g.AddOffset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field max_id: %w", err)
		}
		g.MaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field min_id: %w", err)
		}
		g.MinID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getReplies#24b581ba: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetRepliesRequest.
var (
	_ bin.Encoder     = &MessagesGetRepliesRequest{}
	_ bin.Decoder     = &MessagesGetRepliesRequest{}
	_ bin.BareEncoder = &MessagesGetRepliesRequest{}
	_ bin.BareDecoder = &MessagesGetRepliesRequest{}
)

// MessagesGetReplies invokes method messages.getReplies#24b581ba returning error if any.
// Get messages in a reply thread
//
// See https://core.telegram.org/method/messages.getReplies for reference.
// Can be used by bots.
func (c *Client) MessagesGetReplies(ctx context.Context, request *MessagesGetRepliesRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
