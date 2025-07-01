package sectionRepository

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type SectionRepository interface {
	FindAll() (v map[int]models.Section, err error)
}
