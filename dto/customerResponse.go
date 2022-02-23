package dto

type CustomerResponse struct {
	ID          string `json:"customer_id"`
	Name        string `json:"fullname"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}
