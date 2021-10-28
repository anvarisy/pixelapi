package models

import (
	"time"

	"gorm.io/gorm"
)

type Stuff struct {
	ID         uint           `gorm:"primaryKey"`
	StuffName  string         `json:"stuff_name" form:"stuff_name"`
	StuffPrice int64          `json:"stuff_price" form:"stuff_price"`
	StuffStock int64          `json:"stuff_stock" form:"stuff_stock"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	Deleted    gorm.DeletedAt `gorm:"index" json:"-"`
}

type StuffSerializer struct {
	StuffName  string `json:"stuff_name" form:"stuff_name"`
	StuffPrice int64  `json:"stuff_price" form:"stuff_price"`
	StuffStock int64  `json:"stuff_stock" form:"stuff_stock"`
}

type StuffID struct {
	ID []string `json:"stuff_id" form:"stuff_id"`
}

func (s *Stuff) CreateStuff(db *gorm.DB) (*Stuff, error) {
	err := db.Create(&s).Error
	if err != nil {
		return &Stuff{}, err
	}
	return s, nil
}

func (s *Stuff) GetAllStuff(db *gorm.DB) (*[]Stuff, error) {
	stuffes := []Stuff{}
	err := db.Model(s).Find(&stuffes).Error
	if err != nil {
		return &[]Stuff{}, err
	}
	return &stuffes, nil
}
