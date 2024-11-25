package model

type Customer struct {
	Base
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Phone    string `json:"phone" gorm:"type:varchar(15);unique"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}

