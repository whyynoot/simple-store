package api

import (
	_ "awesomeProject/docs"
	"awesomeProject/internal/category_service"
	"awesomeProject/internal/config"
	"awesomeProject/internal/order_service"
	"awesomeProject/internal/product_service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
)

type Server struct {
	eng             *gin.Engine
	service         product_service.ProductServiceInterface
	categoryService category_service.CategoryServiceInterface
	requestService  order_service.OrderServiceInterface
	cfg             config.Cfg
}

func NewServer(cfg config.Cfg) *Server {
	server := &Server{
		eng:             gin.Default(),
		service:         product_service.NewProductService(cfg),
		categoryService: category_service.NewCategorySerice(cfg),
		requestService:  order_service.NewOrderService(cfg),
		cfg:             cfg,
	}

	server.eng.LoadHTMLGlob("template/*")
	server.eng.Static("/image", "./resources")

	server.eng.GET("/ping", server.Ping)
	server.eng.GET("/", server.GetAll)
	server.eng.GET("/product/:name", server.GetProduct)
	server.eng.GET("/admin", server.AdminAll)

	// API Product Handlers
	server.eng.DELETE("/api/product/delete/:name", server.DeleteProduct)
	server.eng.GET("/api/product/:name", server.GetAPIProduct)
	server.eng.POST("/api/product/create/", server.CreateProduct)
	server.eng.GET("/api/products", server.GetProducts)

	// API Category Handlers
	server.eng.GET("/api/categories", server.GetCategories)
	server.eng.GET("/api/category/:category/products", server.GetCategory)
	server.eng.DELETE("/api/category/delete/:category", server.DeleteCategory)
	server.eng.POST("/api/category/create", server.CreateCategory)

	// API Request Handler
	server.eng.POST("/api/request/create", server.CreateRequest)
	server.eng.GET("api/requests", server.GetRequests)

	// Swagger
	server.eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return server
}

func (server *Server) StartServer() {
	log.Println("Server start up")

	err := server.eng.Run()
	if err != nil {
		log.Fatal("Server down", err)
	}
}
