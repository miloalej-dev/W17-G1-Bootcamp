package memory

import "github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"

type EmployeeRepo struct {
	db map[int]models.Employee
}

func NewEmployeeRepo(db map[int]models.Employee) *EmployeeRepo {
	defaultDb := make(map[int]models.Employee)
	if db != nil {
		defaultDb = db
	}
	return &EmployeeRepo{db: defaultDb}
}

func (em *EmployeeRepo) FindAll() (e map[int]models.Employee, err error) {
	e = make(map[int]models.Employee)
	for k, emp := range em.db {
		e[k] = emp
	}
	return
}

func (em *EmployeeRepo) FindById(id int) (models.Employee, error) {
	//TODO implement me
	panic("implement me")
}

func (em *EmployeeRepo) Create(models.Employee) (models.Employee, error) {
	//TODO implement me
	panic("implement me")
}

func (em *EmployeeRepo) Update(models.Employee) (models.Employee, error) {
	//TODO implement me
	panic("implement me")
}

func (em *EmployeeRepo) PartialUpdate(id int, fields map[string]interface{}) (models.Employee, error) {
	//TODO implement me
	panic("implement me")
}

func (em *EmployeeRepo) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
