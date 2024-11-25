package model

type Category struct {
	Base
	Name string `json:"name" gorm:"type:varchar(100);not null"`
}

