package main

import (
	"context"
	"gitee.com/wuxiansheng/utils-framework/pkg/monitor/exampls/monitor"
	"testing"
	"time"
)

func TestNextStdChunk(t *testing.T) {
	monitor.Init()
	monitor.TimeCost(context.Background())()
	monitor.TimeCost(context.Background())()
	monitor.TimeCost(context.Background())()
	foo()
	defer monitor.TimeCost(context.Background())()
	time.Sleep(time.Second * 10)
}

func foo() {
	monitor.TimeCost(context.Background())()
}
