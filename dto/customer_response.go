package dto

type CustomerResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code"`
	Status      string `json:"status"`
}
