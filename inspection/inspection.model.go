package inspection

import "time"

type ReguralCarInspection struct {
	Name              string    `gorm:"name"`
	DateInspectionCar time.Time `gorm:"date_inspection_time"`
	CarMilage         int       `gorm:"car_milage"`
	NextCarMilage     int       `gorm:"next_car_milage"` //nastepny sugerowany przeglad !!! 10 tys dla bezyny 15 tys dla dizla edytowalne w ustawieniach sumowac z CarMilage
	Description       string    `gorm:"description"`
	TotalPrice        float32   `gorm:"total_price"`
	//zdjecia paragonu
	//struktura tablica z wymienionymi podzespoalmi
}

type ReguralCarInspectionRequest struct {
	Name              string  `json:"name"`
	DateInspectionCar string  `json:"dateInspectionTime"`
	CarMilage         int     `json:"carMilage"`
	NextCarMilage     int     `json:"nextCarMilage"`
	Description       string  `json:"description"`
	TotalPrice        float32 `json:"totalPrice"`
}
