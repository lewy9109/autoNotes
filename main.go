package main

import (
	"fmt"
	"github.com/lewy9109/autoNotes/pkg-Inspection/inspection"
	inspectController "github.com/lewy9109/autoNotes/pkg-Inspection/controller"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&inspection.ReguralCarInspection{})
	if err != nil {
		log.Fatal(err)
	}

	inspectionRepo := inspection.GetInceptionRepository(db)
	inspectionService := inspection.GetInceptionSercvice(inspectionRepo)

	startHttpServer(inspectionService)
}

func startHttpServer(inspectionService inspection.InseptionServceInterface) {

	inspectCarController := inspectController.GetInspectionControllerInterface(inspectionService)
	router := gin.Default()
	fmt.Println("Starting HTTP on port 8080 ...")

	inspectCar := router.Group("/inspect")
	{
		inspectCar.POST("/", inspectCarController.CreateInseption)
		inspectCar.GET("/", inspectCarController.GetListInspections)
		inspectCar.GET("/:id", inspectCarController.GetInspectionById)
	}

	err := router.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
