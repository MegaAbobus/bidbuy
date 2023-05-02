package entities

type Order struct {
	ID      uint
	UserID  uint
	Address string
	Price   float32
	Status  string
}
