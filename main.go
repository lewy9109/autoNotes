package main

import (
	"fmt"
	"github/lewy9109/autoNotes/inspection"
	"log"

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

	// carInspection := inspection.ReguralCarInspectionRequest{
	// 	Name:              "dziewiaty oleju",
	// 	TotalPrice:        350,
	// 	CarMilage:         190000,
	// 	DateInspectionCar: "2022-04-09",
	// }

	// err = inspectionService.CreateRegularCarInspection(carInspection)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	result, err := inspectionService.GetListRegularCarInceptions()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)
}
