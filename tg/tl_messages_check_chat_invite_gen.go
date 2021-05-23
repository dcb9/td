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

// MessagesCheckChatInviteRequest represents TL type `messages.checkChatInvite#3eadb1bb`.
// Check the validity of a chat invite link and get basic info about it
//
// See https://core.telegram.org/method/messages.checkChatInvite for reference.
type MessagesCheckChatInviteRequest struct {
	// Invite hash in t.me/joinchat/hash
	Hash string
}

// MessagesCheckChatInviteRequestTypeID is TL type id of MessagesCheckChatInviteRequest.
const MessagesCheckChatInviteRequestTypeID = 0x3eadb1bb

func (c *MessagesCheckChatInviteRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Hash == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *MessagesCheckChatInviteRequest) String() string {
	if c == nil {
		return "MessagesCheckChatInviteRequest(nil)"
	}
	type Alias MessagesCheckChatInviteRequest
	return fmt.Sprintf("MessagesCheckChatInviteRequest%+v", Alias(*c))
}

// FillFrom fills MessagesCheckChatInviteRequest from given interface.
func (c *MessagesCheckChatInviteRequest) FillFrom(from interface {
	GetHash() (value string)
}) {
	c.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesCheckChatInviteRequest) TypeID() uint32 {
	return MessagesCheckChatInviteRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesCheckChatInviteRequest) TypeName() string {
	return "messages.checkChatInvite"
}

// TypeInfo returns info about TL type.
func (c *MessagesCheckChatInviteRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.checkChatInvite",
		ID:   MessagesCheckChatInviteRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *MessagesCheckChatInviteRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.checkChatInvite#3eadb1bb as nil")
	}
	b.PutID(MessagesCheckChatInviteRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *MessagesCheckChatInviteRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.checkChatInvite#3eadb1bb as nil")
	}
	b.PutString(c.Hash)
	return nil
}

// GetHash returns value of Hash field.
func (c *MessagesCheckChatInviteRequest) GetHash() (value string) {
	return c.Hash
}

// Decode implements bin.Decoder.
func (c *MessagesCheckChatInviteRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.checkChatInvite#3eadb1bb to nil")
	}
	if err := b.ConsumeID(MessagesCheckChatInviteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.checkChatInvite#3eadb1bb: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *MessagesCheckChatInviteRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.checkChatInvite#3eadb1bb to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.checkChatInvite#3eadb1bb: field hash: %w", err)
		}
		c.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesCheckChatInviteRequest.
var (
	_ bin.Encoder     = &MessagesCheckChatInviteRequest{}
	_ bin.Decoder     = &MessagesCheckChatInviteRequest{}
	_ bin.BareEncoder = &MessagesCheckChatInviteRequest{}
	_ bin.BareDecoder = &MessagesCheckChatInviteRequest{}
)

// MessagesCheckChatInvite invokes method messages.checkChatInvite#3eadb1bb returning error if any.
// Check the validity of a chat invite link and get basic info about it
//
// Possible errors:
//  400 INVITE_HASH_EMPTY: The invite hash is empty
//  400 INVITE_HASH_EXPIRED: The invite link has expired
//  400 INVITE_HASH_INVALID: The invite hash is invalid
//
// See https://core.telegram.org/method/messages.checkChatInvite for reference.
func (c *Client) MessagesCheckChatInvite(ctx context.Context, hash string) (ChatInviteClass, error) {
	var result ChatInviteBox

	request := &MessagesCheckChatInviteRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ChatInvite, nil
}
