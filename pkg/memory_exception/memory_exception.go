package memory_exception

import (
	"context"
	"github.com/pkg/profile"
	"runtime"
	"time"
)

/*
内存监控 程序内存暴涨之后立即可以查到问题。
实现原理根据:检测内存容量，超过内存阀值（memoryThreshold），同时满足时间间隔（intervalTime）
会输出 go prof 中 Memory alloc 的内容。
*/

var (
	memoryThreshold uint64 // 内存阀值 用于输出的内存限制
	intervalTime    int64  // 时间间隔 用于输出的时间控制
)

const MB = uint64(1024 * 1024)

// Init 服务器启动
func Init(context context.Context, dirName string, memoryThreshold, samplingIntervalTime uint64, memoryExceptionDetection bool) interface{} {
	if memoryExceptionDetection {
		done := make(chan struct{})
		go MemoryException(context, dirName, memoryThreshold, samplingIntervalTime, done)
		return func() {
			close(done)
		}
	}
	return nil
}

// MemoryException 内存分析
func MemoryException(context context.Context, dirName string, mt, sit uint64, done chan struct{}) {
	memoryThreshold = mt * MB
	for {
		c := time.Tick(time.Second)
		select {
		case <-c:
			if OutputRule() {
				intervalTime = time.Now().Unix() + int64(sit)
				filename := time.Now().Format(time.RFC3339)
				profile.Start(profile.MemProfileAllocs, profile.ProfilePath(dirName+filename), profile.NoShutdownHook).Stop()
			}
		case <-done:
			break
		}
	}
}

// OutputRule 输出规则，超过内存阀值（memoryThreshold），同时满足时间间隔（intervalTime） 即满足条件
func OutputRule() bool {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	t := time.Now().Unix()
	if ms.HeapInuse > memoryThreshold &&
		t > intervalTime { // 计算堆使用中的内存即可，栈内存几乎不会产生过多消耗，其中巨块会分配到堆上。
		return true
	}
	return false
}
