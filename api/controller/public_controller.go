package controller

import (
	"desa-sragen/domain"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (cx *Controller) Env()  string {
	return "http://localhost:8080/"
}

func (cx *Controller) Dashboard(c *gin.Context) {


	c.HTML(200, "index.html", gin.H{
		"url":cx.Env(),
	})
}

func (cx *Controller) Login(c *gin.Context) {

	c.HTML(200, "login.html", gin.H{
		"url":cx.Env(),
	})
}


func (cx *Controller) Register(c *gin.Context) {


	c.HTML(200, "register.html", gin.H{
		"url":cx.Env(),
	})
}

func (cx *Controller) LoginHandler(c *gin.Context)  {
	var req domain.LoginParam
	if c.ShouldBindJSON(&req) != nil {
		
		c.HTML(200, "login.html", gin.H{
			"error_messages": "data tidak boleh kosong",
		})
	}

	c.HTML(200, "admin_dashboard.html", gin.H{
		"url":cx.Env(),
	})
}