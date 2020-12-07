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

// AccountGetAllSecureValuesRequest represents TL type `account.getAllSecureValues#b288bc7d`.
// Get all saved Telegram Passport documents, for more info see the passport docs »
//
// See https://core.telegram.org/method/account.getAllSecureValues for reference.
type AccountGetAllSecureValuesRequest struct {
}

// AccountGetAllSecureValuesRequestTypeID is TL type id of AccountGetAllSecureValuesRequest.
const AccountGetAllSecureValuesRequestTypeID = 0xb288bc7d

// Encode implements bin.Encoder.
func (g *AccountGetAllSecureValuesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getAllSecureValues#b288bc7d as nil")
	}
	b.PutID(AccountGetAllSecureValuesRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetAllSecureValuesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getAllSecureValues#b288bc7d to nil")
	}
	if err := b.ConsumeID(AccountGetAllSecureValuesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getAllSecureValues#b288bc7d: %w", err)
	}

	return nil
}

// Ensuring interfaces in compile-time for AccountGetAllSecureValuesRequest.
var (
	_ bin.Encoder = &AccountGetAllSecureValuesRequest{}
	_ bin.Decoder = &AccountGetAllSecureValuesRequest{}
)

// AccountGetAllSecureValues invokes method account.getAllSecureValues#b288bc7d returning error if any.
// Get all saved Telegram Passport documents, for more info see the passport docs »
//
// See https://core.telegram.org/method/account.getAllSecureValues for reference.
func (c *Client) AccountGetAllSecureValues(ctx context.Context, request *AccountGetAllSecureValuesRequest) (*SecureValueVector, error) {
	var result SecureValueVector
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
