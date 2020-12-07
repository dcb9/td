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

// ChannelsGetAdminLogRequest represents TL type `channels.getAdminLog#33ddf480`.
// Get the admin log of a channel/supergroup
//
// See https://core.telegram.org/method/channels.getAdminLog for reference.
type ChannelsGetAdminLogRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Channel
	Channel InputChannelClass
	// Search query, can be empty
	Q string
	// Event filter
	//
	// Use SetEventsFilter and GetEventsFilter helpers.
	EventsFilter ChannelAdminLogEventsFilter
	// Only show events from these admins
	//
	// Use SetAdmins and GetAdmins helpers.
	Admins []InputUserClass
	// Maximum ID of message to return (see pagination)
	MaxID int64
	// Minimum ID of message to return (see pagination)
	MinID int64
	// Maximum number of results to return, see pagination
	Limit int
}

// ChannelsGetAdminLogRequestTypeID is TL type id of ChannelsGetAdminLogRequest.
const ChannelsGetAdminLogRequestTypeID = 0x33ddf480

// Encode implements bin.Encoder.
func (g *ChannelsGetAdminLogRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getAdminLog#33ddf480 as nil")
	}
	b.PutID(ChannelsGetAdminLogRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getAdminLog#33ddf480: field flags: %w", err)
	}
	if g.Channel == nil {
		return fmt.Errorf("unable to encode channels.getAdminLog#33ddf480: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getAdminLog#33ddf480: field channel: %w", err)
	}
	b.PutString(g.Q)
	if g.Flags.Has(0) {
		if err := g.EventsFilter.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.getAdminLog#33ddf480: field events_filter: %w", err)
		}
	}
	if g.Flags.Has(1) {
		b.PutVectorHeader(len(g.Admins))
		for idx, v := range g.Admins {
			if v == nil {
				return fmt.Errorf("unable to encode channels.getAdminLog#33ddf480: field admins element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode channels.getAdminLog#33ddf480: field admins element with index %d: %w", idx, err)
			}
		}
	}
	b.PutLong(g.MaxID)
	b.PutLong(g.MinID)
	b.PutInt(g.Limit)
	return nil
}

// SetEventsFilter sets value of EventsFilter conditional field.
func (g *ChannelsGetAdminLogRequest) SetEventsFilter(value ChannelAdminLogEventsFilter) {
	g.Flags.Set(0)
	g.EventsFilter = value
}

// GetEventsFilter returns value of EventsFilter conditional field and
// boolean which is true if field was set.
func (g *ChannelsGetAdminLogRequest) GetEventsFilter() (value ChannelAdminLogEventsFilter, ok bool) {
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.EventsFilter, true
}

// SetAdmins sets value of Admins conditional field.
func (g *ChannelsGetAdminLogRequest) SetAdmins(value []InputUserClass) {
	g.Flags.Set(1)
	g.Admins = value
}

// GetAdmins returns value of Admins conditional field and
// boolean which is true if field was set.
func (g *ChannelsGetAdminLogRequest) GetAdmins() (value []InputUserClass, ok bool) {
	if !g.Flags.Has(1) {
		return value, false
	}
	return g.Admins, true
}

// Decode implements bin.Decoder.
func (g *ChannelsGetAdminLogRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getAdminLog#33ddf480 to nil")
	}
	if err := b.ConsumeID(ChannelsGetAdminLogRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getAdminLog#33ddf480: %w", err)
	}

	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.getAdminLog#33ddf480: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.getAdminLog#33ddf480: field channel: %w", err)
		}
		g.Channel = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getAdminLog#33ddf480: field q: %w", err)
		}
		g.Q = value
	}
	if g.Flags.Has(0) {
		if err := g.EventsFilter.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.getAdminLog#33ddf480: field events_filter: %w", err)
		}
	}
	if g.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getAdminLog#33ddf480: field admins: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode channels.getAdminLog#33ddf480: field admins: %w", err)
			}
			g.Admins = append(g.Admins, value)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getAdminLog#33ddf480: field max_id: %w", err)
		}
		g.MaxID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getAdminLog#33ddf480: field min_id: %w", err)
		}
		g.MinID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getAdminLog#33ddf480: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsGetAdminLogRequest.
var (
	_ bin.Encoder = &ChannelsGetAdminLogRequest{}
	_ bin.Decoder = &ChannelsGetAdminLogRequest{}
)

// ChannelsGetAdminLog invokes method channels.getAdminLog#33ddf480 returning error if any.
// Get the admin log of a channel/supergroup
//
// See https://core.telegram.org/method/channels.getAdminLog for reference.
func (c *Client) ChannelsGetAdminLog(ctx context.Context, request *ChannelsGetAdminLogRequest) (*ChannelsAdminLogResults, error) {
	var result ChannelsAdminLogResults
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
