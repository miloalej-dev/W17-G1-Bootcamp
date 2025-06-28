package models

type Locality struct {
	name string
	Province
	Country
}

func NewLocality(name string, province *Province, country *Country) *Locality {
	return &Locality{
		name:     name,
		Province: *province,
		Country:  *country,
	}
}
