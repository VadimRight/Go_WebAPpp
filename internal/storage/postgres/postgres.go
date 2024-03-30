package postgres

import (
	"fmt"

	"github.com/VadimRight/Go_WebApp/internal/config"
	"github.com/VadimRight/Go_WebApp/internal/lib/logger/sl"
	"github.com/VadimRight/Go_WebApp/models"

	"github.com/google/uuid"
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

func AddURL(url string, site_name string) *gorm.DB {
	var db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{}) if err != nil {
		sl.Error(err)
	}
	id := uuid.New()
	new_url := models.URL{Id: id, Url: url, Site: site_name}
	result := db.Create(new_url)
	return result

}

func TestAddUrl() *gorm.DB {
	var db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		sl.Error(err)
	}
	id := uuid.New()
	url := models.URL{Id: id, Url: "test.com", Site: "test"}
	result := db.Create(url)
	defer db.Delete(url)
	fmt.Printf("\nTest session is done!\n")
	return result
}

