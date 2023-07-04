package main

import (
	"fmt"
	"github.com/lewy9109/autoNotes/pkg-Inspection/inspection"
	inspectController "github.com/lewy9109/autoNotes/pkg-Inspection/controller"
	"github.com/lewy9109/autoNotes/pkg-User/controller/userController"
	"github.com/lewy9109/autoNotes/pkg-User/user"
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


	startHttpServer(db)
}

func startHttpServer(db *gorm.DB) {

	inspectionRepo := inspection.GetInceptionRepository(db)
	inspectionService := inspection.GetInceptionSercvice(inspectionRepo)
	inspectCarController := inspectController.GetInspectionControllerInterface(inspectionService)

	router := gin.Default()
	fmt.Println("Starting HTTP on port 8080 ...")

	inspectCarGroup := router.Group("/inspect")
	{
		inspectCarGroup.POST("/", inspectCarController.CreateInseption)
		inspectCarGroup.GET("/", inspectCarController.GetListInspections)
		inspectCarGroup.GET("/:id", inspectCarController.GetInspectionById)
	}

	userInfra := user.DefaultUserInfraStructure(db)
	userService := user.DefalutUserService(userInfra, "secretToken")

	userServer := userController.DefalutUserServer(userService)


	groupUser := router.Group("/user/", userServer.Authorize)
	{
		groupUser.GET("/", userServer.GetInfo)
	}

	router.POST("/users", userServer.CreateUser)
	router.POST("/login", userServer.LoginUser)

	err := router.Run(":9889")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
