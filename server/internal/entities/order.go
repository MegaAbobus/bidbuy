package entities

type Order struct {
	ID      uint    `json:"id"`
	UserID  uint    `json:"userId"`
	Address string  `json:"address"`
	Price   float32 `json:"price"`
	Status  string  `json:"status"`
}
