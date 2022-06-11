package monitor

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

type Monitor struct {
}

var (
	tick *time.Ticker
	m    sync.Map
	buf  []string
)

func mRangeCallBack(key, value any) bool {
	ss := value.(*SyncSlice)
	totalTime := time.Duration(0)
	ssCount := ss.Count()
	if ssCount > 0 {
		sss := ss.Clear()
		for _, v := range sss {
			totalTime += v.(time.Duration)
		}
		buf = append(buf, fmt.Sprintf(`[%s] runs for %d times, with an average running time of %v.`, key, ssCount, totalTime/time.Duration(ssCount)))
	}
	return true
}

func Run() {
	if mFlag&Close != 0 {
		return
	}
	if mFlag&Periodic == 0 {
		return
	}
	go func() {
		for {
			select {
			case <-tick.C:
				m.Range(mRangeCallBack)
				if len(buf) > 0 {
					for _, v := range buf {
						tccb(v)
					}
					buf = buf[:0]
				}
			}
		}
	}()
}

// TimeCost ,The argument skip is the number of stack frames to ascend, with 0 identifying the caller of Caller.
func (this *Monitor) TimeCost(ctx context.Context, skip int) func() {
	if mFlag&Close != 0 {
		return func() {}
	}
	start := time.Now()
	return func() {
		pc, _, lineNo, ok := runtime.Caller(skip)
		if !ok {
			return
		}
		path_name := runtime.FuncForPC(pc).Name()
		_, funcName := filepath.Split(path_name)
		tc := time.Since(start)

		if mFlag&Func != 0 {
			tccb(fmt.Sprintf(`[%s:%d] [time cost = %v]`, funcName, lineNo, tc))
		}
		inter, OK := m.Load(path_name)
		if OK != true {
			inter = &SyncSlice{}
			m.Store(path_name, inter)
		}
		ss := inter.(*SyncSlice)
		ss.Add(tc)
	}
}
