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

// TestCallEmptyRequest represents TL type `testCallEmpty#da9c4a16`.
type TestCallEmptyRequest struct {
}

// TestCallEmptyRequestTypeID is TL type id of TestCallEmptyRequest.
const TestCallEmptyRequestTypeID = 0xda9c4a16

// Ensuring interfaces in compile-time for TestCallEmptyRequest.
var (
	_ bin.Encoder     = &TestCallEmptyRequest{}
	_ bin.Decoder     = &TestCallEmptyRequest{}
	_ bin.BareEncoder = &TestCallEmptyRequest{}
	_ bin.BareDecoder = &TestCallEmptyRequest{}
)

func (t *TestCallEmptyRequest) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestCallEmptyRequest) String() string {
	if t == nil {
		return "TestCallEmptyRequest(nil)"
	}
	type Alias TestCallEmptyRequest
	return fmt.Sprintf("TestCallEmptyRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestCallEmptyRequest) TypeID() uint32 {
	return TestCallEmptyRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*TestCallEmptyRequest) TypeName() string {
	return "testCallEmpty"
}

// TypeInfo returns info about TL type.
func (t *TestCallEmptyRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "testCallEmpty",
		ID:   TestCallEmptyRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *TestCallEmptyRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testCallEmpty#da9c4a16 as nil")
	}
	b.PutID(TestCallEmptyRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TestCallEmptyRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testCallEmpty#da9c4a16 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TestCallEmptyRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testCallEmpty#da9c4a16 to nil")
	}
	if err := b.ConsumeID(TestCallEmptyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode testCallEmpty#da9c4a16: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TestCallEmptyRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testCallEmpty#da9c4a16 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TestCallEmptyRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode testCallEmpty#da9c4a16 as nil")
	}
	b.ObjStart()
	b.PutID("testCallEmpty")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TestCallEmptyRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode testCallEmpty#da9c4a16 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("testCallEmpty"); err != nil {
				return fmt.Errorf("unable to decode testCallEmpty#da9c4a16: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// TestCallEmpty invokes method testCallEmpty#da9c4a16 returning error if any.
func (c *Client) TestCallEmpty(ctx context.Context) error {
	var ok Ok

	request := &TestCallEmptyRequest{}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}