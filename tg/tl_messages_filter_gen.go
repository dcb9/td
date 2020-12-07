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

// InputMessagesFilterEmpty represents TL type `inputMessagesFilterEmpty#57e2f66c`.
// Filter is absent.
//
// See https://core.telegram.org/constructor/inputMessagesFilterEmpty for reference.
type InputMessagesFilterEmpty struct {
}

// InputMessagesFilterEmptyTypeID is TL type id of InputMessagesFilterEmpty.
const InputMessagesFilterEmptyTypeID = 0x57e2f66c

// Encode implements bin.Encoder.
func (i *InputMessagesFilterEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagesFilterEmpty#57e2f66c as nil")
	}
	b.PutID(InputMessagesFilterEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagesFilterEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagesFilterEmpty#57e2f66c to nil")
	}
	if err := b.ConsumeID(InputMessagesFilterEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagesFilterEmpty#57e2f66c: %w", err)
	}

	return nil
}

// construct implements constructor of MessagesFilterClass.
func (i InputMessagesFilterEmpty) construct() MessagesFilterClass { return &i }

// Ensuring interfaces in compile-time for InputMessagesFilterEmpty.
var (
	_ bin.Encoder = &InputMessagesFilterEmpty{}
	_ bin.Decoder = &InputMessagesFilterEmpty{}

	_ MessagesFilterClass = &InputMessagesFilterEmpty{}
)

// InputMessagesFilterPhotos represents TL type `inputMessagesFilterPhotos#9609a51c`.
// Filter for messages containing photos.
//
// See https://core.telegram.org/constructor/inputMessagesFilterPhotos for reference.
type InputMessagesFilterPhotos struct {
}

// InputMessagesFilterPhotosTypeID is TL type id of InputMessagesFilterPhotos.
const InputMessagesFilterPhotosTypeID = 0x9609a51c

// Encode implements bin.Encoder.
func (i *InputMessagesFilterPhotos) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagesFilterPhotos#9609a51c as nil")
	}
	b.PutID(InputMessagesFilterPhotosTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagesFilterPhotos) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagesFilterPhotos#9609a51c to nil")
	}
	if err := b.ConsumeID(InputMessagesFilterPhotosTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagesFilterPhotos#9609a51c: %w", err)
	}

	return nil
}

// construct implements constructor of MessagesFilterClass.
func (i InputMessagesFilterPhotos) construct() MessagesFilterClass { return &i }

// Ensuring interfaces in compile-time for InputMessagesFilterPhotos.
var (
	_ bin.Encoder = &InputMessagesFilterPhotos{}
	_ bin.Decoder = &InputMessagesFilterPhotos{}

	_ MessagesFilterClass = &InputMessagesFilterPhotos{}
)

// InputMessagesFilterVideo represents TL type `inputMessagesFilterVideo#9fc00e65`.
// Filter for messages containing videos.
//
// See https://core.telegram.org/constructor/inputMessagesFilterVideo for reference.
type InputMessagesFilterVideo struct {
}

// InputMessagesFilterVideoTypeID is TL type id of InputMessagesFilterVideo.
const InputMessagesFilterVideoTypeID = 0x9fc00e65

// Encode implements bin.Encoder.
func (i *InputMessagesFilterVideo) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagesFilterVideo#9fc00e65 as nil")
	}
	b.PutID(InputMessagesFilterVideoTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagesFilterVideo) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagesFilterVideo#9fc00e65 to nil")
	}
	if err := b.ConsumeID(InputMessagesFilterVideoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagesFilterVideo#9fc00e65: %w", err)
	}

	return nil
}

// construct implements constructor of MessagesFilterClass.
func (i InputMessagesFilterVideo) construct() MessagesFilterClass { return &i }

// Ensuring interfaces in compile-time for InputMessagesFilterVideo.
var (
	_ bin.Encoder = &InputMessagesFilterVideo{}
	_ bin.Decoder = &InputMessagesFilterVideo{}

	_ MessagesFilterClass = &InputMessagesFilterVideo{}
)

// InputMessagesFilterPhotoVideo represents TL type `inputMessagesFilterPhotoVideo#56e9f0e4`.
// Filter for messages containing photos or videos.
//
// See https://core.telegram.org/constructor/inputMessagesFilterPhotoVideo for reference.
type InputMessagesFilterPhotoVideo struct {
}

// InputMessagesFilterPhotoVideoTypeID is TL type id of InputMessagesFilterPhotoVideo.
const InputMessagesFilterPhotoVideoTypeID = 0x56e9f0e4

// Encode implements bin.Encoder.
func (i *InputMessagesFilterPhotoVideo) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagesFilterPhotoVideo#56e9f0e4 as nil")
	}
	b.PutID(InputMessagesFilterPhotoVideoTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagesFilterPhotoVideo) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagesFilterPhotoVideo#56e9f0e4 to nil")
	}
	if err := b.ConsumeID(InputMessagesFilterPhotoVideoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagesFilterPhotoVideo#56e9f0e4: %w", err)
	}

	return nil
}

// construct implements constructor of MessagesFilterClass.
func (i InputMessagesFilterPhotoVideo) construct() MessagesFilterClass { return &i }

// Ensuring interfaces in compile-time for InputMessagesFilterPhotoVideo.
var (
	_ bin.Encoder = &InputMessagesFilterPhotoVideo{}
	_ bin.Decoder = &InputMessagesFilterPhotoVideo{}

	_ MessagesFilterClass = &InputMessagesFilterPhotoVideo{}
)

// InputMessagesFilterDocument represents TL type `inputMessagesFilterDocument#9eddf188`.
// Filter for messages containing documents.
//
// See https://core.telegram.org/constructor/inputMessagesFilterDocument for reference.
type InputMessagesFilterDocument struct {
}

// InputMessagesFilterDocumentTypeID is TL type id of InputMessagesFilterDocument.
const InputMessagesFilterDocumentTypeID = 0x9eddf188

// Encode implements bin.Encoder.
func (i *InputMessagesFilterDocument) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagesFilterDocument#9eddf188 as nil")
	}
	b.PutID(InputMessagesFilterDocumentTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagesFilterDocument) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagesFilterDocument#9eddf188 to nil")
	}
	if err := b.ConsumeID(InputMessagesFilterDocumentTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagesFilterDocument#9eddf188: %w", err)
	}

	return nil
}

// construct implements constructor of MessagesFilterClass.
func (i InputMessagesFilterDocument) construct() MessagesFilterClass { return &i }

// Ensuring interfaces in compile-time for InputMessagesFilterDocument.
var (
	_ bin.Encoder = &InputMessagesFilterDocument{}
	_ bin.Decoder = &InputMessagesFilterDocument{}

	_ MessagesFilterClass = &InputMessagesFilterDocument{}
)

// InputMessagesFilterUrl represents TL type `inputMessagesFilterUrl#7ef0dd87`.
// Return only messages containing URLs
//
// See https://core.telegram.org/constructor/inputMessagesFilterUrl for reference.
type InputMessagesFilterUrl struct {
}

// InputMessagesFilterUrlTypeID is TL type id of InputMessagesFilterUrl.
const InputMessagesFilterUrlTypeID = 0x7ef0dd87

// Encode implements bin.Encoder.
func (i *InputMessagesFilterUrl) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagesFilterUrl#7ef0dd87 as nil")
	}
	b.PutID(InputMessagesFilterUrlTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagesFilterUrl) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagesFilterUrl#7ef0dd87 to nil")
	}
	if err := b.ConsumeID(InputMessagesFilterUrlTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagesFilterUrl#7ef0dd87: %w", err)
	}

	return nil
}

// construct implements constructor of MessagesFilterClass.
func (i InputMessagesFilterUrl) construct() MessagesFilterClass { return &i }

// Ensuring interfaces in compile-time for InputMessagesFilterUrl.
var (
	_ bin.Encoder = &InputMessagesFilterUrl{}
	_ bin.Decoder = &InputMessagesFilterUrl{}

	_ MessagesFilterClass = &InputMessagesFilterUrl{}
)

// InputMessagesFilterGif represents TL type `inputMessagesFilterGif#ffc86587`.
// Return only messages containing gifs
//
// See https://core.telegram.org/constructor/inputMessagesFilterGif for reference.
type InputMessagesFilterGif struct {
}

// InputMessagesFilterGifTypeID is TL type id of InputMessagesFilterGif.
const InputMessagesFilterGifTypeID = 0xffc86587

// Encode implements bin.Encoder.
func (i *InputMessagesFilterGif) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagesFilterGif#ffc86587 as nil")
	}
	b.PutID(InputMessagesFilterGifTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagesFilterGif) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagesFilterGif#ffc86587 to nil")
	}
	if err := b.ConsumeID(InputMessagesFilterGifTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagesFilterGif#ffc86587: %w", err)
	}

	return nil
}

// construct implements constructor of MessagesFilterClass.
func (i InputMessagesFilterGif) construct() MessagesFilterClass { return &i }

// Ensuring interfaces in compile-time for InputMessagesFilterGif.
var (
	_ bin.Encoder = &InputMessagesFilterGif{}
	_ bin.Decoder = &InputMessagesFilterGif{}

	_ MessagesFilterClass = &InputMessagesFilterGif{}
)

// InputMessagesFilterVoice represents TL type `inputMessagesFilterVoice#50f5c392`.
// Return only messages containing voice notes
//
// See https://core.telegram.org/constructor/inputMessagesFilterVoice for reference.
type InputMessagesFilterVoice struct {
}

// InputMessagesFilterVoiceTypeID is TL type id of InputMessagesFilterVoice.
const InputMessagesFilterVoiceTypeID = 0x50f5c392

// Encode implements bin.Encoder.
func (i *InputMessagesFilterVoice) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagesFilterVoice#50f5c392 as nil")
	}
	b.PutID(InputMessagesFilterVoiceTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagesFilterVoice) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagesFilterVoice#50f5c392 to nil")
	}
	if err := b.ConsumeID(InputMessagesFilterVoiceTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagesFilterVoice#50f5c392: %w", err)
	}

	return nil
}

// construct implements constructor of MessagesFilterClass.
func (i InputMessagesFilterVoice) construct() MessagesFilterClass { return &i }

// Ensuring interfaces in compile-time for InputMessagesFilterVoice.
var (
	_ bin.Encoder = &InputMessagesFilterVoice{}
	_ bin.Decoder = &InputMessagesFilterVoice{}

	_ MessagesFilterClass = &InputMessagesFilterVoice{}
)

// InputMessagesFilterMusic represents TL type `inputMessagesFilterMusic#3751b49e`.
// Return only messages containing audio files
//
// See https://core.telegram.org/constructor/inputMessagesFilterMusic for reference.
type InputMessagesFilterMusic struct {
}

// InputMessagesFilterMusicTypeID is TL type id of InputMessagesFilterMusic.
const InputMessagesFilterMusicTypeID = 0x3751b49e

// Encode implements bin.Encoder.
func (i *InputMessagesFilterMusic) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagesFilterMusic#3751b49e as nil")
	}
	b.PutID(InputMessagesFilterMusicTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagesFilterMusic) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagesFilterMusic#3751b49e to nil")
	}
	if err := b.ConsumeID(InputMessagesFilterMusicTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagesFilterMusic#3751b49e: %w", err)
	}

	return nil
}

// construct implements constructor of MessagesFilterClass.
func (i InputMessagesFilterMusic) construct() MessagesFilterClass { return &i }

// Ensuring interfaces in compile-time for InputMessagesFilterMusic.
var (
	_ bin.Encoder = &InputMessagesFilterMusic{}
	_ bin.Decoder = &InputMessagesFilterMusic{}

	_ MessagesFilterClass = &InputMessagesFilterMusic{}
)

// InputMessagesFilterChatPhotos represents TL type `inputMessagesFilterChatPhotos#3a20ecb8`.
// Return only chat photo changes
//
// See https://core.telegram.org/constructor/inputMessagesFilterChatPhotos for reference.
type InputMessagesFilterChatPhotos struct {
}

// InputMessagesFilterChatPhotosTypeID is TL type id of InputMessagesFilterChatPhotos.
const InputMessagesFilterChatPhotosTypeID = 0x3a20ecb8

// Encode implements bin.Encoder.
func (i *InputMessagesFilterChatPhotos) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagesFilterChatPhotos#3a20ecb8 as nil")
	}
	b.PutID(InputMessagesFilterChatPhotosTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagesFilterChatPhotos) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagesFilterChatPhotos#3a20ecb8 to nil")
	}
	if err := b.ConsumeID(InputMessagesFilterChatPhotosTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagesFilterChatPhotos#3a20ecb8: %w", err)
	}

	return nil
}

// construct implements constructor of MessagesFilterClass.
func (i InputMessagesFilterChatPhotos) construct() MessagesFilterClass { return &i }

// Ensuring interfaces in compile-time for InputMessagesFilterChatPhotos.
var (
	_ bin.Encoder = &InputMessagesFilterChatPhotos{}
	_ bin.Decoder = &InputMessagesFilterChatPhotos{}

	_ MessagesFilterClass = &InputMessagesFilterChatPhotos{}
)

// InputMessagesFilterPhoneCalls represents TL type `inputMessagesFilterPhoneCalls#80c99768`.
// Return only phone calls
//
// See https://core.telegram.org/constructor/inputMessagesFilterPhoneCalls for reference.
type InputMessagesFilterPhoneCalls struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Return only missed phone calls
	Missed bool
}

// InputMessagesFilterPhoneCallsTypeID is TL type id of InputMessagesFilterPhoneCalls.
const InputMessagesFilterPhoneCallsTypeID = 0x80c99768

// Encode implements bin.Encoder.
func (i *InputMessagesFilterPhoneCalls) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagesFilterPhoneCalls#80c99768 as nil")
	}
	b.PutID(InputMessagesFilterPhoneCallsTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputMessagesFilterPhoneCalls#80c99768: field flags: %w", err)
	}
	return nil
}

// SetMissed sets value of Missed conditional field.
func (i *InputMessagesFilterPhoneCalls) SetMissed(value bool) {
	if value {
		i.Flags.Set(0)
	} else {
		i.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (i *InputMessagesFilterPhoneCalls) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagesFilterPhoneCalls#80c99768 to nil")
	}
	if err := b.ConsumeID(InputMessagesFilterPhoneCallsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagesFilterPhoneCalls#80c99768: %w", err)
	}

	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputMessagesFilterPhoneCalls#80c99768: field flags: %w", err)
		}
	}
	i.Missed = i.Flags.Has(0)
	return nil
}

// construct implements constructor of MessagesFilterClass.
func (i InputMessagesFilterPhoneCalls) construct() MessagesFilterClass { return &i }

// Ensuring interfaces in compile-time for InputMessagesFilterPhoneCalls.
var (
	_ bin.Encoder = &InputMessagesFilterPhoneCalls{}
	_ bin.Decoder = &InputMessagesFilterPhoneCalls{}

	_ MessagesFilterClass = &InputMessagesFilterPhoneCalls{}
)

// InputMessagesFilterRoundVoice represents TL type `inputMessagesFilterRoundVoice#7a7c17a4`.
// Return only round videos and voice notes
//
// See https://core.telegram.org/constructor/inputMessagesFilterRoundVoice for reference.
type InputMessagesFilterRoundVoice struct {
}

// InputMessagesFilterRoundVoiceTypeID is TL type id of InputMessagesFilterRoundVoice.
const InputMessagesFilterRoundVoiceTypeID = 0x7a7c17a4

// Encode implements bin.Encoder.
func (i *InputMessagesFilterRoundVoice) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagesFilterRoundVoice#7a7c17a4 as nil")
	}
	b.PutID(InputMessagesFilterRoundVoiceTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagesFilterRoundVoice) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagesFilterRoundVoice#7a7c17a4 to nil")
	}
	if err := b.ConsumeID(InputMessagesFilterRoundVoiceTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagesFilterRoundVoice#7a7c17a4: %w", err)
	}

	return nil
}

// construct implements constructor of MessagesFilterClass.
func (i InputMessagesFilterRoundVoice) construct() MessagesFilterClass { return &i }

// Ensuring interfaces in compile-time for InputMessagesFilterRoundVoice.
var (
	_ bin.Encoder = &InputMessagesFilterRoundVoice{}
	_ bin.Decoder = &InputMessagesFilterRoundVoice{}

	_ MessagesFilterClass = &InputMessagesFilterRoundVoice{}
)

// InputMessagesFilterRoundVideo represents TL type `inputMessagesFilterRoundVideo#b549da53`.
// Return only round videos
//
// See https://core.telegram.org/constructor/inputMessagesFilterRoundVideo for reference.
type InputMessagesFilterRoundVideo struct {
}

// InputMessagesFilterRoundVideoTypeID is TL type id of InputMessagesFilterRoundVideo.
const InputMessagesFilterRoundVideoTypeID = 0xb549da53

// Encode implements bin.Encoder.
func (i *InputMessagesFilterRoundVideo) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagesFilterRoundVideo#b549da53 as nil")
	}
	b.PutID(InputMessagesFilterRoundVideoTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagesFilterRoundVideo) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagesFilterRoundVideo#b549da53 to nil")
	}
	if err := b.ConsumeID(InputMessagesFilterRoundVideoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagesFilterRoundVideo#b549da53: %w", err)
	}

	return nil
}

// construct implements constructor of MessagesFilterClass.
func (i InputMessagesFilterRoundVideo) construct() MessagesFilterClass { return &i }

// Ensuring interfaces in compile-time for InputMessagesFilterRoundVideo.
var (
	_ bin.Encoder = &InputMessagesFilterRoundVideo{}
	_ bin.Decoder = &InputMessagesFilterRoundVideo{}

	_ MessagesFilterClass = &InputMessagesFilterRoundVideo{}
)

// InputMessagesFilterMyMentions represents TL type `inputMessagesFilterMyMentions#c1f8e69a`.
// Return only messages where the current user was mentioned.
//
// See https://core.telegram.org/constructor/inputMessagesFilterMyMentions for reference.
type InputMessagesFilterMyMentions struct {
}

// InputMessagesFilterMyMentionsTypeID is TL type id of InputMessagesFilterMyMentions.
const InputMessagesFilterMyMentionsTypeID = 0xc1f8e69a

// Encode implements bin.Encoder.
func (i *InputMessagesFilterMyMentions) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagesFilterMyMentions#c1f8e69a as nil")
	}
	b.PutID(InputMessagesFilterMyMentionsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagesFilterMyMentions) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagesFilterMyMentions#c1f8e69a to nil")
	}
	if err := b.ConsumeID(InputMessagesFilterMyMentionsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagesFilterMyMentions#c1f8e69a: %w", err)
	}

	return nil
}

// construct implements constructor of MessagesFilterClass.
func (i InputMessagesFilterMyMentions) construct() MessagesFilterClass { return &i }

// Ensuring interfaces in compile-time for InputMessagesFilterMyMentions.
var (
	_ bin.Encoder = &InputMessagesFilterMyMentions{}
	_ bin.Decoder = &InputMessagesFilterMyMentions{}

	_ MessagesFilterClass = &InputMessagesFilterMyMentions{}
)

// InputMessagesFilterGeo represents TL type `inputMessagesFilterGeo#e7026d0d`.
// Return only messages containing geolocations
//
// See https://core.telegram.org/constructor/inputMessagesFilterGeo for reference.
type InputMessagesFilterGeo struct {
}

// InputMessagesFilterGeoTypeID is TL type id of InputMessagesFilterGeo.
const InputMessagesFilterGeoTypeID = 0xe7026d0d

// Encode implements bin.Encoder.
func (i *InputMessagesFilterGeo) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagesFilterGeo#e7026d0d as nil")
	}
	b.PutID(InputMessagesFilterGeoTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagesFilterGeo) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagesFilterGeo#e7026d0d to nil")
	}
	if err := b.ConsumeID(InputMessagesFilterGeoTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagesFilterGeo#e7026d0d: %w", err)
	}

	return nil
}

// construct implements constructor of MessagesFilterClass.
func (i InputMessagesFilterGeo) construct() MessagesFilterClass { return &i }

// Ensuring interfaces in compile-time for InputMessagesFilterGeo.
var (
	_ bin.Encoder = &InputMessagesFilterGeo{}
	_ bin.Decoder = &InputMessagesFilterGeo{}

	_ MessagesFilterClass = &InputMessagesFilterGeo{}
)

// InputMessagesFilterContacts represents TL type `inputMessagesFilterContacts#e062db83`.
// Return only messages containing contacts
//
// See https://core.telegram.org/constructor/inputMessagesFilterContacts for reference.
type InputMessagesFilterContacts struct {
}

// InputMessagesFilterContactsTypeID is TL type id of InputMessagesFilterContacts.
const InputMessagesFilterContactsTypeID = 0xe062db83

// Encode implements bin.Encoder.
func (i *InputMessagesFilterContacts) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagesFilterContacts#e062db83 as nil")
	}
	b.PutID(InputMessagesFilterContactsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagesFilterContacts) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagesFilterContacts#e062db83 to nil")
	}
	if err := b.ConsumeID(InputMessagesFilterContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagesFilterContacts#e062db83: %w", err)
	}

	return nil
}

// construct implements constructor of MessagesFilterClass.
func (i InputMessagesFilterContacts) construct() MessagesFilterClass { return &i }

// Ensuring interfaces in compile-time for InputMessagesFilterContacts.
var (
	_ bin.Encoder = &InputMessagesFilterContacts{}
	_ bin.Decoder = &InputMessagesFilterContacts{}

	_ MessagesFilterClass = &InputMessagesFilterContacts{}
)

// InputMessagesFilterPinned represents TL type `inputMessagesFilterPinned#1bb00451`.
// Fetch only pinned messages
//
// See https://core.telegram.org/constructor/inputMessagesFilterPinned for reference.
type InputMessagesFilterPinned struct {
}

// InputMessagesFilterPinnedTypeID is TL type id of InputMessagesFilterPinned.
const InputMessagesFilterPinnedTypeID = 0x1bb00451

// Encode implements bin.Encoder.
func (i *InputMessagesFilterPinned) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessagesFilterPinned#1bb00451 as nil")
	}
	b.PutID(InputMessagesFilterPinnedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessagesFilterPinned) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessagesFilterPinned#1bb00451 to nil")
	}
	if err := b.ConsumeID(InputMessagesFilterPinnedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessagesFilterPinned#1bb00451: %w", err)
	}

	return nil
}

// construct implements constructor of MessagesFilterClass.
func (i InputMessagesFilterPinned) construct() MessagesFilterClass { return &i }

// Ensuring interfaces in compile-time for InputMessagesFilterPinned.
var (
	_ bin.Encoder = &InputMessagesFilterPinned{}
	_ bin.Decoder = &InputMessagesFilterPinned{}

	_ MessagesFilterClass = &InputMessagesFilterPinned{}
)

// MessagesFilterClass represents MessagesFilter generic type.
//
// See https://core.telegram.org/type/MessagesFilter for reference.
//
// Example:
//  g, err := DecodeMessagesFilter(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputMessagesFilterEmpty: // inputMessagesFilterEmpty#57e2f66c
//  case *InputMessagesFilterPhotos: // inputMessagesFilterPhotos#9609a51c
//  case *InputMessagesFilterVideo: // inputMessagesFilterVideo#9fc00e65
//  case *InputMessagesFilterPhotoVideo: // inputMessagesFilterPhotoVideo#56e9f0e4
//  case *InputMessagesFilterDocument: // inputMessagesFilterDocument#9eddf188
//  case *InputMessagesFilterUrl: // inputMessagesFilterUrl#7ef0dd87
//  case *InputMessagesFilterGif: // inputMessagesFilterGif#ffc86587
//  case *InputMessagesFilterVoice: // inputMessagesFilterVoice#50f5c392
//  case *InputMessagesFilterMusic: // inputMessagesFilterMusic#3751b49e
//  case *InputMessagesFilterChatPhotos: // inputMessagesFilterChatPhotos#3a20ecb8
//  case *InputMessagesFilterPhoneCalls: // inputMessagesFilterPhoneCalls#80c99768
//  case *InputMessagesFilterRoundVoice: // inputMessagesFilterRoundVoice#7a7c17a4
//  case *InputMessagesFilterRoundVideo: // inputMessagesFilterRoundVideo#b549da53
//  case *InputMessagesFilterMyMentions: // inputMessagesFilterMyMentions#c1f8e69a
//  case *InputMessagesFilterGeo: // inputMessagesFilterGeo#e7026d0d
//  case *InputMessagesFilterContacts: // inputMessagesFilterContacts#e062db83
//  case *InputMessagesFilterPinned: // inputMessagesFilterPinned#1bb00451
//  default: panic(v)
//  }
type MessagesFilterClass interface {
	bin.Encoder
	bin.Decoder
	construct() MessagesFilterClass
}

// DecodeMessagesFilter implements binary de-serialization for MessagesFilterClass.
func DecodeMessagesFilter(buf *bin.Buffer) (MessagesFilterClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputMessagesFilterEmptyTypeID:
		// Decoding inputMessagesFilterEmpty#57e2f66c.
		v := InputMessagesFilterEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", err)
		}
		return &v, nil
	case InputMessagesFilterPhotosTypeID:
		// Decoding inputMessagesFilterPhotos#9609a51c.
		v := InputMessagesFilterPhotos{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", err)
		}
		return &v, nil
	case InputMessagesFilterVideoTypeID:
		// Decoding inputMessagesFilterVideo#9fc00e65.
		v := InputMessagesFilterVideo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", err)
		}
		return &v, nil
	case InputMessagesFilterPhotoVideoTypeID:
		// Decoding inputMessagesFilterPhotoVideo#56e9f0e4.
		v := InputMessagesFilterPhotoVideo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", err)
		}
		return &v, nil
	case InputMessagesFilterDocumentTypeID:
		// Decoding inputMessagesFilterDocument#9eddf188.
		v := InputMessagesFilterDocument{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", err)
		}
		return &v, nil
	case InputMessagesFilterUrlTypeID:
		// Decoding inputMessagesFilterUrl#7ef0dd87.
		v := InputMessagesFilterUrl{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", err)
		}
		return &v, nil
	case InputMessagesFilterGifTypeID:
		// Decoding inputMessagesFilterGif#ffc86587.
		v := InputMessagesFilterGif{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", err)
		}
		return &v, nil
	case InputMessagesFilterVoiceTypeID:
		// Decoding inputMessagesFilterVoice#50f5c392.
		v := InputMessagesFilterVoice{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", err)
		}
		return &v, nil
	case InputMessagesFilterMusicTypeID:
		// Decoding inputMessagesFilterMusic#3751b49e.
		v := InputMessagesFilterMusic{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", err)
		}
		return &v, nil
	case InputMessagesFilterChatPhotosTypeID:
		// Decoding inputMessagesFilterChatPhotos#3a20ecb8.
		v := InputMessagesFilterChatPhotos{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", err)
		}
		return &v, nil
	case InputMessagesFilterPhoneCallsTypeID:
		// Decoding inputMessagesFilterPhoneCalls#80c99768.
		v := InputMessagesFilterPhoneCalls{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", err)
		}
		return &v, nil
	case InputMessagesFilterRoundVoiceTypeID:
		// Decoding inputMessagesFilterRoundVoice#7a7c17a4.
		v := InputMessagesFilterRoundVoice{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", err)
		}
		return &v, nil
	case InputMessagesFilterRoundVideoTypeID:
		// Decoding inputMessagesFilterRoundVideo#b549da53.
		v := InputMessagesFilterRoundVideo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", err)
		}
		return &v, nil
	case InputMessagesFilterMyMentionsTypeID:
		// Decoding inputMessagesFilterMyMentions#c1f8e69a.
		v := InputMessagesFilterMyMentions{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", err)
		}
		return &v, nil
	case InputMessagesFilterGeoTypeID:
		// Decoding inputMessagesFilterGeo#e7026d0d.
		v := InputMessagesFilterGeo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", err)
		}
		return &v, nil
	case InputMessagesFilterContactsTypeID:
		// Decoding inputMessagesFilterContacts#e062db83.
		v := InputMessagesFilterContacts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", err)
		}
		return &v, nil
	case InputMessagesFilterPinnedTypeID:
		// Decoding inputMessagesFilterPinned#1bb00451.
		v := InputMessagesFilterPinned{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesFilterClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesFilter boxes the MessagesFilterClass providing a helper.
type MessagesFilterBox struct {
	MessagesFilter MessagesFilterClass
}

// Decode implements bin.Decoder for MessagesFilterBox.
func (b *MessagesFilterBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesFilterBox to nil")
	}
	v, err := DecodeMessagesFilter(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessagesFilter = v
	return nil
}

// Encode implements bin.Encode for MessagesFilterBox.
func (b *MessagesFilterBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.MessagesFilter == nil {
		return fmt.Errorf("unable to encode MessagesFilterClass as nil")
	}
	return b.MessagesFilter.Encode(buf)
}
