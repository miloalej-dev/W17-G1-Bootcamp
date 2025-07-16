package service

import "errors"

var (
	ErrProductIdConflict = errors.New("the product with the given id does not exists")
	ErrProductNotFound   = errors.New("product not found")
)
