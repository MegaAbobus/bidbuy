package entities

type Product struct {
	ID          uint    `json:"id"`
	Description string  `json:"description"`
	Body        string  `json:"body"`
	IsAvailable bool    `json:"isAvailable"`
	Price       float32 `json:"price"`
	Rating      float32 `json:"rating"`
}
