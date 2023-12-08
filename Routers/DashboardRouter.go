package Routers

import (
	"github.com/ArminaFarhad/DashboardManagment/Controllers"
	"github.com/gin-gonic/gin"
)

type DashboardRouteController struct {
	dashboardController Controllers.DashboardController
}

func NewRouteDashboardController(dashboardController Controllers.DashboardController) DashboardRouteController {
	return DashboardRouteController{dashboardController}
}

func (dc *DashboardRouteController) DashboardRoute(rg *gin.RouterGroup) {

	router := rg.Group("posts")
	router.POST("/", dc.dashboardController.CreateDashboard)
	router.GET("/", dc.dashboardController.GetAllDashboards)
	router.POST("/", dc.dashboardController.AddDashboardWidget)
	router.GET("/:dashboardId", dc.dashboardController.GetDashboardWidget)
}
