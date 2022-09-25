package wait

import (
	"context"
	"time"
)

type T interface {
	Fail()
}

func makeCtx(timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	return ctx, cancel
}

func For1s() (context.Context, context.CancelFunc) {
	return makeCtx(1 * time.Second)
}

func For5s() (context.Context, context.CancelFunc) {
	return makeCtx(5 * time.Second)
}

func ForCtx(ctx context.Context) ContextFunc {
	return func() (context.Context, context.CancelFunc) {
		return context.WithCancel(ctx)
	}
}

func ForTimeout(timeout time.Duration) ContextFunc {
	return func() (context.Context, context.CancelFunc) {
		return makeCtx(timeout)
	}
}

type Status struct {
	iteration int

	Interval time.Duration
	Complete bool
	Err      error
}

func (s *Status) Iteration() int {
	return s.iteration
}

type ContextFunc func() (context.Context, context.CancelFunc)

type Func func(*Status)

func On(t T, c ContextFunc, f Func) {
	ctx, cancel := c()
	defer cancel()

	s := new(Status)
	_ = ctx
	for {
		s.iteration++
		f(s)
	}
}
