package controller

import (
	"desa-sragen/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cx *Controller) AdminDashboard(c *gin.Context) {


	c.HTML(200, "admin-dashboard.html", gin.H{
		"url": cx.Config.ServerUrl,
	})
}

func (cx *Controller) PageAdmin(c *gin.Context)  {
	c.HTML(200, "daftar-admin.html", gin.H{
		"url": cx.Config.ServerUrl,
	})
}

func (cx *Controller) ListUsers(c *gin.Context)  {
	c.HTML(200, "daftar-admin.html", gin.H{
		"url": cx.Config.ServerUrl,
	})
}

func (cx *Controller) GetAllAdmin(c *gin.Context)  {
	res, err := cx.Repo.GetAllAdmin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.JsonResponse(http.StatusInternalServerError, err.Error(), domain.Empty{}))
		return
	}

	admins := []domain.AuthResponse{}

	if len(res) > 0 {
		for _, val := range res {
			admin := domain.AuthResponse{
				Id: val.AuthId,
				Username: val.Username,
				Email: val.Email,
				Role: domain.RoleH(val.RoleId),
				IsActive: domain.StatusActive(val.IsActive),
				CreatedAt: domain.TimeToString(val.CreatedAt),
				UpdatedAt: domain.TimeToString(val.UpdatedAt),
			}

			admins = append(admins, admin)
		}
	}

	c.JSON(200, domain.JsonResponse(200, "Success", admins))
}

func (cx *Controller) TambahAdmin(c *gin.Context)  {
	var req domain.RegisterParam
	if c.ShouldBindJSON(&req) != nil {
		c.JSON(400, "bad request")
		return
	}


	checkUser, err := cx.Repo.GetUser(req.Username)
	if err != nil || checkUser != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse(404, "username sudah digunakan", domain.Empty{}))
		return
	}

	pass := domain.HashAndSalt([]byte("123456"))

	data := domain.Auth{
		Username: req.Username,
		Email: req.Email,
		Password: pass,
		RoleId: 2,
		IsActive: true,
		CreatedAt: domain.TimeNow(),
		UpdatedAt: domain.TimeNow(),
	}

	if err := cx.Repo.CreateUser(data);err != nil {
		c.JSON(http.StatusInternalServerError, domain.JsonResponse(500, err.Error(), domain.Empty{}))
		return
	}

	c.JSON(201, domain.JsonResponse(201, "Success", domain.Empty{}))
}

func (cx *Controller) InActiveAdmin(c *gin.Context)  {
	id := domain.StringToInt(c.Param("id"))
	
	if err := cx.Repo.InActiveAuth(id, false);err != nil {
		c.JSON(http.StatusInternalServerError, domain.JsonResponse(500, err.Error(), domain.Empty{}))
		return
	}

	c.JSON(200, domain.JsonResponse(200, "Success", domain.Empty{}))
}

func (cx *Controller) ActivationAdmin(c *gin.Context)  {
	id := domain.StringToInt(c.Param("id"))
	
	if err := cx.Repo.InActiveAuth(id, true);err != nil {
		c.JSON(http.StatusInternalServerError, domain.JsonResponse(500, err.Error(), domain.Empty{}))
		return
	}

	c.JSON(200, domain.JsonResponse(200, "Success", domain.Empty{}))
}

func (cx *Controller) RedirectUrl(c *gin.Context)  {
	tipe := c.Param("type")
	token := c.Query("token")

	if tipe == "" || token == "" {
		c.HTML(500, "page-404.html", gin.H{})
		return
	}

	c.Header("Authorization", "Bearer "+token)
	redirectUrl := cx.Env()+tipe
	fmt.Print(token)
	fmt.Print(redirectUrl)
	c.Header("Location", redirectUrl)
	c.AbortWithStatus(307)
}