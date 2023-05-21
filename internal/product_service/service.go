package product_service

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/database"
	"context"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"log"
)

type ProductService struct {
	DataBaseClient database.Client
	//Products []Product
}

func NewProductService(cfg config.Cfg) ProductServiceInterface {
	db, err := database.NewClient(context.Background(), cfg.DataBaseURL)
	if err != nil {
		log.Fatal(err)
	}
	return &ProductService{
		DataBaseClient: db,
		//Products: []Product{
		//{
		// Name:
		//}
		//}
	}
}

func (p *ProductService) GetAll() ([]UserProduct, error) {
	query := sqlbuilder.Select(
		"name_ru",
		"api_name",
		"description",
		"image_path",
		"price").From("public.product").String()
	productRows, err := p.DataBaseClient.Query(context.Background(), query)
	if err != nil {
		log.Println("Failed to query rows", err)
		return nil, err
	}

	var Products []Product

	for productRows.Next() {
		var product Product
		if err := productRows.Scan(&product.Name, &product.ApiName, &product.Description,
			&product.Image, &product.Price); err != nil {
			return nil, err
		}
		Products = append(Products, product)
	}

	products := make([]UserProduct, 0, len(Products))
	for _, product := range Products {
		products = append(products, product.ToUser())
	}

	//products := make([]UserProduct, 0, len(p.Products))
	//for _, product := range p.Products {
	//	products = append(products, product.ToUser())
	//}

	return products, nil
}

func (p *ProductService) GetProduct(name string) (UserProduct, error) {
	sb := sqlbuilder.NewSelectBuilder()
	query := sb.Select(
		"name_ru",
		"api_name",
		"description",
		"image_path",
		"price").From("public.product").Where(fmt.Sprintf("api_name = '%s'", name)).String()
	productRows := p.DataBaseClient.QueryRow(context.Background(), query)

	var product Product
	if err := productRows.Scan(&product.Name, &product.ApiName, &product.Description,
		&product.Image, &product.Price); err != nil {
		return UserProduct{}, err
	}

	return product.ToUser(), nil
}

func (p *ProductService) CreateProduct(apiName, name, description, image string, price int) error {
	builder := sqlbuilder.PostgreSQL.NewInsertBuilder()
	query, args := builder.InsertInto("public.product").Cols(
		"name_ru",
		"api_name",
		"description",
		"image_path",
		"price").Values(name, apiName, description,
		image, price).Build()

	_, err := p.DataBaseClient.Query(context.Background(), query, args...)
	if err != nil {
		log.Println("error creating product", err)
		return err
	}

	return nil
}

func (p *ProductService) DeleteProduct(name string) error {
	sb := sqlbuilder.PostgreSQL.NewDeleteBuilder()
	query, args := sb.DeleteFrom("public.product").Where(sb.Equal("api_name", name)).Build()

	_, err := p.DataBaseClient.Query(context.Background(), query, args...)
	if err != nil {
		return err
	}

	return nil
}
