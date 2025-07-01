package sectionLoader

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

// Loader is an interface that represents the loader
type SectionLoader interface {
	Load() (v map[int]models.Section, err error)
}
