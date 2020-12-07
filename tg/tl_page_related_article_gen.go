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

// PageRelatedArticle represents TL type `pageRelatedArticle#b390dc08`.
// Related article
//
// See https://core.telegram.org/constructor/pageRelatedArticle for reference.
type PageRelatedArticle struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// URL of article
	URL string
	// Webpage ID of generated IV preview
	WebpageID int64
	// Title
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// Description
	//
	// Use SetDescription and GetDescription helpers.
	Description string
	// ID of preview photo
	//
	// Use SetPhotoID and GetPhotoID helpers.
	PhotoID int64
	// Author name
	//
	// Use SetAuthor and GetAuthor helpers.
	Author string
	// Date of pubblication
	//
	// Use SetPublishedDate and GetPublishedDate helpers.
	PublishedDate int
}

// PageRelatedArticleTypeID is TL type id of PageRelatedArticle.
const PageRelatedArticleTypeID = 0xb390dc08

// Encode implements bin.Encoder.
func (p *PageRelatedArticle) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageRelatedArticle#b390dc08 as nil")
	}
	b.PutID(PageRelatedArticleTypeID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode pageRelatedArticle#b390dc08: field flags: %w", err)
	}
	b.PutString(p.URL)
	b.PutLong(p.WebpageID)
	if p.Flags.Has(0) {
		b.PutString(p.Title)
	}
	if p.Flags.Has(1) {
		b.PutString(p.Description)
	}
	if p.Flags.Has(2) {
		b.PutLong(p.PhotoID)
	}
	if p.Flags.Has(3) {
		b.PutString(p.Author)
	}
	if p.Flags.Has(4) {
		b.PutInt(p.PublishedDate)
	}
	return nil
}

// SetTitle sets value of Title conditional field.
func (p *PageRelatedArticle) SetTitle(value string) {
	p.Flags.Set(0)
	p.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (p *PageRelatedArticle) GetTitle() (value string, ok bool) {
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.Title, true
}

// SetDescription sets value of Description conditional field.
func (p *PageRelatedArticle) SetDescription(value string) {
	p.Flags.Set(1)
	p.Description = value
}

// GetDescription returns value of Description conditional field and
// boolean which is true if field was set.
func (p *PageRelatedArticle) GetDescription() (value string, ok bool) {
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.Description, true
}

// SetPhotoID sets value of PhotoID conditional field.
func (p *PageRelatedArticle) SetPhotoID(value int64) {
	p.Flags.Set(2)
	p.PhotoID = value
}

// GetPhotoID returns value of PhotoID conditional field and
// boolean which is true if field was set.
func (p *PageRelatedArticle) GetPhotoID() (value int64, ok bool) {
	if !p.Flags.Has(2) {
		return value, false
	}
	return p.PhotoID, true
}

// SetAuthor sets value of Author conditional field.
func (p *PageRelatedArticle) SetAuthor(value string) {
	p.Flags.Set(3)
	p.Author = value
}

// GetAuthor returns value of Author conditional field and
// boolean which is true if field was set.
func (p *PageRelatedArticle) GetAuthor() (value string, ok bool) {
	if !p.Flags.Has(3) {
		return value, false
	}
	return p.Author, true
}

// SetPublishedDate sets value of PublishedDate conditional field.
func (p *PageRelatedArticle) SetPublishedDate(value int) {
	p.Flags.Set(4)
	p.PublishedDate = value
}

// GetPublishedDate returns value of PublishedDate conditional field and
// boolean which is true if field was set.
func (p *PageRelatedArticle) GetPublishedDate() (value int, ok bool) {
	if !p.Flags.Has(4) {
		return value, false
	}
	return p.PublishedDate, true
}

// Decode implements bin.Decoder.
func (p *PageRelatedArticle) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageRelatedArticle#b390dc08 to nil")
	}
	if err := b.ConsumeID(PageRelatedArticleTypeID); err != nil {
		return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: %w", err)
	}

	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: field url: %w", err)
		}
		p.URL = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: field webpage_id: %w", err)
		}
		p.WebpageID = value
	}
	if p.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: field title: %w", err)
		}
		p.Title = value
	}
	if p.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: field description: %w", err)
		}
		p.Description = value
	}
	if p.Flags.Has(2) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: field photo_id: %w", err)
		}
		p.PhotoID = value
	}
	if p.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: field author: %w", err)
		}
		p.Author = value
	}
	if p.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode pageRelatedArticle#b390dc08: field published_date: %w", err)
		}
		p.PublishedDate = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PageRelatedArticle.
var (
	_ bin.Encoder = &PageRelatedArticle{}
	_ bin.Decoder = &PageRelatedArticle{}
)
