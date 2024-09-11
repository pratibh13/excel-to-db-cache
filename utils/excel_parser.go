package utils

import (
	"assignment/models"
	"fmt"
	"io"

	"github.com/xuri/excelize/v2"
)

// ParseExcel parses the Excel file and returns a list of records.
func ParseExcel(file io.Reader) ([]models.Record, error) {
	var records []models.Record

	// Open the Excel file
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open the Excel file: %v", err)
	}
	fmt.Println(xlsx)

	// List all sheet names for debugging
	sheetNames := xlsx.GetSheetMap()
	fmt.Println("Available sheets:")
	var sheetname string
	for _, sheetName := range sheetNames {
		fmt.Println(sheetName)
		sheetname = sheetName
	}

	// Check if the specified sheet exists
	sheetName := sheetname // Make sure this matches your sheet name
	_, err = xlsx.GetSheetIndex(sheetName)
	if err != nil {
		return nil, fmt.Errorf("error retrieving sheet index: %v", err)
	}

	// Get all rows from the sheet
	rows, err := xlsx.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to get rows from sheet '%s': %v", sheetName, err)
	}

	// Print rows for debugging
	fmt.Println("Rows retrieved from Excel:")
	for i, row := range rows {
		fmt.Printf("Row %d: %v\n", i+1, row)
	}

	// Ensure there are at least two rows (one for the header, one for the data)
	if len(rows) < 2 {
		return nil, fmt.Errorf("invalid Excel format: no data found")
	}

	// Iterate over rows (skipping the header row)
	for i, row := range rows[1:] { // Skipping the header row (rows[0])
		if len(row) < 10 { // Check if the row has the required number of columns
			fmt.Printf("Skipping row %d: expected 10 columns, got %d\n", i+2, len(row))
			continue
		}

		// Parse each row into a Record struct
		record := models.Record{
			FirstName: row[0],
			LastName:  row[1],
			Company:   row[2],
			Address:   row[3],
			City:      row[4],
			Country:   row[5],
			Postal:    row[6],
			Phone:     row[7],
			Email:     row[8],
			Web:       row[9],
		}

		// Append the record to the records slice
		records = append(records, record)
	}

	// If no valid rows were found, return an error
	if len(records) == 0 {
		return nil, fmt.Errorf("no valid rows found in the Excel file")
	}

	return records, nil
}
