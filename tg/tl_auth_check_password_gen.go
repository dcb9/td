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

// AuthCheckPasswordRequest represents TL type `auth.checkPassword#d18b4d16`.
// Try logging to an account protected by a 2FA password.
//
// See https://core.telegram.org/method/auth.checkPassword for reference.
type AuthCheckPasswordRequest struct {
	// The account's password (see SRP)
	Password InputCheckPasswordSRPClass
}

// AuthCheckPasswordRequestTypeID is TL type id of AuthCheckPasswordRequest.
const AuthCheckPasswordRequestTypeID = 0xd18b4d16

// Encode implements bin.Encoder.
func (c *AuthCheckPasswordRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode auth.checkPassword#d18b4d16 as nil")
	}
	b.PutID(AuthCheckPasswordRequestTypeID)
	if c.Password == nil {
		return fmt.Errorf("unable to encode auth.checkPassword#d18b4d16: field password is nil")
	}
	if err := c.Password.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.checkPassword#d18b4d16: field password: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *AuthCheckPasswordRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode auth.checkPassword#d18b4d16 to nil")
	}
	if err := b.ConsumeID(AuthCheckPasswordRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.checkPassword#d18b4d16: %w", err)
	}

	{
		value, err := DecodeInputCheckPasswordSRP(b)
		if err != nil {
			return fmt.Errorf("unable to decode auth.checkPassword#d18b4d16: field password: %w", err)
		}
		c.Password = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthCheckPasswordRequest.
var (
	_ bin.Encoder = &AuthCheckPasswordRequest{}
	_ bin.Decoder = &AuthCheckPasswordRequest{}
)

// AuthCheckPassword invokes method auth.checkPassword#d18b4d16 returning error if any.
// Try logging to an account protected by a 2FA password.
//
// See https://core.telegram.org/method/auth.checkPassword for reference.
func (c *Client) AuthCheckPassword(ctx context.Context, request *AuthCheckPasswordRequest) (AuthAuthorizationClass, error) {
	var result AuthAuthorizationBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Authorization, nil
}
