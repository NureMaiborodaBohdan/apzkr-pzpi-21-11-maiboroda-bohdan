package service

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"apz-pzpi-21-11-maiboroda-bohdan-task2/pkg/repository"
)

type TestResultService struct {
	repo repository.TestResult
}

func NewTestResultService(repo repository.TestResult) *TestResultService {
	return &TestResultService{repo: repo}
}
func (s *TestResultService) GetAll() ([]AlcoSafe.TestResult, error) {
	return s.repo.GetAll()
}
func (s *TestResultService) GetByID(testResultID int) (AlcoSafe.TestResult, error) {
	return s.repo.GetByID(testResultID)
}

func (s *TestResultService) Create(userID int, testresult AlcoSafe.TestResult) (int, error) {
	return s.repo.Create(userID, testresult)
}

func (s *TestResultService) GetUserTestResult(userID int) ([]AlcoSafe.TestResult, error) {
	return s.repo.GetUserTestResult(userID)
}

func (s *TestResultService) Delete(testresultID int) error {
	return s.repo.Delete(testresultID)
}
