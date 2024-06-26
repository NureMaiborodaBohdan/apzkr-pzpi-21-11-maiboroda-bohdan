package repository

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type AdminMysql struct {
	db *sqlx.DB
}

func NewAdminMysql(db *sqlx.DB) *AdminMysql {
	return &AdminMysql{db: db}
}

func (r *AdminMysql) GetUserByID(userID int) (AlcoSafe.User, error) {
	var user AlcoSafe.User
	query := fmt.Sprintf("SELECT UserID, Username, Email, Role, Name, Surname, Patronymic, Sex FROM User WHERE UserID=?")
	err := r.db.Get(&user, query, userID)
	if err != nil {
		log.Printf("Error fetching user by ID %d: %s", userID, err)
	}
	return user, err
}

func (r *AdminMysql) CreateUser(user AlcoSafe.User) (int, error) {
	checkQuery := fmt.Sprintf("SELECT COUNT(*) FROM User WHERE email=? OR username=?")
	var count int
	err := r.db.Get(&count, checkQuery, user.Email, user.Username)
	if err != nil {
		return 0, err
	}

	if count > 0 {
		return 0, fmt.Errorf("user with this email or username already exists")
	}

	query := fmt.Sprintf("INSERT INTO User (Username, Email, Password, Role, Name, Surname, Patronymic, CompanyID, Sex) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")

	result, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.Role, user.Name, user.Surname, user.Patronymic, user.CompanyID, user.Sex)
	if err != nil {
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok && mysqlErr.Number == 1062 {
			return 0, fmt.Errorf("user with this email or username already exists")
		}
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
func (r *AdminMysql) GetAllUsers() ([]AlcoSafe.User, error) {
	var users []AlcoSafe.User
	query := fmt.Sprintf("SELECT UserID, Username, Email, Role, Name, Surname, Patronymic, CompanyID, Sex FROM User")
	err := r.db.Select(&users, query)
	if err != nil {
		log.Printf("Error fetching all users: %s", err)
	}
	return users, err
}
func (r *AdminMysql) UpdateUser(UserID int, input AlcoSafe.UpdateUserInput, user AlcoSafe.User) error {
	existingUser, err := r.GetUserByID(UserID)
	if err != nil {
		return err
	}

	if input.Email != nil {
		checkEmailQuery := "SELECT UserID FROM User WHERE Email=? AND UserID != ?"
		var otherUserID int
		err := r.db.Get(&otherUserID, checkEmailQuery, *input.Email, UserID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if otherUserID != 0 {
			return errors.New("user with this email already exists")
		}

		existingUser.Email = *input.Email
	}
	if input.Password != nil {
		existingUser.Password = *input.Password
	}
	if input.Username != nil {
		checkUsernameQuery := "SELECT UserID FROM User WHERE Username=? AND UserID != ?"
		var otherUserID int
		err := r.db.Get(&otherUserID, checkUsernameQuery, *input.Username, UserID)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if otherUserID != 0 {
			return errors.New("user with this username already exists")
		}

		existingUser.Username = *input.Username
	}

	if input.Role != nil {
		existingUser.Role = *input.Role
	}

	if input.Name != nil {
		existingUser.Name = *input.Name
	}

	if input.Surname != nil {
		existingUser.Surname = *input.Surname
	}

	if input.Patronymic != nil {
		existingUser.Patronymic = *input.Patronymic
	}

	if input.CompanyID != nil {
		existingUser.CompanyID = input.CompanyID
	}

	if input.Sex != nil {
		existingUser.Sex = *input.Sex
	}

	query := fmt.Sprintf("UPDATE User SET Username=?, Email=?, Password=?, Role=?, Name=?, Surname=?, Patronymic=?, CompanyID=?, Sex=? WHERE UserID=?")
	_, err = r.db.Exec(query, existingUser.Username, existingUser.Email, existingUser.Password, existingUser.Role, existingUser.Name, existingUser.Surname, existingUser.Patronymic, existingUser.CompanyID, existingUser.Sex, UserID)
	if err != nil {
		return err
	}

	return nil
}
func (r *AdminMysql) Delete(UserID int) error {
	query := fmt.Sprintf("DELETE FROM User WHERE UserID = ? ")
	_, err := r.db.Exec(query, UserID)
	return err
}
func (r *AdminMysql) BackupData(backupPath string) error {
	cmd := exec.Command(
		"docker",
		"exec",
		"alcosafe-db",
		"mysqldump",
		"-u",
		"root",
		"-pROOT",
		"mysql",
	)

	outputFile, err := os.Create(backupPath)
	if err != nil {
		return fmt.Errorf("error creating backup file: %s", err)
	}
	defer outputFile.Close()

	cmd.Stdout = outputFile

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running mysqldump: %s", err)
	}

	return nil
}
func (r *AdminMysql) RestoreData(backupPath string) error {
	backupPath = "backup.sql"

	containerBackupPath := "/var/lib/mysql/backup.sql"
	cmd := exec.Command(
		"docker",
		"cp",
		backupPath,
		"alcosafe-db:"+containerBackupPath,
	)
	if err := cmd.Run(); err != nil {
		logrus.Errorf("Failed to copy dump file to container: %s", err)
		return err
	}

	dumpContent, err := ioutil.ReadFile(backupPath)
	if err != nil {
		logrus.Errorf("Failed to read dump file: %s", err)
		return err
	}

	mysqlCmd := exec.Command(
		"docker",
		"exec",
		"-i",
		"alcosafe-db",
		"mysql",
		"-u",
		"root",
		"-pROOT",
		"mysql",
	)

	mysqlCmd.Stdin = bytes.NewReader(dumpContent)
	mysqlCmd.Stdout = os.Stdout
	mysqlCmd.Stderr = os.Stderr

	if err := mysqlCmd.Run(); err != nil {
		logrus.Errorf("Failed to restore data in container: %s", err)
		return err
	}

	return nil
}
func (r *AdminMysql) ExportData(exportPath string) error {
	logger := logrus.New()

	file := xlsx.NewFile()

	tables := map[string]string{
		"User":          "SELECT * FROM User;",
		"AccessControl": "SELECT * FROM AccessControl;",
		"Notification":  "SELECT * FROM Notification;",
		"TestResult":    "SELECT * FROM TestResult;",
		"Company":       "SELECT * FROM Company;",
		"Location":      "SELECT * FROM Location;",
	}

	for tableName, query := range tables {
		rows, err := r.db.Query(query)
		if err != nil {
			logger.Errorf("Error querying table %s: %s", tableName, err)
			return fmt.Errorf("error querying table %s: %w", tableName, err)
		}
		defer rows.Close()

		columns, err := rows.Columns()
		if err != nil {
			logger.Errorf("Error fetching columns for table %s: %s", tableName, err)
			return fmt.Errorf("error fetching columns for table %s: %w", tableName, err)
		}

		sheet, err := file.AddSheet(tableName)
		if err != nil {
			logger.Errorf("Error adding sheet %s: %s", tableName, err)
			return fmt.Errorf("error adding sheet %s: %w", tableName, err)
		}

		headerRow := sheet.AddRow()
		for _, column := range columns {
			cell := headerRow.AddCell()
			cell.Value = column
		}

		for rows.Next() {
			rowData := make([]sql.NullString, len(columns))
			valuePointers := make([]interface{}, len(columns))
			for i := range rowData {
				valuePointers[i] = &rowData[i]
			}

			err := rows.Scan(valuePointers...)
			if err != nil {
				logger.Errorf("Error scanning rows for table %s: %s", tableName, err)
				return fmt.Errorf("error scanning rows for table %s: %w", tableName, err)
			}

			row := sheet.AddRow()
			for _, value := range rowData {
				cell := row.AddCell()
				if value.Valid {
					cell.Value = value.String
				} else {
					cell.Value = "NULL"
				}
			}
		}
	}

	err := file.Save(exportPath)
	if err != nil {
		logger.Errorf("Error saving Excel file: %s", err)
		return fmt.Errorf("error saving Excel file: %w", err)
	}

	return nil
}

func (r *AdminMysql) ImportData(importPath string) error {
	file, err := xlsx.OpenFile(importPath)
	if err != nil {
		logrus.Errorf("Error opening Excel file: %s", err)
		return fmt.Errorf("error opening Excel file: %w", err)
	}

	for _, sheet := range file.Sheets {
		tableName := sheet.Name

		rows := sheet.Rows
		if len(rows) < 2 {
			continue
		}

		columns := make([]string, len(rows[0].Cells))
		for i, cell := range rows[0].Cells {
			columns[i] = cell.String()
		}

		query := fmt.Sprintf("INSERT IGNORE INTO %s (%s) VALUES ", tableName, strings.Join(columns, ",")) // or use INSERT ... ON DUPLICATE KEY UPDATE

		var valueStrings []string
		var valueArgs []interface{}

		for _, row := range rows[1:] {
			var values []interface{}

			for _, cell := range row.Cells {
				values = append(values, cell.Value)
			}

			placeholders := make([]string, len(columns))
			for i := range placeholders {
				placeholders[i] = "?"
			}

			valueStrings = append(valueStrings, "("+strings.Join(placeholders, ",")+")")
			valueArgs = append(valueArgs, values...)
		}

		query += strings.Join(valueStrings, ",")
		_, err := r.db.Exec(query, valueArgs...)
		if err != nil {
			logrus.Errorf("Error inserting data into table %s: %s", tableName, err)
			return fmt.Errorf("error inserting data into table %s: %w", tableName, err)
		}
	}

	return nil
}
