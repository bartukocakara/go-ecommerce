package service

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/entity"
	"ecommerce/internal/repository"
)

type CategoryService interface {
	GetCategories() ([]entity.Category, error)
	GetCategoryByID(id uint) (*entity.Category, error)
	CreateCategory(createDto dto.CreateCategoryDto) (*entity.Category, error)
	UpdateCategory(id uint, updateDto dto.UpdateCategoryDto) (*entity.Category, error)
	DeleteCategory(id uint) error
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *categoryService) GetCategories() ([]entity.Category, error) {
	return s.categoryRepository.GetCategories()
}

func (s *categoryService) GetCategoryByID(id uint) (*entity.Category, error) {
	return s.categoryRepository.GetCategoryByID(id)
}

func (s *categoryService) CreateCategory(createDto dto.CreateCategoryDto) (*entity.Category, error) {
	category := &entity.Category{
		Name: createDto.Name,
	}

	err := s.categoryRepository.CreateCategory(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) UpdateCategory(id uint, updateDto dto.UpdateCategoryDto) (*entity.Category, error) {
	category, err := s.categoryRepository.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	category.Name = updateDto.Name

	err = s.categoryRepository.UpdateCategory(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) DeleteCategory(id uint) error {
	category, err := s.categoryRepository.GetCategoryByID(id)
	if err != nil {
		return err
	}

	return s.categoryRepository.DeleteCategory(category)
}
