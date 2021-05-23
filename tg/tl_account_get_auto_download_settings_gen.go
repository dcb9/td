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
)

// AccountGetAutoDownloadSettingsRequest represents TL type `account.getAutoDownloadSettings#56da0b3f`.
// Get media autodownload settings
//
// See https://core.telegram.org/method/account.getAutoDownloadSettings for reference.
type AccountGetAutoDownloadSettingsRequest struct {
}

// AccountGetAutoDownloadSettingsRequestTypeID is TL type id of AccountGetAutoDownloadSettingsRequest.
const AccountGetAutoDownloadSettingsRequestTypeID = 0x56da0b3f

func (g *AccountGetAutoDownloadSettingsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetAutoDownloadSettingsRequest) String() string {
	if g == nil {
		return "AccountGetAutoDownloadSettingsRequest(nil)"
	}
	type Alias AccountGetAutoDownloadSettingsRequest
	return fmt.Sprintf("AccountGetAutoDownloadSettingsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetAutoDownloadSettingsRequest) TypeID() uint32 {
	return AccountGetAutoDownloadSettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetAutoDownloadSettingsRequest) TypeName() string {
	return "account.getAutoDownloadSettings"
}

// TypeInfo returns info about TL type.
func (g *AccountGetAutoDownloadSettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getAutoDownloadSettings",
		ID:   AccountGetAutoDownloadSettingsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetAutoDownloadSettingsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getAutoDownloadSettings#56da0b3f as nil")
	}
	b.PutID(AccountGetAutoDownloadSettingsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetAutoDownloadSettingsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getAutoDownloadSettings#56da0b3f as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetAutoDownloadSettingsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getAutoDownloadSettings#56da0b3f to nil")
	}
	if err := b.ConsumeID(AccountGetAutoDownloadSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getAutoDownloadSettings#56da0b3f: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetAutoDownloadSettingsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getAutoDownloadSettings#56da0b3f to nil")
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetAutoDownloadSettingsRequest.
var (
	_ bin.Encoder     = &AccountGetAutoDownloadSettingsRequest{}
	_ bin.Decoder     = &AccountGetAutoDownloadSettingsRequest{}
	_ bin.BareEncoder = &AccountGetAutoDownloadSettingsRequest{}
	_ bin.BareDecoder = &AccountGetAutoDownloadSettingsRequest{}
)

// AccountGetAutoDownloadSettings invokes method account.getAutoDownloadSettings#56da0b3f returning error if any.
// Get media autodownload settings
//
// See https://core.telegram.org/method/account.getAutoDownloadSettings for reference.
func (c *Client) AccountGetAutoDownloadSettings(ctx context.Context) (*AccountAutoDownloadSettings, error) {
	var result AccountAutoDownloadSettings

	request := &AccountGetAutoDownloadSettingsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
