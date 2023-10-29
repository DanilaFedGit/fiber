package main

import (
	"github.com/DanilaFedGit/fiber/handler"
	"github.com/DanilaFedGit/fiber/models"
	"github.com/DanilaFedGit/fiber/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	config := &storage.Config{
		Host:     os.Getenv("DB_HOSt"),
		Port:     os.Getenv("DB_Port"),
		DBName:   os.Getenv("DB_DBName"),
		UserName: os.Getenv("DB_User"),
		Password: os.Getenv("DB_Password"),
		SSLMode:  os.Getenv("DB_SSLMode"),
	}
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal(err)
	}
	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal(err)
	}
	r := handler.Repository{DataBase: db}
	r.SetupRouters(app)
	app.Listen(":8080")
}
