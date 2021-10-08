package main

import (
	"github.com/gin-gonic/gin"
	"github.com/youthlin/pub/handlers"
)

func register(r *gin.Engine) {
	// middlewares
	r.Use(handlers.Trace)
	r.Use(handlers.I18n)

	// route
	r.Any("/ping", handlers.Ping)

}
