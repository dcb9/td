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

// SavedPhoneContact represents TL type `savedPhoneContact#1142bd56`.
// Saved contact
//
// See https://core.telegram.org/constructor/savedPhoneContact for reference.
type SavedPhoneContact struct {
	// Phone number
	Phone string
	// First name
	FirstName string
	// Last name
	LastName string
	// Date added
	Date int
}

// SavedPhoneContactTypeID is TL type id of SavedPhoneContact.
const SavedPhoneContactTypeID = 0x1142bd56

// Encode implements bin.Encoder.
func (s *SavedPhoneContact) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode savedPhoneContact#1142bd56 as nil")
	}
	b.PutID(SavedPhoneContactTypeID)
	b.PutString(s.Phone)
	b.PutString(s.FirstName)
	b.PutString(s.LastName)
	b.PutInt(s.Date)
	return nil
}

// Decode implements bin.Decoder.
func (s *SavedPhoneContact) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode savedPhoneContact#1142bd56 to nil")
	}
	if err := b.ConsumeID(SavedPhoneContactTypeID); err != nil {
		return fmt.Errorf("unable to decode savedPhoneContact#1142bd56: %w", err)
	}

	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode savedPhoneContact#1142bd56: field phone: %w", err)
		}
		s.Phone = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode savedPhoneContact#1142bd56: field first_name: %w", err)
		}
		s.FirstName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode savedPhoneContact#1142bd56: field last_name: %w", err)
		}
		s.LastName = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode savedPhoneContact#1142bd56: field date: %w", err)
		}
		s.Date = value
	}
	return nil
}

// Ensuring interfaces in compile-time for SavedPhoneContact.
var (
	_ bin.Encoder = &SavedPhoneContact{}
	_ bin.Decoder = &SavedPhoneContact{}
)
