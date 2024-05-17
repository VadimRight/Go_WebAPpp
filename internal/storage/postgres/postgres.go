package postgres

import (
	"fmt"

	"github.com/VadimRight/Url-Saver/internal/config"
	"github.com/VadimRight/Url-Saver/models"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//	"database/sql"
)

var cfg = config.MustLoad()
var dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Postgres_User, cfg.Postgres_Password, cfg.Postgres_Host, cfg.Postgres_Port, cfg.Postgres_Name)


type GORMStorage struct {
	db *gorm.DB
}

func (g *GORMStorage) InitDB() (*GORMStorage, error) {
	const op = "storage.Posgres.New"
	var db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	errMigration := db.AutoMigrate(
		&models.URL{},
	)
	if errMigration != nil {
		return nil, fmt.Errorf("%s: %w", op, errMigration)
	}
	return &GORMStorage{db: db}, nil
}

func (g *GORMStorage) GetURL(id uuid.UUID) (*GORMStorage, error) {	
	const op = "storage.Posgres.New"
	var db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	url := models.URL{}
	url_id := []string{}
	query := db.First(&url, url_id)
	return &GORMStorage{db: query}, nil
}

func (g *GORMStorage) SaveURL(urltosave string, alias_name string) (string, error) {
	const op = "storage.Posgres.New"
	var db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	id := uuid.New()
	new_url := models.URL{Id: id, Url: urltosave, Alias: alias_name}
	db.Create(new_url)
	uuid_string := id.String()
	return uuid_string, nil

}

func (g *GORMStorage)TestAddUrl() (*GORMStorage, error) {
	const op = "storage.Posgres.New"
	var db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)	
	}
	id := uuid.New()
	url := models.URL{Id: id, Url: "test.com", Alias: "test"}
	result := db.Create(url)
	defer db.Delete(url)
	fmt.Printf("\nTest session is done!\n")
	return &GORMStorage{db: result}, nil
}

func (g *GORMStorage) DeleteURL(id uuid.UUID) (*GORMStorage, error) {
	const op = "storage.Posgres.New"
	var db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)	
	}
	db.Delete(&models.URL{}, id)
	return &GORMStorage{db: db}, nil
}
