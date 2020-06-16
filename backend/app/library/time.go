package library

import (
	"context"
	"time"
)

var (
	fixedTime *time.Time = nil
)

func TimeNow(ctx context.Context) time.Time {
	v := ctx.Value("now")
	if v != nil {
		return v.(time.Time)
	}

	if fixedTime != nil {
		return *fixedTime
	}
	return time.Now()
}

func SetTimeNow(ctx context.Context, now time.Time) context.Context {
	return context.WithValue(ctx, "now", now)
}

func FixTime() {
	now := time.Now()
	fixedTime = &now
}

func FlowTime() {
	fixedTime = nil
}
