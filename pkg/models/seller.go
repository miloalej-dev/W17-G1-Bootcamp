package models

// Seller is a struct that represents a seller
type Seller struct {
	CID       int
	Name      string
	Address   string
	Telephone string
	Locality  *Locality
}

// NewSeller is a function that creates a new seller
func NewSeller(cid int, name string, address string, telephone string, locality *Locality) *Seller {
	return &Seller{
		CID:       cid,
		Name:      name,
		Address:   address,
		Telephone: telephone,
		Locality:  locality,
	}
}
