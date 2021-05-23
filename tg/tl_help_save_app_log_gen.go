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

// HelpSaveAppLogRequest represents TL type `help.saveAppLog#6f02f748`.
// Saves logs of application on the server.
//
// See https://core.telegram.org/method/help.saveAppLog for reference.
type HelpSaveAppLogRequest struct {
	// List of input events
	Events []InputAppEvent
}

// HelpSaveAppLogRequestTypeID is TL type id of HelpSaveAppLogRequest.
const HelpSaveAppLogRequestTypeID = 0x6f02f748

func (s *HelpSaveAppLogRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Events == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *HelpSaveAppLogRequest) String() string {
	if s == nil {
		return "HelpSaveAppLogRequest(nil)"
	}
	type Alias HelpSaveAppLogRequest
	return fmt.Sprintf("HelpSaveAppLogRequest%+v", Alias(*s))
}

// FillFrom fills HelpSaveAppLogRequest from given interface.
func (s *HelpSaveAppLogRequest) FillFrom(from interface {
	GetEvents() (value []InputAppEvent)
}) {
	s.Events = from.GetEvents()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpSaveAppLogRequest) TypeID() uint32 {
	return HelpSaveAppLogRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpSaveAppLogRequest) TypeName() string {
	return "help.saveAppLog"
}

// TypeInfo returns info about TL type.
func (s *HelpSaveAppLogRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.saveAppLog",
		ID:   HelpSaveAppLogRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Events",
			SchemaName: "events",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *HelpSaveAppLogRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode help.saveAppLog#6f02f748 as nil")
	}
	b.PutID(HelpSaveAppLogRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *HelpSaveAppLogRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode help.saveAppLog#6f02f748 as nil")
	}
	b.PutVectorHeader(len(s.Events))
	for idx, v := range s.Events {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.saveAppLog#6f02f748: field events element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetEvents returns value of Events field.
func (s *HelpSaveAppLogRequest) GetEvents() (value []InputAppEvent) {
	return s.Events
}

// Decode implements bin.Decoder.
func (s *HelpSaveAppLogRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode help.saveAppLog#6f02f748 to nil")
	}
	if err := b.ConsumeID(HelpSaveAppLogRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.saveAppLog#6f02f748: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *HelpSaveAppLogRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode help.saveAppLog#6f02f748 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.saveAppLog#6f02f748: field events: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value InputAppEvent
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode help.saveAppLog#6f02f748: field events: %w", err)
			}
			s.Events = append(s.Events, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpSaveAppLogRequest.
var (
	_ bin.Encoder     = &HelpSaveAppLogRequest{}
	_ bin.Decoder     = &HelpSaveAppLogRequest{}
	_ bin.BareEncoder = &HelpSaveAppLogRequest{}
	_ bin.BareDecoder = &HelpSaveAppLogRequest{}
)

// HelpSaveAppLog invokes method help.saveAppLog#6f02f748 returning error if any.
// Saves logs of application on the server.
//
// See https://core.telegram.org/method/help.saveAppLog for reference.
func (c *Client) HelpSaveAppLog(ctx context.Context, events []InputAppEvent) (bool, error) {
	var result BoolBox

	request := &HelpSaveAppLogRequest{
		Events: events,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
