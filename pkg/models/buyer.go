package models

type BuyerAtributtes struct {
	cardNumberId string
	firstName    string
	LastName     string
}

type Buyer struct {
	id int
	BuyerAtributtes
}

type BuyerDoc struct {
	id           int    `json:"id"`
	cardNumberId string `json:"card_number_id"`
	firstName    string `json:"first_name"`
	LastName     string `json:"Last_name"`
}
