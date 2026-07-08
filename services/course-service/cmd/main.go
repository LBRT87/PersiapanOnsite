package main

import (
	"fmt"
	"log"

	"github.com/LBRT87/PersiapanOnsite/services/course-service/config"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error while Loading .env : %v", err)
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error while Parsing .env : %v", err)
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v port=%v name=%v sslmode=disabled TimeZone=Asia/Jakarta", cfg.DbHost, cfg.DbUser, cfg.DbPass, cfg.DbPort, cfg.DbName)

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatalf("Error while connecting to database : %v", err)
	}

	_ = db
	_ = cfg
}