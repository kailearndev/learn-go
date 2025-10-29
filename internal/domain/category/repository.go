package category

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(limit, offset int) ([]Category, error)
	FindByID(id uuid.UUID) (Category, error)
	Create(p *Category) error
	Update(p *Category) error
	Delete(id uuid.UUID) error
	Count() (int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll(limit, offset int) ([]Category, error) {
	var categories []Category
	if err := r.db.Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
func (r *repository) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&Category{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
func (r *repository) FindByID(id uuid.UUID) (Category, error) {
	var category Category
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		return Category{}, err
	}
	return category, nil
}

func (r *repository) Create(p *Category) error {
	return r.db.Create(p).Error
}
func (r *repository) Update(p *Category) error {
	return r.db.Save(p).Error
}
func (r *repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Category{}, "id = ?", id).Error
}
