package models

type Locality struct {
	Id         int    `json:"locality_id" gorm:"primaryKey"`
	Locality   string `json:"locality_name"`
	ProvinceId int    `json:"province_id" gorm:"column:province_id"`
}

type LocalityDoc struct {
	Id       int    `json:"locality_id" gorm:"primaryKey"`
	Locality string `json:"locality_name"`
	Province string `json:"province_name,omitempty"`
	Country  string `json:"country_name,omitempty"`
}

type LocalitySellerCount struct {
	LocalityDoc
	SellerCount *int `json:"sellers_count,omitempty"`
}

// Used for carriers report
type LocalityCarrierCount struct {
	LocalityID    int    `gorm:"column:locality_id"`
	LocalityName  string `gorm:"column:locality_name"`
	TotalCarriers int    `gorm:"column:total_carriers"`
}

type Province struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Province  string `json:"province"`
	CountryId int    `json:"country_id" gorm:"column:country_id"`
}

func (Province) TableName() string {
	return "provinces"
}

func (Locality) TableName() string {
	return "localities"
}
