package Controllers

import (
	"net/http"
	"strings"

	"github.com/ArminaFarhad/DashboardManagment/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DashboardController struct {
	DB *gorm.DB
}

func NewDashboardController(DB *gorm.DB) DashboardController {
	return DashboardController{DB}
}

func (DC *DashboardController) GetAllDashboards(ctx *gin.Context) {

	var dashboard []Models.Dashboard
	result := DC.DB.Find(&dashboard)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": dashboard})
}

func (DC *DashboardController) CreateDashboard(ctx *gin.Context) {
	var payload *Models.CreateDashboard

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	newDashboard := Models.CreateDashboard{
		Name: payload.Name,
	}

	result := DC.DB.Create(&newDashboard)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

}

func (DC *DashboardController) GetDashboardWidget(ctx *gin.Context) {
	dashboardId := ctx.Param("dashboardId")

	var dashboardWidgets []Models.DashboardWidget
	result := DC.DB.Find(&dashboardWidgets, "dashboardId =", dashboardId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": dashboardWidgets})
}

func (DC *DashboardController) AddDashboardWidget(ctx *gin.Context) {
	var payload []*Models.DashboardWidget
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	dashboardId := ctx.Param("dashboardId")
	dashboardWidgets := DC.DB.Delete(&Models.DashboardWidget{}, "DashboardId = ?", dashboardId)
	if dashboardWidgets.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	for i := 0; i < len(payload); i++ {
		newWidget := Models.AddDashboardWidget{
			WidgetId:    payload[i].WidgetId,
			DashboardId: payload[i].DashboardId,
			Position:    payload[i].Position,
		}
		result := DC.DB.Create(&newWidget)
		if result.Error != nil {
			if strings.Contains(result.Error.Error(), "duplicate key") {
				ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
				return
			}
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
			return
		}
	}

}
