package service

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"apz-pzpi-21-11-maiboroda-bohdan-task2/pkg/repository"
)

type Authorization interface {
	CreateAdmin(user AlcoSafe.User) (int, error)
	CreateUser(user AlcoSafe.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Company interface {
	Create(company AlcoSafe.Company) (int, error)
	GetAll() ([]AlcoSafe.Company, error)
	GetByID(companyID int) (AlcoSafe.Company, error)
	Delete(companyID int) error
	Update(CompanyID int, input AlcoSafe.UpdateCompany) error
}

type Location interface {
	Create(location AlcoSafe.Location) (int, error)
	GetByID(locationID int) (AlcoSafe.Location, error)
	Delete(locationID int) error
	GetAll() ([]AlcoSafe.Location, error)
	Update(LocationID int, input AlcoSafe.UpdateLocation) error
}

type TestResult interface {
	Create(userID int, testresult AlcoSafe.TestResult) (int, error)
	GetUserTestResult(userID int) ([]AlcoSafe.TestResult, error)
	Delete(testResultID int) error
	GetAll() ([]AlcoSafe.TestResult, error)
	GetByID(testResultID int) (AlcoSafe.TestResult, error)
}

type AccessControl interface {
	GetUserAccessControl(userID int) ([]AlcoSafe.AccessControl, error)
}

type Notification interface {
	Create(notification AlcoSafe.Notification) (int, error)
	GetAllUserNotification(userID int) ([]AlcoSafe.Notification, error)
	Delete(notificationID int) error
	GetAll() ([]AlcoSafe.Notification, error)
	GetByID(notificationID int) (AlcoSafe.Notification, error)
}

type Admin interface {
	GetUserByID(userID int) (AlcoSafe.User, error)
	CreateUser(user AlcoSafe.User) (int, error)
	GetAllUsers() ([]AlcoSafe.User, error)
	Delete(UserID int) error
	UpdateUser(UserID int, input AlcoSafe.UpdateUserInput, user AlcoSafe.User) error
	BackupData(backupPath string) error
	RestoreData(backupPath string) error
	ImportData(backupPath string) error
	ExportData(backupPath string) error
}

type Service struct {
	Authorization
	Company
	Location
	TestResult
	AccessControl
	Notification
	Admin
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Admin:         NewAdminService(repos.Admin),
		Company:       NewCompanyService(repos.Company),
		Location:      NewLocationService(repos.Location),
		TestResult:    NewTestResultService(repos.TestResult),
		Notification:  NewNotificationService(repos.Notification),
		AccessControl: NewAccessControlService(repos.AccessControl),
	}
}
