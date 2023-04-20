package ginx

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type HandlerFunc func(c *Context) error

func W(h HandlerFunc) gin.HandlerFunc {
	handler := func(c *gin.Context) {
		ctx := &Context{
			Context: c,
		}
		if err := h(ctx); err != nil {
			log.Errorln(err)
			_ = ctx.Error(err)
			c.Abort()
		} else {
			c.Next()
		}
	}
	return handler
}
