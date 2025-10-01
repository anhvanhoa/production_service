package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("dev.config.yml")
	viper.ReadInConfig()
	dbURL := viper.GetString("url_db")
	dbOptions, err := pg.ParseURL(dbURL)
	if err != nil {
		log.Fatal("Failed to parse database URL:", err)
	}

	var seedFiles []string

	// Connect to database
	db := pg.Connect(dbOptions)

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Please provide an argument: up or down")
	}
	if args[0] == "up" {
		seedFiles = []string{
			"migrations/seed/000004_seed_harvest_records.up.sql",
			"migrations/seed/000005_seed_pest_disease_records.up.sql",
		}
	}
	if args[0] == "down" {
		seedFiles = []string{
			"migrations/seed/000005_seed_pest_disease_records.down.sql",
			"migrations/seed/000004_seed_harvest_records.down.sql",
		}
	}

	for _, file := range seedFiles {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatal("Failed to read file:", err)
		}
		_, err = db.Exec(string(content))
		if err != nil {
			log.Fatal("Failed to execute file:", err)
		}
	}
	fmt.Println("Seed data inserted successfully.")
	defer db.Close()
}
