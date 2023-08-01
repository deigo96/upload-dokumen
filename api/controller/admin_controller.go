package controller

import (
	"desa-sragen/domain"
	"net/http"
	"path/filepath"

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
		c.HTML(500, "page-404.html", gin.H{
			"url": cx.Env(),
		})
		return
	}

	c.Header("Authorization", "Bearer "+token)
	redirectUrl := cx.Env()+tipe
	c.Header("Location", redirectUrl)
	c.AbortWithStatus(307)
}

func (cx *Controller) AksiDokumen(c *gin.Context)  {
	updatedBy := c.GetInt("userId")
	status := int8(domain.StringToInt(c.PostForm("status")))
	idPengajuan := domain.StringToInt(c.PostForm("id_pengajuan"))
	alasanPenolakan := c.PostForm("alasan_penolakan")
	var extDoc string

	if updatedBy == 0 {
		c.Redirect(307, "/login")
		return
	}

	if status != 1 && status != 2 {
		c.JSON(400, domain.JsonResponse(400, "status tidak sesuai", domain.Empty{}))
		return
	}

	if idPengajuan == 0{
		c.JSON(400, domain.JsonResponse(400, "id pengajuan tidak sesuai", domain.Empty{}))
		return
	}

	req := domain.Aksi {
		Status: status,
		UpdatedAt: domain.TimeNow(),
		UpdatedBy: updatedBy,
	}

	if status == 1 {
		fileDoc, err := c.FormFile("document")

		if fileDoc == nil || err != nil {
			c.JSON(400, domain.JsonResponse(400, "file dokumen tidak boleh kosong", domain.Empty{}))
			return
		}


		if fileDoc.Size > maxFileSize {
			c.JSON(http.StatusBadRequest, domain.JsonResponse(http.StatusBadRequest, "Ukuran file max 2 MB", domain.Empty{}))
			return
		}

		if !isAllowedExtension(fileDoc.Filename, []string{".pdf"}) {
			c.JSON(http.StatusBadRequest, domain.JsonResponse(http.StatusBadRequest, "Format file harus pdf", domain.Empty{}))
			return
		}

		extDoc = domain.GenerateUuid() + filepath.Ext(fileDoc.Filename)
		destination := filepath.Join("./uploads/dokumen", extDoc)
		if err := c.SaveUploadedFile(fileDoc, destination); err != nil {
			c.JSON(http.StatusInternalServerError, domain.JsonResponse(http.StatusInternalServerError, "Failed to save file", domain.Empty{}))
			return
		}

		req.Dokumen = &extDoc
	}

	if status == 2 {
		req.AlasanPenolakan = &alasanPenolakan
	}

	if err := cx.Repo.UpdateStatus(idPengajuan, req); err != nil {
		c.JSON(500, domain.JsonResponse(500, "Ada kesalahan, mohon dicoba lagi, "+err.Error(), gin.H{}))
		return
	}

	c.JSON(200, domain.JsonResponse(200, "Data berhasil disimpan", domain.Empty{}))
	
}