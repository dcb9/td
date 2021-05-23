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

// MessagesGetHistoryRequest represents TL type `messages.getHistory#dcbb8260`.
// Gets back the conversation history with one interlocutor / within a chat
//
// See https://core.telegram.org/method/messages.getHistory for reference.
type MessagesGetHistoryRequest struct {
	// Target peer
	Peer InputPeerClass
	// Only return messages starting from the specified message ID
	OffsetID int
	// Only return messages sent before the specified date
	OffsetDate int
	// Number of list elements to be skipped, negative values are also accepted.
	AddOffset int
	// Number of results to return
	Limit int
	// If a positive value was transferred, the method will return only messages with IDs
	// less than max_id
	MaxID int
	// If a positive value was transferred, the method will return only messages with IDs
	// more than min_id
	MinID int
	// Result hash¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Hash int
}

// MessagesGetHistoryRequestTypeID is TL type id of MessagesGetHistoryRequest.
const MessagesGetHistoryRequestTypeID = 0xdcbb8260

func (g *MessagesGetHistoryRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
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
func (g *MessagesGetHistoryRequest) String() string {
	if g == nil {
		return "MessagesGetHistoryRequest(nil)"
	}
	type Alias MessagesGetHistoryRequest
	return fmt.Sprintf("MessagesGetHistoryRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetHistoryRequest from given interface.
func (g *MessagesGetHistoryRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetOffsetID() (value int)
	GetOffsetDate() (value int)
	GetAddOffset() (value int)
	GetLimit() (value int)
	GetMaxID() (value int)
	GetMinID() (value int)
	GetHash() (value int)
}) {
	g.Peer = from.GetPeer()
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
func (*MessagesGetHistoryRequest) TypeID() uint32 {
	return MessagesGetHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetHistoryRequest) TypeName() string {
	return "messages.getHistory"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getHistory",
		ID:   MessagesGetHistoryRequestTypeID,
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
func (g *MessagesGetHistoryRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getHistory#dcbb8260 as nil")
	}
	b.PutID(MessagesGetHistoryRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getHistory#dcbb8260 as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getHistory#dcbb8260: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getHistory#dcbb8260: field peer: %w", err)
	}
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
func (g *MessagesGetHistoryRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// GetOffsetID returns value of OffsetID field.
func (g *MessagesGetHistoryRequest) GetOffsetID() (value int) {
	return g.OffsetID
}

// GetOffsetDate returns value of OffsetDate field.
func (g *MessagesGetHistoryRequest) GetOffsetDate() (value int) {
	return g.OffsetDate
}

// GetAddOffset returns value of AddOffset field.
func (g *MessagesGetHistoryRequest) GetAddOffset() (value int) {
	return g.AddOffset
}

// GetLimit returns value of Limit field.
func (g *MessagesGetHistoryRequest) GetLimit() (value int) {
	return g.Limit
}

// GetMaxID returns value of MaxID field.
func (g *MessagesGetHistoryRequest) GetMaxID() (value int) {
	return g.MaxID
}

// GetMinID returns value of MinID field.
func (g *MessagesGetHistoryRequest) GetMinID() (value int) {
	return g.MinID
}

// GetHash returns value of Hash field.
func (g *MessagesGetHistoryRequest) GetHash() (value int) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *MessagesGetHistoryRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getHistory#dcbb8260 to nil")
	}
	if err := b.ConsumeID(MessagesGetHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getHistory#dcbb8260 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: field offset_date: %w", err)
		}
		g.OffsetDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: field add_offset: %w", err)
		}
		g.AddOffset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: field max_id: %w", err)
		}
		g.MaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: field min_id: %w", err)
		}
		g.MinID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getHistory#dcbb8260: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetHistoryRequest.
var (
	_ bin.Encoder     = &MessagesGetHistoryRequest{}
	_ bin.Decoder     = &MessagesGetHistoryRequest{}
	_ bin.BareEncoder = &MessagesGetHistoryRequest{}
	_ bin.BareDecoder = &MessagesGetHistoryRequest{}
)

// MessagesGetHistory invokes method messages.getHistory#dcbb8260 returning error if any.
// Gets back the conversation history with one interlocutor / within a chat
//
// Possible errors:
//  401 AUTH_KEY_PERM_EMPTY: The temporary auth key must be binded to the permanent auth key to use these methods.
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  400 CONNECTION_DEVICE_MODEL_EMPTY: Device model empty
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.getHistory for reference.
func (c *Client) MessagesGetHistory(ctx context.Context, request *MessagesGetHistoryRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
