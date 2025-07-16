package models

// Seller is a struct that represents a seller company
type Seller struct {
	Id         int    `json:"id" gorm:"primaryKey"` // Id is the identifier of the seller company
	Name       string `json:"name"`                 // Name is the name of the seller company
	Address    string `json:"address"`              // Address is the address of the seller company
	Telephone  string `json:"telephone"`            // Telephone is the telephone number of the seller company
	LocalityId int    `json:"locality_id"`          // LocalityId is the locality id of the seller company
}

// NewSeller is a function that creates a new seller
func NewSeller(id int, name string, address string, telephone string) *Seller {
	return &Seller{
		Id:        id,
		Name:      name,
		Address:   address,
		Telephone: telephone,
	}
}
