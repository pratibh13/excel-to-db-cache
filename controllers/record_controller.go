package controllers

import (
	"assignment/config"
	"assignment/models"
	"assignment/services"
	"assignment/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ImportExcel(c *gin.Context) {
	// Retrieve the file from the form
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("Error retrieving file: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload error"})
		return
	}

	// Open the file
	f, err := file.Open()
	if err != nil {
		log.Printf("Error opening file: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error opening file"})
		return
	}
	defer f.Close()

	// Parse the Excel file
	records, err := utils.ParseExcel(f)
	if err != nil {
		log.Printf("Error parsing Excel file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Excel format"})
		return
	}

	// Import records into the database
	if err := services.ImportRecords(records); err != nil {
		log.Printf("Error importing records: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error importing data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data imported successfully"})
}

func ViewRecords(c *gin.Context) {
	records, err := services.GetRecords()
	if err != nil {
		log.Printf("Error fetching records: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching records"})
		return
	}
	c.JSON(http.StatusOK, records)
}

func UpdateRecord(c *gin.Context) {
	var record models.Record
	idStr := c.Param("id")

	// Convert ID from string to uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Printf("Error parsing ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	record.ID = uint(id)

	// Bind the request body to the record struct
	if err := c.BindJSON(&record); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Update the record in the database using the global DB connection
	if err := services.UpdateRecordInDB(config.DB, record); err != nil {
		log.Printf("Error updating record in database in record_controller: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating record in database"})
		return
	}

	// Update the record in Redis cache
	if err := services.UpdateRecordInCache(&record); err != nil {
		log.Printf("Error updating record in cache: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating record in cache"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record updated successfully"})
}
