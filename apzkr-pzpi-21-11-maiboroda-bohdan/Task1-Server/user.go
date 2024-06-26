package AlcoSafe

import (
	"errors"
	"regexp"
)

type User struct {
	UserID     int    `json:"UserID" db:"UserID"`
	Username   string `json:"Username" binding:"required" db:"Username"`
	Email      string `json:"Email" binding:"required" db:"Email"`
	Password   string `json:"Password" binding:"required" db:"Password"`
	Role       string `json:"Role" db:"Role"`
	Name       string `json:"Name" db:"Name"`
	Surname    string `json:"Surname" db:"Surname"`
	Patronymic string `json:"Patronymic" db:"Patronymic"`
	CompanyID  *int   `json:"CompanyID" db:"CompanyID"`
	Sex        string `json:"Sex" db:"Sex"`
}

type UpdateUserInput struct {
	UserID     *int    `json:"UserID"`
	Username   *string `json:"Username"`
	Email      *string `json:"Email"`
	Password   *string `json:"Password"`
	Role       *string `json:"Role"`
	Name       *string `json:"Name"`
	Surname    *string `json:"Surname"`
	Patronymic *string `json:"Patronymic"`
	CompanyID  *int    `json:"CompanyID"`
	Sex        *string `json:"Sex"`
}

func (i UpdateUserInput) Validate() error {
	if i.Username == nil && i.Email == nil && i.Password == nil && i.Role == nil &&
		i.Name == nil && i.Surname == nil && i.Patronymic == nil && i.CompanyID == nil && i.Sex == nil {
		return errors.New("Update structure has no value")
	}
	return nil
}
func (u *User) ValidateEmail() error {
	if u.Email != "" {
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		match, _ := regexp.MatchString(emailRegex, u.Email)
		if !match {
			return errors.New("invalid email format")
		}
	}
	return nil
}
func (u *User) ValidatePassword() error {
	if u.Password != "" {
		if len(u.Password) < 8 {
			return errors.New("password should be at least 8 characters long")
		}

		letterRegex := regexp.MustCompile(`[a-zA-Z]`)
		if !letterRegex.MatchString(u.Password) {
			return errors.New("password should contain at least 1 letter")
		}

		digitRegex := regexp.MustCompile(`\d`)
		if !digitRegex.MatchString(u.Password) {
			return errors.New("password should contain at least 1 digit")
		}

		specialCharRegex := regexp.MustCompile(`[^a-zA-Z0-9\s]`)
		if !specialCharRegex.MatchString(u.Password) {
			return errors.New("password should contain at least 1 special character")
		}
	}
	return nil
}
