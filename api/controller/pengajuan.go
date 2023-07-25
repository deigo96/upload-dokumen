package controller

import (
	"desa-sragen/domain"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (cx *Controller) PengajuanDomisili(c *gin.Context) {

	c.HTML(200, "pengajuan-surat-domisili.html", gin.H{})
}

func (cx *Controller) ListDokumen(c *gin.Context) {
	param := c.Query("nik")
	if param == "" {
		c.JSON(http.StatusBadRequest, "nik tidak boleh kosong")
		return
	}

	c.HTML(200, "list-document.html", gin.H{})
}

func (cx *Controller) KirimPengajuan(c *gin.Context)  {
	fileFoto, err := c.FormFile("foto") 
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request file foto: ", err)
			return
	}

	fileKtp, err := c.FormFile("ktp") 
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request file ktp: ", err)
			return
	}

	auth_id := c.PostForm("auth_id")
	nik := c.PostForm("nik")
	phone := c.PostForm("phone")
	jenis_kelamin := c.PostForm("jenis_kelamin")
	agama := c.PostForm("agama")
	pekerjaan := c.PostForm("pekerjaan")
	tempat_lahir := c.PostForm("tempat_lahir")
	tanggal_lahir := c.PostForm("tanggal_lahir")
	alamat := c.PostForm("alamat")
	alasan := c.PostForm("alasan")
	jenis_pengajuan := c.PostForm("jenis_pengajuan")

	extFoto := domain.GenerateUuid() + filepath.Ext(fileFoto.Filename)
	extKtp := domain.GenerateUuid() + filepath.Ext(fileKtp.Filename)

	// Save the uploaded file to a specific location (e.g., ./uploads)
	destination := filepath.Join("./uploads/foto", extFoto)
	if err := c.SaveUploadedFile(fileFoto, destination); err != nil {
		c.String(http.StatusInternalServerError, "Failed to save file")
		return
	}

	destinationKtp := filepath.Join("./uploads/ktp", extKtp)
	if err := c.SaveUploadedFile(fileKtp, destinationKtp); err != nil {
		c.String(http.StatusInternalServerError, "Failed to save file")
		return
	}

	// Parse JSON data from the request body
	

	jsonData := domain.PengajuanParam{
		AuthId: int64(domain.StringToInt(auth_id)),
		Nik: nik,
		Phone: phone,
		JenisKelamin: jenis_kelamin,
		Agama: agama,
		Pekerjaan: pekerjaan,
		TempatLahir: tempat_lahir,
		TanggalLahir: tanggal_lahir,
		Alamat: alamat,
		Alasan: alasan,
		JenisPengajuan: int8(domain.StringToInt(jenis_pengajuan)),
	}

	jsonData.CreatedAt = domain.TimeNow()
	jsonData.Foto = extFoto
	jsonData.Ktp = extKtp

	if err := cx.Repo.CreatePengajuan(jsonData); err != nil {
		c.JSON(500, domain.JsonResponse(500, err.Error(), domain.Empty{}))
		return
	}

	c.JSON(201, domain.JsonResponse(201, "success", domain.Empty{}))
}