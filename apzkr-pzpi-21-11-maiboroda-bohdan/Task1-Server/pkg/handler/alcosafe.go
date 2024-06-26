package handler

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func (h *Handlers) getCompanies(c *gin.Context) {
	companies, err := h.service.Company.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companies)
}

func (h *Handlers) getCompanyByID(c *gin.Context) {
	companyID, err := strconv.Atoi(c.Param("companyID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid company ID"})
		return
	}
	company, err := h.service.Company.GetByID(companyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, company)
}

func (h *Handlers) createCompany(c *gin.Context) {
	var company AlcoSafe.Company
	if err := c.BindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	id, err := h.service.Company.Create(company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"CompanyID": id,
	})
}

func (h *Handlers) deleteCompany(c *gin.Context) {
	companyID, err := strconv.Atoi(c.Param("companyID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid company ID"})
		return
	}
	if err := h.service.Company.Delete(companyID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted"})
}

func (h *Handlers) updateCompany(c *gin.Context) {
	CompanyID := c.Param("companyID")
	id, err := strconv.Atoi(CompanyID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	var input AlcoSafe.UpdateCompany
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Company.Update(id, input)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			newErrorResponse(c, http.StatusConflict, "Company with this name already exists")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}
func (h *Handlers) getLocations(c *gin.Context) {
	location, err := h.service.Location.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, location)
}

func (h *Handlers) getLocationByID(c *gin.Context) {
	locationID, err := strconv.Atoi(c.Param("locationID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid location ID"})
		return
	}
	location, err := h.service.Location.GetByID(locationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, location)
}

func (h *Handlers) createLocation(c *gin.Context) {
	var location AlcoSafe.Location
	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	id, err := h.service.Location.Create(location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handlers) deleteLocation(c *gin.Context) {
	locationID, err := strconv.Atoi(c.Param("locationID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid location ID"})
		return
	}
	if err := h.service.Location.Delete(locationID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Location deleted"})
}

func (h *Handlers) updateLocation(c *gin.Context) {
	LocationID := c.Param("locationID")
	id, err := strconv.Atoi(LocationID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	var input AlcoSafe.UpdateLocation
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Location.Update(id, input)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			newErrorResponse(c, http.StatusConflict, "Location with this name already exists")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

func (h *Handlers) getAllNotification(c *gin.Context) {
	notification, err := h.service.Notification.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notification)
}
func (h *Handlers) getNotificationByID(c *gin.Context) {
	notification, err := strconv.Atoi(c.Param("notificationID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	notificationRes, err := h.service.Notification.GetByID(notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notificationRes)
}

func (h *Handlers) createNotification(c *gin.Context) {
	var notification AlcoSafe.Notification
	if err := c.BindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	id, err := h.service.Notification.Create(notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handlers) deleteNotification(c *gin.Context) {
	notificationID, err := strconv.Atoi(c.Param("notificationID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := h.service.Notification.Delete(notificationID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Notification deleted"})
}

func (h *Handlers) getTestResult(c *gin.Context) {
	testResult, err := h.service.TestResult.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, testResult)
}
func (h *Handlers) getTestResultByID(c *gin.Context) {
	testResultID, err := strconv.Atoi(c.Param("testID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	company, err := h.service.TestResult.GetByID(testResultID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, company)
}

func (h *Handlers) deleteTestResult(c *gin.Context) {
	testResultID, err := strconv.Atoi(c.Param("testID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid testResultID"})
		return
	}
	if err := h.service.TestResult.Delete(testResultID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Test result deleted"})
}

func (h *Handlers) getUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}
	user, err := h.service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handlers) getUsers(c *gin.Context) {
	users, err := h.service.Admin.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handlers) createUser(c *gin.Context) {
	var user AlcoSafe.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	id, err := h.service.Admin.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"UserID": id,
	})
}

func (h *Handlers) updateUser(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid UserID")
		return
	}

	var input AlcoSafe.UpdateUserInput
	var userInput AlcoSafe.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.Admin.UpdateUser(userID, input, userInput)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

func (h *Handlers) deleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}
	if err := h.service.Admin.Delete(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func (h *Handlers) getUserTestResults(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}
	testResults, err := h.service.TestResult.GetUserTestResult(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, testResults)
}

func (h *Handlers) createUserTestResult(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}

	var testresult AlcoSafe.TestResult
	if err := c.ShouldBindJSON(&testresult); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	testresult.UserID = userID

	id, err := h.service.TestResult.Create(userID, testresult)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handlers) getUserNotifications(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}

	notifications, err := h.service.Notification.GetAllUserNotification(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch notifications"})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

func (h *Handlers) getUserAccessControls(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	accessControls, err := h.service.GetUserAccessControl(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"UserID":         userID,
		"accessControls": accessControls,
	})
}

func (h *Handlers) backupData(c *gin.Context) {
	backupPath := "backup.sql"

	file, err := os.Create(backupPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create backup file",
		})
		return
	}
	defer file.Close()

	if err := h.service.Admin.BackupData(backupPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to backup data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data backup successful",
	})
}

func (h *Handlers) restoreData(c *gin.Context) {
	restorePath := "backup.sql"

	if err := h.service.Admin.RestoreData(restorePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to restore data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data restore successful",
	})
}

func (h *Handlers) exportData(c *gin.Context) {
	exportPath := "export.xlsx"

	if err := h.service.Admin.ExportData(exportPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to export data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data export successful"})
}

func (h *Handlers) importData(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uploadPath := file.Filename
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	if err := h.service.Admin.ImportData(uploadPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data imported successfully"})
}
