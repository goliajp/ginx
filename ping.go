package ginx

import (
	"time"
)

func Ping(ctx *Context) error {
	return ctx.Success("pong")
}

func Status(runningStartedAt time.Time) func(ctx *Context) error {
	return func(ctx *Context) error {
		var rsp struct {
			Running string `json:"running"`
		}
		rsp.Running = time.Since(runningStartedAt).String()
		return ctx.Success(rsp)
	}
}
