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
	"time"
)

type InboundOrderRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo *InboundOrderRepository
}

func (s *InboundOrderRepositoryTestSuite) SetupSuite() {
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
	s.repo = NewInboundOrderRepository(gormDB)
}

func (s *InboundOrderRepositoryTestSuite) TestFindAll_Success() {
	// Arrange
	orderDate1 := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	orderDate2 := time.Date(2024, 2, 20, 14, 30, 0, 0, time.UTC)

	expectedOrders := []models.InboundOrder{
		{
			Id:             1,
			OrderNumber:    "ORD-001",
			OrderDate:      orderDate1,
			EmployeeId:     1,
			ProductBatchId: 1,
			WarehouseId:    1,
		},
		{
			Id:             2,
			OrderNumber:    "ORD-002",
			OrderDate:      orderDate2,
			EmployeeId:     2,
			ProductBatchId: 2,
			WarehouseId:    2,
		},
	}

	rows := s.mock.NewRows([]string{"id", "order_number", "order_date", "employee_id", "product_batch_id", "warehouse_id"}).
		AddRow(expectedOrders[0].Id, expectedOrders[0].OrderNumber, expectedOrders[0].OrderDate,
			expectedOrders[0].EmployeeId, expectedOrders[0].ProductBatchId, expectedOrders[0].WarehouseId).
		AddRow(expectedOrders[1].Id, expectedOrders[1].OrderNumber, expectedOrders[1].OrderDate,
			expectedOrders[1].EmployeeId, expectedOrders[1].ProductBatchId, expectedOrders[1].WarehouseId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `inbound_orders`")).WillReturnRows(rows)

	// Act
	orders, err := s.repo.FindAll()

	// Assert
	s.NoError(err)
	s.Len(orders, 2)
	s.Equal(expectedOrders[0].OrderNumber, orders[0].OrderNumber)
	s.Equal(expectedOrders[1].OrderNumber, orders[1].OrderNumber)
}

func (s *InboundOrderRepositoryTestSuite) TestFindAll_DatabaseError() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `inbound_orders`")).WillReturnError(sql.ErrConnDone)

	// Act
	orders, err := s.repo.FindAll()

	// Assert
	s.Error(err)
	s.Nil(orders)
	s.Equal(sql.ErrConnDone, err)
}

func (s *InboundOrderRepositoryTestSuite) TestFindById_Success() {
	// Arrange
	orderDate := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	expectedOrder := models.InboundOrder{
		Id:             1,
		OrderNumber:    "ORD-001",
		OrderDate:      orderDate,
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}

	rows := s.mock.NewRows([]string{"id", "order_number", "order_date", "employee_id", "product_batch_id", "warehouse_id"}).
		AddRow(expectedOrder.Id, expectedOrder.OrderNumber, expectedOrder.OrderDate,
			expectedOrder.EmployeeId, expectedOrder.ProductBatchId, expectedOrder.WarehouseId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `inbound_orders` WHERE `inbound_orders`.`id` = ? ORDER BY `inbound_orders`.`id` LIMIT ?")).
		WithArgs(1, 1).WillReturnRows(rows)

	// Act
	order, err := s.repo.FindById(1)

	// Assert
	s.NoError(err)
	s.Equal(expectedOrder.Id, order.Id)
	s.Equal(expectedOrder.OrderNumber, order.OrderNumber)
}

func (s *InboundOrderRepositoryTestSuite) TestFindById_NotFound() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `inbound_orders` WHERE `inbound_orders`.`id` = ? ORDER BY `inbound_orders`.`id` LIMIT ?")).
		WithArgs(999, 1).WillReturnError(gorm.ErrRecordNotFound)

	// Act
	order, err := s.repo.FindById(999)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal(models.InboundOrder{}, order)
}

func (s *InboundOrderRepositoryTestSuite) TestFindById_DatabaseError() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `inbound_orders` WHERE `inbound_orders`.`id` = ? ORDER BY `inbound_orders`.`id` LIMIT ?")).
		WithArgs(1, 1).WillReturnError(sql.ErrConnDone)

	// Act
	order, err := s.repo.FindById(1)

	// Assert
	s.Error(err)
	s.Equal(models.InboundOrder{}, order)
	s.Equal(sql.ErrConnDone, err)
}

func (s *InboundOrderRepositoryTestSuite) TestCreate_Success() {
	// Arrange
	orderDate := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	newOrder := models.InboundOrder{
		OrderNumber:    "ORD-003",
		OrderDate:      orderDate,
		EmployeeId:     3,
		ProductBatchId: 3,
		WarehouseId:    3,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `inbound_orders` (`order_number`,`order_date`,`employee_id`,`product_batch_id`,`warehouse_id`) VALUES (?,?,?,?,?)")).
		WithArgs(newOrder.OrderNumber, newOrder.OrderDate, newOrder.EmployeeId, newOrder.ProductBatchId, newOrder.WarehouseId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	createdOrder, err := s.repo.Create(newOrder)

	// Assert
	s.NoError(err)
	s.Equal(newOrder.OrderNumber, createdOrder.OrderNumber)
	s.Equal(1, createdOrder.Id)
}

func (s *InboundOrderRepositoryTestSuite) TestCreate_ForeignKeyViolated() {
	// Arrange
	orderDate := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	newOrder := models.InboundOrder{
		OrderNumber:    "ORD-003",
		OrderDate:      orderDate,
		EmployeeId:     999, // Invalid employee ID
		ProductBatchId: 3,
		WarehouseId:    3,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `inbound_orders` (`order_number`,`order_date`,`employee_id`,`product_batch_id`,`warehouse_id`) VALUES (?,?,?,?,?)")).
		WithArgs(newOrder.OrderNumber, newOrder.OrderDate, newOrder.EmployeeId, newOrder.ProductBatchId, newOrder.WarehouseId).
		WillReturnError(gorm.ErrForeignKeyViolated)
	s.mock.ExpectRollback()

	// Act
	createdOrder, err := s.repo.Create(newOrder)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrForeignKeyViolation, err)
	s.Equal(models.InboundOrder{}, createdOrder)
}

func (s *InboundOrderRepositoryTestSuite) TestCreate_DuplicatedKey() {
	// Arrange
	orderDate := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	newOrder := models.InboundOrder{
		OrderNumber:    "ORD-001", // Duplicate order number
		OrderDate:      orderDate,
		EmployeeId:     3,
		ProductBatchId: 3,
		WarehouseId:    3,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `inbound_orders` (`order_number`,`order_date`,`employee_id`,`product_batch_id`,`warehouse_id`) VALUES (?,?,?,?,?)")).
		WithArgs(newOrder.OrderNumber, newOrder.OrderDate, newOrder.EmployeeId, newOrder.ProductBatchId, newOrder.WarehouseId).
		WillReturnError(gorm.ErrDuplicatedKey)
	s.mock.ExpectRollback()

	// Act
	createdOrder, err := s.repo.Create(newOrder)

	// Assert
	s.Error(err)
	s.Equal("inbound order number already exists, must be unique", err.Error())
	s.Equal(models.InboundOrder{}, createdOrder)
}

func (s *InboundOrderRepositoryTestSuite) TestCreate_DatabaseError() {
	// Arrange
	orderDate := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	newOrder := models.InboundOrder{
		OrderNumber:    "ORD-003",
		OrderDate:      orderDate,
		EmployeeId:     3,
		ProductBatchId: 3,
		WarehouseId:    3,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `inbound_orders` (`order_number`,`order_date`,`employee_id`,`product_batch_id`,`warehouse_id`) VALUES (?,?,?,?,?)")).
		WithArgs(newOrder.OrderNumber, newOrder.OrderDate, newOrder.EmployeeId, newOrder.ProductBatchId, newOrder.WarehouseId).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	createdOrder, err := s.repo.Create(newOrder)

	// Assert
	s.Error(err)
	s.Equal(models.InboundOrder{}, createdOrder)
	s.Equal(sql.ErrConnDone, err)
}

func (s *InboundOrderRepositoryTestSuite) TestUpdate_Success() {
	// Arrange
	orderDate := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	existingOrder := models.InboundOrder{
		Id:             1,
		OrderNumber:    "ORD-001-UPDATED",
		OrderDate:      orderDate,
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `inbound_orders` SET `order_number`=?,`order_date`=?,`employee_id`=?,`product_batch_id`=?,`warehouse_id`=? WHERE `id` = ?")).
		WithArgs(existingOrder.OrderNumber, existingOrder.OrderDate, existingOrder.EmployeeId, existingOrder.ProductBatchId, existingOrder.WarehouseId, existingOrder.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	updatedOrder, err := s.repo.Update(existingOrder)

	// Assert
	s.NoError(err)
	s.Equal(existingOrder.OrderNumber, updatedOrder.OrderNumber)
	s.Equal(existingOrder.Id, updatedOrder.Id)
}

func (s *InboundOrderRepositoryTestSuite) TestUpdate_ForeignKeyViolated() {
	// Arrange
	orderDate := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	existingOrder := models.InboundOrder{
		Id:             1,
		OrderNumber:    "ORD-001-UPDATED",
		OrderDate:      orderDate,
		EmployeeId:     999, // Invalid employee ID
		ProductBatchId: 1,
		WarehouseId:    1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `inbound_orders` SET `order_number`=?,`order_date`=?,`employee_id`=?,`product_batch_id`=?,`warehouse_id`=? WHERE `id` = ?")).
		WithArgs(existingOrder.OrderNumber, existingOrder.OrderDate, existingOrder.EmployeeId, existingOrder.ProductBatchId, existingOrder.WarehouseId, existingOrder.Id).
		WillReturnError(gorm.ErrForeignKeyViolated)
	s.mock.ExpectRollback()

	// Act
	updatedOrder, err := s.repo.Update(existingOrder)

	// Assert
	s.Error(err)
	s.Equal(models.InboundOrder{}, updatedOrder)
	s.Equal(repository.ErrForeignKeyViolation, err)
}

func (s *InboundOrderRepositoryTestSuite) TestUpdate_DuplicatedKey() {
	// Arrange
	orderDate := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	existingOrder := models.InboundOrder{
		Id:             1,
		OrderNumber:    "ORD-002", // Duplicate order number
		OrderDate:      orderDate,
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `inbound_orders` SET `order_number`=?,`order_date`=?,`employee_id`=?,`product_batch_id`=?,`warehouse_id`=? WHERE `id` = ?")).
		WithArgs(existingOrder.OrderNumber, existingOrder.OrderDate, existingOrder.EmployeeId, existingOrder.ProductBatchId, existingOrder.WarehouseId, existingOrder.Id).
		WillReturnError(gorm.ErrDuplicatedKey)
	s.mock.ExpectRollback()

	// Act
	updatedOrder, err := s.repo.Update(existingOrder)

	// Assert
	s.Error(err)
	s.Equal(models.InboundOrder{}, updatedOrder)
	s.Equal("inbound order number already exists, must be unique", err.Error())
}

func (s *InboundOrderRepositoryTestSuite) TestUpdate_DatabaseError() {
	// Arrange
	orderDate := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	existingOrder := models.InboundOrder{
		Id:             1,
		OrderNumber:    "ORD-001-UPDATED",
		OrderDate:      orderDate,
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `inbound_orders` SET `order_number`=?,`order_date`=?,`employee_id`=?,`product_batch_id`=?,`warehouse_id`=? WHERE `id` = ?")).
		WithArgs(existingOrder.OrderNumber, existingOrder.OrderDate, existingOrder.EmployeeId, existingOrder.ProductBatchId, existingOrder.WarehouseId, existingOrder.Id).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	updatedOrder, err := s.repo.Update(existingOrder)

	// Assert
	s.Error(err)
	s.Equal(models.InboundOrder{}, updatedOrder)
	s.Equal(sql.ErrConnDone, err)
}

func (s *InboundOrderRepositoryTestSuite) TestPartialUpdate_Success() {
	// Arrange
	orderID := 1
	fields := map[string]interface{}{
		"order_number": "ORD-001-PARTIAL",
		"employee_id":  2,
	}

	orderDate := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	expectedOrder := models.InboundOrder{
		Id:             orderID,
		OrderNumber:    "ORD-001-OLD",
		OrderDate:      orderDate,
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}

	// First query to find the order
	rows := s.mock.NewRows([]string{"id", "order_number", "order_date", "employee_id", "product_batch_id", "warehouse_id"}).
		AddRow(expectedOrder.Id, expectedOrder.OrderNumber, expectedOrder.OrderDate, expectedOrder.EmployeeId, expectedOrder.ProductBatchId, expectedOrder.WarehouseId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `inbound_orders` WHERE `inbound_orders`.`id` = ? ORDER BY `inbound_orders`.`id` LIMIT ?")).
		WithArgs(orderID, 1).WillReturnRows(rows)

	// Update query
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `inbound_orders` SET `employee_id`=?,`order_number`=? WHERE `id` = ?")).
		WithArgs(fields["employee_id"], fields["order_number"], orderID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	updatedOrder, err := s.repo.PartialUpdate(orderID, fields)

	// Assert
	s.NoError(err)
	s.Equal(orderID, updatedOrder.Id)
}

func (s *InboundOrderRepositoryTestSuite) TestPartialUpdate_NotFound() {
	// Arrange
	orderID := 999
	fields := map[string]interface{}{
		"order_number": "ORD-999",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `inbound_orders` WHERE `inbound_orders`.`id` = ? ORDER BY `inbound_orders`.`id` LIMIT ?")).
		WithArgs(orderID, 1).WillReturnError(gorm.ErrRecordNotFound)

	// Act
	updatedOrder, err := s.repo.PartialUpdate(orderID, fields)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal(models.InboundOrder{}, updatedOrder)
}

func (s *InboundOrderRepositoryTestSuite) TestPartialUpdate_ForeignKeyViolated() {
	// Arrange
	orderID := 1
	fields := map[string]interface{}{
		"employee_id": 999, // Invalid employee ID
	}

	orderDate := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	expectedOrder := models.InboundOrder{
		Id:             orderID,
		OrderNumber:    "ORD-001",
		OrderDate:      orderDate,
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}

	// First query to find the order
	rows := s.mock.NewRows([]string{"id", "order_number", "order_date", "employee_id", "product_batch_id", "warehouse_id"}).
		AddRow(expectedOrder.Id, expectedOrder.OrderNumber, expectedOrder.OrderDate, expectedOrder.EmployeeId, expectedOrder.ProductBatchId, expectedOrder.WarehouseId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `inbound_orders` WHERE `inbound_orders`.`id` = ? ORDER BY `inbound_orders`.`id` LIMIT ?")).
		WithArgs(orderID, 1).WillReturnRows(rows)

	// Update query with foreign key violation
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `inbound_orders` SET `employee_id`=? WHERE `id` = ?")).
		WithArgs(fields["employee_id"], orderID).
		WillReturnError(gorm.ErrForeignKeyViolated)
	s.mock.ExpectRollback()

	// Act
	updatedOrder, err := s.repo.PartialUpdate(orderID, fields)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrForeignKeyViolation, err)
	s.Equal(models.InboundOrder{}, updatedOrder)
}

func (s *InboundOrderRepositoryTestSuite) TestPartialUpdate_DuplicatedKey() {
	// Arrange
	orderID := 1
	fields := map[string]interface{}{
		"order_number": "ORD-002", // Duplicate order number
	}

	orderDate := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	expectedOrder := models.InboundOrder{
		Id:             orderID,
		OrderNumber:    "ORD-001",
		OrderDate:      orderDate,
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}

	// First query to find the order
	rows := s.mock.NewRows([]string{"id", "order_number", "order_date", "employee_id", "product_batch_id", "warehouse_id"}).
		AddRow(expectedOrder.Id, expectedOrder.OrderNumber, expectedOrder.OrderDate, expectedOrder.EmployeeId, expectedOrder.ProductBatchId, expectedOrder.WarehouseId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `inbound_orders` WHERE `inbound_orders`.`id` = ? ORDER BY `inbound_orders`.`id` LIMIT ?")).
		WithArgs(orderID, 1).WillReturnRows(rows)

	// Update query with duplicate key violation
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `inbound_orders` SET `order_number`=? WHERE `id` = ?")).
		WithArgs(fields["order_number"], orderID).
		WillReturnError(gorm.ErrDuplicatedKey)
	s.mock.ExpectRollback()

	// Act
	updatedOrder, err := s.repo.PartialUpdate(orderID, fields)

	// Assert
	s.Error(err)
	s.Equal("inbound order number already exists, must be unique", err.Error())
	s.Equal(models.InboundOrder{}, updatedOrder)
}

func (s *InboundOrderRepositoryTestSuite) TestPartialUpdate_DatabaseError() {
	// Arrange
	orderID := 1
	fields := map[string]interface{}{
		"order_number": "ORD-001-UPDATED",
	}

	orderDate := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	expectedOrder := models.InboundOrder{
		Id:             orderID,
		OrderNumber:    "ORD-001",
		OrderDate:      orderDate,
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}

	// First query to find the order
	rows := s.mock.NewRows([]string{"id", "order_number", "order_date", "employee_id", "product_batch_id", "warehouse_id"}).
		AddRow(expectedOrder.Id, expectedOrder.OrderNumber, expectedOrder.OrderDate, expectedOrder.EmployeeId, expectedOrder.ProductBatchId, expectedOrder.WarehouseId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `inbound_orders` WHERE `inbound_orders`.`id` = ? ORDER BY `inbound_orders`.`id` LIMIT ?")).
		WithArgs(orderID, 1).WillReturnRows(rows)

	// Update query with database error
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `inbound_orders` SET `order_number`=? WHERE `id` = ?")).
		WithArgs(fields["order_number"], orderID).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	updatedOrder, err := s.repo.PartialUpdate(orderID, fields)

	// Assert
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.InboundOrder{}, updatedOrder)
}

func (s *InboundOrderRepositoryTestSuite) TestDelete_Success() {
	// Arrange
	orderID := 1
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `inbound_orders` WHERE `inbound_orders`.`id` = ?")).
		WithArgs(orderID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	err := s.repo.Delete(orderID)

	// Assert
	s.NoError(err)
}

func (s *InboundOrderRepositoryTestSuite) TestDelete_NotFound() {
	// Arrange
	orderID := 999

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `inbound_orders` WHERE `inbound_orders`.`id` = ?")).
		WithArgs(orderID).
		WillReturnResult(sqlmock.NewResult(1, 0)) // 0 rows affected
	s.mock.ExpectCommit()

	// Act
	err := s.repo.Delete(orderID)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
}

// Run the test suite
func TestInboundOrderRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(InboundOrderRepositoryTestSuite))
}
