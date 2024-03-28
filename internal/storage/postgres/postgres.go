package postgres

import (
	"fmt"

	"github.com/VadimRight/Go_WebApp/internal/config"
	"github.com/VadimRight/Go_WebApp/internal/lib/logger/sl"
	"github.com/VadimRight/Go_WebApp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	cfg := config.MustLoad()
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Postgres_User, cfg.Postgres_Password, cfg.Postgres_Host, cfg.Postgres_Port, cfg.Postgres_Name)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		sl.Error(err)
	}
	db.AutoMigrate(
		&models.User{},
	)
	return db
}
