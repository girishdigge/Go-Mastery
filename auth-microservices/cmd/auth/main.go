package main

import (
	"log"
	"os"

	"github.com/girishdigge/Go-Mastery/auth-microservices/config"
	"github.com/girishdigge/Go-Mastery/auth-microservices/pkg/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Starting the application.")
	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	log.Printf("Success parsed config file: %v", cfg.Server.AppVersion)
	// db, err := postgres.NewPsqlDB(cfg)
	// if err != nil {
	// 	log.Fatalf("NewPsqlDB: %v", err)
	// }
}
