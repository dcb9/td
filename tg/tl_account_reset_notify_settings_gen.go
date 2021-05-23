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

// AccountResetNotifySettingsRequest represents TL type `account.resetNotifySettings#db7e1747`.
// Resets all notification settings from users and groups.
//
// See https://core.telegram.org/method/account.resetNotifySettings for reference.
type AccountResetNotifySettingsRequest struct {
}

// AccountResetNotifySettingsRequestTypeID is TL type id of AccountResetNotifySettingsRequest.
const AccountResetNotifySettingsRequestTypeID = 0xdb7e1747

func (r *AccountResetNotifySettingsRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *AccountResetNotifySettingsRequest) String() string {
	if r == nil {
		return "AccountResetNotifySettingsRequest(nil)"
	}
	type Alias AccountResetNotifySettingsRequest
	return fmt.Sprintf("AccountResetNotifySettingsRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountResetNotifySettingsRequest) TypeID() uint32 {
	return AccountResetNotifySettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountResetNotifySettingsRequest) TypeName() string {
	return "account.resetNotifySettings"
}

// TypeInfo returns info about TL type.
func (r *AccountResetNotifySettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.resetNotifySettings",
		ID:   AccountResetNotifySettingsRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *AccountResetNotifySettingsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetNotifySettings#db7e1747 as nil")
	}
	b.PutID(AccountResetNotifySettingsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AccountResetNotifySettingsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetNotifySettings#db7e1747 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountResetNotifySettingsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetNotifySettings#db7e1747 to nil")
	}
	if err := b.ConsumeID(AccountResetNotifySettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.resetNotifySettings#db7e1747: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AccountResetNotifySettingsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetNotifySettings#db7e1747 to nil")
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountResetNotifySettingsRequest.
var (
	_ bin.Encoder     = &AccountResetNotifySettingsRequest{}
	_ bin.Decoder     = &AccountResetNotifySettingsRequest{}
	_ bin.BareEncoder = &AccountResetNotifySettingsRequest{}
	_ bin.BareDecoder = &AccountResetNotifySettingsRequest{}
)

// AccountResetNotifySettings invokes method account.resetNotifySettings#db7e1747 returning error if any.
// Resets all notification settings from users and groups.
//
// See https://core.telegram.org/method/account.resetNotifySettings for reference.
func (c *Client) AccountResetNotifySettings(ctx context.Context) (bool, error) {
	var result BoolBox

	request := &AccountResetNotifySettingsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
