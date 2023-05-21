package api

import (
	"awesomeProject/internal/category_service"
	"awesomeProject/internal/order_service"
	"awesomeProject/internal/product_service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (server *Server) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (server *Server) GetAll(c *gin.Context) {
	products, err := server.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println("Error on product handler", err)
		return
	}

	c.HTML(http.StatusOK, "all.page.tmpl", gin.H{
		"Products": products,
	})
}

func (server *Server) GetProduct(c *gin.Context) {
	productName := c.Param("name")

	product, err := server.service.GetProduct(productName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		log.Println("Error on product handler", err, productName)
		return
	}

	c.HTML(http.StatusOK, "product.page.tmpl", product)
}

func (server *Server) AdminAll(c *gin.Context) {
	products, err := server.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println("Error on product handler", err)
		return
	}

	c.HTML(http.StatusOK, "admin.page.tmpl", gin.H{
		"Products": products,
	})
}

// GetProducts ... Get all products
// @Summary Getting all products info
// @Tags Product
// @Success 200 {object} ProductsResponse
// @Router /api/products [get]
func (server *Server) GetProducts(c *gin.Context) {
	products, err := server.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println("Error on product handler", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
		"status":   "ok",
	})
}

// GetAPIProduct ... Get product
// @Summary Getting product info
// @Tags Product
// @Param name path string true "name"
// @Success 200 {object} ProductResponse
// @Router /api/product/{name} [get]
func (server *Server) GetAPIProduct(c *gin.Context) {
	productName := c.Param("name")

	product, err := server.service.GetProduct(productName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		log.Println("Error on product handler", err, productName)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
		"status":  "ok",
	})
}

// DeleteProduct ... Delete product by name
// @Summary Deletion of product
// @Tags Product
// @Param name path string true "name"
// @Success 200 {object} BaseResponse
// @Router /api/product/delete/{name} [delete]
func (server *Server) DeleteProduct(c *gin.Context) {
	productName := c.Param("name")

	err := server.service.DeleteProduct(productName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		log.Println("Error on product handler", err, productName)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// CreateProduct ... Create Product
// @Summary Creation of product
// @Tags Product
// @Param request body product_service.ProductDTO true "product parameters"
// @Success 200 {object} BaseResponse
// @Router /api/product/create/ [post]
func (server *Server) CreateProduct(c *gin.Context) {
	product := &product_service.ProductDTO{}
	err := c.BindJSON(product)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		log.Println("error getting data from request", err)
		return
	}

	err = server.service.CreateProduct(product.ApiName, product.Name, product.Description, product.Image,
		product.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println("Error on creating product handler", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "created",
	})
}

// GetCategories ... Get all categories
// @Summary Getting all categories
// @Tags Category
// @Success 200 {object} CategoriesResponse
// @Router /api/categories [get]
func (server *Server) GetCategories(c *gin.Context) {
	categories, err := server.categoryService.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println("Error on category handler", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
		"status":     "ok",
	})
}

// GetCategory ... Get products for category
// @Summary Get products for category
// @Tags Category
// @Param category path string true "category"
// @Success 200 {object} ProductsResponse
// @Router /api/category/{category}/products [get]
func (server *Server) GetCategory(c *gin.Context) {
	categoryName := c.Param("category")
	products, err := server.categoryService.GetCategory(categoryName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println("Error on category handler", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
		"status":   "ok",
	})
}

// DeleteCategory ... Delete category by name
// @Summary delete category by name
// @Tags Category
// @Param category path string true "category"
// @Success 200 {object} BaseResponse
// @Router /api/category/delete/{category} [delete]
func (server *Server) DeleteCategory(c *gin.Context) {
	categoryName := c.Param("category")
	err := server.categoryService.DeleteCategory(categoryName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println("Error on category handler", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// CreateCategory ... Create category by body
// @Summary Create category by body
// @Tags Category
// @Param request body category_service.CategoryDTO true "category dto"
// @Success 200 {object} BaseResponse
// @Router /api/category/create [post]
func (server *Server) CreateCategory(c *gin.Context) {
	category := &category_service.CategoryDTO{}
	err := c.BindJSON(category)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		log.Println("error getting data from request", err)
		return
	}

	err = server.categoryService.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println("Error on category handler", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// CreateRequest ... Creating new request
// @Summary Creating new request
// @Tags Request
// @Param request body order_service.Request true "request DTO"
// @Success 200 {object} BaseResponse
// @Router /api/request/create [post]
func (server *Server) CreateRequest(c *gin.Context) {
	request := &order_service.Request{}
	err := c.BindJSON(request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		log.Println("error getting data from request", err)
		return
	}

	err = server.requestService.CreateRequest(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println("Error on request handler", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// GetRequests ... Getting all requests
// @Summary Getting all requests
// @Tags Request
// @Success 200 {object} BaseResponse
// @Router /api/requests [get]
func (server *Server) GetRequests(c *gin.Context) {
	requests, err := server.requestService.GetRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		log.Println("Error on request handler", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"requests": requests,
		"status":   "ok",
	})
}
