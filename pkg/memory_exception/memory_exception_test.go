package memory_exception

import (
	"context"
	"github.com/thatched-cottage/utils-framework/pkg/log"
	"runtime"
	"testing"
	"time"
)

var ms runtime.MemStats
var onece bool

func TestMemoryException(t *testing.T) {
	Init(context.Background(), "./", 500, 1, true)
	{
		onece = true
		thisBug(&[]byte{})
	}
	var b []byte
	b = append(b, "如果程序OOM了 会做什么操作？ 我们要怎么查？"...)
	for true {
		time.Sleep(time.Second / 2)
		b = append(b, b...)
		foo()
	}
}
func foo() {
	runtime.ReadMemStats(&ms)
	log.Debugf("Sys Memory:%dMB", ms.Sys/MB)
	log.Debugf("StackInuse Memory:%dMB", ms.StackInuse/MB)
	log.Debugf("StackSys Memory:%dMB", ms.StackSys/MB)
	log.Debugf("TotalAlloc Memory:%dMB", ms.TotalAlloc/MB)
	log.Debugf("HeapInuse Memory:%dMB", ms.HeapInuse/MB)
	log.Debugf("HeapAlloc Memory:%dMB", ms.HeapAlloc/MB)
	log.Debugf("HeapObjects Memory:%dMB", ms.HeapObjects/MB)
}

func thisBug(i *[]byte) {
	time.Sleep(time.Microsecond)
	if ms.HeapAlloc < 1024*MB && onece {
		*i = append(*i, "如果程序OOM了 会做什么操作？ 我们要怎么查？"...)
		go thisBug(i)
	} else {
		runtime.GC()
		onece = false
	}
}
