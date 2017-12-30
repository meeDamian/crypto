package scheduler

import (
	"net/http"
	"time"
)

// TODO: everything

type (
	Params struct {
		Timeout         time.Duration
		Unit            time.Duration
		RequestsPerUnit int
	}

	Scheduler struct {
		Params
		List []string
	}
)

func New(p Params) Scheduler {
	return Scheduler{
		Params: p,
		List:   []string{},
	}
}

func (s Scheduler) Schedule(req *http.Request) (res *http.Response, err error) {
	return
}
