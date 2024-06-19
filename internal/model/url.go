package model

import (
	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Link string `json:"url" gorm:"size:200;not null;"`
	Hash string `json:"hash" gorm:"size:20;not null;"`
}

func (u *Url) BeforeCreate(tx *gorm.DB) (err error) {
	sid, _ := shortid.New(1, shortid.DefaultABC, 1234)
	hash, err := sid.Generate()
	if err != nil {
		return
	}
	u.Hash = "bale.ai/" + hash
	return nil
}
