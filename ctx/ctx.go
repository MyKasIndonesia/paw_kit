// Package ctx consists of functions related to context utilization.
package ctx

import (
	"context"
	"time"
)

type contextKey string

// BgWithTimeout return context.Background with Timeout
func BgWithTimeout(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d)
}
