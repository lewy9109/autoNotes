package inspection

import (
	"errors"
)

var (
	ErrEmptyField = errors.New("err-empty-field")
)

type InseptionServceInterface interface {
	CreateRegularCarInspection(carInspectioon ReguralCarInspection) error
	GetListRegularCarInceptions() (*[]ReguralCarInspection, error)
	GetRegularCarInceptions(id int) (*ReguralCarInspection, error)
}

type inseptionService struct {
	repository InspectionRepositoryInterface
}

func GetInceptionSercvice(repo InspectionRepositoryInterface) InseptionServceInterface {
	return &inseptionService{
		repository: repo,
	}
}

func (is inseptionService) CreateRegularCarInspection(carInspection ReguralCarInspection) error {

	if carInspection.Name == "" {
		return ErrEmptyField
	}

	carInspection.NextCarMilage = carInspection.CarMilage + 10000

	err := is.repository.CreateRegularCarInspection(carInspection)
	if err != nil {
		return err
	}
	return nil
}

func (is inseptionService) GetListRegularCarInceptions() (*[]ReguralCarInspection, error) {

	//TODO: add offset and limit from request param or header param
	result, err := is.repository.GetListRegularCarInceptions(30, 0)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (is inseptionService) GetRegularCarInceptions(id int) (*ReguralCarInspection, error) {
	result, err := is.repository.GetReguralCarInspectionById(id)

	if err != nil {
		return nil, err
	}
	return result, nil
}
