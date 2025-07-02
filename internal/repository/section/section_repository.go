package repository

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type SectionRepository interface {
	FindAll() (v map[int]models.Section, err error)
	FindByID(id int) (models.Section, error)
	Add(s models.Section) (models.Section, error)
	Update(s models.Section) (models.Section, error)
	Delete(id int) (models.Section, error)
}
