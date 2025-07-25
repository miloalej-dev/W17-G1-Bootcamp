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
	"time"
)

type PurchaseOrderTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo *PurchaseOrderRepository
}

func (s *PurchaseOrderTestSuite) SetupSuite() {
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
	s.repo = NewPurchaseOrderRepository(gormDB)

}

func (s *PurchaseOrderTestSuite) SetupTest() {
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
	s.repo = NewPurchaseOrderRepository(gormDB)

}

func (s *PurchaseOrderTestSuite) TestFindAll_Success() {
	purchaseOrder := []models.PurchaseOrder{
		{
			Id:            1,
			OrderNumber:   "PO-20250715-020",
			OrderDate:     time.Now(),
			TracingCode:   "TRC004",
			BuyerID:       1,
			WarehouseID:   2,
			CarrierID:     1,
			OrderStatusID: 1,
		}, {
			Id:            2,
			OrderNumber:   "PO-20250715-021",
			OrderDate:     time.Now(),
			TracingCode:   "TRC005",
			BuyerID:       1,
			WarehouseID:   2,
			CarrierID:     1,
			OrderStatusID: 1,
		},
	}
	rows := s.mock.NewRows([]string{
		"id", "order_number", "order_date", "tracing_code",
		"buyer_id", "warehouse_id", "carrier_id", "order_status_id",
	}).
		AddRow(
			purchaseOrder[0].Id,
			purchaseOrder[0].OrderNumber,
			purchaseOrder[0].OrderDate,
			purchaseOrder[0].TracingCode,
			purchaseOrder[0].BuyerID,
			purchaseOrder[0].WarehouseID,
			purchaseOrder[0].CarrierID,
			purchaseOrder[0].OrderStatusID,
		).
		AddRow(
			purchaseOrder[1].Id,
			purchaseOrder[1].OrderNumber,
			purchaseOrder[1].OrderDate,
			purchaseOrder[1].TracingCode,
			purchaseOrder[1].BuyerID,
			purchaseOrder[1].WarehouseID,
			purchaseOrder[1].CarrierID,
			purchaseOrder[1].OrderStatusID,
		)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `purchase_orders`")).WillReturnRows(rows)

	// Act
	po, err := s.repo.FindAll()

	// Assert
	s.NoError(err)
	s.Len(po, 2)
	s.Equal(purchaseOrder[0].OrderNumber, po[0].OrderNumber)
	s.Equal(purchaseOrder[1].OrderNumber, po[1].OrderNumber)
}

func (s *PurchaseOrderTestSuite) TestFindAll_DatabaseError() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `purchase_orders`")).WillReturnError(sql.ErrConnDone)
	// Act
	purchaseOrder, err := s.repo.FindAll()

	// Assert
	s.Error(err)
	s.Nil(purchaseOrder)
	s.Equal(sql.ErrConnDone, err)

}

func (s *PurchaseOrderTestSuite) TestFindById_Success() {
	// Arrange
	purchaseOrder := models.PurchaseOrder{
		Id:            1,
		OrderNumber:   "PO-20250715-020",
		OrderDate:     time.Now(),
		TracingCode:   "TRC004",
		BuyerID:       1,
		WarehouseID:   2,
		CarrierID:     1,
		OrderStatusID: 1,
	}

	rows := s.mock.NewRows([]string{
		"id", "order_number", "order_date", "tracing_code",
		"buyer_id", "warehouse_id", "carrier_id", "order_status_id",
	}).
		AddRow(
			purchaseOrder.Id,
			purchaseOrder.OrderNumber,
			purchaseOrder.OrderDate,
			purchaseOrder.TracingCode,
			purchaseOrder.BuyerID,
			purchaseOrder.WarehouseID,
			purchaseOrder.CarrierID,
			purchaseOrder.OrderStatusID,
		)
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `purchase_orders` WHERE `purchase_orders`.`id` = ? ORDER BY `purchase_orders`.`id` LIMIT ?")).
		WithArgs(1, 1).WillReturnRows(rows)

	// Act
	po, err := s.repo.FindById(1)

	// Assert
	s.NoError(err)
	s.Equal(purchaseOrder.Id, po.Id)
	s.Equal(purchaseOrder.OrderNumber, po.OrderNumber)

}

func (s *PurchaseOrderTestSuite) TestFindById_NotFound() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `purchase_orders` WHERE `purchase_orders`.`id` = ? ORDER BY `purchase_orders`.`id` LIMIT ?",
	)).WithArgs(999, 1).WillReturnError(repository.ErrEntityNotFound)
	// Act
	po, err := s.repo.FindById(999)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal(models.PurchaseOrder{}, po)
}

func (s *PurchaseOrderTestSuite) TestCreate_Success() {
	orderDetail := []models.OrderDetail{
		{
			Quantity:         10,
			CleanLinesStatus: "OK",
			Temperature:      5.5,
			ProductRecordID:  100,
		},
	}
	purchaseOrder := models.PurchaseOrder{
		OrderNumber:   "PO-20250715-020",
		OrderDate:     time.Now(),
		TracingCode:   "TRC004",
		BuyerID:       1,
		WarehouseID:   2,
		CarrierID:     1,
		OrderStatusID: 1,
		OrderDetails:  &orderDetail,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `purchase_orders` (`order_number`,`order_date`,`tracing_code`,`buyer_id`,`warehouse_id`,`carrier_id`,`order_status_id`) VALUES (?,?,?,?,?,?,?)")).
		WithArgs(purchaseOrder.OrderNumber, purchaseOrder.OrderDate, purchaseOrder.TracingCode, purchaseOrder.BuyerID, purchaseOrder.WarehouseID, purchaseOrder.CarrierID, purchaseOrder.OrderStatusID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	s.mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `order_details` (`quantity`,`clean_lines_status`,`temperature`,`product_record_id`,`purchase_order_id`) VALUES (?,?,?,?,?)")).
		WithArgs(
			orderDetail[0].Quantity,
			orderDetail[0].CleanLinesStatus,
			orderDetail[0].Temperature,
			orderDetail[0].ProductRecordID,
			sqlmock.AnyArg(), // O usa 1 si tu código setea el id a mano en el struct
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	s.mock.ExpectCommit()

	createdPO, err := s.repo.Create(purchaseOrder)

	s.NoError(err)
	s.Equal(purchaseOrder.OrderNumber, createdPO.OrderNumber)
	// No asegures el Id aquí porque queda 0
	// s.NotZero(createdPO.OrderNumber) // sólo para asegurar que algo se retornó
}

func (s *PurchaseOrderTestSuite) TestCreate_ForeignKeyViolated() {
	// Arrange
	orderDetail := []models.OrderDetail{}
	purchaseOrder := models.PurchaseOrder{
		OrderNumber:   "PO-20250715-020",
		OrderDate:     time.Now(),
		TracingCode:   "TRC004",
		BuyerID:       1,
		WarehouseID:   2,
		CarrierID:     1,
		OrderStatusID: 1,
		OrderDetails:  &orderDetail,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `purchase_orders` (`order_number`,`order_date`,`tracing_code`,`buyer_id`,`warehouse_id`,`carrier_id`,`order_status_id`) VALUES (?,?,?,?,?,?,?)")).
		WithArgs(purchaseOrder.OrderNumber, purchaseOrder.OrderDate, purchaseOrder.TracingCode, purchaseOrder.BuyerID, purchaseOrder.WarehouseID, purchaseOrder.CarrierID, purchaseOrder.OrderStatusID).
		WillReturnError(repository.ErrForeignKeyViolation)
	s.mock.ExpectRollback()

	// Act
	createdPO, err := s.repo.Create(purchaseOrder)

	// Assert
	s.Error(err)
	s.Equal(models.PurchaseOrder{}, createdPO)

}

func (s *PurchaseOrderTestSuite) TestCreate_BeginTransactionFails() {
	// Arrange
	orderDetail := []models.OrderDetail{}
	purchaseOrder := models.PurchaseOrder{
		OrderNumber:   "PO-20250715-020",
		OrderDate:     time.Now(),
		TracingCode:   "TRC004",
		BuyerID:       1,
		WarehouseID:   2,
		CarrierID:     1,
		OrderStatusID: 1,
		OrderDetails:  &orderDetail,
	}
	s.mock.ExpectBegin().WillReturnError(errors.New("begin failed"))

	// Act
	createdPO, err := s.repo.Create(purchaseOrder)

	// Assert
	s.Error(err)
	s.EqualError(err, "begin failed")
	s.Equal(models.PurchaseOrder{}, createdPO)
}
func (s *PurchaseOrderTestSuite) TestCreate_OrderDetailFails() {
	// Arrange
	detail := models.OrderDetail{
		Quantity:         10,
		CleanLinesStatus: "OK",
		Temperature:      5.5,
		ProductRecordID:  100,
	}
	orderDetails := []models.OrderDetail{detail}

	purchaseOrder := models.PurchaseOrder{
		OrderNumber:   "PO-20250715-020",
		OrderDate:     time.Now(),
		TracingCode:   "TRC004",
		BuyerID:       1,
		WarehouseID:   2,
		CarrierID:     1,
		OrderStatusID: 1,
		OrderDetails:  &orderDetails,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `purchase_orders` (`order_number`,`order_date`,`tracing_code`,`buyer_id`,`warehouse_id`,`carrier_id`,`order_status_id`) VALUES (?,?,?,?,?,?,?)")).
		WithArgs(
			purchaseOrder.OrderNumber,
			purchaseOrder.OrderDate,
			purchaseOrder.TracingCode,
			purchaseOrder.BuyerID,
			purchaseOrder.WarehouseID,
			purchaseOrder.CarrierID,
			purchaseOrder.OrderStatusID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	s.mock.ExpectExec("INSERT INTO `order_details`").
		WithArgs(detail.Quantity, detail.CleanLinesStatus, detail.Temperature, detail.ProductRecordID, 1).
		WillReturnError(errors.New("order detail insert failed"))

	s.mock.ExpectRollback()

	// Act
	createdPO, err := s.repo.Create(purchaseOrder)

	// Assert
	s.Error(err)
	s.Contains(err.Error(), "order detail insert failed")
	s.Equal(models.PurchaseOrder{}, createdPO)
}

func (s *PurchaseOrderTestSuite) TestCreate_OrderDetailsNil() {
	// Arrange
	purchaseOrder := models.PurchaseOrder{
		OrderNumber:   "PO-20250715-999",
		OrderDate:     time.Now(),
		TracingCode:   "TRC-ORDERDETAILS-NIL",
		BuyerID:       1,
		WarehouseID:   2,
		CarrierID:     1,
		OrderStatusID: 1,
		OrderDetails:  nil, // <-- nil aquí
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `purchase_orders` (`order_number`,`order_date`,`tracing_code`,`buyer_id`,`warehouse_id`,`carrier_id`,`order_status_id`) VALUES (?,?,?,?,?,?,?)"),
	).WithArgs(
		purchaseOrder.OrderNumber,
		purchaseOrder.OrderDate,
		purchaseOrder.TracingCode,
		purchaseOrder.BuyerID,
		purchaseOrder.WarehouseID,
		purchaseOrder.CarrierID,
		purchaseOrder.OrderStatusID,
	).WillReturnResult(sqlmock.NewResult(1, 1)) // Simulas que el insert fue bien
	s.mock.ExpectRollback() // Porque el rollback SÍ ocurre si ordersDetails == nil

	// Act
	createdPO, err := s.repo.Create(purchaseOrder)

	// Assert
	s.Error(err)
	s.Equal(models.PurchaseOrder{}, createdPO)

}

func (s *PurchaseOrderTestSuite) TestUpdate_Success() {
	po := models.PurchaseOrder{
		Id:            1,
		OrderNumber:   "1234",
		OrderDate:     time.Now(),
		TracingCode:   "TRC004",
		BuyerID:       1,
		WarehouseID:   2,
		CarrierID:     1,
		OrderStatusID: 1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `purchase_orders` SET `order_number`=?,`order_date`=?,`tracing_code`=?,`buyer_id`=?,`warehouse_id`=?,`carrier_id`=?,`order_status_id`=? WHERE `id` = ?",
	)).WithArgs(
		po.OrderNumber, po.OrderDate, po.TracingCode, po.BuyerID,
		po.WarehouseID, po.CarrierID, po.OrderStatusID, po.Id,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	result, err := s.repo.Update(po)
	s.NoError(err)
	s.Equal(po, result)
}

func (s *PurchaseOrderTestSuite) TestCreate_CommitError() {
	// 1. Setup tu modelo
	po := models.PurchaseOrder{
		// ...pon aquí los campos necesarios, incluyendo OrderDetails válidos...
		OrderDetails: &[]models.OrderDetail{
			{
				Quantity:         1,
				CleanLinesStatus: "OK",
				Temperature:      1.0,
				ProductRecordID:  1,
			},
		},
	}

	// 2. Espera el BEGIN de la transacción
	s.mock.ExpectBegin()

	// 3. Espera el INSERT del purchase_order (ajusta tabla/nombre/campos según tu modelo real)
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `purchase_orders`")).
		WithArgs( /* pon aquí los expected args o sqlmock.AnyArg() si no importa */ ).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// 4. Espera el INSERT del order_detail (igual, ajusta los args; repite si tienes más)
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `order_details`")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// 5. Espera un Commit que FALLE
	s.mock.ExpectCommit().WillReturnError(errors.New("commit failed"))

	// 6. Espera el Rollback que hace tu código al fallar el commit
	s.mock.ExpectRollback()

	// 7. Ejecuta el método
	_, err := s.repo.Create(po)

	// 8. Asegúrate que el error es correcto
	s.Error(err)
	s.Contains(err.Error(), "commit failed")

}

func (s *PurchaseOrderTestSuite) TestUpdate_Error() {
	po := models.PurchaseOrder{
		Id:          999, // ID que simulamos que no existe
		OrderNumber: "XXXX",
	}

	// Simula que el UPDATE falla
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `purchase_orders` SET `order_number`=?,`order_date`=?,`tracing_code`=?,`buyer_id`=?,`warehouse_id`=?,`carrier_id`=?,`order_status_id`=? WHERE `id` = ?",
	)).WithArgs(
		po.OrderNumber, po.OrderDate, po.TracingCode, po.BuyerID,
		po.WarehouseID, po.CarrierID, po.OrderStatusID, po.Id,
	).WillReturnError(errors.New("update failed"))
	s.mock.ExpectRollback()

	// Ejecutar
	result, err := s.repo.Update(po)

	// Verificar
	s.Error(err)
	s.EqualError(err, "update failed")
	s.Equal(models.PurchaseOrder{}, result)
}

func (s *PurchaseOrderTestSuite) TestPartialUpdate_Success() {
	id := 1
	now := time.Now()

	fields := map[string]interface{}{
		"order_number":    "9999",
		"tracing_code":    "TRACK123",
		"buyer_id":        float64(3),
		"order_status_id": float64(1),
		"order_date":      now,
		"warehouse_id":    float64(1),
		"carrier_id":      float64(1),
	}

	// Mock del SELECT previo
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `purchase_orders` WHERE `purchase_orders`.`id` = ? ORDER BY `purchase_orders`.`id` LIMIT ?",
	)).
		WithArgs(id, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "order_number", "order_date", "tracing_code",
			"buyer_id", "warehouse_id", "carrier_id", "order_status_id",
		}).AddRow(
			id, "1234", now, "OLD", 2, 0, 0, 0,
		))

	// Mock del UPDATE
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `purchase_orders` SET `order_number`=?,`order_date`=?,`tracing_code`=?,`buyer_id`=?,`warehouse_id`=?,`carrier_id`=?,`order_status_id`=? WHERE `id` = ?",
	)).WithArgs(
		"9999",     // order_number
		now,        // order_date
		"TRACK123", // tracing_code
		3,          // buyer_id
		1,          // warehouse_id
		1,          // carrier_id
		1,          // order_status_id
		id,         // WHERE id = ?
	).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Ejecutar función
	result, err := s.repo.PartialUpdate(id, fields)

	// Validaciones
	s.NoError(err)
	s.Equal("9999", result.OrderNumber)
	s.Equal("TRACK123", result.TracingCode)
	s.Equal(3, result.BuyerID)
	s.Equal(1, result.OrderStatusID)
	s.Equal(1, result.WarehouseID)
	s.Equal(1, result.CarrierID)
	s.WithinDuration(now, result.OrderDate, time.Second)
}

func (s *PurchaseOrderTestSuite) TestPartialUpdate_SaveError() {
	id := 1
	fields := map[string]interface{}{
		"order_number":    "9999",
		"tracing_code":    "TRACK123",
		"buyer_id":        float64(3),
		"order_status_id": float64(1),
	}

	// Simula que sí encuentra el registro en el SELECT
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `purchase_orders` WHERE `purchase_orders`.`id` = ? ORDER BY `purchase_orders`.`id` LIMIT ?",
	)).
		WithArgs(id, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "order_number", "order_date", "tracing_code",
			"buyer_id", "warehouse_id", "carrier_id", "order_status_id",
		}).
			AddRow(id, "1234", time.Now(), "OLD", 2, 0, 0, 0))

	// Simula que el `Save()` falla
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `purchase_orders` SET `order_number`=?,`order_date`=?,`tracing_code`=?,`buyer_id`=?,`warehouse_id`=?,`carrier_id`=?,`order_status_id`=? WHERE `id` = ?",
	)).
		WithArgs("9999", sqlmock.AnyArg(), "TRACK123", 3, 0, 0, 1, id).
		WillReturnError(errors.New("update failed"))
	s.mock.ExpectRollback()

	// Ejecuta
	result, err := s.repo.PartialUpdate(id, fields)

	// Verifica
	s.Error(err)
	s.EqualError(err, "update failed")
	s.Equal(models.PurchaseOrder{}, result)
}
func (s *PurchaseOrderTestSuite) TestPartialUpdate_FirstFails() {
	id := 999
	fields := map[string]interface{}{
		"order_number": "NOPE",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `purchase_orders` WHERE `purchase_orders`.`id` = ? ORDER BY `purchase_orders`.`id` LIMIT ?",
	)).WithArgs(id, sqlmock.AnyArg()).
		WillReturnError(errors.New("record not found"))

	result, err := s.repo.PartialUpdate(id, fields)

	s.Error(err)
	s.EqualError(err, "record not found")
	s.Equal(models.PurchaseOrder{}, result)
}

func (s *PurchaseOrderTestSuite) TestDelete_Success() {
	id := 1

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"DELETE FROM `purchase_orders` WHERE `purchase_orders`.`id` = ?",
	)).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.repo.Delete(id)
	s.NoError(err)
}

func (s *PurchaseOrderTestSuite) TestFindByBuyerId_Success() {
	buyerID := 1

	rows := sqlmock.NewRows([]string{"id", "order_number", "buyer_id"}).
		AddRow(1, "ORD-001", buyerID).
		AddRow(2, "ORD-002", buyerID)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `purchase_orders` WHERE buyer_id = ?")).
		WithArgs(buyerID).WillReturnRows(rows)

	result, err := s.repo.FindByBuyerId(buyerID)
	s.NoError(err)
	s.Len(result, 2)
	s.Equal("ORD-001", result[0].OrderNumber)
}

func (s *PurchaseOrderTestSuite) TestFindByBuyerId_Fail() {
	buyerID := 999 // <-- el que vas a buscar

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `purchase_orders` WHERE buyer_id = ?")).
		WithArgs(buyerID).
		WillReturnError(repository.ErrEntityNotFound)
	// Act
	po, err := s.repo.FindByBuyerId(buyerID)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal([]models.PurchaseOrder{}, po)

}

// Run the test suite
func TestOrderPurchaseRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PurchaseOrderTestSuite))
}
