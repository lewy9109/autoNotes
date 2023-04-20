package cache

import (
	"github/lewy9109/autoNotes/inspection/controller"
)

type InspectionCarCache interface {
	Set(key string, value *controller.GetInspectionResponse)
	Get(key string) *controller.GetInspectionResponse
}
