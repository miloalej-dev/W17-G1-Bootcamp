package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type configuration struct {
	user     string
	password string
	host     string
	port     string
	database string
}

func configure() (*configuration, error) {
	// Set environment variables or load from .env file
	_ = os.Setenv("DB_USER", "frescos_user")
	_ = os.Setenv("DB_PASSWORD", "password")
	_ = os.Setenv("DB_HOST", "database")
	_ = os.Setenv("DB_NAME", "frescos")
	_ = os.Setenv("DB_PORT", "3306")

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	// Default values
	if host == "" {
		host = "database"
	}
	if port == "" {
		port = "3306"
	}
	if database == "" {
		database = "frescos"
	}

	if user == "" {
		return nil, fmt.Errorf("user enviroment variable is required")
	}

	if password == "" {
		return nil, fmt.Errorf("password enviroment variable is required")
	}

	return &configuration{
		user:     user,
		password: password,
		host:     host,
		port:     port,
		database: database,
	}, nil
}

// NewConnection creates a new GORM database connection
func NewConnection() (*gorm.DB, error) {
	// Create a new configuration struct
	config, err := configure()

	if err != nil {
		return nil, err
	}

	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.user,
		config.password,
		config.host,
		config.port,
		config.database,
	)

	var db *gorm.DB
	maxRetries := 10
	retryDelay := 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // Enable logging
			NowFunc: func() time.Time {
				return time.Now().Local()
			},
		})
		if err == nil {
			// Connection successful
			log.Println("Successfully connected to the database")
			break
		}
		log.Printf("Failed to connect to database (attempt %d/%d): %v. Retrying in %v...", i+1, maxRetries, err, retryDelay)
		time.Sleep(retryDelay)
	}

	// If the connection is still nil after all retries, return the final error.
	if db == nil {
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}
	// Connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
