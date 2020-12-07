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

// AccountCheckUsernameRequest represents TL type `account.checkUsername#2714d86c`.
// Validates a username and checks availability.
//
// See https://core.telegram.org/method/account.checkUsername for reference.
type AccountCheckUsernameRequest struct {
	// usernameAccepted characters: A-z (case-insensitive), 0-9 and underscores.Length: 5-32 characters.
	Username string
}

// AccountCheckUsernameRequestTypeID is TL type id of AccountCheckUsernameRequest.
const AccountCheckUsernameRequestTypeID = 0x2714d86c

// Encode implements bin.Encoder.
func (c *AccountCheckUsernameRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.checkUsername#2714d86c as nil")
	}
	b.PutID(AccountCheckUsernameRequestTypeID)
	b.PutString(c.Username)
	return nil
}

// Decode implements bin.Decoder.
func (c *AccountCheckUsernameRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.checkUsername#2714d86c to nil")
	}
	if err := b.ConsumeID(AccountCheckUsernameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.checkUsername#2714d86c: %w", err)
	}

	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.checkUsername#2714d86c: field username: %w", err)
		}
		c.Username = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountCheckUsernameRequest.
var (
	_ bin.Encoder = &AccountCheckUsernameRequest{}
	_ bin.Decoder = &AccountCheckUsernameRequest{}
)

// AccountCheckUsername invokes method account.checkUsername#2714d86c returning error if any.
// Validates a username and checks availability.
//
// See https://core.telegram.org/method/account.checkUsername for reference.
func (c *Client) AccountCheckUsername(ctx context.Context, request *AccountCheckUsernameRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
