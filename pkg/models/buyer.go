package models

type BuyerAtributtes struct {
	CardNumberId string
	FirstName    string
	LastName     string
}

type Buyer struct {
	Id int
	BuyerAtributtes
}

type BuyerDoc struct {
	Id           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"Last_name"`
}
