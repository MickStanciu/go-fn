package models

type Flower struct {
	ProductId    int     `json:"productId"`
	Category     string  `json:"category"`
	Price        float64 `json:"price"`
	Instructions string  `json:"instructions"`
	Photo        string  `json:"photo,omitempty"`
	Name         string  `json:"name"`
}

type FlowerResponse struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Photo *string `json:"photo,omitempty"`
}
