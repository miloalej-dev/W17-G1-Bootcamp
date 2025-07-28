package database

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type ProductBatchRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo *ProductBatchRepository
}

func (p *ProductBatchRepositoryTestSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, p.mock, err = sqlmock.New()
	if err != nil {
		p.T().Fatal(err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		p.T().Fatal(err)
	}

	p.db = gormDB
	p.repo = NewProductBatchRepository(gormDB)
}

func (p *ProductBatchRepositoryTestSuite) TestCreate_Success() {
	// Arrange
	newBatch := models.ProductBatch{
		BatchNumber:        40,
		CurrentQuantity:    200,
		CurrentTemperature: 20,
		DueDate:            "2022-04-04",
		InitialQuantity:    10,
		ManufacturingDate:  "2020-04-04",
		ManufacturingHour:  10,
		MinimumTemperature: 5,
		SectionId:          1,
		ProductId:          1,
	}

	p.mock.ExpectBegin()
	p.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `product_batches` (`batch_number`,`current_quantity`,`current_temperature`,`due_date`,`initial_quantity`,`manufacturing_date`,`manufacturing_hour`,`minimum_temperature`,`section_id`,`product_id`) VALUES (?,?,?,?,?,?,?,?,?,?)")).
		WithArgs(newBatch.BatchNumber, newBatch.CurrentQuantity, newBatch.CurrentTemperature, newBatch.DueDate, newBatch.InitialQuantity, newBatch.ManufacturingDate, newBatch.ManufacturingHour, newBatch.MinimumTemperature, newBatch.SectionId, newBatch.ProductId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	p.mock.ExpectCommit()

	// Act
	createdBatch, err := p.repo.Create(newBatch)

	// Assert
	p.NoError(err)
	p.Equal(newBatch.BatchNumber, createdBatch.BatchNumber)
	p.Equal(1, createdBatch.Id)
}

func (p *ProductBatchRepositoryTestSuite) TestCreate_ForeignKeyViolated() {
	// Arrange
	newBatch := models.ProductBatch{
		BatchNumber:        40,
		CurrentQuantity:    200,
		CurrentTemperature: 20,
		DueDate:            "2022-04-04",
		InitialQuantity:    10,
		ManufacturingDate:  "2020-04-04",
		ManufacturingHour:  10,
		MinimumTemperature: 5,
		SectionId:          1,
		ProductId:          1,
	}

	p.mock.ExpectBegin()
	p.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `product_batches` (`batch_number`,`current_quantity`,`current_temperature`,`due_date`,`initial_quantity`,`manufacturing_date`,`manufacturing_hour`,`minimum_temperature`,`section_id`,`product_id`) VALUES (?,?,?,?,?,?,?,?,?,?)")).
		WithArgs(newBatch.BatchNumber, newBatch.CurrentQuantity, newBatch.CurrentTemperature, newBatch.DueDate, newBatch.InitialQuantity, newBatch.ManufacturingDate, newBatch.ManufacturingHour, newBatch.MinimumTemperature, newBatch.SectionId, newBatch.ProductId).
		WillReturnError(gorm.ErrForeignKeyViolated)
	p.mock.ExpectRollback()

	// Act
	createdBatch, err := p.repo.Create(newBatch)

	// Assert
	p.Error(err)
	p.Equal(models.ProductBatch{}, createdBatch)
}

func (p *ProductBatchRepositoryTestSuite) TestCreate_GenericDataBaseError() {
	// Arrange
	newBatch := models.ProductBatch{
		BatchNumber:        40,
		CurrentQuantity:    200,
		CurrentTemperature: 20,
		DueDate:            "2022-04-04",
		InitialQuantity:    10,
		ManufacturingDate:  "2020-04-04",
		ManufacturingHour:  10,
		MinimumTemperature: 5,
		SectionId:          1,
		ProductId:          1,
	}

	p.mock.ExpectBegin()
	p.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `product_batches` (`batch_number`,`current_quantity`,`current_temperature`,`due_date`,`initial_quantity`,`manufacturing_date`,`manufacturing_hour`,`minimum_temperature`,`section_id`,`product_id`) VALUES (?,?,?,?,?,?,?,?,?,?)")).
		WithArgs(newBatch.BatchNumber, newBatch.CurrentQuantity, newBatch.CurrentTemperature, newBatch.DueDate, newBatch.InitialQuantity, newBatch.ManufacturingDate, newBatch.ManufacturingHour, newBatch.MinimumTemperature, newBatch.SectionId, newBatch.ProductId).
		WillReturnError(gorm.ErrInvalidValue)
	p.mock.ExpectRollback()

	// Act
	createdBatch, err := p.repo.Create(newBatch)

	// Assert
	p.Error(err)
	p.Equal(models.ProductBatch{}, createdBatch)
}
func (p *ProductBatchRepositoryTestSuite) TestFindAll_PanicsWhenCalled() {
	// Assert
	p.PanicsWithValue("method FindAll not implemented for ProductBatchRepository", func() {
		_, _ = p.repo.FindAll()
	})
}
func (p *ProductBatchRepositoryTestSuite) TestUpdate_PanicsWhenCalled() {
	// Assert
	p.PanicsWithValue("method Update not implemented for ProductBatchRepository", func() {
		_, _ = p.repo.Update(models.ProductBatch{})
	})
}
func (p *ProductBatchRepositoryTestSuite) TestFindById_PanicsWhenCalled() {
	// Assert
	p.PanicsWithValue("method FindById not implemented for ProductBatchRepository", func() {
		_, _ = p.repo.FindById(1)
	})
}
func (p *ProductBatchRepositoryTestSuite) TestPartialUpdate_PanicsWhenCalled() {
	// Assert
	p.PanicsWithValue("method PartialUpdate not implemented for ProductBatchRepository", func() {
		_, _ = p.repo.PartialUpdate(1, map[string]interface{}{})
	})
}
func (p *ProductBatchRepositoryTestSuite) TestDelete_PanicsWhenCalled() {
	// Assert
	p.PanicsWithValue("method Delete not implemented for ProductBatchRepository", func() {
		_ = p.repo.Delete(1)
	})
}

// Run the test suite
func TestProductBatchRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProductBatchRepositoryTestSuite))
}
