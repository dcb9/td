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

// ContactStatus represents TL type `contactStatus#d3680c61`.
// Contact status: online / offline.
//
// See https://core.telegram.org/constructor/contactStatus for reference.
type ContactStatus struct {
	// User identifier
	UserID int
	// Online status
	Status UserStatusClass
}

// ContactStatusTypeID is TL type id of ContactStatus.
const ContactStatusTypeID = 0xd3680c61

// Encode implements bin.Encoder.
func (c *ContactStatus) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode contactStatus#d3680c61 as nil")
	}
	b.PutID(ContactStatusTypeID)
	b.PutInt(c.UserID)
	if c.Status == nil {
		return fmt.Errorf("unable to encode contactStatus#d3680c61: field status is nil")
	}
	if err := c.Status.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contactStatus#d3680c61: field status: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ContactStatus) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode contactStatus#d3680c61 to nil")
	}
	if err := b.ConsumeID(ContactStatusTypeID); err != nil {
		return fmt.Errorf("unable to decode contactStatus#d3680c61: %w", err)
	}

	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contactStatus#d3680c61: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := DecodeUserStatus(b)
		if err != nil {
			return fmt.Errorf("unable to decode contactStatus#d3680c61: field status: %w", err)
		}
		c.Status = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactStatus.
var (
	_ bin.Encoder = &ContactStatus{}
	_ bin.Decoder = &ContactStatus{}
)
