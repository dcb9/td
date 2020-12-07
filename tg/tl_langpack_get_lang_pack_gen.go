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

// LangpackGetLangPackRequest represents TL type `langpack.getLangPack#f2f2330a`.
// Get localization pack strings
//
// See https://core.telegram.org/method/langpack.getLangPack for reference.
type LangpackGetLangPackRequest struct {
	// Language pack name
	LangPack string
	// Language code
	LangCode string
}

// LangpackGetLangPackRequestTypeID is TL type id of LangpackGetLangPackRequest.
const LangpackGetLangPackRequestTypeID = 0xf2f2330a

// Encode implements bin.Encoder.
func (g *LangpackGetLangPackRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode langpack.getLangPack#f2f2330a as nil")
	}
	b.PutID(LangpackGetLangPackRequestTypeID)
	b.PutString(g.LangPack)
	b.PutString(g.LangCode)
	return nil
}

// Decode implements bin.Decoder.
func (g *LangpackGetLangPackRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode langpack.getLangPack#f2f2330a to nil")
	}
	if err := b.ConsumeID(LangpackGetLangPackRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode langpack.getLangPack#f2f2330a: %w", err)
	}

	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getLangPack#f2f2330a: field lang_pack: %w", err)
		}
		g.LangPack = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getLangPack#f2f2330a: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for LangpackGetLangPackRequest.
var (
	_ bin.Encoder = &LangpackGetLangPackRequest{}
	_ bin.Decoder = &LangpackGetLangPackRequest{}
)

// LangpackGetLangPack invokes method langpack.getLangPack#f2f2330a returning error if any.
// Get localization pack strings
//
// See https://core.telegram.org/method/langpack.getLangPack for reference.
func (c *Client) LangpackGetLangPack(ctx context.Context, request *LangpackGetLangPackRequest) (*LangPackDifference, error) {
	var result LangPackDifference
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
