package main

import (
	"desa-sragen/api/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.LoadHTMLGlob("templates/*.html")

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders: []string{"X-Total-Count"},
	}))

	r.Static("/assets", "./templates/assets")
	r.Static("/admin", "./templates/admin")

	route := r.Group("")
	router.Setup(route)

	r.Run()
}