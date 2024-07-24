package postgres

import (
	"log"

	"github.com/nuhmanudheent/hotel-booking/payment-service/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := "host=localhost user=microservice_project password=Nuhman@456 dbname=user_service port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect with postgres......")
	}
	db.AutoMigrate(domain.PaymentDetails{})
	return db
}
