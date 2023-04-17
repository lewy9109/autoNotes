package inspection

import "time"

type ReguralCarInspection struct {
	Name              string
	DateInspectionCar time.Time
	CarMilage         int
	NextCarMilage     int //nastepny sugerowany przeglad !!! 10 tys dla bezyny 15 tys dla dizla edytowalne w ustawieniach sumowac z CarMilage
	Description       string
	TotalPrice        float32
	//zdjecia paragonu
	//struktura tablica z wymienionymi podzespoalmi
}
