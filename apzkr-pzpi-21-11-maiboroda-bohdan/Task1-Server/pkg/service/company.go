package service

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"apz-pzpi-21-11-maiboroda-bohdan-task2/pkg/repository"
)

type CompanyService struct {
	repo repository.Company
}

func NewCompanyService(repo repository.Company) *CompanyService {
	return &CompanyService{repo: repo}
}

func (s *CompanyService) Create(company AlcoSafe.Company) (int, error) {
	return s.repo.Create(company)
}

func (s *CompanyService) GetAll() ([]AlcoSafe.Company, error) {
	return s.repo.GetAll()
}

func (s *CompanyService) GetByID(companyID int) (AlcoSafe.Company, error) {
	return s.repo.GetByID(companyID)
}

func (s *CompanyService) Delete(companyID int) error {
	return s.repo.Delete(companyID)
}
func (s *CompanyService) Update(CompanyID int, input AlcoSafe.UpdateCompany) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(CompanyID, input)
}
