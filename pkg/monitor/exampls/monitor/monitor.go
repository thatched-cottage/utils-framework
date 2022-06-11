package monitor

import (
	"context"
	"gitee.com/wuxiansheng/utils-framework/pkg/monitor/monitor"
	"log"
	"time"
)

func Init() {
	monitor.Init(monitor.All, OutStringHook, 1*time.Second)
}

func OutStringHook(s string) {
	log.Print(s)
}

func TimeCost(ctx context.Context) func() {
	return monitor.M.TimeCost(ctx, 1)
}
