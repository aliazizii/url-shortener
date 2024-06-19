package store

import (
	"errors"
	"log"

	"github.com/aliazizii/url-shotener/internal/model"
	"gorm.io/gorm"
)

type SQL struct {
	DB *gorm.DB
}

func NewSQL(db *gorm.DB) SQL {
	if err := db.AutoMigrate(&model.Url{}); err != nil {
		log.Fatal(err)
	}
	return SQL{
		DB: db,
	}
}

func (sql SQL) FindByHash(Hash string) (string, error) {
	var url *model.Url
	err := sql.DB.Where("hash = ?", Hash).Take(&url).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("Url not found")
		}
		return "", err
	}
	return url.Link, nil
}

func (sql SQL) Insert(url model.Url) (string, error) {
	err := sql.DB.Create(&url).Error
	if err != nil {
		return "", err
	}
	return url.Hash, err
}
