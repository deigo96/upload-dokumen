package main

import (
	"desa-sragen/api/router"
	"desa-sragen/bootrstrap"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	app := bootrstrap.App()

	db := app.Db
	defer app.CloseDBConnection()

	r := gin.Default()
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
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(200, "page-404.html", gin.H{
		
		})
	})
	router.Setup(app.Env, *db, route)

	r.Run(app.Env.ServerHost+":"+app.Env.ServerPort)
}