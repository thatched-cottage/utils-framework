package monitor

import "context"

type iMonitor interface {
	TimeCost(ctx context.Context, skip int) func()
}

var M iMonitor
