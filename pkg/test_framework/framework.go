package test_framework

import (
	"github.com/thatched-cottage/utils-framework/pkg/log"
	"runtime/debug"
)

type TestFrameworkFunc func(...interface{})

func TCase(testFF TestFrameworkFunc, args ...interface{}) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		log.Debugf("e")
		if err := recover(); err != nil {
			log.Debugf("error:%v", err) // 这里的err其实就是panic传入的内容
			log.Debugf("[PANIC] panic = %s, stack = %s", err, string(debug.Stack()))
		}
		log.Debugf("f")
	}()

	testFF(args)
}
