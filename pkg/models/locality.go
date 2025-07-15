package models

type Locality struct {
	Id          int     `json:"locality_id" gorm:"primaryKey"`
	Locality    *string `json:"locality_name"`
	Province    *string `json:"province_name,omitempty"`
	Country     *string `json:"country_name,omitempty"`
	SellerCount int     `json:"sellers_count" gorm:"default:0"`
}

func NewLocality(id int, locality, province, country *string) *Locality {
	return &Locality{
		Id:       id,
		Locality: locality,
		Province: province,
		Country:  country,
	}
}

func (Locality) TableName() string {
	return "localities"
}
