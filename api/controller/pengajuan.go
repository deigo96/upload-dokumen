package controller

import (
	"net/http"

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
