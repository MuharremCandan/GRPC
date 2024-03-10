package db

import (
	"fmt"
	"log"
	"test-grpc-project/pkg/config"
	"test-grpc-project/pkg/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb(config *config.Config) *gorm.DB {
	// PostgreSQL connection parameters
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Pass,
		config.Database.Name)
	// Open a connection to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the PostgreSQL database!")

	// Migrate the schema (create the 'users' table)
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
