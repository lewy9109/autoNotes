package main

import (
	"fmt"
	"github/lewy9109/autoNotes/inspection"
	inspectController "github/lewy9109/autoNotes/inspection/controller"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db := startDB()

	inspectionRepo := inspection.GetInceptionRepository(db)
	inspectionService := inspection.GetInceptionSercvice(inspectionRepo)

	startHttpServer(inspectionService)
}

func startDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Error connection")

	}
	err = db.AutoMigrate(&inspection.ReguralCarInspection{})
	if err != nil {
		log.Fatal(err)
		panic("Error automigrate")
	}

	return db
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
