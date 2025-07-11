package _default

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type SectionDefault struct {
	// rp is the repository that will be used by the service
	rp repository.SectionRepository
}

func NewSectionDefault(rp repository.SectionRepository) *SectionDefault {
	return &SectionDefault{rp: rp}
}

func (s *SectionDefault) RetrieveAll() (v []models.Section, err error) {
	return s.rp.FindAll()
}

func (s *SectionDefault) Retrieve(id int) (models.Section, error) {
	return s.rp.FindById(id)
}

func (s *SectionDefault) Register(ss models.Section) (models.Section, error) {

	return s.rp.Create(ss)
}
func (s *SectionDefault) Modify(ss models.Section) (models.Section, error) {
	return s.rp.Update(ss)
}
func (s *SectionDefault) PartialModify(id int, fields map[string]any) (models.Section, error) {
	return s.rp.PartialUpdate(id, fields)
}
func (s *SectionDefault) Remove(id int) error {
	return s.rp.Delete(id)
}
