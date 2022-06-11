package monitor

import (
	"sync"
)

type tcCollBack func(string)

var (
	tccb tcCollBack
)

type SyncSlice struct {
	l  sync.Mutex
	is []interface{}
}

func (s *SyncSlice) Add(i interface{}) {
	s.l.Lock()
	defer s.l.Unlock()
	s.is = append(s.is, i)
}

func (s *SyncSlice) Count() int {
	return len(s.is)
}

func (s *SyncSlice) Clear() []interface{} {
	s.l.Lock()
	defer s.l.Unlock()
	newIs := s.is
	s.is = s.is[:0]
	return newIs
}
