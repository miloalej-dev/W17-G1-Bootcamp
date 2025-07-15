package models

type Locality struct {
	Id       int     `json:"locality_id" gorm:"primaryKey"`
	Locality *string `json:"locality_name"`
	Province *string `json:"province_name,omitempty"`
	Country  *string `json:"country_name,omitempty"`
}

type LocalitySellerCount struct {
	Locality
	SellerCount *int `json:"sellers_count,omitempty"`
}

func (Locality) TableName() string {
	return "localities"
}
