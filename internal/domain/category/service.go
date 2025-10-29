package category

import (
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	ListCategories(limit, offset int) ([]Category, int64, error)
	GetCategoryByID(id uuid.UUID) (Category, error)
	CreateCategory(p CategoryRequest) (Category, error)
	UpdateCategory(id uuid.UUID, input CategoryRequest) (Category, error)
	DeleteCategory(id uuid.UUID) error
	CountCategories() (int64, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) ListCategories(limit, offset int) ([]Category, int64, error) {
	items, err := s.repo.FindAll(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	count, err := s.repo.Count()
	if err != nil {
		return nil, 0, err
	}
	return items, count, nil
}

func (s *service) GetCategoryByID(id uuid.UUID) (Category, error) {

	existing, err := s.repo.FindByID(id)
	if err != nil {
		return Category{}, err
	}
	return existing, nil
}

func (s *service) CreateCategory(req CategoryRequest) (Category, error) {
	// validation cơ bản
	if req.Name == "" {
		return Category{}, errors.New("tên danh mục không được để trống")
	}

	// map DTO -> model
	category := Category{
		Name:        req.Name,
		Description: req.Description,
	}

	// gọi repo
	if err := s.repo.Create(&category); err != nil {
		return Category{}, err
	}

	return category, nil
}

func (s *service) UpdateCategory(id uuid.UUID, p CategoryRequest) (Category, error) {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return Category{}, err
	}
	existing.Name = p.Name
	existing.Description = p.Description
	existing.Description = p.Description
	if err := s.repo.Update(&existing); err != nil {
		return Category{}, err
	}
	return existing, nil
}
func (s *service) DeleteCategory(id uuid.UUID) error {
	// optionally: check existence first

	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *service) CountCategories() (int64, error) {
	return s.repo.Count()
}
