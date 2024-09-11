package repository

import (
	"assignment/models"
	"database/sql"
	"fmt"
)

// InsertRecord inserts a record into the records table.
func InsertRecord(db *sql.DB, record models.Record) error {
	query := `
    INSERT INTO records (first_name, last_name, company_name, address, city, country, postal, phone, email, web)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
	_, err := db.Exec(query, record.FirstName, record.LastName, record.Company, record.Address, record.City, record.Country, record.Postal, record.Phone, record.Email, record.Web)

	return err
}

// GetAllRecords retrieves all records from the records table.
func GetAllRecords(db *sql.DB) ([]models.Record, error) {
	query := `
    SELECT id, first_name, last_name, company_name, address, city, country, postal, phone, email, web
    FROM records
    `
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []models.Record
	for rows.Next() {
		var record models.Record
		if err := rows.Scan(&record.ID, &record.FirstName, &record.LastName, &record.Company, &record.Address, &record.City, &record.Country, &record.Postal, &record.Phone, &record.Email, &record.Web); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

// UpdateRecord updates a record in the database
func UpdateRecord(db *sql.DB, record models.Record) error {
	query := `
    UPDATE records
    SET first_name = ?, last_name = ?, company_name = ?, address = ?, city = ?, country = ?, postal = ?, phone = ?, email = ?, web = ?
    WHERE id = ?
    `
	_, err := db.Exec(query, record.FirstName, record.LastName, record.Company, record.Address, record.City, record.Country, record.Postal, record.Phone, record.Email, record.Web, record.ID)
	if err != nil {
		fmt.Println("error updating record in record_repository: ", err)
		return err
	}
	return nil
}
