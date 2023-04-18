package controller

import (
	"github/lewy9109/autoNotes/inspection"
	insopectService "github/lewy9109/autoNotes/inspection"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type InspectionControllerInterface interface {
	CreateInseption(c *gin.Context)
	GetInspectionById(c *gin.Context)
	GetListInspections(c *gin.Context)
}

type inspectionController struct {
	service inspection.InseptionServceInterface
}

func GetInspectionControllerInterface(service inspection.InseptionServceInterface) InspectionControllerInterface {
	return &inspectionController{service: service}
}

func (ic *inspectionController) CreateInseption(c *gin.Context) {
	inspect := CreateInseptionRequest{}

	err := c.BindJSON(&inspect)
	if err != nil {
		c.JSON(http.StatusBadRequest, CreateInseptionResponse{})
		return
	}
	const shortForm = "2006-01-02"
	dateInception, _ := time.Parse(shortForm, "2022-04-09")

	createInspect := insopectService.ReguralCarInspection{
		Name:              inspect.Name,
		TotalPrice:        inspect.TotalPrice,
		Description:       inspect.Description,
		DateInspectionCar: dateInception,
		CarMilage:         inspect.CarMilage,
	}

	err = ic.service.CreateRegularCarInspection(createInspect)

	if err != nil {
		c.JSON(http.StatusBadRequest, CreateInseptionResponse{})
		return
	}

	c.JSON(http.StatusCreated, inspect)
}
func (ic *inspectionController) GetInspectionById(c *gin.Context) {
	// id := c.Param("id")

	c.JSON(http.StatusBadRequest, CreateInseptionResponse{})
	return
}
func (ic *inspectionController) GetListInspections(c *gin.Context) {
	result, err := ic.service.GetListRegularCarInceptions()

	if err != nil {
		c.JSON(http.StatusInternalServerError, GetInspectionResponse{})
		return
	}

	c.JSON(http.StatusOK, result)
}
