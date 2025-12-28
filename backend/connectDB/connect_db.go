package connectDB

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, func()) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Could not load .env file")
	}

	dbConn, ok := os.LookupEnv("DATABASE_URL")

	if !ok {
		log.Fatal("Error: DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error: Could not open database connection: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	closer := func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Error getting database instance: %v\n", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing the database: %v\n", err)
		}
	}
	return db, closer
}
