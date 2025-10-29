package product

import (
	"kai-shop-be/internal/domain/user" // ✅ dùng model user trong project

	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Product struct {
	ID            uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name          string         `gorm:"size:255; not null" json:"name"`
	Price         float64        `gorm:"not null" json:"price"`
	Stock         int            `gorm:"not null" json:"stock"`
	SKU           string         `gorm:"size:100; unique; not null" json:"sku"`
	Description   string         `gorm:"type:text" json:"description"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	ImageURLs     pq.StringArray `gorm:"type:text[]" json:"image_urls"` // 👈 thêm mảng ảnh
	CategoryID    *uuid.UUID     `gorm:"type:uuid" json:"category_id"`  // 👈 thêm FK,
	UserID        uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	CreatedByName string         `gorm:"size:100" json:"created_by"`
	User          user.User      `gorm:"foreignKey:UserID;references:ID" json:"-"` // 👈 ẩn khỏi JSON
}
