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

// PrivacyKeyStatusTimestamp represents TL type `privacyKeyStatusTimestamp#bc2eab30`.
// Whether we can see the last online timestamp
//
// See https://core.telegram.org/constructor/privacyKeyStatusTimestamp for reference.
type PrivacyKeyStatusTimestamp struct {
}

// PrivacyKeyStatusTimestampTypeID is TL type id of PrivacyKeyStatusTimestamp.
const PrivacyKeyStatusTimestampTypeID = 0xbc2eab30

// Encode implements bin.Encoder.
func (p *PrivacyKeyStatusTimestamp) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyStatusTimestamp#bc2eab30 as nil")
	}
	b.PutID(PrivacyKeyStatusTimestampTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyStatusTimestamp) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyStatusTimestamp#bc2eab30 to nil")
	}
	if err := b.ConsumeID(PrivacyKeyStatusTimestampTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyStatusTimestamp#bc2eab30: %w", err)
	}

	return nil
}

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyStatusTimestamp) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyStatusTimestamp.
var (
	_ bin.Encoder = &PrivacyKeyStatusTimestamp{}
	_ bin.Decoder = &PrivacyKeyStatusTimestamp{}

	_ PrivacyKeyClass = &PrivacyKeyStatusTimestamp{}
)

// PrivacyKeyChatInvite represents TL type `privacyKeyChatInvite#500e6dfa`.
// Whether the user can be invited to chats
//
// See https://core.telegram.org/constructor/privacyKeyChatInvite for reference.
type PrivacyKeyChatInvite struct {
}

// PrivacyKeyChatInviteTypeID is TL type id of PrivacyKeyChatInvite.
const PrivacyKeyChatInviteTypeID = 0x500e6dfa

// Encode implements bin.Encoder.
func (p *PrivacyKeyChatInvite) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyChatInvite#500e6dfa as nil")
	}
	b.PutID(PrivacyKeyChatInviteTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyChatInvite) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyChatInvite#500e6dfa to nil")
	}
	if err := b.ConsumeID(PrivacyKeyChatInviteTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyChatInvite#500e6dfa: %w", err)
	}

	return nil
}

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyChatInvite) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyChatInvite.
var (
	_ bin.Encoder = &PrivacyKeyChatInvite{}
	_ bin.Decoder = &PrivacyKeyChatInvite{}

	_ PrivacyKeyClass = &PrivacyKeyChatInvite{}
)

// PrivacyKeyPhoneCall represents TL type `privacyKeyPhoneCall#3d662b7b`.
// Whether the user accepts phone calls
//
// See https://core.telegram.org/constructor/privacyKeyPhoneCall for reference.
type PrivacyKeyPhoneCall struct {
}

// PrivacyKeyPhoneCallTypeID is TL type id of PrivacyKeyPhoneCall.
const PrivacyKeyPhoneCallTypeID = 0x3d662b7b

// Encode implements bin.Encoder.
func (p *PrivacyKeyPhoneCall) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyPhoneCall#3d662b7b as nil")
	}
	b.PutID(PrivacyKeyPhoneCallTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyPhoneCall) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyPhoneCall#3d662b7b to nil")
	}
	if err := b.ConsumeID(PrivacyKeyPhoneCallTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyPhoneCall#3d662b7b: %w", err)
	}

	return nil
}

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyPhoneCall) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyPhoneCall.
var (
	_ bin.Encoder = &PrivacyKeyPhoneCall{}
	_ bin.Decoder = &PrivacyKeyPhoneCall{}

	_ PrivacyKeyClass = &PrivacyKeyPhoneCall{}
)

// PrivacyKeyPhoneP2P represents TL type `privacyKeyPhoneP2P#39491cc8`.
// Whether P2P connections in phone calls are allowed
//
// See https://core.telegram.org/constructor/privacyKeyPhoneP2P for reference.
type PrivacyKeyPhoneP2P struct {
}

// PrivacyKeyPhoneP2PTypeID is TL type id of PrivacyKeyPhoneP2P.
const PrivacyKeyPhoneP2PTypeID = 0x39491cc8

// Encode implements bin.Encoder.
func (p *PrivacyKeyPhoneP2P) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyPhoneP2P#39491cc8 as nil")
	}
	b.PutID(PrivacyKeyPhoneP2PTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyPhoneP2P) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyPhoneP2P#39491cc8 to nil")
	}
	if err := b.ConsumeID(PrivacyKeyPhoneP2PTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyPhoneP2P#39491cc8: %w", err)
	}

	return nil
}

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyPhoneP2P) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyPhoneP2P.
var (
	_ bin.Encoder = &PrivacyKeyPhoneP2P{}
	_ bin.Decoder = &PrivacyKeyPhoneP2P{}

	_ PrivacyKeyClass = &PrivacyKeyPhoneP2P{}
)

// PrivacyKeyForwards represents TL type `privacyKeyForwards#69ec56a3`.
// Whether messages forwarded from the user will be anonymously forwarded
//
// See https://core.telegram.org/constructor/privacyKeyForwards for reference.
type PrivacyKeyForwards struct {
}

// PrivacyKeyForwardsTypeID is TL type id of PrivacyKeyForwards.
const PrivacyKeyForwardsTypeID = 0x69ec56a3

// Encode implements bin.Encoder.
func (p *PrivacyKeyForwards) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyForwards#69ec56a3 as nil")
	}
	b.PutID(PrivacyKeyForwardsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyForwards) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyForwards#69ec56a3 to nil")
	}
	if err := b.ConsumeID(PrivacyKeyForwardsTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyForwards#69ec56a3: %w", err)
	}

	return nil
}

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyForwards) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyForwards.
var (
	_ bin.Encoder = &PrivacyKeyForwards{}
	_ bin.Decoder = &PrivacyKeyForwards{}

	_ PrivacyKeyClass = &PrivacyKeyForwards{}
)

// PrivacyKeyProfilePhoto represents TL type `privacyKeyProfilePhoto#96151fed`.
// Whether the profile picture of the user is visible
//
// See https://core.telegram.org/constructor/privacyKeyProfilePhoto for reference.
type PrivacyKeyProfilePhoto struct {
}

// PrivacyKeyProfilePhotoTypeID is TL type id of PrivacyKeyProfilePhoto.
const PrivacyKeyProfilePhotoTypeID = 0x96151fed

// Encode implements bin.Encoder.
func (p *PrivacyKeyProfilePhoto) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyProfilePhoto#96151fed as nil")
	}
	b.PutID(PrivacyKeyProfilePhotoTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyProfilePhoto) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyProfilePhoto#96151fed to nil")
	}
	if err := b.ConsumeID(PrivacyKeyProfilePhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyProfilePhoto#96151fed: %w", err)
	}

	return nil
}

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyProfilePhoto) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyProfilePhoto.
var (
	_ bin.Encoder = &PrivacyKeyProfilePhoto{}
	_ bin.Decoder = &PrivacyKeyProfilePhoto{}

	_ PrivacyKeyClass = &PrivacyKeyProfilePhoto{}
)

// PrivacyKeyPhoneNumber represents TL type `privacyKeyPhoneNumber#d19ae46d`.
// Whether the user allows us to see his phone number
//
// See https://core.telegram.org/constructor/privacyKeyPhoneNumber for reference.
type PrivacyKeyPhoneNumber struct {
}

// PrivacyKeyPhoneNumberTypeID is TL type id of PrivacyKeyPhoneNumber.
const PrivacyKeyPhoneNumberTypeID = 0xd19ae46d

// Encode implements bin.Encoder.
func (p *PrivacyKeyPhoneNumber) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyPhoneNumber#d19ae46d as nil")
	}
	b.PutID(PrivacyKeyPhoneNumberTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyPhoneNumber) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyPhoneNumber#d19ae46d to nil")
	}
	if err := b.ConsumeID(PrivacyKeyPhoneNumberTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyPhoneNumber#d19ae46d: %w", err)
	}

	return nil
}

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyPhoneNumber) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyPhoneNumber.
var (
	_ bin.Encoder = &PrivacyKeyPhoneNumber{}
	_ bin.Decoder = &PrivacyKeyPhoneNumber{}

	_ PrivacyKeyClass = &PrivacyKeyPhoneNumber{}
)

// PrivacyKeyAddedByPhone represents TL type `privacyKeyAddedByPhone#42ffd42b`.
// Whether people can add you to their contact list by your phone number
//
// See https://core.telegram.org/constructor/privacyKeyAddedByPhone for reference.
type PrivacyKeyAddedByPhone struct {
}

// PrivacyKeyAddedByPhoneTypeID is TL type id of PrivacyKeyAddedByPhone.
const PrivacyKeyAddedByPhoneTypeID = 0x42ffd42b

// Encode implements bin.Encoder.
func (p *PrivacyKeyAddedByPhone) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyKeyAddedByPhone#42ffd42b as nil")
	}
	b.PutID(PrivacyKeyAddedByPhoneTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyKeyAddedByPhone) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyKeyAddedByPhone#42ffd42b to nil")
	}
	if err := b.ConsumeID(PrivacyKeyAddedByPhoneTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyKeyAddedByPhone#42ffd42b: %w", err)
	}

	return nil
}

// construct implements constructor of PrivacyKeyClass.
func (p PrivacyKeyAddedByPhone) construct() PrivacyKeyClass { return &p }

// Ensuring interfaces in compile-time for PrivacyKeyAddedByPhone.
var (
	_ bin.Encoder = &PrivacyKeyAddedByPhone{}
	_ bin.Decoder = &PrivacyKeyAddedByPhone{}

	_ PrivacyKeyClass = &PrivacyKeyAddedByPhone{}
)

// PrivacyKeyClass represents PrivacyKey generic type.
//
// See https://core.telegram.org/type/PrivacyKey for reference.
//
// Example:
//  g, err := DecodePrivacyKey(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *PrivacyKeyStatusTimestamp: // privacyKeyStatusTimestamp#bc2eab30
//  case *PrivacyKeyChatInvite: // privacyKeyChatInvite#500e6dfa
//  case *PrivacyKeyPhoneCall: // privacyKeyPhoneCall#3d662b7b
//  case *PrivacyKeyPhoneP2P: // privacyKeyPhoneP2P#39491cc8
//  case *PrivacyKeyForwards: // privacyKeyForwards#69ec56a3
//  case *PrivacyKeyProfilePhoto: // privacyKeyProfilePhoto#96151fed
//  case *PrivacyKeyPhoneNumber: // privacyKeyPhoneNumber#d19ae46d
//  case *PrivacyKeyAddedByPhone: // privacyKeyAddedByPhone#42ffd42b
//  default: panic(v)
//  }
type PrivacyKeyClass interface {
	bin.Encoder
	bin.Decoder
	construct() PrivacyKeyClass
}

// DecodePrivacyKey implements binary de-serialization for PrivacyKeyClass.
func DecodePrivacyKey(buf *bin.Buffer) (PrivacyKeyClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PrivacyKeyStatusTimestampTypeID:
		// Decoding privacyKeyStatusTimestamp#bc2eab30.
		v := PrivacyKeyStatusTimestamp{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	case PrivacyKeyChatInviteTypeID:
		// Decoding privacyKeyChatInvite#500e6dfa.
		v := PrivacyKeyChatInvite{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	case PrivacyKeyPhoneCallTypeID:
		// Decoding privacyKeyPhoneCall#3d662b7b.
		v := PrivacyKeyPhoneCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	case PrivacyKeyPhoneP2PTypeID:
		// Decoding privacyKeyPhoneP2P#39491cc8.
		v := PrivacyKeyPhoneP2P{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	case PrivacyKeyForwardsTypeID:
		// Decoding privacyKeyForwards#69ec56a3.
		v := PrivacyKeyForwards{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	case PrivacyKeyProfilePhotoTypeID:
		// Decoding privacyKeyProfilePhoto#96151fed.
		v := PrivacyKeyProfilePhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	case PrivacyKeyPhoneNumberTypeID:
		// Decoding privacyKeyPhoneNumber#d19ae46d.
		v := PrivacyKeyPhoneNumber{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	case PrivacyKeyAddedByPhoneTypeID:
		// Decoding privacyKeyAddedByPhone#42ffd42b.
		v := PrivacyKeyAddedByPhone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PrivacyKeyClass: %w", bin.NewUnexpectedID(id))
	}
}

// PrivacyKey boxes the PrivacyKeyClass providing a helper.
type PrivacyKeyBox struct {
	PrivacyKey PrivacyKeyClass
}

// Decode implements bin.Decoder for PrivacyKeyBox.
func (b *PrivacyKeyBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PrivacyKeyBox to nil")
	}
	v, err := DecodePrivacyKey(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PrivacyKey = v
	return nil
}

// Encode implements bin.Encode for PrivacyKeyBox.
func (b *PrivacyKeyBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PrivacyKey == nil {
		return fmt.Errorf("unable to encode PrivacyKeyClass as nil")
	}
	return b.PrivacyKey.Encode(buf)
}
