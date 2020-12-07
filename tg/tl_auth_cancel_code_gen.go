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

// AuthCancelCodeRequest represents TL type `auth.cancelCode#1f040578`.
// Cancel the login verification code
//
// See https://core.telegram.org/method/auth.cancelCode for reference.
type AuthCancelCodeRequest struct {
	// Phone number
	PhoneNumber string
	// Phone code hash from auth.sendCode
	PhoneCodeHash string
}

// AuthCancelCodeRequestTypeID is TL type id of AuthCancelCodeRequest.
const AuthCancelCodeRequestTypeID = 0x1f040578

// Encode implements bin.Encoder.
func (c *AuthCancelCodeRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode auth.cancelCode#1f040578 as nil")
	}
	b.PutID(AuthCancelCodeRequestTypeID)
	b.PutString(c.PhoneNumber)
	b.PutString(c.PhoneCodeHash)
	return nil
}

// Decode implements bin.Decoder.
func (c *AuthCancelCodeRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode auth.cancelCode#1f040578 to nil")
	}
	if err := b.ConsumeID(AuthCancelCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.cancelCode#1f040578: %w", err)
	}

	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.cancelCode#1f040578: field phone_number: %w", err)
		}
		c.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.cancelCode#1f040578: field phone_code_hash: %w", err)
		}
		c.PhoneCodeHash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthCancelCodeRequest.
var (
	_ bin.Encoder = &AuthCancelCodeRequest{}
	_ bin.Decoder = &AuthCancelCodeRequest{}
)

// AuthCancelCode invokes method auth.cancelCode#1f040578 returning error if any.
// Cancel the login verification code
//
// See https://core.telegram.org/method/auth.cancelCode for reference.
func (c *Client) AuthCancelCode(ctx context.Context, request *AuthCancelCodeRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
