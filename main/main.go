package main

import (
	// "fmt"
	// "github.com/lewy9109/autoNotes/pkg-User/controller/userController"
	// "github.com/lewy9109/autoNotes/pkg-User/user"
	// "log"

	// "github.com/gin-gonic/gin"
	// "gorm.io/driver/sqlite"
	// "gorm.io/gorm"
)

func main() {
	// fmt.Println("start")
	// db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = db.AutoMigrate(&user.User{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// userInfra := user.DefaultUserInfraStructure(db)
	// userService := user.DefalutUserService(userInfra, "secretToken")

	// userServer := userController.DefalutUserServer(userService)

	// server := gin.Default()

	// group := server.Group("/user/", userServer.Authorize)
	// {
	// 	group.GET("/", userServer.GetInfo)
	// }

	// server.POST("/users", userServer.CreateUser)
	// server.POST("/login", userServer.LoginUser)

	// err = server.Run()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}
