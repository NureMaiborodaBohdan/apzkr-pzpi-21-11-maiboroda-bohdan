package service

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"apz-pzpi-21-11-maiboroda-bohdan-task2/pkg/repository"
)

type LocationService struct {
	repo repository.Location
}

func NewLocationService(repo repository.Location) *LocationService {
	return &LocationService{repo: repo}
}

func (s *LocationService) Create(location AlcoSafe.Location) (int, error) {
	return s.repo.Create(location)
}
func (s *LocationService) GetAll() ([]AlcoSafe.Location, error) {
	return s.repo.GetAll()
}
func (s *LocationService) GetByID(locationID int) (AlcoSafe.Location, error) {
	return s.repo.GetByID(locationID)
}

func (s *LocationService) Delete(locationID int) error {
	return s.repo.Delete(locationID)
}

func (s *LocationService) Update(LocationID int, input AlcoSafe.UpdateLocation) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(LocationID, input)
}
