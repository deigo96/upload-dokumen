package router

import (
	"desa-sragen/api/controller"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	cx := &controller.Controller{}

	// adminRoute := router.Group("admin")
	publicRoute := router.Group("")

	publicRoute.GET("/", cx.Dashboard)
	publicRoute.GET("/login", cx.Login)
	publicRoute.GET("/register", cx.Register)
	publicRoute.GET("/loginHandler", cx.LoginHandler)

	publicRoute.GET("/dashboard", cx.AdminDashboard)
	publicRoute.GET("/surat-domisili", cx.SuratDomisili)
	publicRoute.GET("/detail-pengajuan/:id", cx.DetailPengajuan)

	publicRoute.GET("/pengajuan-surat-domisili", cx.PengajuanDomisili)
	publicRoute.GET("/list-dokumen", cx.ListDokumen)
}
