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

// MessagesSentEncryptedMessage represents TL type `messages.sentEncryptedMessage#560f8935`.
// Message without file attachemts sent to an encrypted file.
//
// See https://core.telegram.org/constructor/messages.sentEncryptedMessage for reference.
type MessagesSentEncryptedMessage struct {
	// Date of sending
	Date int
}

// MessagesSentEncryptedMessageTypeID is TL type id of MessagesSentEncryptedMessage.
const MessagesSentEncryptedMessageTypeID = 0x560f8935

// Encode implements bin.Encoder.
func (s *MessagesSentEncryptedMessage) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sentEncryptedMessage#560f8935 as nil")
	}
	b.PutID(MessagesSentEncryptedMessageTypeID)
	b.PutInt(s.Date)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSentEncryptedMessage) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sentEncryptedMessage#560f8935 to nil")
	}
	if err := b.ConsumeID(MessagesSentEncryptedMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sentEncryptedMessage#560f8935: %w", err)
	}

	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sentEncryptedMessage#560f8935: field date: %w", err)
		}
		s.Date = value
	}
	return nil
}

// construct implements constructor of MessagesSentEncryptedMessageClass.
func (s MessagesSentEncryptedMessage) construct() MessagesSentEncryptedMessageClass { return &s }

// Ensuring interfaces in compile-time for MessagesSentEncryptedMessage.
var (
	_ bin.Encoder = &MessagesSentEncryptedMessage{}
	_ bin.Decoder = &MessagesSentEncryptedMessage{}

	_ MessagesSentEncryptedMessageClass = &MessagesSentEncryptedMessage{}
)

// MessagesSentEncryptedFile represents TL type `messages.sentEncryptedFile#9493ff32`.
// Message with a file enclosure sent to a protected chat
//
// See https://core.telegram.org/constructor/messages.sentEncryptedFile for reference.
type MessagesSentEncryptedFile struct {
	// Sending date
	Date int
	// Attached file
	File EncryptedFileClass
}

// MessagesSentEncryptedFileTypeID is TL type id of MessagesSentEncryptedFile.
const MessagesSentEncryptedFileTypeID = 0x9493ff32

// Encode implements bin.Encoder.
func (s *MessagesSentEncryptedFile) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sentEncryptedFile#9493ff32 as nil")
	}
	b.PutID(MessagesSentEncryptedFileTypeID)
	b.PutInt(s.Date)
	if s.File == nil {
		return fmt.Errorf("unable to encode messages.sentEncryptedFile#9493ff32: field file is nil")
	}
	if err := s.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sentEncryptedFile#9493ff32: field file: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSentEncryptedFile) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sentEncryptedFile#9493ff32 to nil")
	}
	if err := b.ConsumeID(MessagesSentEncryptedFileTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sentEncryptedFile#9493ff32: %w", err)
	}

	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sentEncryptedFile#9493ff32: field date: %w", err)
		}
		s.Date = value
	}
	{
		value, err := DecodeEncryptedFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sentEncryptedFile#9493ff32: field file: %w", err)
		}
		s.File = value
	}
	return nil
}

// construct implements constructor of MessagesSentEncryptedMessageClass.
func (s MessagesSentEncryptedFile) construct() MessagesSentEncryptedMessageClass { return &s }

// Ensuring interfaces in compile-time for MessagesSentEncryptedFile.
var (
	_ bin.Encoder = &MessagesSentEncryptedFile{}
	_ bin.Decoder = &MessagesSentEncryptedFile{}

	_ MessagesSentEncryptedMessageClass = &MessagesSentEncryptedFile{}
)

// MessagesSentEncryptedMessageClass represents messages.SentEncryptedMessage generic type.
//
// See https://core.telegram.org/type/messages.SentEncryptedMessage for reference.
//
// Example:
//  g, err := DecodeMessagesSentEncryptedMessage(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *MessagesSentEncryptedMessage: // messages.sentEncryptedMessage#560f8935
//  case *MessagesSentEncryptedFile: // messages.sentEncryptedFile#9493ff32
//  default: panic(v)
//  }
type MessagesSentEncryptedMessageClass interface {
	bin.Encoder
	bin.Decoder
	construct() MessagesSentEncryptedMessageClass
}

// DecodeMessagesSentEncryptedMessage implements binary de-serialization for MessagesSentEncryptedMessageClass.
func DecodeMessagesSentEncryptedMessage(buf *bin.Buffer) (MessagesSentEncryptedMessageClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesSentEncryptedMessageTypeID:
		// Decoding messages.sentEncryptedMessage#560f8935.
		v := MessagesSentEncryptedMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesSentEncryptedMessageClass: %w", err)
		}
		return &v, nil
	case MessagesSentEncryptedFileTypeID:
		// Decoding messages.sentEncryptedFile#9493ff32.
		v := MessagesSentEncryptedFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesSentEncryptedMessageClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesSentEncryptedMessageClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesSentEncryptedMessage boxes the MessagesSentEncryptedMessageClass providing a helper.
type MessagesSentEncryptedMessageBox struct {
	SentEncryptedMessage MessagesSentEncryptedMessageClass
}

// Decode implements bin.Decoder for MessagesSentEncryptedMessageBox.
func (b *MessagesSentEncryptedMessageBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesSentEncryptedMessageBox to nil")
	}
	v, err := DecodeMessagesSentEncryptedMessage(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SentEncryptedMessage = v
	return nil
}

// Encode implements bin.Encode for MessagesSentEncryptedMessageBox.
func (b *MessagesSentEncryptedMessageBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SentEncryptedMessage == nil {
		return fmt.Errorf("unable to encode MessagesSentEncryptedMessageClass as nil")
	}
	return b.SentEncryptedMessage.Encode(buf)
}
