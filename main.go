package main

import (
	"contact-list/api/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userHandler := user.Init()

	r.GET("/contact/:id", userHandler.Get)
	r.GET("/contacts", userHandler.Gets)
	r.POST("/contact", userHandler.Create)
	r.DELETE("/contact", userHandler.Delete)
	r.PUT("/contact", userHandler.Update)

	r.Run()
}
