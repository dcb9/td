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
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// PhotosUploadProfilePhotoRequest represents TL type `photos.uploadProfilePhoto#93c9a51`.
// Updates current user profile photo.
//
// See https://core.telegram.org/method/photos.uploadProfilePhoto for reference.
type PhotosUploadProfilePhotoRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Fallback field of PhotosUploadProfilePhotoRequest.
	Fallback bool
	// File saved in parts by means of upload.saveFilePart¹ method
	//
	// Links:
	//  1) https://core.telegram.org/method/upload.saveFilePart
	//
	// Use SetFile and GetFile helpers.
	File InputFileClass
	// Animated profile picture¹ video
	//
	// Links:
	//  1) https://core.telegram.org/api/files#animated-profile-pictures
	//
	// Use SetVideo and GetVideo helpers.
	Video InputFileClass
	// Floating point UNIX timestamp in seconds, indicating the frame of the video that
	// should be used as static preview.
	//
	// Use SetVideoStartTs and GetVideoStartTs helpers.
	VideoStartTs float64
	// VideoEmojiMarkup field of PhotosUploadProfilePhotoRequest.
	//
	// Use SetVideoEmojiMarkup and GetVideoEmojiMarkup helpers.
	VideoEmojiMarkup VideoSizeClass
}

// PhotosUploadProfilePhotoRequestTypeID is TL type id of PhotosUploadProfilePhotoRequest.
const PhotosUploadProfilePhotoRequestTypeID = 0x93c9a51

// Ensuring interfaces in compile-time for PhotosUploadProfilePhotoRequest.
var (
	_ bin.Encoder     = &PhotosUploadProfilePhotoRequest{}
	_ bin.Decoder     = &PhotosUploadProfilePhotoRequest{}
	_ bin.BareEncoder = &PhotosUploadProfilePhotoRequest{}
	_ bin.BareDecoder = &PhotosUploadProfilePhotoRequest{}
)

func (u *PhotosUploadProfilePhotoRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.Fallback == false) {
		return false
	}
	if !(u.File == nil) {
		return false
	}
	if !(u.Video == nil) {
		return false
	}
	if !(u.VideoStartTs == 0) {
		return false
	}
	if !(u.VideoEmojiMarkup == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *PhotosUploadProfilePhotoRequest) String() string {
	if u == nil {
		return "PhotosUploadProfilePhotoRequest(nil)"
	}
	type Alias PhotosUploadProfilePhotoRequest
	return fmt.Sprintf("PhotosUploadProfilePhotoRequest%+v", Alias(*u))
}

// FillFrom fills PhotosUploadProfilePhotoRequest from given interface.
func (u *PhotosUploadProfilePhotoRequest) FillFrom(from interface {
	GetFallback() (value bool)
	GetFile() (value InputFileClass, ok bool)
	GetVideo() (value InputFileClass, ok bool)
	GetVideoStartTs() (value float64, ok bool)
	GetVideoEmojiMarkup() (value VideoSizeClass, ok bool)
}) {
	u.Fallback = from.GetFallback()
	if val, ok := from.GetFile(); ok {
		u.File = val
	}

	if val, ok := from.GetVideo(); ok {
		u.Video = val
	}

	if val, ok := from.GetVideoStartTs(); ok {
		u.VideoStartTs = val
	}

	if val, ok := from.GetVideoEmojiMarkup(); ok {
		u.VideoEmojiMarkup = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotosUploadProfilePhotoRequest) TypeID() uint32 {
	return PhotosUploadProfilePhotoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotosUploadProfilePhotoRequest) TypeName() string {
	return "photos.uploadProfilePhoto"
}

// TypeInfo returns info about TL type.
func (u *PhotosUploadProfilePhotoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photos.uploadProfilePhoto",
		ID:   PhotosUploadProfilePhotoRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Fallback",
			SchemaName: "fallback",
			Null:       !u.Flags.Has(3),
		},
		{
			Name:       "File",
			SchemaName: "file",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "Video",
			SchemaName: "video",
			Null:       !u.Flags.Has(1),
		},
		{
			Name:       "VideoStartTs",
			SchemaName: "video_start_ts",
			Null:       !u.Flags.Has(2),
		},
		{
			Name:       "VideoEmojiMarkup",
			SchemaName: "video_emoji_markup",
			Null:       !u.Flags.Has(4),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *PhotosUploadProfilePhotoRequest) SetFlags() {
	if !(u.Fallback == false) {
		u.Flags.Set(3)
	}
	if !(u.File == nil) {
		u.Flags.Set(0)
	}
	if !(u.Video == nil) {
		u.Flags.Set(1)
	}
	if !(u.VideoStartTs == 0) {
		u.Flags.Set(2)
	}
	if !(u.VideoEmojiMarkup == nil) {
		u.Flags.Set(4)
	}
}

// Encode implements bin.Encoder.
func (u *PhotosUploadProfilePhotoRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode photos.uploadProfilePhoto#93c9a51 as nil")
	}
	b.PutID(PhotosUploadProfilePhotoRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *PhotosUploadProfilePhotoRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode photos.uploadProfilePhoto#93c9a51 as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photos.uploadProfilePhoto#93c9a51: field flags: %w", err)
	}
	if u.Flags.Has(0) {
		if u.File == nil {
			return fmt.Errorf("unable to encode photos.uploadProfilePhoto#93c9a51: field file is nil")
		}
		if err := u.File.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.uploadProfilePhoto#93c9a51: field file: %w", err)
		}
	}
	if u.Flags.Has(1) {
		if u.Video == nil {
			return fmt.Errorf("unable to encode photos.uploadProfilePhoto#93c9a51: field video is nil")
		}
		if err := u.Video.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.uploadProfilePhoto#93c9a51: field video: %w", err)
		}
	}
	if u.Flags.Has(2) {
		b.PutDouble(u.VideoStartTs)
	}
	if u.Flags.Has(4) {
		if u.VideoEmojiMarkup == nil {
			return fmt.Errorf("unable to encode photos.uploadProfilePhoto#93c9a51: field video_emoji_markup is nil")
		}
		if err := u.VideoEmojiMarkup.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.uploadProfilePhoto#93c9a51: field video_emoji_markup: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *PhotosUploadProfilePhotoRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode photos.uploadProfilePhoto#93c9a51 to nil")
	}
	if err := b.ConsumeID(PhotosUploadProfilePhotoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode photos.uploadProfilePhoto#93c9a51: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *PhotosUploadProfilePhotoRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode photos.uploadProfilePhoto#93c9a51 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode photos.uploadProfilePhoto#93c9a51: field flags: %w", err)
		}
	}
	u.Fallback = u.Flags.Has(3)
	if u.Flags.Has(0) {
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode photos.uploadProfilePhoto#93c9a51: field file: %w", err)
		}
		u.File = value
	}
	if u.Flags.Has(1) {
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode photos.uploadProfilePhoto#93c9a51: field video: %w", err)
		}
		u.Video = value
	}
	if u.Flags.Has(2) {
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode photos.uploadProfilePhoto#93c9a51: field video_start_ts: %w", err)
		}
		u.VideoStartTs = value
	}
	if u.Flags.Has(4) {
		value, err := DecodeVideoSize(b)
		if err != nil {
			return fmt.Errorf("unable to decode photos.uploadProfilePhoto#93c9a51: field video_emoji_markup: %w", err)
		}
		u.VideoEmojiMarkup = value
	}
	return nil
}

// SetFallback sets value of Fallback conditional field.
func (u *PhotosUploadProfilePhotoRequest) SetFallback(value bool) {
	if value {
		u.Flags.Set(3)
		u.Fallback = true
	} else {
		u.Flags.Unset(3)
		u.Fallback = false
	}
}

// GetFallback returns value of Fallback conditional field.
func (u *PhotosUploadProfilePhotoRequest) GetFallback() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(3)
}

// SetFile sets value of File conditional field.
func (u *PhotosUploadProfilePhotoRequest) SetFile(value InputFileClass) {
	u.Flags.Set(0)
	u.File = value
}

// GetFile returns value of File conditional field and
// boolean which is true if field was set.
func (u *PhotosUploadProfilePhotoRequest) GetFile() (value InputFileClass, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.File, true
}

// SetVideo sets value of Video conditional field.
func (u *PhotosUploadProfilePhotoRequest) SetVideo(value InputFileClass) {
	u.Flags.Set(1)
	u.Video = value
}

// GetVideo returns value of Video conditional field and
// boolean which is true if field was set.
func (u *PhotosUploadProfilePhotoRequest) GetVideo() (value InputFileClass, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(1) {
		return value, false
	}
	return u.Video, true
}

// SetVideoStartTs sets value of VideoStartTs conditional field.
func (u *PhotosUploadProfilePhotoRequest) SetVideoStartTs(value float64) {
	u.Flags.Set(2)
	u.VideoStartTs = value
}

// GetVideoStartTs returns value of VideoStartTs conditional field and
// boolean which is true if field was set.
func (u *PhotosUploadProfilePhotoRequest) GetVideoStartTs() (value float64, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(2) {
		return value, false
	}
	return u.VideoStartTs, true
}

// SetVideoEmojiMarkup sets value of VideoEmojiMarkup conditional field.
func (u *PhotosUploadProfilePhotoRequest) SetVideoEmojiMarkup(value VideoSizeClass) {
	u.Flags.Set(4)
	u.VideoEmojiMarkup = value
}

// GetVideoEmojiMarkup returns value of VideoEmojiMarkup conditional field and
// boolean which is true if field was set.
func (u *PhotosUploadProfilePhotoRequest) GetVideoEmojiMarkup() (value VideoSizeClass, ok bool) {
	if u == nil {
		return
	}
	if !u.Flags.Has(4) {
		return value, false
	}
	return u.VideoEmojiMarkup, true
}

// PhotosUploadProfilePhoto invokes method photos.uploadProfilePhoto#93c9a51 returning error if any.
// Updates current user profile photo.
//
// Possible errors:
//
//	400 ALBUM_PHOTOS_TOO_MANY: You have uploaded too many profile photos, delete some before retrying.
//	400 FILE_PARTS_INVALID: The number of file parts is invalid.
//	400 IMAGE_PROCESS_FAILED: Failure while processing image.
//	400 PHOTO_CROP_FILE_MISSING: Photo crop file missing.
//	400 PHOTO_CROP_SIZE_SMALL: Photo is too small.
//	400 PHOTO_EXT_INVALID: The extension of the photo is invalid.
//	400 PHOTO_FILE_MISSING: Profile photo file missing.
//	400 PHOTO_INVALID: Photo invalid.
//	400 STICKER_MIME_INVALID:
//	400 VIDEO_FILE_INVALID: The specified video file is invalid.
//
// See https://core.telegram.org/method/photos.uploadProfilePhoto for reference.
func (c *Client) PhotosUploadProfilePhoto(ctx context.Context, request *PhotosUploadProfilePhotoRequest) (*PhotosPhoto, error) {
	var result PhotosPhoto

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
