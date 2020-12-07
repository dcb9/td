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

// AccountReportPeerRequest represents TL type `account.reportPeer#ae189d5f`.
// Report a peer for violation of telegram's Terms of Service
//
// See https://core.telegram.org/method/account.reportPeer for reference.
type AccountReportPeerRequest struct {
	// The peer to report
	Peer InputPeerClass
	// The reason why this peer is being reported
	Reason ReportReasonClass
}

// AccountReportPeerRequestTypeID is TL type id of AccountReportPeerRequest.
const AccountReportPeerRequestTypeID = 0xae189d5f

// Encode implements bin.Encoder.
func (r *AccountReportPeerRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.reportPeer#ae189d5f as nil")
	}
	b.PutID(AccountReportPeerRequestTypeID)
	if r.Peer == nil {
		return fmt.Errorf("unable to encode account.reportPeer#ae189d5f: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.reportPeer#ae189d5f: field peer: %w", err)
	}
	if r.Reason == nil {
		return fmt.Errorf("unable to encode account.reportPeer#ae189d5f: field reason is nil")
	}
	if err := r.Reason.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.reportPeer#ae189d5f: field reason: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountReportPeerRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.reportPeer#ae189d5f to nil")
	}
	if err := b.ConsumeID(AccountReportPeerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.reportPeer#ae189d5f: %w", err)
	}

	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.reportPeer#ae189d5f: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		value, err := DecodeReportReason(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.reportPeer#ae189d5f: field reason: %w", err)
		}
		r.Reason = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountReportPeerRequest.
var (
	_ bin.Encoder = &AccountReportPeerRequest{}
	_ bin.Decoder = &AccountReportPeerRequest{}
)

// AccountReportPeer invokes method account.reportPeer#ae189d5f returning error if any.
// Report a peer for violation of telegram's Terms of Service
//
// See https://core.telegram.org/method/account.reportPeer for reference.
func (c *Client) AccountReportPeer(ctx context.Context, request *AccountReportPeerRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
