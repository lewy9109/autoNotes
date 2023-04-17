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

func GetInceptionSercvice(repo InspectionRepositoryInterface) InseptionServceInterface {
	return &inseptionService{
		repository: repo,
	}
}

func (is inseptionService) CreateRegularCarInspection(carInspection ReguralCarInspectionRequest) error {

	if carInspection.Name == "" {
		return ErrEmptyField
	}

	if carInspection.DateInspectionCar == "" {
		return ErrEmptyField
	}

	//TODO :: 10 000 dodac config z ustawieniami co ile przeglad
	nextCarMilage := carInspection.CarMilage + 10000
	const shortForm = "2006-01-02"
	dateInception, err := time.Parse(shortForm, carInspection.DateInspectionCar)

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
