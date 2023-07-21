package controller

import (
	"desa-sragen/domain"

	"github.com/gin-gonic/gin"
)

func (cx *Controller) SuratDomisili(c *gin.Context) {
	param := c.Query("id")
	if param != "" {
		c.HTML(200, "detail-pengajuan.html", gin.H{
	})
		return
	}


	data := []domain.Pengajuan{
		{
			Id: "asfas-asf-fasvzxv",
			Name: "deigo",
			Job: "Backend",
			Age: 24,
			CreatedAt: "2023-06-15",
		},{
			Id: "xvxcv-asf-fasvzxv",
			Name: "jonathan",
			Job: "Devloper",
			Age: 24,
			CreatedAt: "2023-06-15",
		},
		{
			Id: "jgfj-afaetsf-qet",
			Name: "Siahaan",
			Job: "Backend Developer",
			Age: 24,
			CreatedAt: "2023-06-15",
		},
	}

	c.HTML(200, "surat-domisili.html", gin.H{
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