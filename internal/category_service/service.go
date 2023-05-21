package category_service

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/database"
	"awesomeProject/internal/product_service"
	"context"
	"github.com/huandu/go-sqlbuilder"
	"log"
)

type CategoryService struct {
	DataBaseClient database.Client
}

func (c CategoryService) GetCategories() ([]Category, error) {
	query := sqlbuilder.Select(
		"id",
		"name_ru",
		"api_name").From("public.category").String()
	categoryRows, err := c.DataBaseClient.Query(context.Background(), query)
	if err != nil {
		log.Println("Failed to query rows", err)
		return nil, err
	}

	var Categories []Category

	for categoryRows.Next() {
		var category Category
		if err := categoryRows.Scan(&category.ID, &category.Name, &category.ApiName); err != nil {
			return nil, err
		}
		Categories = append(Categories, category)
	}

	return Categories, nil
}

func (c CategoryService) GetCategory(name string) ([]product_service.UserProduct, error) {
	builder := sqlbuilder.PostgreSQL.NewSelectBuilder()
	query, args := builder.Select(
		"p.name_ru",
		"p.api_name",
		"p.description",
		"p.image_path",
		"p.price").From("public.product p").Join("public.category cat", "cat.id = p.category_id").Where(builder.Equal("cat.api_name", name)).Build()
	productRows, err := c.DataBaseClient.Query(context.Background(), query, args...)
	if err != nil {
		log.Println("error getting products by category", err)
		return nil, err
	}

	var Products []product_service.UserProduct

	for productRows.Next() {
		var product product_service.UserProduct
		if err := productRows.Scan(&product.Name, &product.ApiName, &product.Description,
			&product.Image, &product.Price); err != nil {
			return nil, err
		}
		Products = append(Products, product)
	}

	return Products, nil
}

func (c CategoryService) DeleteCategory(name string) error {
	sb := sqlbuilder.PostgreSQL.NewDeleteBuilder()
	query, args := sb.DeleteFrom("public.category").Where(sb.Equal("api_name", name)).Build()

	_, err := c.DataBaseClient.Query(context.Background(), query, args...)
	if err != nil {
		log.Println("error deleting category", err)
		return err
	}

	return nil
}

func (c CategoryService) CreateCategory(category *CategoryDTO) error {
	builder := sqlbuilder.PostgreSQL.NewInsertBuilder()
	query, args := builder.InsertInto("public.category").Cols(
		"name_ru",
		"api_name").Values(category.Name, category.ApiName).Build()

	_, err := c.DataBaseClient.Query(context.Background(), query, args...)
	if err != nil {
		log.Println("error creating category", err)
		return err
	}

	return nil
}

func (c CategoryService) AddProductToCategory(apiName string) error {
	//TODO implement me
	panic("implement me")
}

func NewCategorySerice(cfg config.Cfg) CategoryServiceInterface {
	db, err := database.NewClient(context.Background(), cfg.DataBaseURL)
	if err != nil {
		log.Fatal(err)
	}
	return &CategoryService{
		DataBaseClient: db,
	}
}
