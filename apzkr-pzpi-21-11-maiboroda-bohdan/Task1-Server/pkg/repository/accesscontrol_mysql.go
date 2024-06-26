package repository

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"github.com/jmoiron/sqlx"
)

type AccessControlMysql struct {
	db *sqlx.DB
}

func NewAccessControlMysql(db *sqlx.DB) *AccessControlMysql {
	return &AccessControlMysql{db: db}
}
func (r *AccessControlMysql) GetUserAccessControl(userID int) ([]AlcoSafe.AccessControl, error) {
	var userAccessControl []AlcoSafe.AccessControl
	query := "SELECT * FROM AccessControl WHERE UserID=?"
	err := r.db.Select(&userAccessControl, query, userID)
	return userAccessControl, err
}
