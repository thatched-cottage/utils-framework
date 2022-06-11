package pprof

import (
	"gitee.com/wuxiansheng/utils-framework/pkg/log"
	"net/http"
	_ "net/http/pprof"
)

func init() {
	go func() {
		log.Error(http.ListenAndServe(":6060", nil))
	}()
}
