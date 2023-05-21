package order_service

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/database"
	"context"
	"github.com/huandu/go-sqlbuilder"
	"log"
)

type OrderService struct {
	DataBaseClient database.Client
}

func (o OrderService) CreateRequest(request *Request) error {
	builder := sqlbuilder.PostgreSQL.NewInsertBuilder()
	query, args := builder.InsertInto("public.order").Cols(
		"phone",
		"description").Values(request.Phone, request.Information).Build()

	_, err := o.DataBaseClient.Query(context.Background(), query, args...)
	if err != nil {
		log.Println("error creating category", err)
		return err
	}

	return nil
}

func (o OrderService) GetRequests() ([]Request, error) {
	query := sqlbuilder.Select(
		"phone",
		"description").From("public.order").String()
	categoryRows, err := o.DataBaseClient.Query(context.Background(), query)
	if err != nil {
		log.Println("Failed to query rows", err)
		return nil, err
	}

	var Requests []Request

	for categoryRows.Next() {
		var category Request
		if err := categoryRows.Scan(&category.Phone, &category.Information); err != nil {
			return nil, err
		}
		Requests = append(Requests, category)
	}

	return Requests, nil
}

func NewOrderService(cfg config.Cfg) OrderServiceInterface {
	db, err := database.NewClient(context.Background(), cfg.DataBaseURL)
	if err != nil {
		log.Fatal(err)
	}
	return &OrderService{
		DataBaseClient: db,
	}
}
