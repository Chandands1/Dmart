package model

import(
	"time"
	"gorm.io/gorm"
	) 


type User struct{
	ID uint `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:100;not null" json:"name"`
	Email string `gorm:"size:150;unique;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"`
	Phone string `gorm:"size:15;unique" json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	
}