package telegram

import (
	"context"

	"go.uber.org/zap"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tgerr"
)

// Invoke invokes raw MTProto RPC method. It sends input and decodes result
// into output. The request also goes through Middleware from Client’s Options.
func (c *Client) Invoke(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
	return c.invoker.Invoke(ctx, input, output)
}

// invokeDirect directly invokes RPC method without middlewares, automatically
// handling datacenter redirects.
func (c *Client) invokeDirect(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
	if err := c.invokeConn(ctx, input, output); err != nil {
		// Handling datacenter migration request.
		if rpcErr, ok := tgerr.As(err); ok && rpcErr.IsCode(303) {
			targetDC := rpcErr.Argument
			log := c.log.With(
				zap.String("error_type", rpcErr.Type),
				zap.Int("target_dc", targetDC),
			)
			// If migration error is FILE_MIGRATE or STATS_MIGRATE, then the method
			// called by authorized client, so we should try to transfer auth to new DC
			// and create new connection.
			if rpcErr.IsOneOf("FILE_MIGRATE", "STATS_MIGRATE") {
				log.Debug("Invoking on target DC")
				return c.invokeSub(ctx, targetDC, input, output)
			}

			// Otherwise we should change primary DC.
			log.Info("Migrating to target DC")
			return c.invokeMigrate(ctx, targetDC, input, output)
		}

		return err
	}

	return nil
}

// directInvoker implements tg.Invoker on Client for invoking methods directly,
// without middlewares.
type directInvoker struct {
	client *Client
}

// Invoke sends input and decodes result into output.
//
// NOTE: Assuming that call contains content message (seqno increment).
func (d directInvoker) Invoke(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
	return d.client.invokeDirect(ctx, input, output)
}

// invokeConn directly invokes RPC call on primary connection without any
// additional handling.
func (c *Client) invokeConn(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
	c.connMux.Lock()
	conn := c.conn
	c.connMux.Unlock()

	return conn.Invoke(ctx, input, output)
}
