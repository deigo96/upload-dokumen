package controller

import (
	"desa-sragen/domain"

	"github.com/gin-gonic/gin"
)

func (cx *Controller) SuratDomisili(c *gin.Context) {
	param := c.Query("id")
	if param != "" {
	documen, err := cx.Repo.GetPengajuanById(domain.StringToInt(param), 1)
	if err != nil {
		c.HTML(200, "500.html", gin.H{})
		return
	}

		c.HTML(200, "detail-pengajuan.html", gin.H{
			"url": cx.Env(),
			"items": documen,
		})
		return
	}


	data, err := cx.Repo.GetAllPengajuan(1)
	if err != nil {
		c.HTML(200, "500.html", gin.H{})
		return
	}

	c.HTML(200, "surat-domisili.html", gin.H{
		"items": data,
		"url": cx.Env(),
	})
}

func (cx *Controller) SuratKartuKeluarga(c *gin.Context) {
	param := c.Query("id")
	if param != "" {
	documen, err := cx.Repo.GetPengajuanById(domain.StringToInt(param), 3)
	if err != nil {
		c.HTML(200, "500.html", gin.H{})
		return
	}

		c.HTML(200, "detail-kk.html", gin.H{
			"url": cx.Env(),
			"items": documen,
		})
		return
	}


	data, err := cx.Repo.GetAllPengajuan(3)
	if err != nil {
		c.HTML(200, "500.html", gin.H{})
		return
	}

	c.HTML(200, "surat-kartu-keluarga.html", gin.H{
		"items": data,
		"url": cx.Env(),
	})
}

func (cx *Controller) SuratKeteranganUsaha(c *gin.Context) {
	param := c.Query("id")
	if param != "" {
	documen, err := cx.Repo.GetPengajuanById(domain.StringToInt(param), 2)
	if err != nil {
		c.HTML(200, "500.html", gin.H{})
		return
	}

		c.HTML(200, "detail-sku.html", gin.H{
			"url": cx.Env(),
			"items": documen,
		})
		return
	}


	data, err := cx.Repo.GetAllPengajuan(2)
	if err != nil {
		c.HTML(200, "500.html", gin.H{})
		return
	}

	c.HTML(200, "surat-keterangan-usaha.html", gin.H{
		"items": data,
		"url": cx.Env(),
	})
}

func (cx *Controller) SuratKelahiran(c *gin.Context) {
	param := c.Query("id")
	if param != "" {
	documen, err := cx.Repo.GetPengajuanById(domain.StringToInt(param), 6)
	if err != nil {
		c.HTML(600, "500.html", gin.H{})
		return
	}

		c.HTML(200, "detail-kelahiran.html", gin.H{
			"url": cx.Env(),
			"items": documen,
		})
		return
	}


	data, err := cx.Repo.GetAllPengajuan(6)
	if err != nil {
		c.HTML(200, "500.html", gin.H{})
		return
	}

	c.HTML(200, "surat-kelahiran.html", gin.H{
		"items": data,
		"url": cx.Env(),
	})
}

func (cx *Controller) DetailPengajuan(c *gin.Context)  {
	param := c.Param("id")
	if param == "" {

	}

	c.HTML(200, "detail-pengajuan", gin.H{
		
	})
}