package entities

type Worker struct {
	ID                  uint   `json:"id"`
	Name                string `json:"name"`
	PasportSeriesNumber string `json:"pasportSeriesNumber"`
	PasportIssued       string `json:"pasportIssued"`
	PhoneNumber         string `json:"phoneNumber"`
}
