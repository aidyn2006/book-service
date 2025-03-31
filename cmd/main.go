package main

import (
	"book-service/internal/handler"
	"book-service/internal/repository"
	"book-service/internal/usecase"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "os"

	_ "github.com/lib/pq"
)

func main() {
	dbHost := "localhost"
	dbPort := 5432
	dbUser := "postgres"
	dbPassword := "Na260206"
	dbName := "book"

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("БД не отвечает:", err)
	}

	log.Println("✅ Успешное подключение к PostgreSQL")

	bookRepo := repository.NewBookRepository(db)
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	bookHandler := handler.NewBookHandler(bookUsecase)

	http.HandleFunc("/book", bookHandler.GetBook)
	http.HandleFunc("/book/add", bookHandler.AddBook)

	log.Println("🚀 Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
