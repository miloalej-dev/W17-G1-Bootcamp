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

type BuyerRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo *BuyerRepository
}

func (s *BuyerRepositoryTestSuite) SetupSuite() {
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
	s.repo = NewBuyerRepository(gormDB)
}

func (s *BuyerRepositoryTestSuite) TestFindAll_Success() {

	expectedBuyers := []models.Buyer{
		{
			Id:           1,
			CardNumberId: "189-58-5819",
			FirstName:    "Donnamarie",
			LastName:     "Sharpless",
		},
		{
			Id:           2,
			CardNumberId: "174-53-5631",
			FirstName:    "john",
			LastName:     "Smith",
		},
	}

	rows := s.mock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"}).
		AddRow(expectedBuyers[0].Id, expectedBuyers[0].CardNumberId, expectedBuyers[0].FirstName,
			expectedBuyers[0].LastName).
		AddRow(expectedBuyers[1].Id, expectedBuyers[1].CardNumberId, expectedBuyers[1].FirstName,
			expectedBuyers[1].LastName)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `buyers`")).WillReturnRows(rows)

	buyers, err := s.repo.FindAll()

	s.NoError(err)
	s.Len(buyers, 2)
	s.Equal(expectedBuyers[0].CardNumberId, buyers[0].CardNumberId)
	s.Equal(expectedBuyers[1].CardNumberId, buyers[1].CardNumberId)
	s.Equal(expectedBuyers[0].FirstName, buyers[0].FirstName)
	s.Equal(expectedBuyers[1].FirstName, buyers[1].FirstName)
	s.Equal(expectedBuyers[0].LastName, buyers[0].LastName)
	s.Equal(expectedBuyers[1].LastName, buyers[1].LastName)
}

func (s *BuyerRepositoryTestSuite) TestFindAll_DatabaseError() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `buyers`")).WillReturnError(sql.ErrConnDone)

	// Act
	buyers, err := s.repo.FindAll()

	// Assert
	s.Error(err)
	s.Nil(buyers)
	s.Equal(sql.ErrConnDone, err)
}

func (s *BuyerRepositoryTestSuite) TestFindById_Success() {
	// Arrange
	expectedBuyer := models.Buyer{
		Id:           1,
		CardNumberId: "189-58-5819",
		FirstName:    "Donnamarie",
		LastName:     "Sharpless",
	}

	rows := s.mock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"}).
		AddRow(expectedBuyer.Id, expectedBuyer.CardNumberId, expectedBuyer.FirstName,
			expectedBuyer.LastName)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `buyers` WHERE `buyers`.`id` = ? ORDER BY `buyers`.`id` LIMIT ?")).
		WithArgs(1, 1).WillReturnRows(rows)

	// Act
	buyer, err := s.repo.FindById(1)

	// Assert
	s.NoError(err)
	s.Equal(expectedBuyer.Id, buyer.Id)
	s.Equal(expectedBuyer.CardNumberId, buyer.CardNumberId)
	s.Equal(expectedBuyer.FirstName, buyer.FirstName)
	s.Equal(expectedBuyer.LastName, buyer.LastName)
}

func (s *BuyerRepositoryTestSuite) TestFindById_NotFound() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `buyers` WHERE `buyers`.`id` = ? ORDER BY `buyers`.`id` LIMIT ?")).
		WithArgs(999, 1).WillReturnError(repository.ErrEntityNotFound)

	// Act
	buyer, err := s.repo.FindById(999)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal(models.Buyer{}, buyer)
}

func (s *BuyerRepositoryTestSuite) TestCreate_Success() {
	// Arrange
	expectedBuyer := models.Buyer{
		CardNumberId: "189-58-5819",
		FirstName:    "Donnamarie",
		LastName:     "Sharpless",
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `buyers` (`card_number_id`,`first_name`,`last_name`) VALUES (?,?,?)")).
		WithArgs(expectedBuyer.CardNumberId, expectedBuyer.FirstName, expectedBuyer.LastName).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	createdBuyer, err := s.repo.Create(expectedBuyer)

	// Assert
	s.NoError(err)
	s.Equal(1, createdBuyer.Id)
	s.Equal(createdBuyer.CardNumberId, expectedBuyer.CardNumberId)
	s.Equal(createdBuyer.FirstName, expectedBuyer.FirstName)
	s.Equal(createdBuyer.LastName, expectedBuyer.LastName)

}

func (s *BuyerRepositoryTestSuite) TestCreate_DatabaseError() {
	// Arrange
	expectedBuyer := models.Buyer{
		CardNumberId: "189-58-5819",
		FirstName:    "Donnamarie",
		LastName:     "Sharpless",
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `buyers` (`card_number_id`,`first_name`,`last_name`) VALUES (?,?,?)")).
		WithArgs(expectedBuyer.CardNumberId, expectedBuyer.FirstName, expectedBuyer.LastName).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act

	createdBuyer, err := s.repo.Create(expectedBuyer)

	// Assert
	s.Error(err)
	s.Equal(models.Buyer{}, createdBuyer)
	s.Equal(sql.ErrConnDone, err)
}

func (s *BuyerRepositoryTestSuite) TestUpdate_Success() {
	// Arrange
	existingBuyer := models.Buyer{
		Id:           1,
		CardNumberId: "189-58-5819",
		FirstName:    "Donnamarie",
		LastName:     "Sharpless",
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `buyers` SET `card_number_id`=?,`first_name`=?,`last_name`=? WHERE `id` = ?")).
		WithArgs(existingBuyer.CardNumberId, existingBuyer.FirstName, existingBuyer.LastName, existingBuyer.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	updatedBuyer, err := s.repo.Update(existingBuyer)

	// Assert
	s.NoError(err)
	s.Equal(existingBuyer.Id, updatedBuyer.Id)
	s.Equal(existingBuyer.CardNumberId, updatedBuyer.CardNumberId)
	s.Equal(existingBuyer.FirstName, updatedBuyer.FirstName)
	s.Equal(existingBuyer.LastName, updatedBuyer.LastName)

}

func (s *BuyerRepositoryTestSuite) TestUpdate_DatabaseError() {
	// Arrange
	existingBuyer := models.Buyer{
		Id:           1,
		CardNumberId: "189-58-5819",
		FirstName:    "Donnamarie",
		LastName:     "Sharpless",
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `buyers` SET `card_number_id`=?,`first_name`=?,`last_name`=? WHERE `id` = ?")).
		WithArgs(existingBuyer.CardNumberId, existingBuyer.FirstName, existingBuyer.LastName, existingBuyer.Id).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	updatedBuyer, err := s.repo.Update(existingBuyer)

	// Assert
	s.Error(err)
	s.Equal(models.Buyer{}, updatedBuyer)
	s.Equal(sql.ErrConnDone, err)

}

func (s *BuyerRepositoryTestSuite) TestPartialUpdate_Success() {
	// Arrange}
	buyerID := 1
	fields := map[string]interface{}{
		"first_name": "Donnamarie",
		"last_name":  "Sharpless",
	}

	expectedBuyer := models.Buyer{
		Id:           buyerID,
		CardNumberId: "189-58-5819",
		FirstName:    "Don",
		LastName:     "Sharp",
	}

	// First query to find the seller
	rows := s.mock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"}).
		AddRow(expectedBuyer.Id, expectedBuyer.CardNumberId, expectedBuyer.FirstName, expectedBuyer.LastName)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `buyers` WHERE `buyers`.`id` = ? ORDER BY `buyers`.`id` LIMIT ?")).
		WithArgs(buyerID, 1).WillReturnRows(rows)

	// Update query
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `buyers` SET `first_name`=?,`last_name`=? WHERE `id` = ?")).
		WithArgs(fields["first_name"], fields["last_name"], buyerID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	updatedBuyer, err := s.repo.PartialUpdate(buyerID, fields)

	// Assert
	s.NoError(err)
	s.Equal(buyerID, updatedBuyer.Id)
	s.Equal(expectedBuyer.CardNumberId, updatedBuyer.CardNumberId)
	s.Equal(fields["first_name"], updatedBuyer.FirstName)
	s.Equal(fields["last_name"], updatedBuyer.LastName)
}

func (s *BuyerRepositoryTestSuite) TestPartialUpdate_NotFound() {
	// Arrange
	buyerID := 999
	fields := map[string]interface{}{
		"first_name": "Donnamarie",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `buyers` WHERE `buyers`.`id` = ? ORDER BY `buyers`.`id` LIMIT ?")).
		WithArgs(buyerID, 1).WillReturnError(gorm.ErrRecordNotFound)

	// Act
	updatedBuyer, err := s.repo.PartialUpdate(buyerID, fields)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal(models.Buyer{}, updatedBuyer)
}

func (s *BuyerRepositoryTestSuite) TestPartialUpdate_DatabaseError() {
	// Arrange
	buyerID := 1
	fields := map[string]interface{}{
		"first_name": "Donnamarie",
	}

	expectedBuyer := models.Buyer{
		Id:           buyerID,
		CardNumberId: "189-58-5819",
		FirstName:    "Don",
		LastName:     "Sharp",
	}

	// First query to find the seller

	rows := s.mock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"}).
		AddRow(expectedBuyer.Id, expectedBuyer.CardNumberId, expectedBuyer.FirstName, expectedBuyer.LastName)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `buyers` WHERE `buyers`.`id` = ? ORDER BY `buyers`.`id` LIMIT ?")).
		WithArgs(buyerID, 1).WillReturnRows(rows)

	// Update query with database error
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `buyers` SET `first_name`=? WHERE `id` = ?")).
		WithArgs(fields["first_name"], buyerID).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	updatedBuyer, err := s.repo.PartialUpdate(buyerID, fields)

	// Assert
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.Buyer{}, updatedBuyer)
}
func (s *BuyerRepositoryTestSuite) TestDelete_Success() {
	// Arrange
	buyerID := 1
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `buyers` WHERE `buyers`.`id` = ?")).
		WithArgs(buyerID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	// Act
	err := s.repo.Delete(buyerID)
	// Assert
	s.NoError(err)
}

func (s *BuyerRepositoryTestSuite) TestDelete_NotFound() {
	// Arrange
	buyerID := 999

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `buyers` WHERE `buyers`.`id` = ?")).
		WithArgs(buyerID).
		WillReturnResult(sqlmock.NewResult(1, 0)) // 0 rows affected
	s.mock.ExpectCommit()

	// Act
	err := s.repo.Delete(buyerID)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
}

func (s *BuyerRepositoryTestSuite) TestFindByPurchaseOrderReport_AllBuyers_Success() {

	expectedBuyerReport := []models.BuyerReport{
		{
			Buyer: models.Buyer{
				Id:           1,
				CardNumberId: "189-58",
				FirstName:    "Donna",
				LastName:     "Sharp",
			},
			PurchaseOrdersCount: 4,
		},
		{
			Buyer: models.Buyer{
				Id:           1,
				CardNumberId: "189-59",
				FirstName:    "Marie",
				LastName:     "less",
			},
			PurchaseOrdersCount: 2,
		},
	}
	// Arrange:
	rows := s.mock.NewRows([]string{
		"id", "card_number_id", "first_name", "last_name", "purchase_orders_count",
	}).AddRow(expectedBuyerReport[0].Id, expectedBuyerReport[0].CardNumberId, expectedBuyerReport[0].FirstName, expectedBuyerReport[0].LastName, expectedBuyerReport[0].PurchaseOrdersCount).
		AddRow(expectedBuyerReport[1].Id, expectedBuyerReport[1].CardNumberId, expectedBuyerReport[1].FirstName, expectedBuyerReport[1].LastName, expectedBuyerReport[1].PurchaseOrdersCount)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT buyers.id, buyers.card_number_id, buyers.first_name, buyers.last_name, COUNT(purchase_orders.id) AS purchase_orders_count FROM `buyers` LEFT JOIN purchase_orders ON purchase_orders.buyer_id = buyers.id GROUP BY `buyers`.`id`")).
		WillReturnRows(rows)

	// Act
	result, err := s.repo.FindByPurchaseOrderReport(0)

	// Assert
	s.NoError(err)
	s.Len(result, 2)
	s.Equal(expectedBuyerReport[0].FirstName, result[0].FirstName)
	s.Equal(expectedBuyerReport[0].PurchaseOrdersCount, result[0].PurchaseOrdersCount)
	s.Equal(expectedBuyerReport[1].FirstName, result[1].FirstName)
	s.Equal(expectedBuyerReport[1].PurchaseOrdersCount, result[1].PurchaseOrdersCount)
}

func (s *BuyerRepositoryTestSuite) TestFindByPurchaseOrderReport_SingleBuyer_Success() {
	id := 1
	expected := models.BuyerReport{
		Buyer: models.Buyer{
			Id:           id,
			CardNumberId: "189-58",
			FirstName:    "Donna",
			LastName:     "Sharp",
		},
		PurchaseOrdersCount: 4,
	}

	// Mock para FindById (nota el LIMIT ?)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `buyers` WHERE `buyers`.`id` = ? ORDER BY `buyers`.`id` LIMIT ?",
	)).WithArgs(id, 1).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "card_number_id", "first_name", "last_name",
		}).AddRow(
			expected.Id, expected.CardNumberId, expected.FirstName, expected.LastName,
		))

	// Mock para query de reporte individual (solo un argumento: id)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT buyers.id, buyers.card_number_id, buyers.first_name, buyers.last_name,  COUNT(purchase_orders.id) AS purchase_orders_count FROM `buyers` LEFT JOIN purchase_orders ON purchase_orders.buyers_id = buyers.id WHERE buyers.id = ? GROUP BY `buyers`.`id`",
	)).WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "card_number_id", "first_name", "last_name", "purchase_orders_count",
		}).AddRow(
			expected.Id, expected.CardNumberId, expected.FirstName, expected.LastName, expected.PurchaseOrdersCount,
		))

	// Act
	result, err := s.repo.FindByPurchaseOrderReport(id)

	// Assert
	s.NoError(err)
	s.Len(result, 1)
	s.Equal(expected.FirstName, result[0].FirstName)
	s.Equal(expected.PurchaseOrdersCount, result[0].PurchaseOrdersCount)
}

func (s *BuyerRepositoryTestSuite) TestFindByPurchaseOrderReport_SingleBuyer_NotFound() {
	id := 42

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `buyers` WHERE `buyers`.`id` = ? ORDER BY `buyers`.`id` LIMIT ?",
	)).WithArgs(id, 1).
		WillReturnRows(sqlmock.NewRows([]string{})) // sin filas

	// Act
	result, err := s.repo.FindByPurchaseOrderReport(id)

	// Assert
	s.Error(err)
	s.Contains(err.Error(), "record not found")
	s.Empty(result)
}

func (s *BuyerRepositoryTestSuite) TestFindByPurchaseOrderReport_SingleBuyer_ScanError() {
	id := 1

	// Mock para FindById exitoso
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `buyers` WHERE `buyers`.`id` = ? ORDER BY `buyers`.`id` LIMIT ?",
	)).WithArgs(id, 1).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "card_number_id", "first_name", "last_name",
		}).AddRow(id, "X", "John", "Doe"))

	// Mock para query del reporte con error
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT buyers.id, buyers.card_number_id, buyers.first_name, buyers.last_name,  COUNT(purchase_orders.id) AS purchase_orders_count FROM `buyers` LEFT JOIN purchase_orders ON purchase_orders.buyers_id = buyers.id WHERE buyers.id = ? GROUP BY `buyers`.`id`",
	)).WithArgs(id).
		WillReturnError(errors.New("scan failed"))

	// Act
	result, err := s.repo.FindByPurchaseOrderReport(id)

	// Assert
	s.Error(err)
	s.Contains(err.Error(), "scan failed")
	s.Empty(result)
}

func (s *BuyerRepositoryTestSuite) TestFindByPurchaseOrderReport_AllBuyers_QueryError() {
	// id == 0 → se ejecuta la rama que consulta todos los buyers
	errExpected := errors.New("db failed")

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT buyers.id, buyers.card_number_id, buyers.first_name, buyers.last_name, COUNT(purchase_orders.id) AS purchase_orders_count FROM `buyers` LEFT JOIN purchase_orders ON purchase_orders.buyer_id = buyers.id GROUP BY `buyers`.`id`",
	)).WillReturnError(errExpected)

	// Act
	result, err := s.repo.FindByPurchaseOrderReport(0)

	// Assert
	s.Error(err)
	s.EqualError(err, "db failed")
	s.Empty(result)
}

// Run the test suite
func TestBuyerRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BuyerRepositoryTestSuite))
}
