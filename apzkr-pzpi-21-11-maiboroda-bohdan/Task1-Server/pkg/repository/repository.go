package repository

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateAdmin(user AlcoSafe.User) (int, error)
	CreateUser(user AlcoSafe.User) (int, error)
	GetUser(username, password string) (AlcoSafe.User, error)
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
	GetAll() ([]AlcoSafe.Location, error)
	Delete(locationID int) error
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
	Create(location AlcoSafe.Notification) (int, error)
	GetAllUserNotification(userID int) ([]AlcoSafe.Notification, error)
	Delete(notificationID int) error
	GetAll() ([]AlcoSafe.Notification, error)
	GetByID(notificationID int) (AlcoSafe.Notification, error)
}
type Admin interface {
	GetUserByID(userID int) (AlcoSafe.User, error)
	CreateUser(user AlcoSafe.User) (int, error)
	GetAllUsers() ([]AlcoSafe.User, error)
	UpdateUser(UserID int, user AlcoSafe.UpdateUserInput, input AlcoSafe.User) error
	Delete(UserID int) error
	BackupData(backupPath string) error
	RestoreData(backupPath string) error
	ImportData(backupPath string) error
	ExportData(backupPath string) error
}
type Repository struct {
	Authorization
	Company
	Location
	TestResult
	AccessControl
	Notification
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
		Admin:         NewAdminMysql(db),
		Company:       NewCompanyMysql(db),
		Location:      NewLocationMysql(db),
		TestResult:    NewTestResultMysql(db),
		Notification:  NewNotificationMysql(db),
		AccessControl: NewAccessControlMysql(db),
	}
}
