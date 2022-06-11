package monitor

import "time"

var mFlag monitorFlag

func Init(l monitorFlag, tcb tcCollBack, t time.Duration) {
	M = &Monitor{}
	mFlag = l
	tccb = tcb
	tick = time.NewTicker(t)
	Run()
}
