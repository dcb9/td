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

// ChatAdminRights represents TL type `chatAdminRights#5fb224d5`.
// Represents the rights of an admin in a channel/supergroup.
//
// See https://core.telegram.org/constructor/chatAdminRights for reference.
type ChatAdminRights struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// If set, allows the admin to modify the description of the channel/supergroup
	ChangeInfo bool
	// If set, allows the admin to post messages in the channel
	PostMessages bool
	// If set, allows the admin to also edit messages from other admins in the channel
	EditMessages bool
	// If set, allows the admin to also delete messages from other admins in the channel
	DeleteMessages bool
	// If set, allows the admin to ban users from the channel/supergroup
	BanUsers bool
	// If set, allows the admin to invite users in the channel/supergroup
	InviteUsers bool
	// If set, allows the admin to pin messages in the channel/supergroup
	PinMessages bool
	// If set, allows the admin to add other admins with the same (or more limited) permissions in the channel/supergroup
	AddAdmins bool
	// Whether this admin is anonymous
	Anonymous bool
	// ManageCall field of ChatAdminRights.
	ManageCall bool
}

// ChatAdminRightsTypeID is TL type id of ChatAdminRights.
const ChatAdminRightsTypeID = 0x5fb224d5

// Encode implements bin.Encoder.
func (c *ChatAdminRights) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAdminRights#5fb224d5 as nil")
	}
	b.PutID(ChatAdminRightsTypeID)
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatAdminRights#5fb224d5: field flags: %w", err)
	}
	return nil
}

// SetChangeInfo sets value of ChangeInfo conditional field.
func (c *ChatAdminRights) SetChangeInfo(value bool) {
	if value {
		c.Flags.Set(0)
	} else {
		c.Flags.Unset(0)
	}
}

// SetPostMessages sets value of PostMessages conditional field.
func (c *ChatAdminRights) SetPostMessages(value bool) {
	if value {
		c.Flags.Set(1)
	} else {
		c.Flags.Unset(1)
	}
}

// SetEditMessages sets value of EditMessages conditional field.
func (c *ChatAdminRights) SetEditMessages(value bool) {
	if value {
		c.Flags.Set(2)
	} else {
		c.Flags.Unset(2)
	}
}

// SetDeleteMessages sets value of DeleteMessages conditional field.
func (c *ChatAdminRights) SetDeleteMessages(value bool) {
	if value {
		c.Flags.Set(3)
	} else {
		c.Flags.Unset(3)
	}
}

// SetBanUsers sets value of BanUsers conditional field.
func (c *ChatAdminRights) SetBanUsers(value bool) {
	if value {
		c.Flags.Set(4)
	} else {
		c.Flags.Unset(4)
	}
}

// SetInviteUsers sets value of InviteUsers conditional field.
func (c *ChatAdminRights) SetInviteUsers(value bool) {
	if value {
		c.Flags.Set(5)
	} else {
		c.Flags.Unset(5)
	}
}

// SetPinMessages sets value of PinMessages conditional field.
func (c *ChatAdminRights) SetPinMessages(value bool) {
	if value {
		c.Flags.Set(7)
	} else {
		c.Flags.Unset(7)
	}
}

// SetAddAdmins sets value of AddAdmins conditional field.
func (c *ChatAdminRights) SetAddAdmins(value bool) {
	if value {
		c.Flags.Set(9)
	} else {
		c.Flags.Unset(9)
	}
}

// SetAnonymous sets value of Anonymous conditional field.
func (c *ChatAdminRights) SetAnonymous(value bool) {
	if value {
		c.Flags.Set(10)
	} else {
		c.Flags.Unset(10)
	}
}

// SetManageCall sets value of ManageCall conditional field.
func (c *ChatAdminRights) SetManageCall(value bool) {
	if value {
		c.Flags.Set(11)
	} else {
		c.Flags.Unset(11)
	}
}

// Decode implements bin.Decoder.
func (c *ChatAdminRights) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAdminRights#5fb224d5 to nil")
	}
	if err := b.ConsumeID(ChatAdminRightsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatAdminRights#5fb224d5: %w", err)
	}

	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatAdminRights#5fb224d5: field flags: %w", err)
		}
	}
	c.ChangeInfo = c.Flags.Has(0)
	c.PostMessages = c.Flags.Has(1)
	c.EditMessages = c.Flags.Has(2)
	c.DeleteMessages = c.Flags.Has(3)
	c.BanUsers = c.Flags.Has(4)
	c.InviteUsers = c.Flags.Has(5)
	c.PinMessages = c.Flags.Has(7)
	c.AddAdmins = c.Flags.Has(9)
	c.Anonymous = c.Flags.Has(10)
	c.ManageCall = c.Flags.Has(11)
	return nil
}

// Ensuring interfaces in compile-time for ChatAdminRights.
var (
	_ bin.Encoder = &ChatAdminRights{}
	_ bin.Decoder = &ChatAdminRights{}
)
