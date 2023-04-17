package inspection

type InspectionInfrastructureInterface interface {
	CreateRegularCarInspection(inspectioon ReguralCarInspection) error
	GetReguralCarInspectionById(id int) (*ReguralCarInspection, error)
	GetListRegularCarInceptions(limit, offset int) (*[]ReguralCarInspection, error)
}

type inspectionInfrastructure struct {
}

func getInceptionInfrastructure() InspectionInfrastructureInterface {
	return &inspectionInfrastructure{}
}

func (i inspectionInfrastructure) CreateRegularCarInspection(inspectioon ReguralCarInspection) error {
	return nil
}

func (i inspectionInfrastructure) GetReguralCarInspectionById(id int) (*ReguralCarInspection, error) {
	return nil, nil
}

func (i inspectionInfrastructure) GetListRegularCarInceptions(limit, offset int) (*[]ReguralCarInspection, error) {
	return nil, nil
}
