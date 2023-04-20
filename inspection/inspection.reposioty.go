package inspection

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	cache "github/lewy9109/autoNotes/cacheRedis"
	"strconv"
	"time"

	"gorm.io/gorm"
)

var (
	ctx            = context.Background()
	keyInspectList = "get-list-regular-inspect"
)

type InspectionRepositoryInterface interface {
	CreateRegularCarInspection(carInspection ReguralCarInspection) error
	GetReguralCarInspectionById(id int) (*ReguralCarInspection, error)
	GetListRegularCarInspections(offset, limit int) (*[]ReguralCarInspection, error)
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

func (i inspectionRepository) GetListRegularCarInspections(offset, limit int) (*[]ReguralCarInspection, error) {
	reguralCarInspectionSlice := []ReguralCarInspection{}
	keyInspectList = keyInspectList + "offset" + strconv.Itoa(offset) + "lim" + strconv.Itoa(limit)
	fmt.Println(keyInspectList)

	rdb := cache.NewRedisClient()
	resultCache := rdb.Get(ctx, keyInspectList)

	if resultCache != "" {
		json.Unmarshal([]byte(resultCache), &reguralCarInspectionSlice)
		return &reguralCarInspectionSlice, nil
	}

	if result := i.db.Limit(limit).Offset(offset).Order("date_inspection_car DESC").Find(&reguralCarInspectionSlice); result.Error != nil {
		return nil, result.Error
	}

	jsonList, err := json.Marshal(reguralCarInspectionSlice)
	if err != nil {
		panic(err)
	}

	rdb.Set(ctx, keyInspectList, string(jsonList), 1*time.Minute)

	return &reguralCarInspectionSlice, nil
}
