package main

import (
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

	carInspection := inspection.ReguralCarInspectionRequest{
		Name:              "Wymiana oleju",
		TotalPrice:        350,
		CarMilage:         190000,
		DateInspectionCar: "2022-04-03",
	}

	kupa := inspectionService.CreateRegularCarInspection(carInspection)
	if kupa != nil {
		log.Fatalln(kupa)
	}
}
