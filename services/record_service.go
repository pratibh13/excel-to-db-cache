package services

import (
	"assignment/config"
	"assignment/models"
	"assignment/repository"
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

// ImportRecords inserts parsed records into the MySQL database and caches them in Redis
func ImportRecords(records []models.Record) error {
	// Insert each record into the MySQL database
	for _, record := range records {
		err := repository.InsertRecord(config.DB, record) // Assuming InsertRecord returns the ID and error
		if err != nil {
			return err
		}
	}

	// Retrieve all records to cache them
	result, err := repository.GetAllRecords(config.DB)
	if err != nil {
		return err
	}
	// fmt.Println(result)

	// Cache the records in Redis after insertion
	err = CacheRecords(result)
	if err != nil {
		return err
	}

	return nil
}

// GetRecords fetches records from Redis cache or the database if the cache is empty
func GetRecords() ([]models.Record, error) {
	// Check cache first
	cachedRecords, err := GetCachedRecords()
	if err == nil {
		return cachedRecords, nil
	}

	// Fetch from DB if cache miss
	records, err := repository.GetAllRecords(config.DB)
	if err != nil {
		return nil, err
	}

	// Cache the records for future requests
	err = CacheRecords(records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

// GetCachedRecords fetches records from Redis cache
func GetCachedRecords() ([]models.Record, error) {
	// Fetch cached records from Redis
	cachedData, err := config.RDB.Get(config.Ctx, "records").Result()
	if err != nil {
		return nil, errors.New("cache miss or error fetching from cache")
	}

	// Unmarshal JSON into Go struct
	var records []models.Record
	err = json.Unmarshal([]byte(cachedData), &records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

// CacheRecords caches the records in Redis
func CacheRecords(records []models.Record) error {
	// Marshal the records into JSON
	jsonData, err := json.Marshal(records)
	if err != nil {
		return err
	}

	// Cache the records in Redis with an expiration time of 5 minutes
	err = config.RDB.Set(config.Ctx, "records", jsonData, 5*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

// UpdateRecordInCache updates a specific record in Redis cache
func UpdateRecordInCache(record *models.Record) error {

	// Retrieve all records to cache them
	result, err := repository.GetAllRecords(config.DB)
	if err != nil {
		return err
	}
	// fmt.Println(result)

	// Cache the records in Redis after insertion
	err = CacheRecords(result)
	if err != nil {
		return err
	}

	return nil
}

// UpdateRecord updates a specific record in the database
func UpdateRecordInDB(db *sql.DB, record models.Record) error {
	// Call the repository function to update the record
	return repository.UpdateRecord(db, record)
}
