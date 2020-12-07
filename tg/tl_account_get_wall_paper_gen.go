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

// AccountGetWallPaperRequest represents TL type `account.getWallPaper#fc8ddbea`.
// Get info about a certain wallpaper
//
// See https://core.telegram.org/method/account.getWallPaper for reference.
type AccountGetWallPaperRequest struct {
	// The wallpaper to get info about
	Wallpaper InputWallPaperClass
}

// AccountGetWallPaperRequestTypeID is TL type id of AccountGetWallPaperRequest.
const AccountGetWallPaperRequestTypeID = 0xfc8ddbea

// Encode implements bin.Encoder.
func (g *AccountGetWallPaperRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getWallPaper#fc8ddbea as nil")
	}
	b.PutID(AccountGetWallPaperRequestTypeID)
	if g.Wallpaper == nil {
		return fmt.Errorf("unable to encode account.getWallPaper#fc8ddbea: field wallpaper is nil")
	}
	if err := g.Wallpaper.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.getWallPaper#fc8ddbea: field wallpaper: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetWallPaperRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getWallPaper#fc8ddbea to nil")
	}
	if err := b.ConsumeID(AccountGetWallPaperRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getWallPaper#fc8ddbea: %w", err)
	}

	{
		value, err := DecodeInputWallPaper(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.getWallPaper#fc8ddbea: field wallpaper: %w", err)
		}
		g.Wallpaper = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetWallPaperRequest.
var (
	_ bin.Encoder = &AccountGetWallPaperRequest{}
	_ bin.Decoder = &AccountGetWallPaperRequest{}
)

// AccountGetWallPaper invokes method account.getWallPaper#fc8ddbea returning error if any.
// Get info about a certain wallpaper
//
// See https://core.telegram.org/method/account.getWallPaper for reference.
func (c *Client) AccountGetWallPaper(ctx context.Context, request *AccountGetWallPaperRequest) (WallPaperClass, error) {
	var result WallPaperBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.WallPaper, nil
}
