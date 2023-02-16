package pprof

import (
	"github.com/thatched-cottage/utils-framework/pkg/log"
	"net/http"
	_ "net/http/pprof"
)

func init() {
	go func() {
		log.Error(http.ListenAndServe(":7070", nil))
	}()
}
