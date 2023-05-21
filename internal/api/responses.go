package api

import (
	"awesomeProject/internal/category_service"
	"awesomeProject/internal/order_service"
	"awesomeProject/internal/product_service"
)

type BaseResponse struct {
	Status string `json:"status"`
}

type ProductResponse struct {
	Product product_service.UserProduct `json:"product"`
	Status  string                      `json:"status"`
}

type ProductsResponse struct {
	Products []product_service.UserProduct `json:"products"`
	Status   string                        `json:"status"`
}

type CategoriesResponse struct {
	Categories []category_service.Category `json:"categories"`
	Status     string                      `json:"status"`
}

type RequestsResponse struct {
	Requests []order_service.Request `json:"requests"`
	Status   string                  `json:"status"`
}
