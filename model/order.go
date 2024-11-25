package model

type Order struct {
	Base
	Customer      Customer      `json:"customer" gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CustomerID    uint          `json:"customer_id" gorm:"index"` // Jika perlu menyimpan ID customer
	TotalPrice    float32       `json:"total_price" gorm:"not null"`
	Discount      float32       `json:"discount" gorm:"default:0"`
	OrderStatus   string        `json:"order_status" gorm:"type:varchar(50);not null"`
	PaymentMethod string        `json:"payment_method" gorm:"type:varchar(50);not null"`
	PaymentStatus string        `json:"payment_status" gorm:"type:varchar(50);not null"`
	Address       string        `json:"address" gorm:"type:text"`
	OrderDetails  []OrderDetail `json:"order_details" gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type OrderDetail struct {
	Base
	OrderID          uint `json:"order_id" gorm:"not null;index"`
	ProductVariantID uint `json:"product_variant_id" gorm:"not null;index"`
	Quantity         int  `json:"quantity" gorm:"not null"`
}
