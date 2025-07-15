package _default

import (
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
)

type SectionService struct {
	// rp is the repository that will be used by the service
	rp repository.Repository[int, models.Section]
}

func NewSectionService(rp repository.Repository[int, models.Section]) *SectionService {
	return &SectionService{rp: rp}
}

func (s *SectionService) RetrieveAll() (v []models.Section, err error) {
	return s.rp.FindAll()
}

func (s *SectionService) Retrieve(id int) (models.Section, error) {
	return s.rp.FindById(id)
}

func (s *SectionService) Register(ss models.Section) (models.Section, error) {

	return s.rp.Create(ss)
}
func (s *SectionService) Modify(ss models.Section) (models.Section, error) {
	return s.rp.Update(ss)
}
func (s *SectionService) PartialModify(id int, fields map[string]any) (models.Section, error) {
	return s.rp.PartialUpdate(id, fields)
}
func (s *SectionService) Remove(id int) error {
	return s.rp.Delete(id)
}
