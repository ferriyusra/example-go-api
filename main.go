package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	Config "example-go-api/config"
	Database "example-go-api/db"

	"example-go-api/logger"
	"example-go-api/route"

	AuthHandler "example-go-api/domain/auth/handler"
	AuthRepository "example-go-api/domain/auth/repository"
	AuthService "example-go-api/domain/auth/service"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	//  load env
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	//  init config
	config, err := Config.New()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Config initialized.")

	//  init database
	db := Database.New(config.RdsUrl)

	log.Printf("Database initialized.")
	defer db.Close()

	// init logger
	logger.Init(config)

	//  repo
	authRepository := AuthRepository.NewRepository(db)

	// service
	authService := AuthService.NewService(authRepository)

	// crm handler
	crmModuleHandler := AuthHandler.NewCrmAuthHandler(authService)

	//  initialise routes
	r := route.NewRouter(crmModuleHandler)

	handlers := cors.Default().Handler(r)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(`:%s`, os.Getenv("APP_PORT")), handlers))
}
