package inspection

import (
	"errors"
	"time"
)

var (
	ErrEmptyField = errors.New("err-empty-field")
)

type InseptionServceInterface interface {
	CreateRegularCarInspection(carInspectioon ReguralCarInspectionRequest) error
	GetListRegularCarInceptions() (*[]ReguralCarInspection, error)
}

type inseptionService struct {
	repository InspectionRepositoryInterface
}

func GetInceptionSercviceInterface() InseptionServceInterface {
	return &inseptionService{}
}

func (is inseptionService) CreateRegularCarInspection(carInspection ReguralCarInspectionRequest) error {

	if carInspection.Name == "" {
		return ErrEmptyField
	}

	if carInspection.DateInspectionCar == "" {
		return ErrEmptyField
	}

	nextCarMilage := carInspection.CarMilage + 10000
	dateInception, err := time.Parse(time.RFC3339, carInspection.DateInspectionCar)

	if err != nil {
		return err
	}

	carInspectionModel := ReguralCarInspection{
		Name:              carInspection.Name,
		Description:       carInspection.Description,
		TotalPrice:        carInspection.TotalPrice,
		CarMilage:         carInspection.CarMilage,
		NextCarMilage:     nextCarMilage,
		DateInspectionCar: dateInception,
	}

	//TODO :: 10 000 dodac config z ustawieniami co ile przeglad

	err = is.repository.CreateRegularCarInspection(carInspectionModel)
	if err != nil {
		return err
	}
	return nil
}

func (is inseptionService) GetListRegularCarInceptions() (*[]ReguralCarInspection, error) {

	result, err := is.repository.GetListRegularCarInceptions(10, 0)

	if err != nil {
		return nil, err
	}

	return result, nil
}
