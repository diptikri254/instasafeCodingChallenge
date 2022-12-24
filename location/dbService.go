package location

type DbService interface {
	SetLocation(location string)
	ResetLocation()
	GetLocation() string
}

type locationDbService struct {
}

func (ls *locationDbService) SetLocation(location string) {
	location = location
}

func (ls *locationDbService) ResetLocation() {
	location = ""
}

func (ls *locationDbService) GetLocation() string {
	return location
}

func NewLocationDbService() DbService {
	return &locationDbService{}
}
