package monitor

import (
	"context"
	"gitee.com/wuxiansheng/utils-framework/pkg/log"
	"gitee.com/wuxiansheng/utils-framework/pkg/monitor/monitor"
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
