package service

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type SectionService interface {
	RetrieveAll() (v []models.Section, err error)
	Retrieve(id int) (models.Section, error)
	Register(s models.Section) (models.Section, error)
	Modify(s models.Section) (models.Section, error)
	PartialModify(id int, fields map[string]any) (models.Section, error)
	Remove(id int) error
}
