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

// PhoneGetGroupCallRequest represents TL type `phone.getGroupCall#c7cb017`.
//
// See https://core.telegram.org/method/phone.getGroupCall for reference.
type PhoneGetGroupCallRequest struct {
	// Call field of PhoneGetGroupCallRequest.
	Call InputGroupCall
}

// PhoneGetGroupCallRequestTypeID is TL type id of PhoneGetGroupCallRequest.
const PhoneGetGroupCallRequestTypeID = 0xc7cb017

func (g *PhoneGetGroupCallRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Call.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PhoneGetGroupCallRequest) String() string {
	if g == nil {
		return "PhoneGetGroupCallRequest(nil)"
	}
	type Alias PhoneGetGroupCallRequest
	return fmt.Sprintf("PhoneGetGroupCallRequest%+v", Alias(*g))
}

// FillFrom fills PhoneGetGroupCallRequest from given interface.
func (g *PhoneGetGroupCallRequest) FillFrom(from interface {
	GetCall() (value InputGroupCall)
}) {
	g.Call = from.GetCall()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneGetGroupCallRequest) TypeID() uint32 {
	return PhoneGetGroupCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneGetGroupCallRequest) TypeName() string {
	return "phone.getGroupCall"
}

// TypeInfo returns info about TL type.
func (g *PhoneGetGroupCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.getGroupCall",
		ID:   PhoneGetGroupCallRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Call",
			SchemaName: "call",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PhoneGetGroupCallRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode phone.getGroupCall#c7cb017 as nil")
	}
	b.PutID(PhoneGetGroupCallRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *PhoneGetGroupCallRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode phone.getGroupCall#c7cb017 as nil")
	}
	if err := g.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.getGroupCall#c7cb017: field call: %w", err)
	}
	return nil
}

// GetCall returns value of Call field.
func (g *PhoneGetGroupCallRequest) GetCall() (value InputGroupCall) {
	return g.Call
}

// Decode implements bin.Decoder.
func (g *PhoneGetGroupCallRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode phone.getGroupCall#c7cb017 to nil")
	}
	if err := b.ConsumeID(PhoneGetGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.getGroupCall#c7cb017: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *PhoneGetGroupCallRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode phone.getGroupCall#c7cb017 to nil")
	}
	{
		if err := g.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.getGroupCall#c7cb017: field call: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneGetGroupCallRequest.
var (
	_ bin.Encoder     = &PhoneGetGroupCallRequest{}
	_ bin.Decoder     = &PhoneGetGroupCallRequest{}
	_ bin.BareEncoder = &PhoneGetGroupCallRequest{}
	_ bin.BareDecoder = &PhoneGetGroupCallRequest{}
)

// PhoneGetGroupCall invokes method phone.getGroupCall#c7cb017 returning error if any.
//
// See https://core.telegram.org/method/phone.getGroupCall for reference.
func (c *Client) PhoneGetGroupCall(ctx context.Context, call InputGroupCall) (*PhoneGroupCall, error) {
	var result PhoneGroupCall

	request := &PhoneGetGroupCallRequest{
		Call: call,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
