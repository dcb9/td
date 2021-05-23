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

// MessagesExportChatInviteRequest represents TL type `messages.exportChatInvite#14b9bcd7`.
// Export an invite link for a chat
//
// See https://core.telegram.org/method/messages.exportChatInvite for reference.
type MessagesExportChatInviteRequest struct {
	// Flags field of MessagesExportChatInviteRequest.
	Flags bin.Fields
	// LegacyRevokePermanent field of MessagesExportChatInviteRequest.
	LegacyRevokePermanent bool
	// Chat
	Peer InputPeerClass
	// ExpireDate field of MessagesExportChatInviteRequest.
	//
	// Use SetExpireDate and GetExpireDate helpers.
	ExpireDate int
	// UsageLimit field of MessagesExportChatInviteRequest.
	//
	// Use SetUsageLimit and GetUsageLimit helpers.
	UsageLimit int
}

// MessagesExportChatInviteRequestTypeID is TL type id of MessagesExportChatInviteRequest.
const MessagesExportChatInviteRequestTypeID = 0x14b9bcd7

func (e *MessagesExportChatInviteRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Flags.Zero()) {
		return false
	}
	if !(e.LegacyRevokePermanent == false) {
		return false
	}
	if !(e.Peer == nil) {
		return false
	}
	if !(e.ExpireDate == 0) {
		return false
	}
	if !(e.UsageLimit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *MessagesExportChatInviteRequest) String() string {
	if e == nil {
		return "MessagesExportChatInviteRequest(nil)"
	}
	type Alias MessagesExportChatInviteRequest
	return fmt.Sprintf("MessagesExportChatInviteRequest%+v", Alias(*e))
}

// FillFrom fills MessagesExportChatInviteRequest from given interface.
func (e *MessagesExportChatInviteRequest) FillFrom(from interface {
	GetLegacyRevokePermanent() (value bool)
	GetPeer() (value InputPeerClass)
	GetExpireDate() (value int, ok bool)
	GetUsageLimit() (value int, ok bool)
}) {
	e.LegacyRevokePermanent = from.GetLegacyRevokePermanent()
	e.Peer = from.GetPeer()
	if val, ok := from.GetExpireDate(); ok {
		e.ExpireDate = val
	}

	if val, ok := from.GetUsageLimit(); ok {
		e.UsageLimit = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesExportChatInviteRequest) TypeID() uint32 {
	return MessagesExportChatInviteRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesExportChatInviteRequest) TypeName() string {
	return "messages.exportChatInvite"
}

// TypeInfo returns info about TL type.
func (e *MessagesExportChatInviteRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.exportChatInvite",
		ID:   MessagesExportChatInviteRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LegacyRevokePermanent",
			SchemaName: "legacy_revoke_permanent",
			Null:       !e.Flags.Has(2),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ExpireDate",
			SchemaName: "expire_date",
			Null:       !e.Flags.Has(0),
		},
		{
			Name:       "UsageLimit",
			SchemaName: "usage_limit",
			Null:       !e.Flags.Has(1),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *MessagesExportChatInviteRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.exportChatInvite#14b9bcd7 as nil")
	}
	b.PutID(MessagesExportChatInviteRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *MessagesExportChatInviteRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.exportChatInvite#14b9bcd7 as nil")
	}
	if !(e.LegacyRevokePermanent == false) {
		e.Flags.Set(2)
	}
	if !(e.ExpireDate == 0) {
		e.Flags.Set(0)
	}
	if !(e.UsageLimit == 0) {
		e.Flags.Set(1)
	}
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.exportChatInvite#14b9bcd7: field flags: %w", err)
	}
	if e.Peer == nil {
		return fmt.Errorf("unable to encode messages.exportChatInvite#14b9bcd7: field peer is nil")
	}
	if err := e.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.exportChatInvite#14b9bcd7: field peer: %w", err)
	}
	if e.Flags.Has(0) {
		b.PutInt(e.ExpireDate)
	}
	if e.Flags.Has(1) {
		b.PutInt(e.UsageLimit)
	}
	return nil
}

// SetLegacyRevokePermanent sets value of LegacyRevokePermanent conditional field.
func (e *MessagesExportChatInviteRequest) SetLegacyRevokePermanent(value bool) {
	if value {
		e.Flags.Set(2)
		e.LegacyRevokePermanent = true
	} else {
		e.Flags.Unset(2)
		e.LegacyRevokePermanent = false
	}
}

// GetLegacyRevokePermanent returns value of LegacyRevokePermanent conditional field.
func (e *MessagesExportChatInviteRequest) GetLegacyRevokePermanent() (value bool) {
	return e.Flags.Has(2)
}

// GetPeer returns value of Peer field.
func (e *MessagesExportChatInviteRequest) GetPeer() (value InputPeerClass) {
	return e.Peer
}

// SetExpireDate sets value of ExpireDate conditional field.
func (e *MessagesExportChatInviteRequest) SetExpireDate(value int) {
	e.Flags.Set(0)
	e.ExpireDate = value
}

// GetExpireDate returns value of ExpireDate conditional field and
// boolean which is true if field was set.
func (e *MessagesExportChatInviteRequest) GetExpireDate() (value int, ok bool) {
	if !e.Flags.Has(0) {
		return value, false
	}
	return e.ExpireDate, true
}

// SetUsageLimit sets value of UsageLimit conditional field.
func (e *MessagesExportChatInviteRequest) SetUsageLimit(value int) {
	e.Flags.Set(1)
	e.UsageLimit = value
}

// GetUsageLimit returns value of UsageLimit conditional field and
// boolean which is true if field was set.
func (e *MessagesExportChatInviteRequest) GetUsageLimit() (value int, ok bool) {
	if !e.Flags.Has(1) {
		return value, false
	}
	return e.UsageLimit, true
}

// Decode implements bin.Decoder.
func (e *MessagesExportChatInviteRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.exportChatInvite#14b9bcd7 to nil")
	}
	if err := b.ConsumeID(MessagesExportChatInviteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.exportChatInvite#14b9bcd7: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *MessagesExportChatInviteRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.exportChatInvite#14b9bcd7 to nil")
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.exportChatInvite#14b9bcd7: field flags: %w", err)
		}
	}
	e.LegacyRevokePermanent = e.Flags.Has(2)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.exportChatInvite#14b9bcd7: field peer: %w", err)
		}
		e.Peer = value
	}
	if e.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.exportChatInvite#14b9bcd7: field expire_date: %w", err)
		}
		e.ExpireDate = value
	}
	if e.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.exportChatInvite#14b9bcd7: field usage_limit: %w", err)
		}
		e.UsageLimit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesExportChatInviteRequest.
var (
	_ bin.Encoder     = &MessagesExportChatInviteRequest{}
	_ bin.Decoder     = &MessagesExportChatInviteRequest{}
	_ bin.BareEncoder = &MessagesExportChatInviteRequest{}
	_ bin.BareDecoder = &MessagesExportChatInviteRequest{}
)

// MessagesExportChatInvite invokes method messages.exportChatInvite#14b9bcd7 returning error if any.
// Export an invite link for a chat
//
// Possible errors:
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.exportChatInvite for reference.
// Can be used by bots.
func (c *Client) MessagesExportChatInvite(ctx context.Context, request *MessagesExportChatInviteRequest) (*ChatInviteExported, error) {
	var result ChatInviteExported

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
