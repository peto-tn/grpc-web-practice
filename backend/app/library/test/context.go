package testutils

import (
	context "context"

	"blog/app/library"
)

func NewTestContext() context.Context {
	ctx := context.Background()
	now := library.TimeNow(ctx)
	return library.SetTimeNow(ctx, now)
}
