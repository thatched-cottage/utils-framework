package monitor

import (
	"context"
	"github.com/thatched-cottage/utils-framework/pkg/log"
	"github.com/thatched-cottage/utils-framework/pkg/monitor/monitor"
	"time"
)

func Init() {
	monitor.Init(monitor.All, OutStringHook, 1*time.Second)
}

func OutStringHook(s string) {
	log.Debugf(s)
}

func TimeCost(ctx context.Context) func() {
	return monitor.M.TimeCost(ctx, 1)
}
