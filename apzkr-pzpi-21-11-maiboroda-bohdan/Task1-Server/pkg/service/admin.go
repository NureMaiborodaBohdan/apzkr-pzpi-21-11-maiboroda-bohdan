package service

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"apz-pzpi-21-11-maiboroda-bohdan-task2/pkg/repository"
)

type AdministratorService struct {
	repo repository.Admin
}

func NewAdminService(repo repository.Admin) *AdministratorService {
	return &AdministratorService{repo: repo}
}

func (s *AdministratorService) GetUserByID(userID int) (AlcoSafe.User, error) {
	return s.repo.GetUserByID(userID)
}

func (s *AdministratorService) CreateUser(user AlcoSafe.User) (int, error) {
	if err := user.ValidatePassword(); err != nil {
		return 0, err
	}
	if err := user.ValidateEmail(); err != nil {
		return 0, err
	}
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}
func (s *AdministratorService) UpdateUser(UserID int, input AlcoSafe.UpdateUserInput, user AlcoSafe.User) error {
	if err := input.Validate(); err != nil {
		return err
	}
	if err := user.ValidatePassword(); err != nil {
		return err
	}
	if err := user.ValidateEmail(); err != nil {
		return err
	}
	user.Password = generatePasswordHash(user.Password)
	return s.repo.UpdateUser(UserID, input, user)
}
func (s *AdministratorService) GetAllUsers() ([]AlcoSafe.User, error) {
	return s.repo.GetAllUsers()
}
func (s *AdministratorService) Delete(UserID int) error {
	return s.repo.Delete(UserID)
}
func (s *AdministratorService) BackupData(backupPath string) error {
	err := s.repo.BackupData(backupPath)
	if err != nil {
		return err
	}

	return nil
}
func (s *AdministratorService) RestoreData(backupPath string) error {
	err := s.repo.RestoreData(backupPath)
	if err != nil {
		return err
	}

	return nil
}
func (s *AdministratorService) ExportData(exportPath string) error {
	err := s.repo.ExportData(exportPath)
	if err != nil {
		return err
	}

	return nil
}
func (s *AdministratorService) ImportData(importPath string) error {
	err := s.repo.ImportData(importPath)
	if err != nil {
		return err
	}

	return nil
}
