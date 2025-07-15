package models

type Locality struct {
	Id       int `gorm:"primaryKey"`
	Locality *string
	Province *string
	Country  *string
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
