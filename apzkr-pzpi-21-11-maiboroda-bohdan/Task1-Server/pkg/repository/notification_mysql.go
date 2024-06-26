package repository

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"github.com/jmoiron/sqlx"
)

type NotificationMysql struct {
	db *sqlx.DB
}

func NewNotificationMysql(db *sqlx.DB) *NotificationMysql {
	return &NotificationMysql{db: db}
}
func (r *NotificationMysql) Create(notification AlcoSafe.Notification) (int, error) {
	query := "INSERT INTO Notification (Message, SentTime, UserID) VALUES (?, NOW(), ?)"
	result, err := r.db.Exec(query, notification.Message, notification.UserID)
	if err != nil {
		return 0, err
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	id := int(id64)
	return id, nil
}
func (r *NotificationMysql) GetAllUserNotification(userID int) ([]AlcoSafe.Notification, error) {
	var userNotification []AlcoSafe.Notification
	query := "SELECT * FROM Notification WHERE UserID=?"
	err := r.db.Select(&userNotification, query, userID)
	return userNotification, err
}
func (r *NotificationMysql) GetAll() ([]AlcoSafe.Notification, error) {
	var notification []AlcoSafe.Notification
	query := "SELECT * FROM Notification"
	err := r.db.Select(&notification, query)
	return notification, err
}
func (r *NotificationMysql) GetByID(locationID int) (AlcoSafe.Notification, error) {
	var notification AlcoSafe.Notification
	query := "SELECT * FROM Notification WHERE NotificationID = ?"
	err := r.db.Get(&notification, query, locationID)
	if err != nil {
		return notification, err
	}
	return notification, nil
}

func (r *NotificationMysql) Delete(notificationID int) error {
	query := "DELETE FROM Notification WHERE NotificationID = ?"
	_, err := r.db.Exec(query, notificationID)
	return err
}
