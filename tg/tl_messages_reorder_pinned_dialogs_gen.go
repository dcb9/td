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

// MessagesReorderPinnedDialogsRequest represents TL type `messages.reorderPinnedDialogs#3b1adf37`.
// Reorder pinned dialogs
//
// See https://core.telegram.org/method/messages.reorderPinnedDialogs for reference.
type MessagesReorderPinnedDialogsRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// If set, dialogs pinned server-side but not present in the order field will be unpinned.
	Force bool
	// Peer folder ID, for more info click here
	FolderID int
	// New dialog order
	Order []InputDialogPeerClass
}

// MessagesReorderPinnedDialogsRequestTypeID is TL type id of MessagesReorderPinnedDialogsRequest.
const MessagesReorderPinnedDialogsRequestTypeID = 0x3b1adf37

// Encode implements bin.Encoder.
func (r *MessagesReorderPinnedDialogsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reorderPinnedDialogs#3b1adf37 as nil")
	}
	b.PutID(MessagesReorderPinnedDialogsRequestTypeID)
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.reorderPinnedDialogs#3b1adf37: field flags: %w", err)
	}
	b.PutInt(r.FolderID)
	b.PutVectorHeader(len(r.Order))
	for idx, v := range r.Order {
		if v == nil {
			return fmt.Errorf("unable to encode messages.reorderPinnedDialogs#3b1adf37: field order element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.reorderPinnedDialogs#3b1adf37: field order element with index %d: %w", idx, err)
		}
	}
	return nil
}

// SetForce sets value of Force conditional field.
func (r *MessagesReorderPinnedDialogsRequest) SetForce(value bool) {
	if value {
		r.Flags.Set(0)
	} else {
		r.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (r *MessagesReorderPinnedDialogsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reorderPinnedDialogs#3b1adf37 to nil")
	}
	if err := b.ConsumeID(MessagesReorderPinnedDialogsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.reorderPinnedDialogs#3b1adf37: %w", err)
	}

	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.reorderPinnedDialogs#3b1adf37: field flags: %w", err)
		}
	}
	r.Force = r.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.reorderPinnedDialogs#3b1adf37: field folder_id: %w", err)
		}
		r.FolderID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.reorderPinnedDialogs#3b1adf37: field order: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputDialogPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.reorderPinnedDialogs#3b1adf37: field order: %w", err)
			}
			r.Order = append(r.Order, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesReorderPinnedDialogsRequest.
var (
	_ bin.Encoder = &MessagesReorderPinnedDialogsRequest{}
	_ bin.Decoder = &MessagesReorderPinnedDialogsRequest{}
)

// MessagesReorderPinnedDialogs invokes method messages.reorderPinnedDialogs#3b1adf37 returning error if any.
// Reorder pinned dialogs
//
// See https://core.telegram.org/method/messages.reorderPinnedDialogs for reference.
func (c *Client) MessagesReorderPinnedDialogs(ctx context.Context, request *MessagesReorderPinnedDialogsRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
