package inspection

import "time"
import "gorm.io/gorm"

type ReguralCarInspection struct {
	gorm.Model
	Name              string    `gorm:"name"`
	DateInspectionCar time.Time `gorm:"date_inspection_car"`
	CarMilage         int       `gorm:"car_milage"`
	NextCarMilage     int       `gorm:"next_car_milage"` //nastepny sugerowany przeglad !!! 10 tys dla bezyny 15 tys dla dizla edytowalne w ustawieniach sumowac z CarMilage
	Description       string    `gorm:"description"`
	TotalPrice        float32   `gorm:"total_price"`
	//zdjecia paragonu
	//struktura tablica z wymienionymi podzespoalmi
}
