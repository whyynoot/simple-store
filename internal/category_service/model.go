package category_service

type Category struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	ApiName string `json:"api_name"`
}

type CategoryDTO struct {
	Name    string `json:"name"`
	ApiName string `json:"api_name"`
}
