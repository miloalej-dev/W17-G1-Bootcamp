package models

type Buyer struct {
	Id           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

// NewBuyer is a function that creates a new buyer
func NewBuyer(id int, cardNumberId string, firstName string, lastName string) *Buyer {
	return &Buyer{
		Id:           id,
		CardNumberId: cardNumberId,
		FirstName:    firstName,
		LastName:     lastName,
	}
}
