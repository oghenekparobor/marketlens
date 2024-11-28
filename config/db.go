package config

import (
	"fmt"
	"log"
	"oghenekparobor/market-lens/models"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // PostgreSQL driver for database/sql
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB

func connect(cfg *PostgresConfig) *gorm.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	// Replace with your actual database credentials
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	return db
}

func InitDB(cfg *PostgresConfig) {
	DB = connect(cfg)

	// Enable the UUID extension
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// Migrate the schema
	DB.AutoMigrate(&models.UserRole{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.UserSession{})
	DB.AutoMigrate(&models.PasswordReset{})
	DB.AutoMigrate(&models.Address{})
	DB.AutoMigrate(&models.OTP{})

	seedUserRoles(DB)
}

// SeedUserRoles inserts default roles into the user_roles table
func seedUserRoles(db *gorm.DB) error {
	roles := []models.UserRole{
		{RoleName: "customer"},
		{RoleName: "carter"},
		{RoleName: "store"},
		{RoleName: "admin"},
	}

	for _, role := range roles {
		if err := db.FirstOrCreate(&role, models.UserRole{RoleName: role.RoleName}).Error; err != nil {
			log.Printf("Error inserting role %s: %v", role.RoleName, err)
			return err
		}
	}

	return nil
}

// // NewPostgresStorage creates a new PostgreSQL connection and sets up migration.
// func InitPostgresStorage(cfg *PostgresConfig) (*sql.DB, error) {
// 	// PostgreSQL connection string
// 	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
// 		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

// 	// Open a new database connection
// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		log.Fatalf("Failed to open database: %v", err)
// 		return nil, err
// 	}

// 	// Ping to verify the connection is working
// 	if err := db.Ping(); err != nil {
// 		log.Fatalf("Failed to ping database: %v", err)
// 		return nil, err
// 	}

// 	log.Println("Connected to PostgreSQL database successfully!")

// 	// Perform database migration
// 	err = runMigrations(dsn)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

// // runMigrations runs database migrations from a file source (local files in the migrations folder)
// func runMigrations(dsn string) error {
// 	m, err := migrate.New(
// 		"file://migrations", // Path to migrations directory
// 		dsn,                 // PostgreSQL connection string for migrations
// 	)

// 	if err != nil {
// 		log.Fatalf("Failed to create migrate instance: %v", err)
// 		return err
// 	}

// 	// Apply all up migrations
// 	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
// 		log.Fatalf("Failed to run migrations: %v", err)
// 		return err
// 	}

// 	log.Println("Database migrations applied successfully!")
// 	return nil
// }
