package database

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type OrderDetailTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo *OrderDetailRepository
}

func (s *OrderDetailTestSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	s.Require().NoError(err)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		TranslateError: true,
	})
	s.Require().NoError(err)

	s.db = gormDB
	s.repo = NewOrderDetailRepository(gormDB)
}

func (s *OrderDetailTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	s.Require().NoError(err)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		TranslateError: true,
	})
	s.Require().NoError(err)

	s.db = gormDB
	s.mock = mock
	s.repo = NewOrderDetailRepository(gormDB)
}

func (s *OrderDetailTestSuite) TestFindAll_Success() {
	orderDetails := []models.OrderDetail{
		{
			Id:               1,
			Quantity:         10,
			CleanLinesStatus: "OK",
			Temperature:      4.5,
			ProductRecordID:  100,
			PurchaseOrderID:  200,
		},
		{
			Id:               2,
			Quantity:         20,
			CleanLinesStatus: "Pending",
			Temperature:      3.8,
			ProductRecordID:  101,
			PurchaseOrderID:  201,
		},
	}

	rows := s.mock.NewRows([]string{
		"id", "quantity", "clean_lines_status", "temperature", "product_record_id", "purchase_order_id",
	}).
		AddRow(
			orderDetails[0].Id,
			orderDetails[0].Quantity,
			orderDetails[0].CleanLinesStatus,
			orderDetails[0].Temperature,
			orderDetails[0].ProductRecordID,
			orderDetails[0].PurchaseOrderID,
		).
		AddRow(
			orderDetails[1].Id,
			orderDetails[1].Quantity,
			orderDetails[1].CleanLinesStatus,
			orderDetails[1].Temperature,
			orderDetails[1].ProductRecordID,
			orderDetails[1].PurchaseOrderID,
		)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `order_details`")).WillReturnRows(rows)

	// Act
	od, err := s.repo.FindAll()

	// Assert
	s.NoError(err)
	s.Len(od, 2)
	s.Equal(orderDetails[0].CleanLinesStatus, od[0].CleanLinesStatus)
	s.Equal(orderDetails[1].CleanLinesStatus, od[1].CleanLinesStatus)
}

func (s *OrderDetailTestSuite) TestFindAll_DatabaseError() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `order_details`")).WillReturnError(sql.ErrConnDone)
	// Act
	orderDetail, err := s.repo.FindAll()

	// Assert
	s.Error(err)
	s.Nil(orderDetail)
	s.Equal(sql.ErrConnDone, err)

}

func (s *OrderDetailTestSuite) TestFindById_Success() {
	orderDetail := models.OrderDetail{
		Id:               1,
		Quantity:         10,
		CleanLinesStatus: "OK",
		Temperature:      4.5,
		ProductRecordID:  100,
		PurchaseOrderID:  200,
	}

	rows := s.mock.NewRows([]string{
		"id", "quantity", "clean_lines_status", "temperature", "product_record_id",
		"purchase_order_id",
	}).AddRow(
		orderDetail.Id,
		orderDetail.Quantity,
		orderDetail.CleanLinesStatus,
		orderDetail.Temperature,
		orderDetail.ProductRecordID,
		orderDetail.PurchaseOrderID)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `order_details` WHERE `order_details`.`id` = ? ORDER BY `order_details`.`id` LIMIT ?")).
		WithArgs(1, 1).WillReturnRows(rows)

	od, err := s.repo.FindById(orderDetail.Id)
	s.NoError(err)
	s.Equal(orderDetail, od)
	s.Equal(orderDetail.Quantity, od.Quantity)
}

func (s *OrderDetailTestSuite) TestFindById_NotFound() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `order_details` WHERE `order_details`.`id` = ? ORDER BY `order_details`.`id` LIMIT ?",
	)).WithArgs(999, 1).WillReturnError(repository.ErrEntityNotFound)
	// Act
	od, err := s.repo.FindById(999)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal(models.OrderDetail{}, od)
}

func (s *OrderDetailTestSuite) TestUpdate_Success() {
	od := models.OrderDetail{
		Id:               1,
		Quantity:         20,
		CleanLinesStatus: "Updated",
		Temperature:      5.5,
		ProductRecordID:  123,
		PurchaseOrderID:  456,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `order_details` SET `quantity`=?,`clean_lines_status`=?,`temperature`=?,`product_record_id`=?,`purchase_order_id`=? WHERE `id` = ?",
	)).WithArgs(
		od.Quantity,
		od.CleanLinesStatus,
		od.Temperature,
		od.ProductRecordID,
		od.PurchaseOrderID,
		od.Id,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	result, err := s.repo.Update(od)

	s.NoError(err)
	s.Equal(od, result)
}

func (s *OrderDetailTestSuite) TestUpdate_Error() {
	od := models.OrderDetail{
		Id:               1,
		Quantity:         10,
		CleanLinesStatus: "ErrorCase",
		Temperature:      4.0,
		ProductRecordID:  99,
		PurchaseOrderID:  88,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `order_details` SET `quantity`=?,`clean_lines_status`=?,`temperature`=?,`product_record_id`=?,`purchase_order_id`=? WHERE `id` = ?",
	)).WithArgs(
		od.Quantity,
		od.CleanLinesStatus,
		od.Temperature,
		od.ProductRecordID,
		od.PurchaseOrderID,
		od.Id,
	).WillReturnError(errors.New("update failed"))
	s.mock.ExpectRollback()

	result, err := s.repo.Update(od)

	s.Error(err)
	s.Contains(err.Error(), "update failed")
	s.Equal(models.OrderDetail{}, result)
}

func (s *OrderDetailTestSuite) TestPartialUpdate_Success() {
	id := 1
	original := models.OrderDetail{
		Id:               id,
		Quantity:         5,
		CleanLinesStatus: "OK",
		Temperature:      3.5,
		ProductRecordID:  101,
		PurchaseOrderID:  201,
	}

	updatedFields := map[string]interface{}{
		"quantity":           10.0,
		"clean_lines_status": "Updated",
		"temperature":        3.5,
		"product_record_id":  102.0,
		"purchase_order_id":  202.0,
	}

	// Mock SELECT exitoso
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `order_details` WHERE `order_details`.`id` = ? ORDER BY `order_details`.`id` LIMIT ?",
	)).WithArgs(id, 1).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "quantity", "clean_lines_status", "temperature", "product_record_id", "purchase_order_id",
		}).AddRow(
			original.Id, original.Quantity, original.CleanLinesStatus,
			original.Temperature, original.ProductRecordID, original.PurchaseOrderID,
		))

	// Expect UPDATE con los valores actualizados
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `order_details` SET `quantity`=?,`clean_lines_status`=?,`temperature`=?,`product_record_id`=?,`purchase_order_id`=? WHERE `id` = ?",
	)).WithArgs(
		10,        // actualizado
		"Updated", // actualizado
		3.5,       // mismo valor
		102,       // actualizado
		202,       // actualizado
		id,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Ejecutar
	result, err := s.repo.PartialUpdate(id, updatedFields)

	// Validar
	s.NoError(err)
	s.Equal(10, result.Quantity)
	s.Equal("Updated", result.CleanLinesStatus)
	s.Equal(3.5, result.Temperature)
	s.Equal(102, result.ProductRecordID)
	s.Equal(202, result.PurchaseOrderID)
}

func (s *OrderDetailTestSuite) TestPartialUpdate_NotFound() {
	id := 999
	fields := map[string]interface{}{
		"quantity": 10.0,
	}

	// Simula que no se encuentra el registro
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `order_details` WHERE `order_details`.`id` = ? ORDER BY `order_details`.`id` LIMIT ?",
	)).WithArgs(id, 1).
		WillReturnRows(sqlmock.NewRows([]string{})) // sin filas

	result, err := s.repo.PartialUpdate(id, fields)

	s.Error(err)
	s.Equal(models.OrderDetail{}, result)
	s.Contains(err.Error(), "record not found")
}

func (s *OrderDetailTestSuite) TestPartialUpdate_SaveError() {
	id := 1
	existing := models.OrderDetail{
		Id:               id,
		Quantity:         3,
		CleanLinesStatus: "OK",
		Temperature:      2.0,
		ProductRecordID:  50,
		PurchaseOrderID:  60,
	}

	fields := map[string]interface{}{
		"temperature": 4.4,
	}

	// Mock SELECT exitoso con LIMIT ?
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `order_details` WHERE `order_details`.`id` = ? ORDER BY `order_details`.`id` LIMIT ?",
	)).WithArgs(id, 1).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "quantity", "clean_lines_status", "temperature", "product_record_id", "purchase_order_id",
		}).AddRow(
			existing.Id, existing.Quantity, existing.CleanLinesStatus,
			existing.Temperature, existing.ProductRecordID, existing.PurchaseOrderID,
		))

	// Mock error al guardar
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `order_details` SET `quantity`=?,`clean_lines_status`=?,`temperature`=?,`product_record_id`=?,`purchase_order_id`=? WHERE `id` = ?",
	)).WithArgs(
		existing.Quantity,
		existing.CleanLinesStatus,
		4.4, // updated temperature
		existing.ProductRecordID,
		existing.PurchaseOrderID,
		id,
	).WillReturnError(errors.New("save failed"))
	s.mock.ExpectRollback()

	// Act
	result, err := s.repo.PartialUpdate(id, fields)

	// Assert
	s.Error(err)
	s.Contains(err.Error(), "save failed")
	s.Equal(models.OrderDetail{}, result)
}

func (s *OrderDetailTestSuite) TestDelete_Success() {
	id := 1

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"DELETE FROM `order_details` WHERE `order_details`.`id` = ?",
	)).WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 1)) // 1 row affected
	s.mock.ExpectCommit()

	err := s.repo.Delete(id)

	s.NoError(err)
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *OrderDetailTestSuite) TestDelete_Error() {
	id := 1

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"DELETE FROM `order_details` WHERE `order_details`.`id` = ?",
	)).WithArgs(id).
		WillReturnError(errors.New("delete failed"))
	s.mock.ExpectRollback()

	err := s.repo.Delete(id)

	s.Error(err)
	s.EqualError(err, "delete failed")
	s.NoError(s.mock.ExpectationsWereMet())
}

// Run the test suite
func TestOrderDetailRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(OrderDetailTestSuite))
}
