package main

import (
	"fmt"
	"log"

	inspectController "github.com/lewy9109/autoNotes/pkg-Inspection/controller"
	"github.com/lewy9109/autoNotes/pkg-Inspection/inspection"
	"github.com/lewy9109/autoNotes/pkg-User/controller/userController"
	"github.com/lewy9109/autoNotes/pkg-User/user"

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

	err = db.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal(err)
	}

	startHttpServer(db)
}

func startHttpServer(db *gorm.DB) {

	inspectionRepo := inspection.GetInceptionRepository(db)
	inspectionService := inspection.GetInceptionSercvice(inspectionRepo)
	inspectCarController := inspectController.GetInspectionControllerInterface(inspectionService)

	userInfra := user.DefaultUserInfraStructure(db)
	userService := user.DefalutUserService(userInfra, "secretToken")
	userServer := userController.DefalutUserServer(userService)

	router := gin.Default()
	fmt.Println("Starting HTTP on port 8080 ...")

	groupUser := router.Group("/user/", userServer.Authorize)
	{
		groupUser.GET("/", userServer.GetInfo)
	}

	router.POST("/users", userServer.CreateUser)
	router.POST("/login", userServer.LoginUser)

	inspectCarGroup := router.Group("/inspect", userServer.Authorize)
	{
		inspectCarGroup.POST("/", inspectCarController.CreateInseption)
		inspectCarGroup.GET("/", inspectCarController.GetListInspections)
		inspectCarGroup.GET("/:id", inspectCarController.GetInspectionById)
	}

	err := router.Run(":9889")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
