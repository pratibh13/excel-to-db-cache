# Data Import and Caching Web Application

This repository contains a web application for importing Excel data into a MySQL database and caching the data in Redis. The application is built using Go and Gin and includes Docker  configurations for deployment.

## Tech Stack

- **Programming Language:** Go
- **Framework:** Gin
- **Database:** MySQL
- **Caching:** Redis
- **Excel Parsing:** `excelize` library
- **Testing:** Go testing package, Testify
- **Containerization:** Docker

## Prerequisites

Before you begin, ensure you have the following installed:

- Docker
- Docker Compose
- Go (for running tests and building)

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/pratibh13/excel-to-db-cache.git
cd excel-to-db-cache
```
## Build and Run the Application Locally

### Set Up Dependencies

Install Go dependencies:
```
go mod tidy
```
### Build the Application
```
go build -o main ./cmd/excel-to-db-cache
```
### Run the Application
```
go run main.go
```

## Run Tests and Benchmark
```
go test ./tests
```
## Docker Setup

### Build Docker Image
```
docker build -t excel-to-db-cache .
```
### Run Docker Container
```
docker run -p 8080:8080 excel-to-db-cache
```

## Docker Compose Setup
```
docker-compose up --build
```
## Available Functions

### `POST /import`

**Description:** Imports data from an Excel file into the application.

**Request:**

- **Content-Type:** `multipart/form-data`
- **Body:** Excel file to be imported.

**Response:**

- **Status Code:** `200 OK` if the data is successfully imported.
- **Response Body:** `"Data imported successfully"`

### `GET /records`

**Description:** Retrieves all records from the database.

**Response:**

- **Status Code:** `200 OK` if records are retrieved successfully.
- **Response Body:** JSON array of records.

### `PUT /records/:id`

**Description:** Updates a specific record identified by `id`.

**Request:**

- **Content-Type:** `application/json`
- **URL Parameters:** `id` (the ID of the record to update)
- **Body:** JSON object with updated record details. 

### CONTRIBUTING
Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or create a pull request.
