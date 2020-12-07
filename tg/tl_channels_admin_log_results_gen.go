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

// ChannelsAdminLogResults represents TL type `channels.adminLogResults#ed8af74d`.
// Admin log events
//
// See https://core.telegram.org/constructor/channels.adminLogResults for reference.
type ChannelsAdminLogResults struct {
	// Admin log events
	Events []ChannelAdminLogEvent
	// Chats mentioned in events
	Chats []ChatClass
	// Users mentioned in events
	Users []UserClass
}

// ChannelsAdminLogResultsTypeID is TL type id of ChannelsAdminLogResults.
const ChannelsAdminLogResultsTypeID = 0xed8af74d

// Encode implements bin.Encoder.
func (a *ChannelsAdminLogResults) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode channels.adminLogResults#ed8af74d as nil")
	}
	b.PutID(ChannelsAdminLogResultsTypeID)
	b.PutVectorHeader(len(a.Events))
	for idx, v := range a.Events {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.adminLogResults#ed8af74d: field events element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(a.Chats))
	for idx, v := range a.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode channels.adminLogResults#ed8af74d: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.adminLogResults#ed8af74d: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(a.Users))
	for idx, v := range a.Users {
		if v == nil {
			return fmt.Errorf("unable to encode channels.adminLogResults#ed8af74d: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.adminLogResults#ed8af74d: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *ChannelsAdminLogResults) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode channels.adminLogResults#ed8af74d to nil")
	}
	if err := b.ConsumeID(ChannelsAdminLogResultsTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.adminLogResults#ed8af74d: %w", err)
	}

	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.adminLogResults#ed8af74d: field events: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ChannelAdminLogEvent
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode channels.adminLogResults#ed8af74d: field events: %w", err)
			}
			a.Events = append(a.Events, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.adminLogResults#ed8af74d: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode channels.adminLogResults#ed8af74d: field chats: %w", err)
			}
			a.Chats = append(a.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.adminLogResults#ed8af74d: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode channels.adminLogResults#ed8af74d: field users: %w", err)
			}
			a.Users = append(a.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsAdminLogResults.
var (
	_ bin.Encoder = &ChannelsAdminLogResults{}
	_ bin.Decoder = &ChannelsAdminLogResults{}
)
