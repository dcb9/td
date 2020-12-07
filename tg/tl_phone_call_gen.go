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

// PhoneCallEmpty represents TL type `phoneCallEmpty#5366c915`.
// Empty constructor
//
// See https://core.telegram.org/constructor/phoneCallEmpty for reference.
type PhoneCallEmpty struct {
	// Call ID
	ID int64
}

// PhoneCallEmptyTypeID is TL type id of PhoneCallEmpty.
const PhoneCallEmptyTypeID = 0x5366c915

// Encode implements bin.Encoder.
func (p *PhoneCallEmpty) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneCallEmpty#5366c915 as nil")
	}
	b.PutID(PhoneCallEmptyTypeID)
	b.PutLong(p.ID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhoneCallEmpty) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneCallEmpty#5366c915 to nil")
	}
	if err := b.ConsumeID(PhoneCallEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneCallEmpty#5366c915: %w", err)
	}

	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallEmpty#5366c915: field id: %w", err)
		}
		p.ID = value
	}
	return nil
}

// construct implements constructor of PhoneCallClass.
func (p PhoneCallEmpty) construct() PhoneCallClass { return &p }

// Ensuring interfaces in compile-time for PhoneCallEmpty.
var (
	_ bin.Encoder = &PhoneCallEmpty{}
	_ bin.Decoder = &PhoneCallEmpty{}

	_ PhoneCallClass = &PhoneCallEmpty{}
)

// PhoneCallWaiting represents TL type `phoneCallWaiting#1b8f4ad1`.
// Incoming phone call
//
// See https://core.telegram.org/constructor/phoneCallWaiting for reference.
type PhoneCallWaiting struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Is this a video call
	Video bool
	// Call ID
	ID int64
	// Access hash
	AccessHash int64
	// Date
	Date int
	// Admin ID
	AdminID int
	// Participant ID
	ParticipantID int
	// Phone call protocol info
	Protocol PhoneCallProtocol
	// When was the phone call received
	//
	// Use SetReceiveDate and GetReceiveDate helpers.
	ReceiveDate int
}

// PhoneCallWaitingTypeID is TL type id of PhoneCallWaiting.
const PhoneCallWaitingTypeID = 0x1b8f4ad1

// Encode implements bin.Encoder.
func (p *PhoneCallWaiting) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneCallWaiting#1b8f4ad1 as nil")
	}
	b.PutID(PhoneCallWaitingTypeID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phoneCallWaiting#1b8f4ad1: field flags: %w", err)
	}
	b.PutLong(p.ID)
	b.PutLong(p.AccessHash)
	b.PutInt(p.Date)
	b.PutInt(p.AdminID)
	b.PutInt(p.ParticipantID)
	if err := p.Protocol.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phoneCallWaiting#1b8f4ad1: field protocol: %w", err)
	}
	if p.Flags.Has(0) {
		b.PutInt(p.ReceiveDate)
	}
	return nil
}

// SetVideo sets value of Video conditional field.
func (p *PhoneCallWaiting) SetVideo(value bool) {
	if value {
		p.Flags.Set(6)
	} else {
		p.Flags.Unset(6)
	}
}

// SetReceiveDate sets value of ReceiveDate conditional field.
func (p *PhoneCallWaiting) SetReceiveDate(value int) {
	p.Flags.Set(0)
	p.ReceiveDate = value
}

// GetReceiveDate returns value of ReceiveDate conditional field and
// boolean which is true if field was set.
func (p *PhoneCallWaiting) GetReceiveDate() (value int, ok bool) {
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.ReceiveDate, true
}

// Decode implements bin.Decoder.
func (p *PhoneCallWaiting) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneCallWaiting#1b8f4ad1 to nil")
	}
	if err := b.ConsumeID(PhoneCallWaitingTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneCallWaiting#1b8f4ad1: %w", err)
	}

	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phoneCallWaiting#1b8f4ad1: field flags: %w", err)
		}
	}
	p.Video = p.Flags.Has(6)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallWaiting#1b8f4ad1: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallWaiting#1b8f4ad1: field access_hash: %w", err)
		}
		p.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallWaiting#1b8f4ad1: field date: %w", err)
		}
		p.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallWaiting#1b8f4ad1: field admin_id: %w", err)
		}
		p.AdminID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallWaiting#1b8f4ad1: field participant_id: %w", err)
		}
		p.ParticipantID = value
	}
	{
		if err := p.Protocol.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phoneCallWaiting#1b8f4ad1: field protocol: %w", err)
		}
	}
	if p.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallWaiting#1b8f4ad1: field receive_date: %w", err)
		}
		p.ReceiveDate = value
	}
	return nil
}

// construct implements constructor of PhoneCallClass.
func (p PhoneCallWaiting) construct() PhoneCallClass { return &p }

// Ensuring interfaces in compile-time for PhoneCallWaiting.
var (
	_ bin.Encoder = &PhoneCallWaiting{}
	_ bin.Decoder = &PhoneCallWaiting{}

	_ PhoneCallClass = &PhoneCallWaiting{}
)

// PhoneCallRequested represents TL type `phoneCallRequested#87eabb53`.
// Requested phone call
//
// See https://core.telegram.org/constructor/phoneCallRequested for reference.
type PhoneCallRequested struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Whether this is a video call
	Video bool
	// Phone call ID
	ID int64
	// Access hash
	AccessHash int64
	// When was the phone call created
	Date int
	// ID of the creator of the phone call
	AdminID int
	// ID of the other participant of the phone call
	ParticipantID int
	// Parameter for key exchange
	GAHash []byte
	// Call protocol info to be passed to libtgvoip
	Protocol PhoneCallProtocol
}

// PhoneCallRequestedTypeID is TL type id of PhoneCallRequested.
const PhoneCallRequestedTypeID = 0x87eabb53

// Encode implements bin.Encoder.
func (p *PhoneCallRequested) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneCallRequested#87eabb53 as nil")
	}
	b.PutID(PhoneCallRequestedTypeID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phoneCallRequested#87eabb53: field flags: %w", err)
	}
	b.PutLong(p.ID)
	b.PutLong(p.AccessHash)
	b.PutInt(p.Date)
	b.PutInt(p.AdminID)
	b.PutInt(p.ParticipantID)
	b.PutBytes(p.GAHash)
	if err := p.Protocol.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phoneCallRequested#87eabb53: field protocol: %w", err)
	}
	return nil
}

// SetVideo sets value of Video conditional field.
func (p *PhoneCallRequested) SetVideo(value bool) {
	if value {
		p.Flags.Set(6)
	} else {
		p.Flags.Unset(6)
	}
}

// Decode implements bin.Decoder.
func (p *PhoneCallRequested) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneCallRequested#87eabb53 to nil")
	}
	if err := b.ConsumeID(PhoneCallRequestedTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneCallRequested#87eabb53: %w", err)
	}

	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phoneCallRequested#87eabb53: field flags: %w", err)
		}
	}
	p.Video = p.Flags.Has(6)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallRequested#87eabb53: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallRequested#87eabb53: field access_hash: %w", err)
		}
		p.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallRequested#87eabb53: field date: %w", err)
		}
		p.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallRequested#87eabb53: field admin_id: %w", err)
		}
		p.AdminID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallRequested#87eabb53: field participant_id: %w", err)
		}
		p.ParticipantID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallRequested#87eabb53: field g_a_hash: %w", err)
		}
		p.GAHash = value
	}
	{
		if err := p.Protocol.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phoneCallRequested#87eabb53: field protocol: %w", err)
		}
	}
	return nil
}

// construct implements constructor of PhoneCallClass.
func (p PhoneCallRequested) construct() PhoneCallClass { return &p }

// Ensuring interfaces in compile-time for PhoneCallRequested.
var (
	_ bin.Encoder = &PhoneCallRequested{}
	_ bin.Decoder = &PhoneCallRequested{}

	_ PhoneCallClass = &PhoneCallRequested{}
)

// PhoneCallAccepted represents TL type `phoneCallAccepted#997c454a`.
// An accepted phone call
//
// See https://core.telegram.org/constructor/phoneCallAccepted for reference.
type PhoneCallAccepted struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Whether this is a video call
	Video bool
	// ID of accepted phone call
	ID int64
	// Access hash of phone call
	AccessHash int64
	// When was the call accepted
	Date int
	// ID of the call creator
	AdminID int
	// ID of the other user in the call
	ParticipantID int
	// B parameter for secure E2E phone call key exchange
	GB []byte
	// Protocol to use for phone call
	Protocol PhoneCallProtocol
}

// PhoneCallAcceptedTypeID is TL type id of PhoneCallAccepted.
const PhoneCallAcceptedTypeID = 0x997c454a

// Encode implements bin.Encoder.
func (p *PhoneCallAccepted) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneCallAccepted#997c454a as nil")
	}
	b.PutID(PhoneCallAcceptedTypeID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phoneCallAccepted#997c454a: field flags: %w", err)
	}
	b.PutLong(p.ID)
	b.PutLong(p.AccessHash)
	b.PutInt(p.Date)
	b.PutInt(p.AdminID)
	b.PutInt(p.ParticipantID)
	b.PutBytes(p.GB)
	if err := p.Protocol.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phoneCallAccepted#997c454a: field protocol: %w", err)
	}
	return nil
}

// SetVideo sets value of Video conditional field.
func (p *PhoneCallAccepted) SetVideo(value bool) {
	if value {
		p.Flags.Set(6)
	} else {
		p.Flags.Unset(6)
	}
}

// Decode implements bin.Decoder.
func (p *PhoneCallAccepted) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneCallAccepted#997c454a to nil")
	}
	if err := b.ConsumeID(PhoneCallAcceptedTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneCallAccepted#997c454a: %w", err)
	}

	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phoneCallAccepted#997c454a: field flags: %w", err)
		}
	}
	p.Video = p.Flags.Has(6)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallAccepted#997c454a: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallAccepted#997c454a: field access_hash: %w", err)
		}
		p.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallAccepted#997c454a: field date: %w", err)
		}
		p.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallAccepted#997c454a: field admin_id: %w", err)
		}
		p.AdminID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallAccepted#997c454a: field participant_id: %w", err)
		}
		p.ParticipantID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallAccepted#997c454a: field g_b: %w", err)
		}
		p.GB = value
	}
	{
		if err := p.Protocol.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phoneCallAccepted#997c454a: field protocol: %w", err)
		}
	}
	return nil
}

// construct implements constructor of PhoneCallClass.
func (p PhoneCallAccepted) construct() PhoneCallClass { return &p }

// Ensuring interfaces in compile-time for PhoneCallAccepted.
var (
	_ bin.Encoder = &PhoneCallAccepted{}
	_ bin.Decoder = &PhoneCallAccepted{}

	_ PhoneCallClass = &PhoneCallAccepted{}
)

// PhoneCall represents TL type `phoneCall#8742ae7f`.
// Phone call
//
// See https://core.telegram.org/constructor/phoneCall for reference.
type PhoneCall struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Whether P2P connection to the other peer is allowed
	P2PAllowed bool
	// Whether this is a video call
	Video bool
	// Call ID
	ID int64
	// Access hash
	AccessHash int64
	// Date of creation of the call
	Date int
	// User ID of the creator of the call
	AdminID int
	// User ID of the other participant in the call
	ParticipantID int
	// Parameter for key exchange
	GAOrB []byte
	// Key fingerprint
	KeyFingerprint int64
	// Call protocol info to be passed to libtgvoip
	Protocol PhoneCallProtocol
	// List of endpoints the user can connect to to exchange call data
	Connections []PhoneConnectionClass
	// When was the call actually started
	StartDate int
}

// PhoneCallTypeID is TL type id of PhoneCall.
const PhoneCallTypeID = 0x8742ae7f

// Encode implements bin.Encoder.
func (p *PhoneCall) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneCall#8742ae7f as nil")
	}
	b.PutID(PhoneCallTypeID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phoneCall#8742ae7f: field flags: %w", err)
	}
	b.PutLong(p.ID)
	b.PutLong(p.AccessHash)
	b.PutInt(p.Date)
	b.PutInt(p.AdminID)
	b.PutInt(p.ParticipantID)
	b.PutBytes(p.GAOrB)
	b.PutLong(p.KeyFingerprint)
	if err := p.Protocol.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phoneCall#8742ae7f: field protocol: %w", err)
	}
	b.PutVectorHeader(len(p.Connections))
	for idx, v := range p.Connections {
		if v == nil {
			return fmt.Errorf("unable to encode phoneCall#8742ae7f: field connections element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode phoneCall#8742ae7f: field connections element with index %d: %w", idx, err)
		}
	}
	b.PutInt(p.StartDate)
	return nil
}

// SetP2PAllowed sets value of P2PAllowed conditional field.
func (p *PhoneCall) SetP2PAllowed(value bool) {
	if value {
		p.Flags.Set(5)
	} else {
		p.Flags.Unset(5)
	}
}

// SetVideo sets value of Video conditional field.
func (p *PhoneCall) SetVideo(value bool) {
	if value {
		p.Flags.Set(6)
	} else {
		p.Flags.Unset(6)
	}
}

// Decode implements bin.Decoder.
func (p *PhoneCall) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneCall#8742ae7f to nil")
	}
	if err := b.ConsumeID(PhoneCallTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneCall#8742ae7f: %w", err)
	}

	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phoneCall#8742ae7f: field flags: %w", err)
		}
	}
	p.P2PAllowed = p.Flags.Has(5)
	p.Video = p.Flags.Has(6)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCall#8742ae7f: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCall#8742ae7f: field access_hash: %w", err)
		}
		p.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCall#8742ae7f: field date: %w", err)
		}
		p.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCall#8742ae7f: field admin_id: %w", err)
		}
		p.AdminID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCall#8742ae7f: field participant_id: %w", err)
		}
		p.ParticipantID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCall#8742ae7f: field g_a_or_b: %w", err)
		}
		p.GAOrB = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCall#8742ae7f: field key_fingerprint: %w", err)
		}
		p.KeyFingerprint = value
	}
	{
		if err := p.Protocol.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phoneCall#8742ae7f: field protocol: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCall#8742ae7f: field connections: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePhoneConnection(b)
			if err != nil {
				return fmt.Errorf("unable to decode phoneCall#8742ae7f: field connections: %w", err)
			}
			p.Connections = append(p.Connections, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCall#8742ae7f: field start_date: %w", err)
		}
		p.StartDate = value
	}
	return nil
}

// construct implements constructor of PhoneCallClass.
func (p PhoneCall) construct() PhoneCallClass { return &p }

// Ensuring interfaces in compile-time for PhoneCall.
var (
	_ bin.Encoder = &PhoneCall{}
	_ bin.Decoder = &PhoneCall{}

	_ PhoneCallClass = &PhoneCall{}
)

// PhoneCallDiscarded represents TL type `phoneCallDiscarded#50ca4de1`.
// Indicates a discarded phone call
//
// See https://core.telegram.org/constructor/phoneCallDiscarded for reference.
type PhoneCallDiscarded struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Whether the server required the user to rate the call
	NeedRating bool
	// Whether the server required the client to send the libtgvoip call debug data
	NeedDebug bool
	// Whether the call was a video call
	Video bool
	// Call ID
	ID int64
	// Why was the phone call discarded
	//
	// Use SetReason and GetReason helpers.
	Reason PhoneCallDiscardReasonClass
	// Duration of the phone call in seconds
	//
	// Use SetDuration and GetDuration helpers.
	Duration int
}

// PhoneCallDiscardedTypeID is TL type id of PhoneCallDiscarded.
const PhoneCallDiscardedTypeID = 0x50ca4de1

// Encode implements bin.Encoder.
func (p *PhoneCallDiscarded) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneCallDiscarded#50ca4de1 as nil")
	}
	b.PutID(PhoneCallDiscardedTypeID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phoneCallDiscarded#50ca4de1: field flags: %w", err)
	}
	b.PutLong(p.ID)
	if p.Flags.Has(0) {
		if p.Reason == nil {
			return fmt.Errorf("unable to encode phoneCallDiscarded#50ca4de1: field reason is nil")
		}
		if err := p.Reason.Encode(b); err != nil {
			return fmt.Errorf("unable to encode phoneCallDiscarded#50ca4de1: field reason: %w", err)
		}
	}
	if p.Flags.Has(1) {
		b.PutInt(p.Duration)
	}
	return nil
}

// SetNeedRating sets value of NeedRating conditional field.
func (p *PhoneCallDiscarded) SetNeedRating(value bool) {
	if value {
		p.Flags.Set(2)
	} else {
		p.Flags.Unset(2)
	}
}

// SetNeedDebug sets value of NeedDebug conditional field.
func (p *PhoneCallDiscarded) SetNeedDebug(value bool) {
	if value {
		p.Flags.Set(3)
	} else {
		p.Flags.Unset(3)
	}
}

// SetVideo sets value of Video conditional field.
func (p *PhoneCallDiscarded) SetVideo(value bool) {
	if value {
		p.Flags.Set(6)
	} else {
		p.Flags.Unset(6)
	}
}

// SetReason sets value of Reason conditional field.
func (p *PhoneCallDiscarded) SetReason(value PhoneCallDiscardReasonClass) {
	p.Flags.Set(0)
	p.Reason = value
}

// GetReason returns value of Reason conditional field and
// boolean which is true if field was set.
func (p *PhoneCallDiscarded) GetReason() (value PhoneCallDiscardReasonClass, ok bool) {
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.Reason, true
}

// SetDuration sets value of Duration conditional field.
func (p *PhoneCallDiscarded) SetDuration(value int) {
	p.Flags.Set(1)
	p.Duration = value
}

// GetDuration returns value of Duration conditional field and
// boolean which is true if field was set.
func (p *PhoneCallDiscarded) GetDuration() (value int, ok bool) {
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.Duration, true
}

// Decode implements bin.Decoder.
func (p *PhoneCallDiscarded) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneCallDiscarded#50ca4de1 to nil")
	}
	if err := b.ConsumeID(PhoneCallDiscardedTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneCallDiscarded#50ca4de1: %w", err)
	}

	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phoneCallDiscarded#50ca4de1: field flags: %w", err)
		}
	}
	p.NeedRating = p.Flags.Has(2)
	p.NeedDebug = p.Flags.Has(3)
	p.Video = p.Flags.Has(6)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallDiscarded#50ca4de1: field id: %w", err)
		}
		p.ID = value
	}
	if p.Flags.Has(0) {
		value, err := DecodePhoneCallDiscardReason(b)
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallDiscarded#50ca4de1: field reason: %w", err)
		}
		p.Reason = value
	}
	if p.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallDiscarded#50ca4de1: field duration: %w", err)
		}
		p.Duration = value
	}
	return nil
}

// construct implements constructor of PhoneCallClass.
func (p PhoneCallDiscarded) construct() PhoneCallClass { return &p }

// Ensuring interfaces in compile-time for PhoneCallDiscarded.
var (
	_ bin.Encoder = &PhoneCallDiscarded{}
	_ bin.Decoder = &PhoneCallDiscarded{}

	_ PhoneCallClass = &PhoneCallDiscarded{}
)

// PhoneCallClass represents PhoneCall generic type.
//
// See https://core.telegram.org/type/PhoneCall for reference.
//
// Example:
//  g, err := DecodePhoneCall(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *PhoneCallEmpty: // phoneCallEmpty#5366c915
//  case *PhoneCallWaiting: // phoneCallWaiting#1b8f4ad1
//  case *PhoneCallRequested: // phoneCallRequested#87eabb53
//  case *PhoneCallAccepted: // phoneCallAccepted#997c454a
//  case *PhoneCall: // phoneCall#8742ae7f
//  case *PhoneCallDiscarded: // phoneCallDiscarded#50ca4de1
//  default: panic(v)
//  }
type PhoneCallClass interface {
	bin.Encoder
	bin.Decoder
	construct() PhoneCallClass
}

// DecodePhoneCall implements binary de-serialization for PhoneCallClass.
func DecodePhoneCall(buf *bin.Buffer) (PhoneCallClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PhoneCallEmptyTypeID:
		// Decoding phoneCallEmpty#5366c915.
		v := PhoneCallEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhoneCallClass: %w", err)
		}
		return &v, nil
	case PhoneCallWaitingTypeID:
		// Decoding phoneCallWaiting#1b8f4ad1.
		v := PhoneCallWaiting{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhoneCallClass: %w", err)
		}
		return &v, nil
	case PhoneCallRequestedTypeID:
		// Decoding phoneCallRequested#87eabb53.
		v := PhoneCallRequested{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhoneCallClass: %w", err)
		}
		return &v, nil
	case PhoneCallAcceptedTypeID:
		// Decoding phoneCallAccepted#997c454a.
		v := PhoneCallAccepted{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhoneCallClass: %w", err)
		}
		return &v, nil
	case PhoneCallTypeID:
		// Decoding phoneCall#8742ae7f.
		v := PhoneCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhoneCallClass: %w", err)
		}
		return &v, nil
	case PhoneCallDiscardedTypeID:
		// Decoding phoneCallDiscarded#50ca4de1.
		v := PhoneCallDiscarded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhoneCallClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PhoneCallClass: %w", bin.NewUnexpectedID(id))
	}
}

// PhoneCall boxes the PhoneCallClass providing a helper.
type PhoneCallBox struct {
	PhoneCall PhoneCallClass
}

// Decode implements bin.Decoder for PhoneCallBox.
func (b *PhoneCallBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PhoneCallBox to nil")
	}
	v, err := DecodePhoneCall(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PhoneCall = v
	return nil
}

// Encode implements bin.Encode for PhoneCallBox.
func (b *PhoneCallBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PhoneCall == nil {
		return fmt.Errorf("unable to encode PhoneCallClass as nil")
	}
	return b.PhoneCall.Encode(buf)
}
