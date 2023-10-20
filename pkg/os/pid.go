package os

/***********************************************
作者 ： wuyang
示例演示 os的 使用
函数|用途
-|-
GetPID()|获取进程ID
***********************************************/

import (
	"os"
	"runtime"
)

func GetPID() int {
	runtime.GC()
	return os.Getpid()
}
