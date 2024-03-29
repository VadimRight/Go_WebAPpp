package postgres

import (
	"fmt"
	// "log"

	"github.com/VadimRight/Go_WebApp/internal/config"
	"github.com/VadimRight/Go_WebApp/internal/lib/logger/sl"
	"github.com/VadimRight/Go_WebApp/models"

	// "github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var cfg = config.MustLoad()
var dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Postgres_User, cfg.Postgres_Password, cfg.Postgres_Host, cfg.Postgres_Port, cfg.Postgres_Name)

func InitDB() *gorm.DB {
	var db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		sl.Error(err)
	}
	db.AutoMigrate(
		&models.URL{},
	)
	return db
}

func GetURL(id string) *gorm.DB {
	var db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		sl.Error(err)
	}
	url := models.URL{}
	url_id := []string{}
	query := db.First(&url, url_id)
	return query
}
