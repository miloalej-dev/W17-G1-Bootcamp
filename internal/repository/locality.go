package repository

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type LocalityRepository interface {
	Repository[int, models.Locality]
}
