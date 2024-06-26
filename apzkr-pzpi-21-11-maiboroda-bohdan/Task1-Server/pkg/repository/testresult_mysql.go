package repository

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type TestResultMysql struct {
	db *sqlx.DB
}

func NewTestResultMysql(db *sqlx.DB) *TestResultMysql {
	return &TestResultMysql{db: db}
}

func (r *TestResultMysql) Create(userID int, testresult AlcoSafe.TestResult) (int, error) {
	var companyID int
	err := r.db.QueryRow("SELECT CompanyID FROM User WHERE UserID = ?", userID).Scan(&companyID)
	if err != nil {
		return 0, err
	}

	var legalLimit float64
	err = r.db.QueryRow("SELECT LegalLimit FROM Company WHERE CompanyID = ?", companyID).Scan(&legalLimit)
	if err != nil {
		return 0, err
	}

	isDrunk := testresult.AlcoholLevel > legalLimit

	query := "INSERT INTO TestResult (UserID, TestTime, AlcoholLevel, IsDrunk, Description) VALUES (?, NOW(), ?, ?, ?)"
	res, err := r.db.Exec(query, userID, testresult.AlcoholLevel, isDrunk, testresult.Description)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	notificationMessage := "Your alcohol level is within the legal limit."
	if isDrunk {
		notificationMessage = "Your alcohol level exceeds the legal limit!"
	}

	notification := AlcoSafe.Notification{
		UserID:   userID,
		Message:  notificationMessage,
		SentTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	_, err = r.CreateNotification(notification)
	if err != nil {
		return 0, err
	}

	access := "Allowed"
	accessTime := time.Now().Format("2006-01-02 15:04:05")
	if isDrunk {
		access = "Rejected"
		soberingHours := testresult.AlcoholLevel / 0.015
		accessTime = time.Now().Add(time.Duration(soberingHours) * time.Hour).Format("2006-01-02 15:04:05")
	}

	accessControl := AlcoSafe.AccessControl{
		UserID:     userID,
		AccessTime: accessTime,
		Access:     access,
	}
	_, err = r.CreateAccessControl(accessControl)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *TestResultMysql) CreateNotification(notification AlcoSafe.Notification) (int, error) {
	query := "INSERT INTO Notification (Message, SentTime, UserID) VALUES (?, ?, ?)"
	res, err := r.db.Exec(query, notification.Message, time.Now(), notification.UserID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *TestResultMysql) CreateAccessControl(accessControl AlcoSafe.AccessControl) (int, error) {
	query := "INSERT INTO AccessControl (UserID, AccessTime, Access) VALUES (?, ?, ?)"
	res, err := r.db.Exec(query, accessControl.UserID, accessControl.AccessTime, accessControl.Access)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *TestResultMysql) GetUserTestResult(userID int) ([]AlcoSafe.TestResult, error) {
	var testresult []AlcoSafe.TestResult
	query := "SELECT * FROM TestResult WHERE UserID=?"
	err := r.db.Select(&testresult, query, userID)
	return testresult, err
}

func (r *TestResultMysql) Delete(testResultID int) error {
	query := "DELETE FROM TestResult WHERE TestID = ?"
	_, err := r.db.Exec(query, testResultID)
	return err
}

func (r *TestResultMysql) GetAll() ([]AlcoSafe.TestResult, error) {
	var testResults []AlcoSafe.TestResult
	query := "SELECT tr.TestID, tr.UserID, tr.TestTime, tr.AlcoholLevel, tr.IsDrunk, tr.Description, u.Username " +
		"FROM TestResult tr " +
		"LEFT JOIN User u ON tr.UserID = u.UserID"
	err := r.db.Select(&testResults, query)
	if err != nil {
		log.Printf("Error fetching test results: %s", err)
		return nil, err
	}
	return testResults, nil
}

func (r *TestResultMysql) GetByID(testResultID int) (AlcoSafe.TestResult, error) {
	var testResult AlcoSafe.TestResult
	query := "SELECT * FROM TestResult WHERE TestResultID = ?"
	err := r.db.Get(&testResultID, query, testResultID)
	if err != nil {
		return testResult, err
	}
	return testResult, nil
}
