package models

import (
	"database/sql"
	"log"
)

// Record represents a row from the Excel file.
type Record struct {
	ID        uint   `json:"id" gorm:"primaryKey"` // Auto-increment primary key
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company_name"`
	Address   string `json:"address"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Postal    string `json:"postal"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Web       string `json:"web"`
}

// Migrate creates the 'records' table if it doesn't exist.
func Migrate(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS records (
        id INT AUTO_INCREMENT PRIMARY KEY,
        first_name VARCHAR(255),
        last_name VARCHAR(255),
        company_name VARCHAR(255),
        address VARCHAR(255),
        city VARCHAR(255),
        country VARCHAR(255),
        postal VARCHAR(50),
        phone VARCHAR(50),
        email VARCHAR(255),
        web VARCHAR(255)
    );
    `

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	} else {
		log.Println("Table 'records' created or already exists")
	}
}
