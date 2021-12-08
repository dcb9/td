// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// PageBlockRelatedArticle represents TL type `pageBlockRelatedArticle#1cae8493`.
type PageBlockRelatedArticle struct {
	// Related article URL
	URL string
	// Article title; may be empty
	Title string
	// Contains information about a related article
	Description string
	// Article photo; may be null
	Photo Photo
	// Article author; may be empty
	Author string
	// Point in time (Unix timestamp) when the article was published; 0 if unknown
	PublishDate int32
}

// PageBlockRelatedArticleTypeID is TL type id of PageBlockRelatedArticle.
const PageBlockRelatedArticleTypeID = 0x1cae8493

// Ensuring interfaces in compile-time for PageBlockRelatedArticle.
var (
	_ bin.Encoder     = &PageBlockRelatedArticle{}
	_ bin.Decoder     = &PageBlockRelatedArticle{}
	_ bin.BareEncoder = &PageBlockRelatedArticle{}
	_ bin.BareDecoder = &PageBlockRelatedArticle{}
)

func (p *PageBlockRelatedArticle) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.URL == "") {
		return false
	}
	if !(p.Title == "") {
		return false
	}
	if !(p.Description == "") {
		return false
	}
	if !(p.Photo.Zero()) {
		return false
	}
	if !(p.Author == "") {
		return false
	}
	if !(p.PublishDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PageBlockRelatedArticle) String() string {
	if p == nil {
		return "PageBlockRelatedArticle(nil)"
	}
	type Alias PageBlockRelatedArticle
	return fmt.Sprintf("PageBlockRelatedArticle%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PageBlockRelatedArticle) TypeID() uint32 {
	return PageBlockRelatedArticleTypeID
}

// TypeName returns name of type in TL schema.
func (*PageBlockRelatedArticle) TypeName() string {
	return "pageBlockRelatedArticle"
}

// TypeInfo returns info about TL type.
func (p *PageBlockRelatedArticle) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pageBlockRelatedArticle",
		ID:   PageBlockRelatedArticleTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Description",
			SchemaName: "description",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
		{
			Name:       "Author",
			SchemaName: "author",
		},
		{
			Name:       "PublishDate",
			SchemaName: "publish_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PageBlockRelatedArticle) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockRelatedArticle#1cae8493 as nil")
	}
	b.PutID(PageBlockRelatedArticleTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PageBlockRelatedArticle) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockRelatedArticle#1cae8493 as nil")
	}
	b.PutString(p.URL)
	b.PutString(p.Title)
	b.PutString(p.Description)
	if err := p.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode pageBlockRelatedArticle#1cae8493: field photo: %w", err)
	}
	b.PutString(p.Author)
	b.PutInt32(p.PublishDate)
	return nil
}

// Decode implements bin.Decoder.
func (p *PageBlockRelatedArticle) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockRelatedArticle#1cae8493 to nil")
	}
	if err := b.ConsumeID(PageBlockRelatedArticleTypeID); err != nil {
		return fmt.Errorf("unable to decode pageBlockRelatedArticle#1cae8493: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PageBlockRelatedArticle) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockRelatedArticle#1cae8493 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pageBlockRelatedArticle#1cae8493: field url: %w", err)
		}
		p.URL = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pageBlockRelatedArticle#1cae8493: field title: %w", err)
		}
		p.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pageBlockRelatedArticle#1cae8493: field description: %w", err)
		}
		p.Description = value
	}
	{
		if err := p.Photo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode pageBlockRelatedArticle#1cae8493: field photo: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pageBlockRelatedArticle#1cae8493: field author: %w", err)
		}
		p.Author = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode pageBlockRelatedArticle#1cae8493: field publish_date: %w", err)
		}
		p.PublishDate = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PageBlockRelatedArticle) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode pageBlockRelatedArticle#1cae8493 as nil")
	}
	b.ObjStart()
	b.PutID("pageBlockRelatedArticle")
	b.FieldStart("url")
	b.PutString(p.URL)
	b.FieldStart("title")
	b.PutString(p.Title)
	b.FieldStart("description")
	b.PutString(p.Description)
	b.FieldStart("photo")
	if err := p.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode pageBlockRelatedArticle#1cae8493: field photo: %w", err)
	}
	b.FieldStart("author")
	b.PutString(p.Author)
	b.FieldStart("publish_date")
	b.PutInt32(p.PublishDate)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PageBlockRelatedArticle) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode pageBlockRelatedArticle#1cae8493 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("pageBlockRelatedArticle"); err != nil {
				return fmt.Errorf("unable to decode pageBlockRelatedArticle#1cae8493: %w", err)
			}
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode pageBlockRelatedArticle#1cae8493: field url: %w", err)
			}
			p.URL = value
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode pageBlockRelatedArticle#1cae8493: field title: %w", err)
			}
			p.Title = value
		case "description":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode pageBlockRelatedArticle#1cae8493: field description: %w", err)
			}
			p.Description = value
		case "photo":
			if err := p.Photo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode pageBlockRelatedArticle#1cae8493: field photo: %w", err)
			}
		case "author":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode pageBlockRelatedArticle#1cae8493: field author: %w", err)
			}
			p.Author = value
		case "publish_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode pageBlockRelatedArticle#1cae8493: field publish_date: %w", err)
			}
			p.PublishDate = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetURL returns value of URL field.
func (p *PageBlockRelatedArticle) GetURL() (value string) {
	return p.URL
}

// GetTitle returns value of Title field.
func (p *PageBlockRelatedArticle) GetTitle() (value string) {
	return p.Title
}

// GetDescription returns value of Description field.
func (p *PageBlockRelatedArticle) GetDescription() (value string) {
	return p.Description
}

// GetPhoto returns value of Photo field.
func (p *PageBlockRelatedArticle) GetPhoto() (value Photo) {
	return p.Photo
}

// GetAuthor returns value of Author field.
func (p *PageBlockRelatedArticle) GetAuthor() (value string) {
	return p.Author
}

// GetPublishDate returns value of PublishDate field.
func (p *PageBlockRelatedArticle) GetPublishDate() (value int32) {
	return p.PublishDate
}