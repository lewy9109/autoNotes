package controller

import (
	"fmt"
	"github.com/lewy9109/autoNotes/pkg-Inspection/inspection"
	insopectService "github.com/lewy9109/autoNotes/pkg-Inspection/inspection"
	"net/http"
	"strconv"
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
	dateInception, _ := time.Parse(shortForm, inspect.DateInspectionCar)

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
	id, _ := strconv.Atoi(c.Param("id"))

	result, err := ic.service.GetRegularCarInceptions(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, CreateInseptionResponse{})
	}

	carInspection := GetInspectionResponse{
		Name:              result.Name,
		DateInspectionCar: result.DateInspectionCar.Format("2006-01-02 15:04:05"),
		Description:       result.Description,
		TotalPrice:        result.TotalPrice,
		NextCarMilage:     result.NextCarMilage,
		CarMilage:         result.CarMilage,
	}

	c.JSON(http.StatusOK, carInspection)
}

func (ic *inspectionController) GetListInspections(c *gin.Context) {
	offset, bool := c.GetQuery("offset")
	if !bool {
		fmt.Println("wszedlem kupa")
		offset = "0"
	}
	limit, bool := c.GetQuery("limit")
	if !bool {
		limit = "5"
	}

	oofsetInt, _ := strconv.Atoi(offset)
	limitInt, _ := strconv.Atoi(limit)

	result, err := ic.service.GetListRegularCarInceptions(oofsetInt, limitInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, GetInspectionResponse{})
		return
	}

	c.JSON(http.StatusOK, result)
}
