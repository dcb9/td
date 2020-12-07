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

// StatsBroadcastStats represents TL type `stats.broadcastStats#bdf78394`.
// Channel statistics.
//
// See https://core.telegram.org/constructor/stats.broadcastStats for reference.
type StatsBroadcastStats struct {
	// Period in consideration
	Period StatsDateRangeDays
	// Follower count change for period in consideration
	Followers StatsAbsValueAndPrev
	// total_viewcount/postcount, for posts posted during the period in consideration (views_per_post). Note that in this case, current refers to the period in consideration (min_date till max_date), and prev refers to the previous period ((min_date - (max_date - min_date)) till min_date).
	ViewsPerPost StatsAbsValueAndPrev
	// total_viewcount/postcount, for posts posted during the period in consideration (views_per_post). Note that in this case, current refers to the period in consideration (min_date till max_date), and prev refers to the previous period ((min_date - (max_date - min_date)) till min_date)
	SharesPerPost StatsAbsValueAndPrev
	// Percentage of subscribers with enabled notifications
	EnabledNotifications StatsPercentValue
	// Channel growth graph (absolute subscriber count)
	GrowthGraph StatsGraphClass
	// Followers growth graph (relative subscriber count)
	FollowersGraph StatsGraphClass
	// Muted users graph (relative)
	MuteGraph StatsGraphClass
	// Views per hour graph (absolute)
	TopHoursGraph StatsGraphClass
	// Interactions graph (absolute)
	InteractionsGraph StatsGraphClass
	// IV interactions graph (absolute)
	IvInteractionsGraph StatsGraphClass
	// Views by source graph (absolute)
	ViewsBySourceGraph StatsGraphClass
	// New followers by source graph (absolute)
	NewFollowersBySourceGraph StatsGraphClass
	// Subscriber language graph (piechart)
	LanguagesGraph StatsGraphClass
	// Recent message interactions
	RecentMessageInteractions []MessageInteractionCounters
}

// StatsBroadcastStatsTypeID is TL type id of StatsBroadcastStats.
const StatsBroadcastStatsTypeID = 0xbdf78394

// Encode implements bin.Encoder.
func (b *StatsBroadcastStats) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode stats.broadcastStats#bdf78394 as nil")
	}
	buf.PutID(StatsBroadcastStatsTypeID)
	if err := b.Period.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field period: %w", err)
	}
	if err := b.Followers.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field followers: %w", err)
	}
	if err := b.ViewsPerPost.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field views_per_post: %w", err)
	}
	if err := b.SharesPerPost.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field shares_per_post: %w", err)
	}
	if err := b.EnabledNotifications.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field enabled_notifications: %w", err)
	}
	if b.GrowthGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field growth_graph is nil")
	}
	if err := b.GrowthGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field growth_graph: %w", err)
	}
	if b.FollowersGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field followers_graph is nil")
	}
	if err := b.FollowersGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field followers_graph: %w", err)
	}
	if b.MuteGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field mute_graph is nil")
	}
	if err := b.MuteGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field mute_graph: %w", err)
	}
	if b.TopHoursGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field top_hours_graph is nil")
	}
	if err := b.TopHoursGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field top_hours_graph: %w", err)
	}
	if b.InteractionsGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field interactions_graph is nil")
	}
	if err := b.InteractionsGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field interactions_graph: %w", err)
	}
	if b.IvInteractionsGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field iv_interactions_graph is nil")
	}
	if err := b.IvInteractionsGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field iv_interactions_graph: %w", err)
	}
	if b.ViewsBySourceGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field views_by_source_graph is nil")
	}
	if err := b.ViewsBySourceGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field views_by_source_graph: %w", err)
	}
	if b.NewFollowersBySourceGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field new_followers_by_source_graph is nil")
	}
	if err := b.NewFollowersBySourceGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field new_followers_by_source_graph: %w", err)
	}
	if b.LanguagesGraph == nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field languages_graph is nil")
	}
	if err := b.LanguagesGraph.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field languages_graph: %w", err)
	}
	buf.PutVectorHeader(len(b.RecentMessageInteractions))
	for idx, v := range b.RecentMessageInteractions {
		if err := v.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode stats.broadcastStats#bdf78394: field recent_message_interactions element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *StatsBroadcastStats) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode stats.broadcastStats#bdf78394 to nil")
	}
	if err := buf.ConsumeID(StatsBroadcastStatsTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: %w", err)
	}

	{
		if err := b.Period.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field period: %w", err)
		}
	}
	{
		if err := b.Followers.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field followers: %w", err)
		}
	}
	{
		if err := b.ViewsPerPost.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field views_per_post: %w", err)
		}
	}
	{
		if err := b.SharesPerPost.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field shares_per_post: %w", err)
		}
	}
	{
		if err := b.EnabledNotifications.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field enabled_notifications: %w", err)
		}
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field growth_graph: %w", err)
		}
		b.GrowthGraph = value
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field followers_graph: %w", err)
		}
		b.FollowersGraph = value
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field mute_graph: %w", err)
		}
		b.MuteGraph = value
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field top_hours_graph: %w", err)
		}
		b.TopHoursGraph = value
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field interactions_graph: %w", err)
		}
		b.InteractionsGraph = value
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field iv_interactions_graph: %w", err)
		}
		b.IvInteractionsGraph = value
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field views_by_source_graph: %w", err)
		}
		b.ViewsBySourceGraph = value
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field new_followers_by_source_graph: %w", err)
		}
		b.NewFollowersBySourceGraph = value
	}
	{
		value, err := DecodeStatsGraph(buf)
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field languages_graph: %w", err)
		}
		b.LanguagesGraph = value
	}
	{
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field recent_message_interactions: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value MessageInteractionCounters
			if err := value.Decode(buf); err != nil {
				return fmt.Errorf("unable to decode stats.broadcastStats#bdf78394: field recent_message_interactions: %w", err)
			}
			b.RecentMessageInteractions = append(b.RecentMessageInteractions, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsBroadcastStats.
var (
	_ bin.Encoder = &StatsBroadcastStats{}
	_ bin.Decoder = &StatsBroadcastStats{}
)
