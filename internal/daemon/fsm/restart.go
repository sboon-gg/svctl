package fsm

import (
	"context"
	"time"
)

const maxRestarts = 5

type restarter struct {
	numRestarts uint8
	timer       *time.Timer
	cancel      context.CancelFunc
}

func (r *restarter) restartTimer() {
	ctx, cancel := context.WithCancel(context.Background())
	r.cancel = cancel

	if r.timer == nil {
		r.timer = time.NewTimer(time.Minute)
	} else {
		r.cancel()
		r.timer.Stop()
		r.timer.Reset(time.Minute)
	}

	go func() {
		select {
		case <-ctx.Done():
			return
		case <-r.timer.C:
			r.Reset()
		}
	}()
}

func (r *restarter) Increment() {
	r.numRestarts++
	r.restartTimer()
}

func (r *restarter) Reset() {
	if r.timer != nil {
		r.cancel()
		r.timer.Stop()
	}

	r.numRestarts = 0
}

func (r *restarter) LimitReached() bool {
	return r.numRestarts >= maxRestarts
}
