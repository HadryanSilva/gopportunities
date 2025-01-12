## JOB OPENINGS API
This is a simple API that allows users to create, read, update and delete job openings. The API is built using Go with Gin.

### API Endpoints
The API has the following endpoints:
1. **GET** /openings - This endpoint returns all job openings.
2. **GET** /openings/:id - This endpoint returns a single job opening.
3. **POST** /openings - This endpoint creates a new job opening.
4. **PUT** /openings/:id - This endpoint updates a job opening.
5. **DELETE** /openings/:id - This endpoint deletes a job opening.

### Runing the application
To run the application, run the following command on the root directory of the project:
```bash
  go run main.go
```

### API Documentation
To generate the API documentation, run the following command:
```bash 
  swag init
```
The API documentation can be found accessing [here](http://localhost:8080/swagger/index.html) before running the application.

### Libraries Used
1. **Gin** - Gin is a web framework written in Go (Golang). It features a martini-like API with much better performance.
2. **Gorm** - The fantastic ORM library for Golang, aims to be developer-friendly.
3. **Swag** - Swag converts Go annotations to Swagger Documentation 2.0.


