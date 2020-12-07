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

// MessagesGetBotCallbackAnswerRequest represents TL type `messages.getBotCallbackAnswer#9342ca07`.
// Press an inline callback button and get a callback answer from the bot
//
// See https://core.telegram.org/method/messages.getBotCallbackAnswer for reference.
type MessagesGetBotCallbackAnswerRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Whether this is a "play game" button
	Game bool
	// Where was the inline keyboard sent
	Peer InputPeerClass
	// ID of the Message with the inline keyboard
	MsgID int
	// Callback data
	//
	// Use SetData and GetData helpers.
	Data []byte
	// For buttons requiring you to verify your identity with your 2FA password, the SRP payload generated using SRP.
	//
	// Use SetPassword and GetPassword helpers.
	Password InputCheckPasswordSRPClass
}

// MessagesGetBotCallbackAnswerRequestTypeID is TL type id of MessagesGetBotCallbackAnswerRequest.
const MessagesGetBotCallbackAnswerRequestTypeID = 0x9342ca07

// Encode implements bin.Encoder.
func (g *MessagesGetBotCallbackAnswerRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getBotCallbackAnswer#9342ca07 as nil")
	}
	b.PutID(MessagesGetBotCallbackAnswerRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getBotCallbackAnswer#9342ca07: field flags: %w", err)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getBotCallbackAnswer#9342ca07: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getBotCallbackAnswer#9342ca07: field peer: %w", err)
	}
	b.PutInt(g.MsgID)
	if g.Flags.Has(0) {
		b.PutBytes(g.Data)
	}
	if g.Flags.Has(2) {
		if g.Password == nil {
			return fmt.Errorf("unable to encode messages.getBotCallbackAnswer#9342ca07: field password is nil")
		}
		if err := g.Password.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.getBotCallbackAnswer#9342ca07: field password: %w", err)
		}
	}
	return nil
}

// SetGame sets value of Game conditional field.
func (g *MessagesGetBotCallbackAnswerRequest) SetGame(value bool) {
	if value {
		g.Flags.Set(1)
	} else {
		g.Flags.Unset(1)
	}
}

// SetData sets value of Data conditional field.
func (g *MessagesGetBotCallbackAnswerRequest) SetData(value []byte) {
	g.Flags.Set(0)
	g.Data = value
}

// GetData returns value of Data conditional field and
// boolean which is true if field was set.
func (g *MessagesGetBotCallbackAnswerRequest) GetData() (value []byte, ok bool) {
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.Data, true
}

// SetPassword sets value of Password conditional field.
func (g *MessagesGetBotCallbackAnswerRequest) SetPassword(value InputCheckPasswordSRPClass) {
	g.Flags.Set(2)
	g.Password = value
}

// GetPassword returns value of Password conditional field and
// boolean which is true if field was set.
func (g *MessagesGetBotCallbackAnswerRequest) GetPassword() (value InputCheckPasswordSRPClass, ok bool) {
	if !g.Flags.Has(2) {
		return value, false
	}
	return g.Password, true
}

// Decode implements bin.Decoder.
func (g *MessagesGetBotCallbackAnswerRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getBotCallbackAnswer#9342ca07 to nil")
	}
	if err := b.ConsumeID(MessagesGetBotCallbackAnswerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: %w", err)
	}

	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: field flags: %w", err)
		}
	}
	g.Game = g.Flags.Has(1)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	if g.Flags.Has(0) {
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: field data: %w", err)
		}
		g.Data = value
	}
	if g.Flags.Has(2) {
		value, err := DecodeInputCheckPasswordSRP(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getBotCallbackAnswer#9342ca07: field password: %w", err)
		}
		g.Password = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetBotCallbackAnswerRequest.
var (
	_ bin.Encoder = &MessagesGetBotCallbackAnswerRequest{}
	_ bin.Decoder = &MessagesGetBotCallbackAnswerRequest{}
)

// MessagesGetBotCallbackAnswer invokes method messages.getBotCallbackAnswer#9342ca07 returning error if any.
// Press an inline callback button and get a callback answer from the bot
//
// See https://core.telegram.org/method/messages.getBotCallbackAnswer for reference.
func (c *Client) MessagesGetBotCallbackAnswer(ctx context.Context, request *MessagesGetBotCallbackAnswerRequest) (*MessagesBotCallbackAnswer, error) {
	var result MessagesBotCallbackAnswer
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
