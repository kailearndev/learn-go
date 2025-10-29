package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(u *User) error
	FindByEmail(email string) (User, error)
	FindByID(id uuid.UUID) (User, error)
	UpdateUser(u *User) error
	DeleteUser(id uuid.UUID) error
	FindAll(limit, offset int) ([]User, error)
	Count() (int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(u *User) error {
	return r.db.Create(u).Error
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

//	func (r *repository) Count() (int64, error) {
//		var count int64
//		if err := r.db.Model(&Product{}).Count(&count).Error; err != nil {
//			return 0, err
//		}
//		return count, nil
//	}
func (r *repository) FindByID(id uuid.UUID) (User, error) {
	var user User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

//	func (r *repository) Create(p *Product) error {
//		return r.db.Create(p).Error
//	}
func (r *repository) UpdateUser(p *User) error {
	return r.db.Save(p).Error
}
func (r *repository) DeleteUser(id uuid.UUID) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}

func (r *repository) FindAll(limit, offset int) ([]User, error) {
	var users []User
	err := r.db.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

func (r *repository) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
