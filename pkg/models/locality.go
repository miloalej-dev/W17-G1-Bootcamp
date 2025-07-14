package models

import "gorm.io/gorm"

type Locality struct {
	gorm.Model
	ID       int
	Locality *string
	Province *string
	Country  *string
}

func NewLocality(id int, locality, province, country *string) *Locality {
	return &Locality{
		ID:       id,
		Locality: locality,
		Province: province,
		Country:  country,
	}
}
