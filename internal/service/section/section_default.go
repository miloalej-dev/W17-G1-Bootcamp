package section

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository/section"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type SectionDefault struct {
	// rp is the repository that will be used by the service
	rp repository.SectionRepository
}

func NewSectionDefault(rp repository.SectionRepository) *SectionDefault {
	return &SectionDefault{rp: rp}
}

func (s *SectionDefault) FindAll() (v map[int]models.Section, err error) {
	return s.rp.FindAll()
}

func (s *SectionDefault) FindByID(id int) (models.Section, error) {
	return s.rp.FindByID(id)
}

func (s *SectionDefault) Add(ss models.Section) (models.Section, error) {

	return s.rp.Add(ss)
}
func (s *SectionDefault) Update(ss models.Section) (models.Section, error) {
	return s.rp.Update(ss)
}
func (s *SectionDefault) Delete(id int) (models.Section, error) {
	return s.rp.Delete(id)
}
