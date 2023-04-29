package ginx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Context) Success(data any, items ...gin.H) error {
	body := gin.H{
		"success": true,
		"data":    data,
	}
	for _, item := range items {
		for k, v := range item {
			body[k] = v
		}
	}
	c.JSON(http.StatusOK, body)
	return nil
}

func (c *Context) Error(err error, items ...gin.H) error {
	body := gin.H{
		"success": false,
		"error":   err.Error(),
	}
	for _, item := range items {
		for k, v := range item {
			body[k] = v
		}
	}
	c.JSON(http.StatusInternalServerError, body)
	return err
}

func (c *Context) Return(data any, err error, items ...gin.H) error {
	if err != nil {
		return c.Error(err, items...)
	}
	return c.Success(data, items...)
}
