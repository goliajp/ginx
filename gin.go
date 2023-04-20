package ginx

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New(middleware ...gin.HandlerFunc) *gin.Engine {
	e := gin.Default()
	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true
	cfg.AllowFiles = true
	cfg.AddExposeHeaders(
		"Content-Length",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers",
		"Cache-Control",
		"Content-Language",
		"Content-Type",
	)
	cfg.AddAllowHeaders(
		"Authorization",
		"X-Requested-With",
	)
	e.Use(cors.New(cfg))
	e.Use(middleware...)
	e.HandleMethodNotAllowed = true
	return e
}
