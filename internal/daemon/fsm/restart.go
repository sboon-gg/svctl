package fsm

import (
	"context"
	"time"
)

const maxRestarts = 5

type restartCtx struct {
	count       int
	resetCancel context.CancelFunc
}

func (r *restartCtx) inc() {
	r.count++

	if r.resetCancel != nil {
		r.resetCancel()
	}
	go r.start()
}

func (r *restartCtx) max() bool {
	return r.count >= maxRestarts
}

func (r *restartCtx) start() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	resetCtx, resetCancel := context.WithCancel(context.Background())
	r.resetCancel = resetCancel

	for {
		select {
		case <-resetCtx.Done():
			return
		case <-ctx.Done():
			r.reset()
			return
		}
	}
}

func (r *restartCtx) reset() {
	r.count = 0
	if r.resetCancel != nil {
		r.resetCancel()
	}
	r.resetCancel = nil
}
