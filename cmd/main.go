package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kanatovnurzhas/crud-postgresql/internal/handlers"
	"github.com/kanatovnurzhas/crud-postgresql/internal/repository"
	"github.com/kanatovnurzhas/crud-postgresql/internal/services"
	"github.com/kanatovnurzhas/crud-postgresql/pkg/database"
	"log"
)

func main() {
	db, err := database.ConnectionDB()
	if err != nil {
		wrappedErr := fmt.Errorf("connection db refused: %w", err)
		log.Println(wrappedErr)
		return
	}
	log.Println("Connection success!")

	err = CreateTable(db)
	if err != nil {
		wrappedErr := fmt.Errorf("create table: %w", err)
		fmt.Println(wrappedErr)
		return
	}
	log.Println("Create table success!")

	sr := repository.NewStudentRepository(db)
	ss := services.NewStudentService(sr)
	sh := handlers.NewStudentHandler(ss)

	server := fiber.New()
	path := server.Group("/api")
	sh.RegisterStudentHandler(path)

	if err := server.Listen(":7777"); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

const studentTable = `
				CREATE TABLE IF NOT EXISTS students (
    				id SERIAL PRIMARY KEY,
    				firstName VARCHAR(255) NOT NULL,
    				secondName VARCHAR(255) NOT NULL,
				    email VARCHAR(150) NOT NULL UNIQUE,
    				age INTEGER NOT NULL
				);
				`

// функция для создания таблицы
func CreateTable(db *sql.DB) error {
	_, err := db.Query(studentTable)
	if err != nil {
		return fmt.Errorf("create table: %w", err)
	}
	return nil
}
