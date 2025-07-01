package section

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type SectionService interface {
	FindAll() (v map[int]models.Section, err error)
	FindByID(id int) (models.Section, error)
	Add(s models.Section) (models.Section, error)
	Update(s models.Section) (models.Section, error)
	Delete(s models.Section) (models.Section, error)
}
