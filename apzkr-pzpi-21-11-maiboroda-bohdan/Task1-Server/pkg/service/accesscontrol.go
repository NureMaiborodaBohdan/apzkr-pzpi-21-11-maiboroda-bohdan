package service

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"apz-pzpi-21-11-maiboroda-bohdan-task2/pkg/repository"
)

type AccessControlService struct {
	repo repository.AccessControl
}

func NewAccessControlService(repo repository.AccessControl) *AccessControlService {
	return &AccessControlService{repo: repo}
}
func (s *AccessControlService) GetUserAccessControl(userID int) ([]AlcoSafe.AccessControl, error) {
	return s.repo.GetUserAccessControl(userID)
}
