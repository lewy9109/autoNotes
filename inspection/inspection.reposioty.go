package inspection

import (
	"errors"

	"gorm.io/gorm"
)

type InspectionRepositoryInterface interface {
	CreateRegularCarInspection(carInspection ReguralCarInspection) error
	GetReguralCarInspectionById(id int) (*ReguralCarInspection, error)
	GetListRegularCarInceptions(offset, limit int) (*[]ReguralCarInspection, error)
}

type inspectionRepository struct {
	db *gorm.DB
}

func GetInceptionRepository(db *gorm.DB) InspectionRepositoryInterface {
	return &inspectionRepository{
		db,
	}
}

func (i inspectionRepository) CreateRegularCarInspection(carInspection ReguralCarInspection) error {
	if result := i.db.Create(&carInspection); result.Error != nil {
		return result.Error
	}

	return nil
}

func (i inspectionRepository) GetReguralCarInspectionById(id int) (*ReguralCarInspection, error) {
	reguralCarInspection := ReguralCarInspection{}

	if result := i.db.First(&reguralCarInspection, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &reguralCarInspection, nil
}

func (i inspectionRepository) GetListRegularCarInceptions(offset, limit int) (*[]ReguralCarInspection, error) {
	reguralCarInspectionSlice := []ReguralCarInspection{}

	if result := i.db.Limit(limit).Offset(offset).Order("date_inspection_car DESC").Find(&reguralCarInspectionSlice); result.Error != nil {
		return nil, result.Error
	}

	return &reguralCarInspectionSlice, nil
}
