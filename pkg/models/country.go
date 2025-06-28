package models

type Country struct {
	Name string
}

func NewCountry(name string) *Country {
	return &Country{
		Name: name,
	}
}
