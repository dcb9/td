// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// AssignAppStoreTransactionRequest represents TL type `assignAppStoreTransaction#86f30bb0`.
type AssignAppStoreTransactionRequest struct {
	// App Store receipt
	Receipt []byte
	// Transaction purpose
	Purpose StorePaymentPurposeClass
}

// AssignAppStoreTransactionRequestTypeID is TL type id of AssignAppStoreTransactionRequest.
const AssignAppStoreTransactionRequestTypeID = 0x86f30bb0

// Ensuring interfaces in compile-time for AssignAppStoreTransactionRequest.
var (
	_ bin.Encoder     = &AssignAppStoreTransactionRequest{}
	_ bin.Decoder     = &AssignAppStoreTransactionRequest{}
	_ bin.BareEncoder = &AssignAppStoreTransactionRequest{}
	_ bin.BareDecoder = &AssignAppStoreTransactionRequest{}
)

func (a *AssignAppStoreTransactionRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Receipt == nil) {
		return false
	}
	if !(a.Purpose == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AssignAppStoreTransactionRequest) String() string {
	if a == nil {
		return "AssignAppStoreTransactionRequest(nil)"
	}
	type Alias AssignAppStoreTransactionRequest
	return fmt.Sprintf("AssignAppStoreTransactionRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AssignAppStoreTransactionRequest) TypeID() uint32 {
	return AssignAppStoreTransactionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AssignAppStoreTransactionRequest) TypeName() string {
	return "assignAppStoreTransaction"
}

// TypeInfo returns info about TL type.
func (a *AssignAppStoreTransactionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "assignAppStoreTransaction",
		ID:   AssignAppStoreTransactionRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Receipt",
			SchemaName: "receipt",
		},
		{
			Name:       "Purpose",
			SchemaName: "purpose",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AssignAppStoreTransactionRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode assignAppStoreTransaction#86f30bb0 as nil")
	}
	b.PutID(AssignAppStoreTransactionRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AssignAppStoreTransactionRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode assignAppStoreTransaction#86f30bb0 as nil")
	}
	b.PutBytes(a.Receipt)
	if a.Purpose == nil {
		return fmt.Errorf("unable to encode assignAppStoreTransaction#86f30bb0: field purpose is nil")
	}
	if err := a.Purpose.Encode(b); err != nil {
		return fmt.Errorf("unable to encode assignAppStoreTransaction#86f30bb0: field purpose: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AssignAppStoreTransactionRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode assignAppStoreTransaction#86f30bb0 to nil")
	}
	if err := b.ConsumeID(AssignAppStoreTransactionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode assignAppStoreTransaction#86f30bb0: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AssignAppStoreTransactionRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode assignAppStoreTransaction#86f30bb0 to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode assignAppStoreTransaction#86f30bb0: field receipt: %w", err)
		}
		a.Receipt = value
	}
	{
		value, err := DecodeStorePaymentPurpose(b)
		if err != nil {
			return fmt.Errorf("unable to decode assignAppStoreTransaction#86f30bb0: field purpose: %w", err)
		}
		a.Purpose = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AssignAppStoreTransactionRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode assignAppStoreTransaction#86f30bb0 as nil")
	}
	b.ObjStart()
	b.PutID("assignAppStoreTransaction")
	b.Comma()
	b.FieldStart("receipt")
	b.PutBytes(a.Receipt)
	b.Comma()
	b.FieldStart("purpose")
	if a.Purpose == nil {
		return fmt.Errorf("unable to encode assignAppStoreTransaction#86f30bb0: field purpose is nil")
	}
	if err := a.Purpose.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode assignAppStoreTransaction#86f30bb0: field purpose: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AssignAppStoreTransactionRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode assignAppStoreTransaction#86f30bb0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("assignAppStoreTransaction"); err != nil {
				return fmt.Errorf("unable to decode assignAppStoreTransaction#86f30bb0: %w", err)
			}
		case "receipt":
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode assignAppStoreTransaction#86f30bb0: field receipt: %w", err)
			}
			a.Receipt = value
		case "purpose":
			value, err := DecodeTDLibJSONStorePaymentPurpose(b)
			if err != nil {
				return fmt.Errorf("unable to decode assignAppStoreTransaction#86f30bb0: field purpose: %w", err)
			}
			a.Purpose = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetReceipt returns value of Receipt field.
func (a *AssignAppStoreTransactionRequest) GetReceipt() (value []byte) {
	if a == nil {
		return
	}
	return a.Receipt
}

// GetPurpose returns value of Purpose field.
func (a *AssignAppStoreTransactionRequest) GetPurpose() (value StorePaymentPurposeClass) {
	if a == nil {
		return
	}
	return a.Purpose
}

// AssignAppStoreTransaction invokes method assignAppStoreTransaction#86f30bb0 returning error if any.
func (c *Client) AssignAppStoreTransaction(ctx context.Context, request *AssignAppStoreTransactionRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}