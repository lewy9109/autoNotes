package controller

type CreateInseptionRequest struct {
	Name              string  `json:"name"`
	DateInspectionCar string  `json:"dateInspectionTime"`
	CarMilage         int     `json:"carMilage"`
	Description       string  `json:"description"`
	TotalPrice        float32 `json:"totalPrice"`
}

type CreateInseptionResponse struct {
}

type GetInspectionResponse struct {
}
