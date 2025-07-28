package database

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type EmployeeRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo *EmployeeRepository
}

func (s *EmployeeRepositoryTestSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	if err != nil {
		s.T().Fatal(err)
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = gormDB
	s.repo = NewEmployeeRepository(s.db)
}

func (s *EmployeeRepositoryTestSuite) TestFindAll_DatabaseError() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `employees`")).
		WillReturnError(sql.ErrConnDone)
	// Act
	es, err := s.repo.FindAll()
	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Nil(es)
}

func (s *EmployeeRepositoryTestSuite) TestFindAll_Success() {
	// Arrange
	employeesExpected := []models.Employee{
		{
			Id:           1,
			CardNumberId: "EMP768",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseId:  1,
		},
		{
			Id:           2,
			CardNumberId: "EMP002",
			FirstName:    "Jane",
			LastName:     "Smith",
			WarehouseId:  2,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"})
	for _, emp := range employeesExpected {
		rows.AddRow(emp.Id, emp.CardNumberId, emp.FirstName, emp.LastName, emp.WarehouseId)
	}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `employees`")).
		WillReturnRows(rows)
	// Act
	employees, err := s.repo.FindAll()
	// Asserts
	s.NoError(err) // espera NO haya error
	s.Equal(employeesExpected, employees)

}

func (s *EmployeeRepositoryTestSuite) TestFindAll_EmptyList() {
	// Arrange
	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"})
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `employees`")).WillReturnRows(rows)
	// Act
	employees, err := s.repo.FindAll()
	// Asserts
	s.NoError(err)
	s.Empty(employees)
}

func (s *EmployeeRepositoryTestSuite) TestFindById_DatabaseError() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `employees` WHERE `employees`.`id` = ? ORDER BY `employees`.`id` LIMIT ?")).
		WithArgs(1, 1).
		WillReturnError(sql.ErrConnDone)

	// Act
	employee, err := s.repo.FindById(1)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.Employee{}, employee)
}

func (s *EmployeeRepositoryTestSuite) TestFindById_Success() {
	// Arrange
	expectedEmployee := models.Employee{
		Id:           1,
		CardNumberId: "EMP001",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseId:  1,
	}

	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"}).
		AddRow(expectedEmployee.Id, expectedEmployee.CardNumberId, expectedEmployee.FirstName, expectedEmployee.LastName, expectedEmployee.WarehouseId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `employees` WHERE `employees`.`id` = ? ORDER BY `employees`.`id` LIMIT ?")).
		WithArgs(1, 1).
		WillReturnRows(rows)

	// Act
	employee, err := s.repo.FindById(1)

	// Asserts
	s.NoError(err)
	s.NotNil(employee)
	s.Equal(expectedEmployee, employee)
}

func (s *EmployeeRepositoryTestSuite) TestFindById_NotFound() {
	// Arrange
	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"})

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `employees` WHERE `employees`.`id` = ? ORDER BY `employees`.`id` LIMIT ?")).
		WithArgs(999, 1).
		WillReturnRows(rows)

	// Act
	employee, err := s.repo.FindById(999)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal(models.Employee{}, employee)
}

func (s *EmployeeRepositoryTestSuite) TestCreate_Success() {
	// Arrange
	newEmployee := models.Employee{
		CardNumberId: "EMP003",
		FirstName:    "Alice",
		LastName:     "Johnson",
		WarehouseId:  1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `employees` (`card_number_id`,`first_name`,`last_name`,`warehouse_id`) VALUES (?,?,?,?)")).
		WithArgs(newEmployee.CardNumberId, newEmployee.FirstName, newEmployee.LastName, newEmployee.WarehouseId).
		WillReturnResult(sqlmock.NewResult(3, 1))
	s.mock.ExpectCommit()

	// Act
	createdEmployee, err := s.repo.Create(newEmployee)

	// Asserts
	s.NoError(err)
	s.Equal(newEmployee.CardNumberId, createdEmployee.CardNumberId)
	s.Equal(newEmployee.FirstName, createdEmployee.FirstName)
	s.Equal(newEmployee.LastName, createdEmployee.LastName)
	s.Equal(newEmployee.WarehouseId, createdEmployee.WarehouseId)
	s.Equal(3, createdEmployee.Id)
}

func (s *EmployeeRepositoryTestSuite) TestCreate_DatabaseError() {
	// Arrange
	newEmployee := models.Employee{
		CardNumberId: "EMP004",
		FirstName:    "Bob",
		LastName:     "Smith",
		WarehouseId:  1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `employees` (`card_number_id`,`first_name`,`last_name`,`warehouse_id`) VALUES (?,?,?,?)")).
		WithArgs(newEmployee.CardNumberId, newEmployee.FirstName, newEmployee.LastName, newEmployee.WarehouseId).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	createdEmployee, err := s.repo.Create(newEmployee)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.Employee{}, createdEmployee)
}

func (s *EmployeeRepositoryTestSuite) TestCreate_ForeignKeyViolation() {
	// Arrange
	newEmployee := models.Employee{
		CardNumberId: "EMP005",
		FirstName:    "Charlie",
		LastName:     "Brown",
		WarehouseId:  999, // Warehouse ID que no existe
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `employees` (`card_number_id`,`first_name`,`last_name`,`warehouse_id`) VALUES (?,?,?,?)")).
		WithArgs(newEmployee.CardNumberId, newEmployee.FirstName, newEmployee.LastName, newEmployee.WarehouseId).
		WillReturnError(gorm.ErrForeignKeyViolated)
	s.mock.ExpectRollback()

	// Act
	createdEmployee, err := s.repo.Create(newEmployee)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrForeignKeyViolation, err)
	s.Equal(models.Employee{}, createdEmployee)
}

func (s *EmployeeRepositoryTestSuite) TestUpdate_Success() {
	// Arrange
	ep := models.Employee{
		Id:           1,
		CardNumberId: "EMP001-UPDATED",
		FirstName:    "John Updated",
		LastName:     "Doe Updated",
		WarehouseId:  2,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `employees` SET `card_number_id`=?,`first_name`=?,`last_name`=?,`warehouse_id`=? WHERE `id` = ?")).
		WithArgs(ep.CardNumberId, ep.FirstName, ep.LastName, ep.WarehouseId, ep.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	updatedEmployee, err := s.repo.Update(ep)

	// Asserts
	s.NoError(err)
	s.Equal(ep.Id, updatedEmployee.Id)
	s.Equal(ep.CardNumberId, updatedEmployee.CardNumberId)
	s.Equal(ep.FirstName, updatedEmployee.FirstName)
	s.Equal(ep.LastName, updatedEmployee.LastName)
	s.Equal(ep.WarehouseId, updatedEmployee.WarehouseId)
}

func (s *EmployeeRepositoryTestSuite) TestUpdate_DatabaseError() {
	// Arrange
	existingEmployee := models.Employee{
		Id:           1,
		CardNumberId: "EMP001-UPDATED",
		FirstName:    "John Updated",
		LastName:     "Doe Updated",
		WarehouseId:  2,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `employees` SET `card_number_id`=?,`first_name`=?,`last_name`=?,`warehouse_id`=? WHERE `id` = ?")).
		WithArgs(existingEmployee.CardNumberId, existingEmployee.FirstName, existingEmployee.LastName, existingEmployee.WarehouseId, existingEmployee.Id).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	updatedEmployee, err := s.repo.Update(existingEmployee)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.Employee{}, updatedEmployee)
}

func (s *EmployeeRepositoryTestSuite) TestUpdate_ForeignKeyViolation() {
	// Arrange
	existingEmployee := models.Employee{
		Id:           1,
		CardNumberId: "EMP001-UPDATED",
		FirstName:    "John Updated",
		LastName:     "Doe Updated",
		WarehouseId:  999, // Warehouse ID que no existe
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `employees` SET `card_number_id`=?,`first_name`=?,`last_name`=?,`warehouse_id`=? WHERE `id` = ?")).
		WithArgs(existingEmployee.CardNumberId, existingEmployee.FirstName, existingEmployee.LastName, existingEmployee.WarehouseId, existingEmployee.Id).
		WillReturnError(gorm.ErrForeignKeyViolated)
	s.mock.ExpectRollback()

	// Act
	updatedEmployee, err := s.repo.Update(existingEmployee)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrForeignKeyViolation, err)
	s.Equal(models.Employee{}, updatedEmployee)
}

func (s *EmployeeRepositoryTestSuite) TestPartialUpdate_Success() {
	// Arrange
	employeeId := 1
	fields := map[string]interface{}{
		"first_name": "John Updated",
		"last_name":  "Doe Updated",
	}

	existingEmployee := models.Employee{
		Id:           1,
		CardNumberId: "EMP001",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseId:  1,
	}

	// Primero busca el empleado existente
	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"}).
		AddRow(existingEmployee.Id, existingEmployee.CardNumberId, existingEmployee.FirstName, existingEmployee.LastName, existingEmployee.WarehouseId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `employees` WHERE `employees`.`id` = ? ORDER BY `employees`.`id` LIMIT ?")).
		WithArgs(employeeId, 1).
		WillReturnRows(rows)

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `employees` SET `first_name`=?,`last_name`=? WHERE `id` = ?")).
		WithArgs(fields["first_name"], fields["last_name"], employeeId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	updatedEmployee, err := s.repo.PartialUpdate(employeeId, fields)

	// Asserts
	s.NoError(err)
	s.Equal(employeeId, updatedEmployee.Id)
	s.Equal(existingEmployee.CardNumberId, updatedEmployee.CardNumberId)
	s.Equal(existingEmployee.WarehouseId, updatedEmployee.WarehouseId)
}

func (s *EmployeeRepositoryTestSuite) TestPartialUpdate_NotFound() {
	// Arrange
	employeeId := 999
	fields := map[string]interface{}{
		"first_name": "John Updated",
	}

	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"})
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `employees` WHERE `employees`.`id` = ? ORDER BY `employees`.`id` LIMIT ?")).
		WithArgs(employeeId, 1).
		WillReturnRows(rows)

	// Act
	updatedEmployee, err := s.repo.PartialUpdate(employeeId, fields)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal(models.Employee{}, updatedEmployee)
}

// Test PartialUpdate - Error genérico durante la búsqueda inicial (NO record not found)
func (s *EmployeeRepositoryTestSuite) TestPartialUpdate_FindDatabaseError() {
	// Arrange
	employeeId := 1
	fields := map[string]interface{}{
		"first_name": "John Updated",
	}

	// Simula un error de conexión durante la búsqueda inicial
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `employees` WHERE `employees`.`id` = ? ORDER BY `employees`.`id` LIMIT ?")).
		WithArgs(employeeId, 1).
		WillReturnError(sql.ErrConnDone) // Error genérico de BD, NO gorm.ErrRecordNotFound

	// Act
	updatedEmployee, err := s.repo.PartialUpdate(employeeId, fields)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err) // Debe retornar el error original
	s.Equal(models.Employee{}, updatedEmployee)
}

// Test PartialUpdate - Error de Foreign Key durante Updates
func (s *EmployeeRepositoryTestSuite) TestPartialUpdate_ForeignKeyViolation() {
	// Arrange
	employeeId := 1
	fields := map[string]interface{}{
		"warehouse_id": 999, // Warehouse ID que no existe
	}

	existingEmployee := models.Employee{
		Id:           1,
		CardNumberId: "EMP001",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseId:  1,
	}

	// Primero busca el empleado existente exitosamente
	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"}).
		AddRow(existingEmployee.Id, existingEmployee.CardNumberId, existingEmployee.FirstName, existingEmployee.LastName, existingEmployee.WarehouseId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `employees` WHERE `employees`.`id` = ? ORDER BY `employees`.`id` LIMIT ?")).
		WithArgs(employeeId, 1).
		WillReturnRows(rows)

	// Falla en la actualización por foreign key
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `employees` SET `warehouse_id`=? WHERE `id` = ?")).
		WithArgs(fields["warehouse_id"], employeeId).
		WillReturnError(gorm.ErrForeignKeyViolated)
	s.mock.ExpectRollback()

	// Act
	updatedEmployee, err := s.repo.PartialUpdate(employeeId, fields)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrForeignKeyViolation, err)
	s.Equal(models.Employee{}, updatedEmployee)
}

func (s *EmployeeRepositoryTestSuite) TestDelete_Success() {
	// Arrange
	employeeID := 1
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `employees` WHERE `employees`.`id` = ?")).
		WithArgs(employeeID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	err := s.repo.Delete(employeeID)

	// Asserts
	s.NoError(err)
}

func (s *EmployeeRepositoryTestSuite) TestDelete_NotFound() {
	// Arrange
	employeeID := 999
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `employees` WHERE `employees`.`id` = ?")).
		WithArgs(employeeID).
		WillReturnResult(sqlmock.NewResult(1, 0)) // 0 rows affected
	s.mock.ExpectCommit()

	// Act
	err := s.repo.Delete(employeeID)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
}

func (s *EmployeeRepositoryTestSuite) TestDelete_DatabaseError() {
	// Arrange
	employeeID := 1
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `employees` WHERE `employees`.`id` = ?")).
		WithArgs(employeeID).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	err := s.repo.Delete(employeeID)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
}
func (s *EmployeeRepositoryTestSuite) TestPartialUpdate_DatabaseError() {
	// Arrange
	employeeId := 1
	fields := map[string]interface{}{
		"first_name": "John Updated",
	}

	existingEmployee := models.Employee{
		Id:           1,
		CardNumberId: "EMP001",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseId:  1,
	}

	// Primero busca el empleado existente exitosamente
	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"}).
		AddRow(existingEmployee.Id, existingEmployee.CardNumberId, existingEmployee.FirstName, existingEmployee.LastName, existingEmployee.WarehouseId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `employees` WHERE `employees`.`id` = ? ORDER BY `employees`.`id` LIMIT ?")).
		WithArgs(employeeId, 1).
		WillReturnRows(rows)

	// Falla en la actualización con un error genérico (NO foreign key)
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `employees` SET `first_name`=? WHERE `id` = ?")).
		WithArgs(fields["first_name"], employeeId).
		WillReturnError(sql.ErrTxDone) // Error genérico diferente a ForeignKey
	s.mock.ExpectRollback()

	// Act
	updatedEmployee, err := s.repo.PartialUpdate(employeeId, fields)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrTxDone, err)
	s.Equal(models.Employee{}, updatedEmployee)
}

func (s *EmployeeRepositoryTestSuite) TestInboundOrdersReport_Success() {
	// Arrange
	expectedReports := []models.EmployeeInboundOrdersReport{
		{
			Employee: models.Employee{
				Id:           1,
				CardNumberId: "EMP001",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseId:  1,
			},
			InboundOrdersCount: 5,
		},
		{
			Employee: models.Employee{
				Id:           2,
				CardNumberId: "EMP002",
				FirstName:    "Jane",
				LastName:     "Smith",
				WarehouseId:  2,
			},
			InboundOrdersCount: 3,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count"})
	for _, report := range expectedReports {
		rows.AddRow(report.Id, report.CardNumberId, report.FirstName, report.LastName, report.WarehouseId, report.InboundOrdersCount)
	}

	expectedQuery := `SELECT
            e.id,
            e.card_number_id,
            e.first_name,
            e.last_name,
			e.warehouse_id,
            COUNT(io.id) AS inbound_orders_count
         FROM employees e LEFT JOIN inbound_orders io ON e.id = io.employee_id LEFT JOIN warehouses w ON e.warehouse_id = w.id GROUP BY e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id ORDER BY e.id`

	s.mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
		WillReturnRows(rows)

	// Act
	reports, err := s.repo.InboundOrdersReport()

	// Asserts
	s.NoError(err)
	s.Equal(expectedReports, reports)
}

// Test InboundOrdersReport - Database Error
func (s *EmployeeRepositoryTestSuite) TestInboundOrdersReport_DatabaseError() {
	// Arrange
	expectedQuery := `SELECT
            e.id,
            e.card_number_id,
            e.first_name,
            e.last_name,
			e.warehouse_id,
            COUNT(io.id) AS inbound_orders_count
         FROM employees e LEFT JOIN inbound_orders io ON e.id = io.employee_id LEFT JOIN warehouses w ON e.warehouse_id = w.id GROUP BY e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id ORDER BY e.id`

	s.mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
		WillReturnError(sql.ErrConnDone)

	// Act
	reports, err := s.repo.InboundOrdersReport()

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Nil(reports)
}

// Test InboundOrdersReport - Empty Result
func (s *EmployeeRepositoryTestSuite) TestInboundOrdersReport_EmptyResult() {
	// Arrange
	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count"})

	expectedQuery := `SELECT
            e.id,
            e.card_number_id,
            e.first_name,
            e.last_name,
			e.warehouse_id,
            COUNT(io.id) AS inbound_orders_count
         FROM employees e LEFT JOIN inbound_orders io ON e.id = io.employee_id LEFT JOIN warehouses w ON e.warehouse_id = w.id GROUP BY e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id ORDER BY e.id`

	s.mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
		WillReturnRows(rows)

	// Act
	reports, err := s.repo.InboundOrdersReport()

	// Asserts
	s.NoError(err)
	s.Empty(reports)
}

// Test InboundOrdersReportById - Success
func (s *EmployeeRepositoryTestSuite) TestInboundOrdersReportById_Success() {
	// Arrange
	employeeId := 1
	expectedReport := models.EmployeeInboundOrdersReport{
		Employee: models.Employee{
			Id:           1,
			CardNumberId: "EMP001",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseId:  1,
		},
		InboundOrdersCount: 5,
	}

	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count"}).
		AddRow(expectedReport.Id, expectedReport.CardNumberId, expectedReport.FirstName, expectedReport.LastName, expectedReport.WarehouseId, expectedReport.InboundOrdersCount)

	expectedQuery := `SELECT
            e.id,
            e.card_number_id,
            e.first_name,
            e.last_name,
			e.warehouse_id,
            COUNT(io.id) AS inbound_orders_count
         FROM employees e LEFT JOIN inbound_orders io ON e.id = io.employee_id LEFT JOIN warehouses w ON e.warehouse_id = w.id WHERE e.id = ? GROUP BY e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id`

	s.mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
		WithArgs(employeeId).
		WillReturnRows(rows)

	// Act
	report, err := s.repo.InboundOrdersReportById(employeeId)

	// Asserts
	s.NoError(err)
	s.Equal(expectedReport, report)
}

// Test InboundOrdersReportById - Database Error
func (s *EmployeeRepositoryTestSuite) TestInboundOrdersReportById_DatabaseError() {
	// Arrange
	employeeId := 1
	expectedQuery := `SELECT
            e.id,
            e.card_number_id,
            e.first_name,
            e.last_name,
			e.warehouse_id,
            COUNT(io.id) AS inbound_orders_count
         FROM employees e LEFT JOIN inbound_orders io ON e.id = io.employee_id LEFT JOIN warehouses w ON e.warehouse_id = w.id WHERE e.id = ? GROUP BY e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id`

	s.mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
		WithArgs(employeeId).
		WillReturnError(sql.ErrConnDone)

	// Act
	report, err := s.repo.InboundOrdersReportById(employeeId)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.EmployeeInboundOrdersReport{}, report)
}

// Test InboundOrdersReportById - Employee Not Found
func (s *EmployeeRepositoryTestSuite) TestInboundOrdersReportById_NotFound() {
	// Arrange
	employeeId := 999
	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count"})

	expectedQuery := `SELECT
            e.id,
            e.card_number_id,
            e.first_name,
            e.last_name,
			e.warehouse_id,
            COUNT(io.id) AS inbound_orders_count
         FROM employees e LEFT JOIN inbound_orders io ON e.id = io.employee_id LEFT JOIN warehouses w ON e.warehouse_id = w.id WHERE e.id = ? GROUP BY e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id`

	s.mock.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
		WithArgs(employeeId).
		WillReturnRows(rows)

	// Act
	report, err := s.repo.InboundOrdersReportById(employeeId)

	// Asserts
	s.Error(err) // The method doesn't return error for empty results, just empty struct
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal(models.EmployeeInboundOrdersReport{}, report)
}

func TestEmployeeRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(EmployeeRepositoryTestSuite))
}
