package ginx

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

func (c *Context) ParamInt(key string) int {
	sv := c.Param(key)
	iv, err := strconv.Atoi(sv)
	if err != nil {
		c.AbortWithError(err)
	}
	return iv
}

func (c *Context) AbortWithError(err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "error": err})
}
