package location

type Service interface {
	SetLocation(location string)
	ResetLocation()
	GetLocation() string
}

type locationService struct {
	dbService DbService
}

func (ls *locationService) SetLocation(location string) {
	ls.dbService.SetLocation(location)
}

func (ls *locationService) ResetLocation() {
	ls.dbService.ResetLocation()
}

func (ls *locationService) GetLocation() string {
	return ls.dbService.GetLocation()
}

func NewLocationService(dbService DbService) Service {
	return &locationService{
		dbService: dbService,
	}
}
