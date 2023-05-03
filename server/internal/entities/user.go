package entities

type User struct {
	ID           uint    `json:"id"`
	PhoneNumber  string  `json:"phoneNumber"`
	Email        string  `json:"email"`
	AmountRansom float32 `json:"amountOfRansom"`
}
