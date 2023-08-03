package controller

import (
	"desa-sragen/bootrstrap"
	"desa-sragen/domain"

	"github.com/gin-gonic/gin"
)

type Controller struct{
	Config *bootrstrap.Env
	Repo domain.Repo
}

func (cx *Controller) Env()  string {
	return cx.Config.ServerUrl
}

func (cx *Controller) Dashboard(c *gin.Context) {


	c.HTML(200, "index.html", gin.H{
		"url":cx.Env(),
	})
}

func (cx *Controller) ValidateToken(c *gin.Context)  {
	c.JSON(200, domain.JsonResponse(200, "Validated", domain.Empty{}))
}

func (cx *Controller) ValidateTokenAdmin(c *gin.Context)  {
	c.JSON(200, domain.JsonResponse(200, "Validated", domain.Empty{}))
}

func (cx *Controller) ValidateTokenSuper(c *gin.Context)  {
	c.JSON(200, domain.JsonResponse(200, "Validated", domain.Empty{}))
}