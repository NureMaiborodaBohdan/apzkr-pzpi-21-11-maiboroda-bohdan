package repository

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"github.com/jmoiron/sqlx"
	"log"
)

type CompanyMysql struct {
	db *sqlx.DB
}

func NewCompanyMysql(db *sqlx.DB) *CompanyMysql {
	return &CompanyMysql{db: db}
}

func (r *CompanyMysql) Create(company AlcoSafe.Company) (int, error) {
	query := "INSERT INTO Company (Name, Description, LegalLimit, LocationID) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, company.Name, company.Description, company.LegalLimit, company.LocationID)
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

func (r *CompanyMysql) GetAll() ([]AlcoSafe.Company, error) {
	var companies []AlcoSafe.Company
	query := "SELECT c.CompanyID, c.Name, c.Description, c.LegalLimit, c.LocationID, " +
		"l.Country, l.City, l.Address, l.PostCode " +
		"FROM Company c " +
		"INNER JOIN Location l ON c.LocationID = l.LocationID"
	err := r.db.Select(&companies, query)
	if err != nil {
		// Log the error or return it based on your application's error handling strategy
		log.Println("Error fetching companies:", err)
		return nil, err
	}
	return companies, nil
}

func (r *CompanyMysql) GetByID(companyID int) (AlcoSafe.Company, error) {
	var company AlcoSafe.Company
	query := "SELECT * FROM Company WHERE CompanyID = ?"
	err := r.db.Get(&company, query, companyID)
	return company, err
}

func (r *CompanyMysql) Delete(companyID int) error {
	query := "DELETE FROM Company WHERE CompanyID = ?"
	_, err := r.db.Exec(query, companyID)
	return err
}
func (r *CompanyMysql) Update(CompanyID int, input AlcoSafe.UpdateCompany) error {
	existingCompany, err := r.GetByID(CompanyID)
	if err != nil {
		return err
	}

	if input.Name != nil {
		existingCompany.Name = *input.Name
	}
	if input.Description != nil {
		existingCompany.Description = *input.Description
	}
	if input.LegalLimit != nil {
		existingCompany.LegalLimit = *input.LegalLimit
	}
	if input.LocationID != nil {
		existingCompany.LocationID = *input.LocationID
	}

	query := "UPDATE Company SET Name=?, Description=?, LegalLimit=?, LocationID=? WHERE CompanyID=?"
	_, err = r.db.Exec(query, existingCompany.Name, existingCompany.Description, existingCompany.LegalLimit, existingCompany.LocationID, CompanyID)
	if err != nil {
		return err
	}

	return nil
}
