package product_service

type ProductServiceInterface interface {
	GetAll() ([]UserProduct, error)
	GetProduct(name string) (UserProduct, error)
	DeleteProduct(name string) error
	CreateProduct(apiName, name, description, image string, price int) error
}
