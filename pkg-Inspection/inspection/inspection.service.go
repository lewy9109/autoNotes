package inspection

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyField = errors.New("err-empty-field")
)

type InseptionServceInterface interface {
	CreateRegularCarInspection(carInspectioon ReguralCarInspection) error
	GetListRegularCarInceptions(offset, limit int) (*[]ReguralCarInspection, error)
	GetRegularCarInceptions(id int) (*ReguralCarInspection, error)
}

type inseptionService struct {
	repository InspectionRepositoryInterface
}

func NewService(repo InspectionRepositoryInterface) InseptionServceInterface {
	return &inseptionService{
		repository: repo,
	}
}

func (is inseptionService) CreateRegularCarInspection(carInspection ReguralCarInspection) error {

	if carInspection.Name == "" {
		return ErrEmptyField
	}

	carInspection.NextCarMilage = carInspection.CarMilage + 10000

	fmt.Println(carInspection)

	err := is.repository.CreateRegularCarInspection(carInspection)
	if err != nil {
		return err
	}
	return nil
}

func (is inseptionService) GetListRegularCarInceptions(offset, limit int) (*[]ReguralCarInspection, error) {

	result, err := is.repository.GetListRegularCarInspections(offset, limit)

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
