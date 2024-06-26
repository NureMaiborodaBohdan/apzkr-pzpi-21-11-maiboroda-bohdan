package service

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"apz-pzpi-21-11-maiboroda-bohdan-task2/pkg/repository"
)

type NotificationService struct {
	repo repository.Notification
}

func NewNotificationService(repo repository.Notification) *NotificationService {
	return &NotificationService{repo: repo}
}
func (s *NotificationService) Create(notification AlcoSafe.Notification) (int, error) {
	return s.repo.Create(notification)
}
func (s *NotificationService) GetAllUserNotification(userID int) ([]AlcoSafe.Notification, error) {
	return s.repo.GetAllUserNotification(userID)
}
func (s *NotificationService) GetAll() ([]AlcoSafe.Notification, error) {
	return s.repo.GetAll()
}
func (s *NotificationService) GetByID(notificationID int) (AlcoSafe.Notification, error) {
	return s.repo.GetByID(notificationID)
}

func (s *NotificationService) Delete(notificationID int) error {
	return s.repo.Delete(notificationID)
}
