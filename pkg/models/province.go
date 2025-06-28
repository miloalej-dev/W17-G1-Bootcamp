package models

type Province struct {
	Name string
}

func NewProvince(name string) *Province {
	return &Province{
		Name: name,
	}
}
