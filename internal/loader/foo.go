package loader

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

// Loader is an interface that represents the loader
type Loader interface {
	Load() (v map[int]models.Foo, err error)
}
