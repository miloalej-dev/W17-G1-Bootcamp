package loaderProduct

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type Loader interface {
	Load() (v map[int]models.Product, err error)
}
