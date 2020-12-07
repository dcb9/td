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

// MessagesExportChatInviteRequest represents TL type `messages.exportChatInvite#df7534c`.
// Export an invite link for a chat
//
// See https://core.telegram.org/method/messages.exportChatInvite for reference.
type MessagesExportChatInviteRequest struct {
	// Chat
	Peer InputPeerClass
}

// MessagesExportChatInviteRequestTypeID is TL type id of MessagesExportChatInviteRequest.
const MessagesExportChatInviteRequestTypeID = 0xdf7534c

// Encode implements bin.Encoder.
func (e *MessagesExportChatInviteRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.exportChatInvite#df7534c as nil")
	}
	b.PutID(MessagesExportChatInviteRequestTypeID)
	if e.Peer == nil {
		return fmt.Errorf("unable to encode messages.exportChatInvite#df7534c: field peer is nil")
	}
	if err := e.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.exportChatInvite#df7534c: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *MessagesExportChatInviteRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.exportChatInvite#df7534c to nil")
	}
	if err := b.ConsumeID(MessagesExportChatInviteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.exportChatInvite#df7534c: %w", err)
	}

	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.exportChatInvite#df7534c: field peer: %w", err)
		}
		e.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesExportChatInviteRequest.
var (
	_ bin.Encoder = &MessagesExportChatInviteRequest{}
	_ bin.Decoder = &MessagesExportChatInviteRequest{}
)

// MessagesExportChatInvite invokes method messages.exportChatInvite#df7534c returning error if any.
// Export an invite link for a chat
//
// See https://core.telegram.org/method/messages.exportChatInvite for reference.
func (c *Client) MessagesExportChatInvite(ctx context.Context, request *MessagesExportChatInviteRequest) (ExportedChatInviteClass, error) {
	var result ExportedChatInviteBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ExportedChatInvite, nil
}
