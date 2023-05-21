package category_service

import "awesomeProject/internal/product_service"

type CategoryServiceInterface interface {
	GetCategories() ([]Category, error)
	GetCategory(name string) ([]product_service.UserProduct, error)
	DeleteCategory(name string) error
	CreateCategory(category *CategoryDTO) error
	AddProductToCategory(apiName string) error
}
