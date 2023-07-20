package controller

import "github.com/gin-gonic/gin"

func (cx *Controller) AdminDashboard(c *gin.Context) {

	c.HTML(200, "admin-dashboard.html", gin.H{})
}