package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/michaelgbenle/ZeinaMFI/internal/models"
	"github.com/michaelgbenle/ZeinaMFI/internal/ports"
	"log"
)

type Postgres struct {
	DB *gorm.DB
}

//NewDB create/returns a new instance of our Database
func NewDB(DB *gorm.DB) ports.Repository {
	return &Postgres{
		DB: DB,
	}
}

//Initialize opens the database, create tables if not created and populate it if its empty and returns a DB
func Initialize(dbURI string) (*gorm.DB, error) {

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	conn.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Blacklist{})
	log.Println("Database connection successful")
	return conn, nil
}
