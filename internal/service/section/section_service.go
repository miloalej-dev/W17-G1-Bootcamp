package section

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type SectionService interface {
	FindAll() (v []models.Section, err error)
	FindByID(id int) (models.Section, error)
	Create(s models.Section) (models.Section, error)
	Update(s models.Section) (models.Section, error)
	Delete(id int) error
}
