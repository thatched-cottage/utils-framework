package pprof

import (
	"gitee.com/wuxiansheng/utils-framework/pkg/log"
	"net/http"
	_ "net/http/pprof"
)

func init() {
	go func() {
		log.Error(http.ListenAndServe(":7070", nil))
	}()
}
