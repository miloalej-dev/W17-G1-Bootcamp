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

type ProductRecordRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo *ProductRecordRepository
}

func (s *ProductRecordRepositoryTestSuite) SetupSuite() {
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
	s.repo = NewProductRecordRepository(gormDB)
}

func (s *ProductRecordRepositoryTestSuite) TestFindAll_Success() {

	expectedProductRecords := []models.ProductRecord{
		{
			Id:            1,
			LastUpdate:    "2022-22-10",
			PurchasePrice: 4.99,
			SalePrice:     5.99,
			ProductId:     1,
		},
		{
			Id:            2,
			LastUpdate:    "2022-22-11",
			PurchasePrice: 4.99,
			SalePrice:     6.99,
			ProductId:     1,
		},
	}

	rows := s.mock.NewRows([]string{"id", "last_update", "purchase_price", "sale_price", "product_id"}).
		AddRow(expectedProductRecords[0].Id, expectedProductRecords[0].LastUpdate, expectedProductRecords[0].PurchasePrice,
			expectedProductRecords[0].SalePrice, expectedProductRecords[0].ProductId).
		AddRow(expectedProductRecords[1].Id, expectedProductRecords[1].LastUpdate, expectedProductRecords[1].PurchasePrice,
			expectedProductRecords[1].SalePrice, expectedProductRecords[1].ProductId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `product_records`")).WillReturnRows(rows)

	productRecord, err := s.repo.FindAll()

	s.NoError(err)
	s.Len(productRecord, 2)
	s.Equal(expectedProductRecords[0].Id, productRecord[0].Id)
	s.Equal(expectedProductRecords[1].Id, productRecord[1].Id)
	s.Equal(expectedProductRecords[0].LastUpdate, productRecord[0].LastUpdate)
	s.Equal(expectedProductRecords[1].LastUpdate, productRecord[1].LastUpdate)
	s.Equal(expectedProductRecords[0].ProductId, productRecord[0].ProductId)
	s.Equal(expectedProductRecords[1].ProductId, productRecord[1].ProductId)
	s.Equal(expectedProductRecords[0].SalePrice, productRecord[0].SalePrice)
	s.Equal(expectedProductRecords[1].SalePrice, productRecord[1].SalePrice)
}

func (s *ProductRecordRepositoryTestSuite) TestFindAll_DatabaseError() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `product_records`")).WillReturnError(sql.ErrConnDone)

	// Act
	productRecord, err := s.repo.FindAll()

	// Assert
	s.Error(err)
	s.Nil(productRecord)
	s.Equal(sql.ErrConnDone, err)
}

func (s *ProductRecordRepositoryTestSuite) TestFindById_Success() {
	// Arrange
	expectedProductRecord := models.ProductRecord{
		Id:            1,
		LastUpdate:    "2022-22-10",
		PurchasePrice: 4.99,
		SalePrice:     5.99,
		ProductId:     1,
	}

	rows := s.mock.NewRows([]string{"id", "last_update", "purchase_price", "sale_price", "product_id"}).
		AddRow(expectedProductRecord.Id, expectedProductRecord.LastUpdate, expectedProductRecord.PurchasePrice,
			expectedProductRecord.SalePrice, expectedProductRecord.ProductId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `product_records` WHERE `product_records`.`id` = ? ORDER BY `product_records`.`id` LIMIT ?")).
		WithArgs(1, 1).WillReturnRows(rows)

	// Act
	productRecord, err := s.repo.FindById(1)

	// Assert
	s.NoError(err)
	s.Equal(expectedProductRecord.Id, productRecord.Id)
	s.Equal(expectedProductRecord.LastUpdate, productRecord.LastUpdate)
	s.Equal(expectedProductRecord.PurchasePrice, productRecord.PurchasePrice)
	s.Equal(expectedProductRecord.ProductId, productRecord.ProductId)
}

func (s *ProductRecordRepositoryTestSuite) TestFindById_NotFound() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `product_records` WHERE `product_records`.`id` = ? ORDER BY `product_records`.`id` LIMIT ?")).
		WithArgs(999, 1).WillReturnError(repository.ErrEntityNotFound)

	// Act
	productRecord, err := s.repo.FindById(999)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal(models.ProductRecord{}, productRecord)
}

func (s *ProductRecordRepositoryTestSuite) TestCreate_Success() {
	// Arrange
	expectedProductRecord := models.ProductRecord{
		LastUpdate:    "2022-22-10",
		PurchasePrice: 4.99,
		SalePrice:     5.99,
		ProductId:     1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `product_records` (`last_update`,`purchase_price`,`sale_price`,`product_id`) VALUES (?,?,?,?)")).
		WithArgs(expectedProductRecord.LastUpdate, expectedProductRecord.PurchasePrice,
			expectedProductRecord.SalePrice, expectedProductRecord.ProductId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act

	createdProductRecord, err := s.repo.Create(expectedProductRecord)

	// Assert
	s.NoError(err)
	s.Equal(1, createdProductRecord.Id)
	s.Equal(createdProductRecord.LastUpdate, expectedProductRecord.LastUpdate)
	s.Equal(createdProductRecord.PurchasePrice, expectedProductRecord.PurchasePrice)
	s.Equal(createdProductRecord.ProductId, expectedProductRecord.ProductId)
}

func (s *ProductRecordRepositoryTestSuite) TestCreate_DatabaseError() {
	// Arrange
	expectedProductRecord := models.ProductRecord{
		LastUpdate:    "2022-22-10",
		PurchasePrice: 4.99,
		SalePrice:     5.99,
		ProductId:     1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `product_records` (`last_update`,`purchase_price`,`sale_price`,`product_id`) VALUES (?,?,?,?)")).
		WithArgs(expectedProductRecord.LastUpdate, expectedProductRecord.PurchasePrice,
			expectedProductRecord.SalePrice, expectedProductRecord.ProductId).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act

	createdProductRecord, err := s.repo.Create(expectedProductRecord)

	// Assert
	s.Error(err)
	s.Equal(models.ProductRecord{}, createdProductRecord)
	s.Equal(sql.ErrConnDone, err)
}

func (s *ProductRecordRepositoryTestSuite) TestUpdate_Success() {
	// Arrange
	expectedProductRecord := models.ProductRecord{
		Id:            1,
		LastUpdate:    "2022-22-10",
		PurchasePrice: 4.99,
		SalePrice:     5.99,
		ProductId:     1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `product_records` SET `last_update`=?,`purchase_price`=?,`sale_price`=?,`product_id`=? WHERE `id` = ?")).
		WithArgs(
			expectedProductRecord.LastUpdate,
			expectedProductRecord.PurchasePrice,
			expectedProductRecord.SalePrice,
			expectedProductRecord.ProductId,
			1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	updatedProductRecord, err := s.repo.Update(expectedProductRecord)

	// Assert

	s.NoError(err)
	s.Equal(expectedProductRecord.Id, updatedProductRecord.Id)
	s.Equal(expectedProductRecord.LastUpdate, updatedProductRecord.LastUpdate)
	s.Equal(expectedProductRecord.PurchasePrice, updatedProductRecord.PurchasePrice)
	s.Equal(expectedProductRecord.SalePrice, updatedProductRecord.SalePrice)
	s.Equal(expectedProductRecord.ProductId, updatedProductRecord.ProductId)

}

func (s *ProductRecordRepositoryTestSuite) TestUpdate_DatabaseError() {
	// Arrange

	expectedProductRecord := models.ProductRecord{
		Id:            1,
		LastUpdate:    "2022-22-10",
		PurchasePrice: 4.99,
		SalePrice:     5.99,
		ProductId:     1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `product_records` SET `last_update`=?,`purchase_price`=?,`sale_price`=?,`product_id`=? WHERE `id` = ?")).
		WithArgs(expectedProductRecord.LastUpdate, expectedProductRecord.PurchasePrice,
			expectedProductRecord.SalePrice, expectedProductRecord.ProductId, 1).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	updatedProductRecord, err := s.repo.Update(expectedProductRecord)

	// Assert
	s.Error(err)
	s.Equal(models.ProductRecord{}, updatedProductRecord)
	s.Equal(sql.ErrConnDone, err)

}

func (s *ProductRecordRepositoryTestSuite) TestPartialUpdate_Success() {
	// Arrange}
	productRecordID := 1
	fields := map[string]interface{}{
		"last_update":    "2022-22-10",
		"purchase_price": 4.99,
	}

	expectedProductRecord := models.ProductRecord{
		LastUpdate:    "2022-22-10",
		PurchasePrice: 4.99,
		SalePrice:     5.99,
		ProductId:     1,
	}

	rows := s.mock.NewRows([]string{"id", "last_update", "purchase_price", "sale_price", "product_id"}).
		AddRow(productRecordID, expectedProductRecord.LastUpdate, expectedProductRecord.PurchasePrice, expectedProductRecord.SalePrice, expectedProductRecord.ProductId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `product_records` WHERE `product_records`.`id` = ? ORDER BY `product_records`.`id` LIMIT ?")).
		WithArgs(productRecordID, 1).WillReturnRows(rows)

	// Update query
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `product_records` SET `last_update`=?,`purchase_price`=? WHERE `id` = ?")).
		WithArgs(fields["last_update"], fields["purchase_price"], productRecordID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	updatedProductRecord, err := s.repo.PartialUpdate(productRecordID, fields)

	// Assert
	s.NoError(err)
	s.Equal(productRecordID, updatedProductRecord.Id)
	s.Equal(fields["last_update"], updatedProductRecord.LastUpdate)
	s.Equal(fields["purchase_price"], updatedProductRecord.PurchasePrice)
}

func (s *ProductRecordRepositoryTestSuite) TestPartialUpdate_NotFound() {
	// Arrange
	productRecordID := 999
	fields := map[string]interface{}{
		"last_update":    "2022-22-10",
		"purchase_price": 4.99,
	}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `product_records` WHERE `product_records`.`id` = ? ORDER BY `product_records`.`id` LIMIT ?")).
		WithArgs(productRecordID, 1).WillReturnError(gorm.ErrRecordNotFound)

	// Act
	updatedProductRecord, err := s.repo.PartialUpdate(productRecordID, fields)

	// Assert
	s.Error(err)
	s.Equal(gorm.ErrRecordNotFound, err)
	s.Equal(models.ProductRecord{}, updatedProductRecord)
}

func (s *ProductRecordRepositoryTestSuite) TestPartialUpdate_DatabaseError() {
	// Arrange
	productRecordID := 1
	fields := map[string]interface{}{
		"last_update":    "2022-22-10",
		"purchase_price": 4.99,
	}

	expectedProductRecord := models.ProductRecord{
		Id:            productRecordID,
		LastUpdate:    "2022-22-10",
		PurchasePrice: 4.99,
		SalePrice:     5.99,
		ProductId:     1,
	}

	rows := s.mock.NewRows([]string{"id", "last_update", "purchase_price", "sale_price", "product_id"}).
		AddRow(expectedProductRecord.Id, expectedProductRecord.LastUpdate, expectedProductRecord.PurchasePrice, expectedProductRecord.SalePrice, expectedProductRecord.ProductId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `product_records` WHERE `product_records`.`id` = ? ORDER BY `product_records`.`id` LIMIT ?")).
		WithArgs(productRecordID, 1).WillReturnRows(rows)

	// Update query with database error
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `product_records` SET `last_update`=?,`purchase_price`=? WHERE `id` = ?")).
		WithArgs(fields["last_update"], fields["purchase_price"], productRecordID).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	updatedProductRecord, err := s.repo.PartialUpdate(productRecordID, fields)

	// Assert
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.ProductRecord{}, updatedProductRecord)
}
func (s *ProductRecordRepositoryTestSuite) TestDelete_Success() {
	// Arrange
	productRecordID := 1
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `product_records` WHERE `product_records`.`id` = ?")).
		WithArgs(productRecordID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	// Act
	err := s.repo.Delete(productRecordID)
	// Assert
	s.NoError(err)
}

func (s *ProductRecordRepositoryTestSuite) TestDelete_NotFound() {
	// Arrange
	productRecordID := 999

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `product_records` WHERE `product_records`.`id` = ?")).
		WithArgs(productRecordID).
		WillReturnResult(sqlmock.NewResult(1, 0)) // 0 rows affected
	s.mock.ExpectCommit()

	// Act
	err := s.repo.Delete(productRecordID)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
}

// Run the test suite
func TestProductRecordRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProductRecordRepositoryTestSuite))
}
