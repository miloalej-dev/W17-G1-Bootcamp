package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	// Default values
	if host == "" {
		host = "localhost"
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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable logging
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		TranslateError: true,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
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
