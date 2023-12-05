package Controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ArminaFarhad/DashboardManagment/Models"
	"gorm.io/gorm"
)

type DashboardController struct {
	DB *gorm.DB
}

func NewDashboardController(DB *gorm.DB) DashboardController {
	return DashboardController{DB}
}

func (DC *DashboardController) CreateDashboard(response http.ResponseWriter, request *http.Request) {
	var httpError = Models.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}
	var dashboard Models.GetOrAddDashboard
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&dashboard)
	defer request.Body.Close()

	if err != nil {
		returnErrorResponse(response, request, httpError)
	} else {
		httpError.Code = http.StatusBadRequest
		if dashboard.Name == "" {
			httpError.Message = "Name can't be empty"
			returnErrorResponse(response, request, httpError)
		}
	}

}

func returnErrorResponse(response http.ResponseWriter, request *http.Request, errorMesage Models.ErrorResponse) {
	httpResponse := &Models.ErrorResponse{Code: errorMesage.Code, Message: errorMesage.Message}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errorMesage.Code)
	response.Write(jsonResponse)
}
