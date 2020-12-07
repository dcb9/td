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

// MessageFwdHeader represents TL type `messageFwdHeader#5f777dce`.
// Info about a forwarded message
//
// See https://core.telegram.org/constructor/messageFwdHeader for reference.
type MessageFwdHeader struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// The ID of the user that originally sent the message
	//
	// Use SetFromID and GetFromID helpers.
	FromID PeerClass
	// The name of the user that originally sent the message
	//
	// Use SetFromName and GetFromName helpers.
	FromName string
	// When was the message originally sent
	Date int
	// ID of the channel message that was forwarded
	//
	// Use SetChannelPost and GetChannelPost helpers.
	ChannelPost int
	// For channels and if signatures are enabled, author of the channel message
	//
	// Use SetPostAuthor and GetPostAuthor helpers.
	PostAuthor string
	// Only for messages forwarded to the current user (inputPeerSelf), full info about the user/channel that originally sent the message
	//
	// Use SetSavedFromPeer and GetSavedFromPeer helpers.
	SavedFromPeer PeerClass
	// Only for messages forwarded to the current user (inputPeerSelf), ID of the message that was forwarded from the original user/channel
	//
	// Use SetSavedFromMsgID and GetSavedFromMsgID helpers.
	SavedFromMsgID int
	// PSA type
	//
	// Use SetPsaType and GetPsaType helpers.
	PsaType string
}

// MessageFwdHeaderTypeID is TL type id of MessageFwdHeader.
const MessageFwdHeaderTypeID = 0x5f777dce

// Encode implements bin.Encoder.
func (m *MessageFwdHeader) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageFwdHeader#5f777dce as nil")
	}
	b.PutID(MessageFwdHeaderTypeID)
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageFwdHeader#5f777dce: field flags: %w", err)
	}
	if m.Flags.Has(0) {
		if m.FromID == nil {
			return fmt.Errorf("unable to encode messageFwdHeader#5f777dce: field from_id is nil")
		}
		if err := m.FromID.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messageFwdHeader#5f777dce: field from_id: %w", err)
		}
	}
	if m.Flags.Has(5) {
		b.PutString(m.FromName)
	}
	b.PutInt(m.Date)
	if m.Flags.Has(2) {
		b.PutInt(m.ChannelPost)
	}
	if m.Flags.Has(3) {
		b.PutString(m.PostAuthor)
	}
	if m.Flags.Has(4) {
		if m.SavedFromPeer == nil {
			return fmt.Errorf("unable to encode messageFwdHeader#5f777dce: field saved_from_peer is nil")
		}
		if err := m.SavedFromPeer.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messageFwdHeader#5f777dce: field saved_from_peer: %w", err)
		}
	}
	if m.Flags.Has(4) {
		b.PutInt(m.SavedFromMsgID)
	}
	if m.Flags.Has(6) {
		b.PutString(m.PsaType)
	}
	return nil
}

// SetFromID sets value of FromID conditional field.
func (m *MessageFwdHeader) SetFromID(value PeerClass) {
	m.Flags.Set(0)
	m.FromID = value
}

// GetFromID returns value of FromID conditional field and
// boolean which is true if field was set.
func (m *MessageFwdHeader) GetFromID() (value PeerClass, ok bool) {
	if !m.Flags.Has(0) {
		return value, false
	}
	return m.FromID, true
}

// SetFromName sets value of FromName conditional field.
func (m *MessageFwdHeader) SetFromName(value string) {
	m.Flags.Set(5)
	m.FromName = value
}

// GetFromName returns value of FromName conditional field and
// boolean which is true if field was set.
func (m *MessageFwdHeader) GetFromName() (value string, ok bool) {
	if !m.Flags.Has(5) {
		return value, false
	}
	return m.FromName, true
}

// SetChannelPost sets value of ChannelPost conditional field.
func (m *MessageFwdHeader) SetChannelPost(value int) {
	m.Flags.Set(2)
	m.ChannelPost = value
}

// GetChannelPost returns value of ChannelPost conditional field and
// boolean which is true if field was set.
func (m *MessageFwdHeader) GetChannelPost() (value int, ok bool) {
	if !m.Flags.Has(2) {
		return value, false
	}
	return m.ChannelPost, true
}

// SetPostAuthor sets value of PostAuthor conditional field.
func (m *MessageFwdHeader) SetPostAuthor(value string) {
	m.Flags.Set(3)
	m.PostAuthor = value
}

// GetPostAuthor returns value of PostAuthor conditional field and
// boolean which is true if field was set.
func (m *MessageFwdHeader) GetPostAuthor() (value string, ok bool) {
	if !m.Flags.Has(3) {
		return value, false
	}
	return m.PostAuthor, true
}

// SetSavedFromPeer sets value of SavedFromPeer conditional field.
func (m *MessageFwdHeader) SetSavedFromPeer(value PeerClass) {
	m.Flags.Set(4)
	m.SavedFromPeer = value
}

// GetSavedFromPeer returns value of SavedFromPeer conditional field and
// boolean which is true if field was set.
func (m *MessageFwdHeader) GetSavedFromPeer() (value PeerClass, ok bool) {
	if !m.Flags.Has(4) {
		return value, false
	}
	return m.SavedFromPeer, true
}

// SetSavedFromMsgID sets value of SavedFromMsgID conditional field.
func (m *MessageFwdHeader) SetSavedFromMsgID(value int) {
	m.Flags.Set(4)
	m.SavedFromMsgID = value
}

// GetSavedFromMsgID returns value of SavedFromMsgID conditional field and
// boolean which is true if field was set.
func (m *MessageFwdHeader) GetSavedFromMsgID() (value int, ok bool) {
	if !m.Flags.Has(4) {
		return value, false
	}
	return m.SavedFromMsgID, true
}

// SetPsaType sets value of PsaType conditional field.
func (m *MessageFwdHeader) SetPsaType(value string) {
	m.Flags.Set(6)
	m.PsaType = value
}

// GetPsaType returns value of PsaType conditional field and
// boolean which is true if field was set.
func (m *MessageFwdHeader) GetPsaType() (value string, ok bool) {
	if !m.Flags.Has(6) {
		return value, false
	}
	return m.PsaType, true
}

// Decode implements bin.Decoder.
func (m *MessageFwdHeader) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageFwdHeader#5f777dce to nil")
	}
	if err := b.ConsumeID(MessageFwdHeaderTypeID); err != nil {
		return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: %w", err)
	}

	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field flags: %w", err)
		}
	}
	if m.Flags.Has(0) {
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field from_id: %w", err)
		}
		m.FromID = value
	}
	if m.Flags.Has(5) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field from_name: %w", err)
		}
		m.FromName = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field date: %w", err)
		}
		m.Date = value
	}
	if m.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field channel_post: %w", err)
		}
		m.ChannelPost = value
	}
	if m.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field post_author: %w", err)
		}
		m.PostAuthor = value
	}
	if m.Flags.Has(4) {
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field saved_from_peer: %w", err)
		}
		m.SavedFromPeer = value
	}
	if m.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field saved_from_msg_id: %w", err)
		}
		m.SavedFromMsgID = value
	}
	if m.Flags.Has(6) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageFwdHeader#5f777dce: field psa_type: %w", err)
		}
		m.PsaType = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessageFwdHeader.
var (
	_ bin.Encoder = &MessageFwdHeader{}
	_ bin.Decoder = &MessageFwdHeader{}
)
