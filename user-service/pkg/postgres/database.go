package postgres

import (
	"log"
	"os"

	"github.com/nuhmanudheent/hotel-booking/user-service/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect with postgres......")
	}
	db.AutoMigrate(domain.User{})
	return db
}
