package models

// Carrier is a struct that represents a carrier in JSON format
type Carrier struct {
	ID				int		`json:"id"`
	CId				string	`json:"cid" gorm:"column:cid"`
	CompanyName		string	`json:"company_name" gorm:"column:name"`
	Address			string	`json:"address"`
	Telephone		string	`json:"telephone"`
	LocalityId		int		`json:"locality_id"`
}

func NewCarrier(id int, cid, company_name, address, telephone string, locality_id int) *Carrier {
	return &Carrier{
		ID: id,
		CId: cid,
		CompanyName: company_name,
		Address: address,
		Telephone: telephone,
		LocalityId: locality_id,
	}
}
