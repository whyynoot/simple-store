package order_service

type OrderServiceInterface interface {
	CreateRequest(*Request) error
	GetRequests() ([]Request, error)
}
