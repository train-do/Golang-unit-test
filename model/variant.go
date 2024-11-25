package model

type Variant struct {
	Base
	Name string `json:"name" gorm:"type:varchar(20);not null"`
}
