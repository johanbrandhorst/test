package wait

import (
	"context"
	"time"
)

func makeCtx(timeout time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	return ctx
}

func For1s() context.Context {
	return makeCtx(1 * time.Second)
}

func For5s() context.Context {
	return makeCtx(5 * time.Second)
}

func OnFunc(ctx context.Context, f func()) {

}
