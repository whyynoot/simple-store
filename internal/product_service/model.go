package product_service

type Product struct {
	ApiName     string
	Name        string
	Description string
	Price       int
	Image       string
	Stock       int
}

func (p *Product) ToUser() UserProduct {
	return UserProduct{
		ApiName:     p.ApiName,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Image:       p.Image,
	}
}

type UserProduct struct {
	ApiName     string `json:"api_name"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
}

type ProductDTO struct {
	ApiName     string `json:"api_name"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
}
