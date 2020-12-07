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

// AccountSaveAutoDownloadSettingsRequest represents TL type `account.saveAutoDownloadSettings#76f36233`.
// Change media autodownload settings
//
// See https://core.telegram.org/method/account.saveAutoDownloadSettings for reference.
type AccountSaveAutoDownloadSettingsRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Whether to save settings in the low data usage preset
	Low bool
	// Whether to save settings in the high data usage preset
	High bool
	// Media autodownload settings
	Settings AutoDownloadSettings
}

// AccountSaveAutoDownloadSettingsRequestTypeID is TL type id of AccountSaveAutoDownloadSettingsRequest.
const AccountSaveAutoDownloadSettingsRequestTypeID = 0x76f36233

// Encode implements bin.Encoder.
func (s *AccountSaveAutoDownloadSettingsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.saveAutoDownloadSettings#76f36233 as nil")
	}
	b.PutID(AccountSaveAutoDownloadSettingsRequestTypeID)
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.saveAutoDownloadSettings#76f36233: field flags: %w", err)
	}
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.saveAutoDownloadSettings#76f36233: field settings: %w", err)
	}
	return nil
}

// SetLow sets value of Low conditional field.
func (s *AccountSaveAutoDownloadSettingsRequest) SetLow(value bool) {
	if value {
		s.Flags.Set(0)
	} else {
		s.Flags.Unset(0)
	}
}

// SetHigh sets value of High conditional field.
func (s *AccountSaveAutoDownloadSettingsRequest) SetHigh(value bool) {
	if value {
		s.Flags.Set(1)
	} else {
		s.Flags.Unset(1)
	}
}

// Decode implements bin.Decoder.
func (s *AccountSaveAutoDownloadSettingsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.saveAutoDownloadSettings#76f36233 to nil")
	}
	if err := b.ConsumeID(AccountSaveAutoDownloadSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.saveAutoDownloadSettings#76f36233: %w", err)
	}

	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.saveAutoDownloadSettings#76f36233: field flags: %w", err)
		}
	}
	s.Low = s.Flags.Has(0)
	s.High = s.Flags.Has(1)
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.saveAutoDownloadSettings#76f36233: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSaveAutoDownloadSettingsRequest.
var (
	_ bin.Encoder = &AccountSaveAutoDownloadSettingsRequest{}
	_ bin.Decoder = &AccountSaveAutoDownloadSettingsRequest{}
)

// AccountSaveAutoDownloadSettings invokes method account.saveAutoDownloadSettings#76f36233 returning error if any.
// Change media autodownload settings
//
// See https://core.telegram.org/method/account.saveAutoDownloadSettings for reference.
func (c *Client) AccountSaveAutoDownloadSettings(ctx context.Context, request *AccountSaveAutoDownloadSettingsRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
