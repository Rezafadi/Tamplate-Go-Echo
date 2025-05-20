package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	CustomGormModel
	Name       string    `json:"name" gorm:"type: varchar(255);"`
	Email      string    `json:"email" gorm:"type: varchar(255);"`
	Password   string    `json:"password" gorm:"type:varchar(255);"`
	Gender     string    `json:"gender" gorm:"type: varchar(2);"`
	TglLahir   time.Time `json:"tgl_lahir" gorm:"type:timestamp;"`
	Phone      string    `json:"phone" gorm:"type: varchar(255);"`
	Image      string    `json:"image" gorm:"type: varchar(255);"`
	Address    string    `json:"address" gorm:"type: text;"`
	IsVerify   bool      `json:"is_verify" gorm:"type: bool;"`
	RoleID     int       `json:"role_id" gorm:"type: int8;"`
	Status     int       `json:"status" gorm:"type: int8;"`
	Prov       int       `json:"prov" gorm:"type: int8;"`
	Kab        int       `json:"kab" gorm:"type: int8;"`
	Kec        int       `json:"kec" gorm:"type: int8;"`
	Kel        string    `json:"kel" gorm:"type: varchar(255);"`
	PostalCode string    `json:"postal_code" gorm:"type: varchar(255);"`
}

type CustomGormModel struct {
	ID        uint            `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
