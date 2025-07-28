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

type ProductRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo *ProductRepository
}

func (p *ProductRepositoryTestSuite) SetupSuite() {
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
	p.repo = NewProductRepository(gormDB)
}

func (p *ProductRepositoryTestSuite) TestFindAll_Success() {
	// Arrange
	expectedProducts := []models.Product{
		{
			Id:                             1,
			ProductCode:                    "mock1",
			Description:                    "mock1",
			Width:                          68.64,
			Height:                         185.04,
			Length:                         185.62,
			NetWeight:                      2.93,
			ExpirationRate:                 8.62,
			RecommendedFreezingTemperature: -33.84,
			FreezingRate:                   0.62,
			ProductTypeId:                  6,
		},
		{
			Id:                             2,
			ProductCode:                    "mock2",
			Description:                    "mock2",
			Width:                          68.64,
			Height:                         185.04,
			Length:                         185.62,
			NetWeight:                      2.93,
			ExpirationRate:                 8.62,
			RecommendedFreezingTemperature: -33.84,
			FreezingRate:                   0.62,
			ProductTypeId:                  6,
		},
	}
	rows := p.mock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id"}).
		AddRow(expectedProducts[0].Id, expectedProducts[0].ProductCode, expectedProducts[0].Description,
			expectedProducts[0].Width, expectedProducts[0].Height, expectedProducts[0].Length, expectedProducts[0].
				NetWeight, expectedProducts[0].ExpirationRate, expectedProducts[0].RecommendedFreezingTemperature,
			expectedProducts[0].FreezingRate, expectedProducts[0].ProductTypeId).
		AddRow(expectedProducts[1].Id, expectedProducts[1].ProductCode, expectedProducts[1].Description,
			expectedProducts[1].Width, expectedProducts[1].Height, expectedProducts[1].Length, expectedProducts[1].
				NetWeight, expectedProducts[1].ExpirationRate, expectedProducts[1].RecommendedFreezingTemperature,
			expectedProducts[1].FreezingRate, expectedProducts[1].ProductTypeId)

	p.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products`")).WillReturnRows(rows)

	// Act
	products, err := p.repo.FindAll()

	// Assert
	p.NoError(err)
	p.Len(products, 2)
	p.Equal(expectedProducts[0], products[0])
	p.Equal(expectedProducts[1], products[1])
}

func (p *ProductRepositoryTestSuite) TestFindAll_DatabaseError() {
	// Arrange
	p.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products`")).WillReturnError(repository.ErrEmptyEntity)

	// Act
	products, err := p.repo.FindAll()

	// Assert
	p.Error(err)
	p.Nil(products)
	p.ErrorIs(repository.ErrEmptyEntity, err)
}

func (p *ProductRepositoryTestSuite) TestFindById_Success() {
	// Arrange
	expectedProducts := []models.Product{
		{
			Id:                             1,
			ProductCode:                    "mock1",
			Description:                    "mock1",
			Width:                          68.64,
			Height:                         185.04,
			Length:                         185.62,
			NetWeight:                      2.93,
			ExpirationRate:                 8.62,
			RecommendedFreezingTemperature: -33.84,
			FreezingRate:                   0.62,
			ProductTypeId:                  6,
		},
	}
	rows := p.mock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id"}).
		AddRow(expectedProducts[0].Id, expectedProducts[0].ProductCode, expectedProducts[0].Description,
			expectedProducts[0].Width, expectedProducts[0].Height, expectedProducts[0].Length, expectedProducts[0].
				NetWeight, expectedProducts[0].ExpirationRate, expectedProducts[0].RecommendedFreezingTemperature,
			expectedProducts[0].FreezingRate, expectedProducts[0].ProductTypeId)

	p.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products` WHERE `products`.`id` = ? ORDER BY `products`.`id` LIMIT ?")).
		WithArgs(1, 1).WillReturnRows(rows)

	// Act
	product, err := p.repo.FindById(1)

	// Assert
	p.NoError(err)
	p.Equal(expectedProducts[0].Id, product.Id)
	p.Equal(expectedProducts[0].ProductCode, product.ProductCode)

}

func (p *ProductRepositoryTestSuite) TestFindById_DataBaseError() {
	// Arrange

	p.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products` WHERE `products`.`id` = ? ORDER BY `products`.`id` LIMIT ?")).
		WithArgs(99, 1).WillReturnError(gorm.ErrRecordNotFound)

	// Act
	product, err := p.repo.FindById(99)

	// Assert
	p.Error(err)
	p.Equal(models.Product{}, product)
	p.ErrorIs(repository.ErrProductNotFound, err)
}

func (p *ProductRepositoryTestSuite) TestFindById_GenericDatabaseError() {
	// Arrange
	p.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products` WHERE `products`.`id` = ? ORDER BY `products`.`id` LIMIT ?")).
		WithArgs(99, 1).
		WillReturnError(sql.ErrConnDone)

	// Act
	product, err := p.repo.FindById(99)

	// Assert
	p.Error(err)
	p.Equal(models.Product{}, product)

	p.ErrorIs(sql.ErrConnDone, err)
}

func (p *ProductRepositoryTestSuite) TestCreate_Success() {
	// Arrange
	newProduct := []models.Product{
		{
			ProductCode:                    "mock1",
			Description:                    "mock1",
			Width:                          68.64,
			Height:                         185.04,
			Length:                         185.62,
			NetWeight:                      2.93,
			ExpirationRate:                 8.62,
			RecommendedFreezingTemperature: -33.84,
			FreezingRate:                   0.62,
			ProductTypeId:                  6,
		},
	}

	p.mock.ExpectBegin()
	p.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `products` (`product_code`,`description`,`width`,`height`,`length`,`net_weight`,`expiration_rate`,`recommended_freezing_temperature`,`freezing_rate`,`product_type_id`,`seller_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).
		WithArgs(newProduct[0].ProductCode, newProduct[0].Description,
			newProduct[0].Width, newProduct[0].Height, newProduct[0].Length, newProduct[0].
				NetWeight, newProduct[0].ExpirationRate, newProduct[0].RecommendedFreezingTemperature,
			newProduct[0].FreezingRate, newProduct[0].ProductTypeId, newProduct[0].SellerId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	p.mock.ExpectCommit()

	// Act
	createdProduct, err := p.repo.Create(newProduct[0])

	// Assert
	p.NoError(err)
	p.Equal(newProduct[0].ProductCode, createdProduct.ProductCode)
	p.Equal(1, createdProduct.Id)
}

func (p *ProductRepositoryTestSuite) TestCreate_GenericDatabaseError() {
	// Arrange
	newProduct := models.Product{
		ProductCode:                    "mock1",
		Description:                    "mock1",
		Width:                          68.64,
		Height:                         185.04,
		Length:                         185.62,
		NetWeight:                      2.93,
		ExpirationRate:                 8.62,
		RecommendedFreezingTemperature: -33.84,
		FreezingRate:                   0.62,
		ProductTypeId:                  6,
	}
	p.mock.ExpectBegin()
	p.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `products` (`product_code`,`description`,`width`,`height`,`length`,`net_weight`,`expiration_rate`,`recommended_freezing_temperature`,`freezing_rate`,`product_type_id`,`seller_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).
		WithArgs(newProduct.ProductCode, newProduct.Description,
			newProduct.Width, newProduct.Height, newProduct.Length, newProduct.
				NetWeight, newProduct.ExpirationRate, newProduct.RecommendedFreezingTemperature,
			newProduct.FreezingRate, newProduct.ProductTypeId, newProduct.SellerId).
		WillReturnError(sql.ErrConnDone)
	p.mock.ExpectRollback()

	// Act
	createdProduct, err := p.repo.Create(newProduct)

	// Assert
	p.Error(err)
	p.Equal(models.Product{}, createdProduct)
	p.ErrorIs(sql.ErrConnDone, err)
}

func (p *ProductRepositoryTestSuite) TestCreate_DatabaseError() {
	// Arrange
	newProduct := models.Product{
		ProductCode:                    "mock1",
		Description:                    "mock1",
		Width:                          68.64,
		Height:                         185.04,
		Length:                         185.62,
		NetWeight:                      2.93,
		ExpirationRate:                 8.62,
		RecommendedFreezingTemperature: -33.84,
		FreezingRate:                   0.62,
		ProductTypeId:                  6,
	}
	p.mock.ExpectBegin()
	p.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `products` (`product_code`,`description`,`width`,`height`,`length`,`net_weight`,`expiration_rate`,`recommended_freezing_temperature`,`freezing_rate`,`product_type_id`,`seller_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).
		WithArgs(newProduct.ProductCode, newProduct.Description,
			newProduct.Width, newProduct.Height, newProduct.Length, newProduct.
				NetWeight, newProduct.ExpirationRate, newProduct.RecommendedFreezingTemperature,
			newProduct.FreezingRate, newProduct.ProductTypeId, newProduct.SellerId).
		WillReturnError(gorm.ErrForeignKeyViolated)
	p.mock.ExpectRollback()

	// Act
	createdProduct, err := p.repo.Create(newProduct)

	// Assert
	p.Error(err)
	p.Equal(models.Product{}, createdProduct)
	p.ErrorIs(repository.ErrForeignKeyViolation, err)
}

func (p *ProductRepositoryTestSuite) TestUpdate_Success() {
	// Arrange
	update := models.Product{
		Id:                             2,
		ProductCode:                    "updatedMock1",
		Description:                    "updatedMock1",
		Width:                          68.64,
		Height:                         185.04,
		Length:                         185.62,
		NetWeight:                      2.93,
		ExpirationRate:                 8.62,
		RecommendedFreezingTemperature: -33.84,
		FreezingRate:                   0.62,
		ProductTypeId:                  6,
	}

	p.mock.ExpectBegin()
	p.mock.ExpectExec(regexp.QuoteMeta("UPDATE `products` SET `product_code`=?,`description`=?,`width`=?,`height`=?,`length`=?,`net_weight`=?,`expiration_rate`=?,`recommended_freezing_temperature`=?,`freezing_rate`=?,`product_type_id`=?,`seller_id`=? WHERE `id` = ?")).
		WithArgs(update.ProductCode, update.Description,
			update.Width, update.Height, update.Length, update.NetWeight, update.ExpirationRate, update.RecommendedFreezingTemperature,
			update.FreezingRate, update.ProductTypeId, update.SellerId, update.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	p.mock.ExpectCommit()

	// Act
	updatedProduct, err := p.repo.Update(update)

	// Assert
	p.NoError(err)
	p.Equal(update.ProductCode, updatedProduct.ProductCode)
	p.Equal(update.Id, updatedProduct.Id)
}

func (p *ProductRepositoryTestSuite) TestUpdate_DatabaseError() {
	// Arrange
	update := models.Product{
		Id:                             2,
		ProductCode:                    "updatedMock1",
		Description:                    "updatedMock1",
		Width:                          68.64,
		Height:                         185.04,
		Length:                         185.62,
		NetWeight:                      2.93,
		ExpirationRate:                 8.62,
		RecommendedFreezingTemperature: -33.84,
		FreezingRate:                   0.62,
		ProductTypeId:                  6,
	}

	p.mock.ExpectBegin()
	p.mock.ExpectExec(regexp.QuoteMeta("UPDATE `products` SET `product_code`=?,`description`=?,`width`=?,`height`=?,`length`=?,`net_weight`=?,`expiration_rate`=?,`recommended_freezing_temperature`=?,`freezing_rate`=?,`product_type_id`=?,`seller_id`=? WHERE `id` = ?")).
		WithArgs(update.ProductCode, update.Description,
			update.Width, update.Height, update.Length, update.NetWeight, update.ExpirationRate, update.RecommendedFreezingTemperature,
			update.FreezingRate, update.ProductTypeId, update.SellerId, update.Id).
		WillReturnError(gorm.ErrForeignKeyViolated)
	p.mock.ExpectRollback()

	// Act
	createdProduct, err := p.repo.Update(update)

	// Assert
	p.Error(err)
	p.Equal(models.Product{}, createdProduct)
	p.ErrorIs(repository.ErrForeignKeyViolation, err)
}

func (p *ProductRepositoryTestSuite) TestUpdate_GenericDatabaseError() {
	// Arrange
	update := models.Product{
		Id:                             2,
		ProductCode:                    "updatedMock1",
		Description:                    "updatedMock1",
		Width:                          68.64,
		Height:                         185.04,
		Length:                         185.62,
		NetWeight:                      2.93,
		ExpirationRate:                 8.62,
		RecommendedFreezingTemperature: -33.84,
		FreezingRate:                   0.62,
		ProductTypeId:                  6,
	}

	p.mock.ExpectBegin()
	p.mock.ExpectExec(regexp.QuoteMeta("UPDATE `products` SET `product_code`=?,`description`=?,`width`=?,`height`=?,`length`=?,`net_weight`=?,`expiration_rate`=?,`recommended_freezing_temperature`=?,`freezing_rate`=?,`product_type_id`=?,`seller_id`=? WHERE `id` = ?")).
		WithArgs(update.ProductCode, update.Description,
			update.Width, update.Height, update.Length, update.NetWeight, update.ExpirationRate, update.RecommendedFreezingTemperature,
			update.FreezingRate, update.ProductTypeId, update.SellerId, update.Id).
		WillReturnError(sql.ErrConnDone)
	p.mock.ExpectRollback()

	// Act
	createdProduct, err := p.repo.Update(update)

	// Assert
	p.Error(err)
	p.Equal(models.Product{}, createdProduct)
	p.ErrorIs(sql.ErrConnDone, err)
}

func (p *ProductRepositoryTestSuite) TestPartialUpdate_Success() {
	// Arrange
	productId := 2
	fields := map[string]interface{}{
		"product_code": "updatedMock1",
		"description":  "updatedMock1",
	}

	expectedProduct := models.Product{
		Id:                             2,
		ProductCode:                    "oldMock1",
		Description:                    "oldMock1",
		Width:                          68.64,
		Height:                         185.04,
		Length:                         185.62,
		NetWeight:                      2.93,
		ExpirationRate:                 8.62,
		RecommendedFreezingTemperature: -33.84,
		FreezingRate:                   0.62,
		ProductTypeId:                  6,
	}

	// First query to find the seller
	rows := p.mock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id"}).
		AddRow(expectedProduct.Id, expectedProduct.ProductCode, expectedProduct.Description,
			expectedProduct.Width, expectedProduct.Height, expectedProduct.Length, expectedProduct.
				NetWeight, expectedProduct.ExpirationRate, expectedProduct.RecommendedFreezingTemperature,
			expectedProduct.FreezingRate, expectedProduct.ProductTypeId)
	p.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products` WHERE `products`.`id` = ? ORDER BY `products`.`id` LIMIT ?")).
		WithArgs(productId, 1).WillReturnRows(rows)

	// Update query
	p.mock.ExpectBegin()
	p.mock.ExpectExec(regexp.QuoteMeta("UPDATE `products` SET `description`=?,`product_code`=? WHERE `id` = ?")).
		WithArgs(fields["description"], fields["product_code"], productId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	p.mock.ExpectCommit()

	// Act
	updatedProduct, err := p.repo.PartialUpdate(productId, fields)

	// Assert
	p.NoError(err)
	p.Equal(productId, updatedProduct.Id)
}

func (p *ProductRepositoryTestSuite) TestPartialUpdate_NotFound() {
	// Arrange
	productID := 999
	fields := map[string]interface{}{}

	p.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products` WHERE `products`.`id` = ? ORDER BY `products`.`id` LIMIT ?")).
		WithArgs(productID, 1).WillReturnError(gorm.ErrRecordNotFound)

	// Act
	updatedProduct, err := p.repo.PartialUpdate(productID, fields)

	// Assert
	p.Error(err)
	p.Equal(repository.ErrProductNotFound, err)
	p.Equal(models.Product{}, updatedProduct)
}

func (p *ProductRepositoryTestSuite) TestPartialUpdate_DatabaseError() {
	// Arrange
	productId := 2
	fields := map[string]interface{}{
		"product_code": "updatedMock1",
	}

	expectedProduct := models.Product{
		Id:                             2,
		ProductCode:                    "oldMock1",
		Description:                    "oldMock1",
		Width:                          68.64,
		Height:                         185.04,
		Length:                         185.62,
		NetWeight:                      2.93,
		ExpirationRate:                 8.62,
		RecommendedFreezingTemperature: -33.84,
		FreezingRate:                   0.62,
		ProductTypeId:                  6,
	}

	// First query to find the seller
	rows := p.mock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id"}).
		AddRow(expectedProduct.Id, expectedProduct.ProductCode, expectedProduct.Description,
			expectedProduct.Width, expectedProduct.Height, expectedProduct.Length, expectedProduct.
				NetWeight, expectedProduct.ExpirationRate, expectedProduct.RecommendedFreezingTemperature,
			expectedProduct.FreezingRate, expectedProduct.ProductTypeId)
	p.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products` WHERE `products`.`id` = ? ORDER BY `products`.`id` LIMIT ?")).
		WithArgs(productId, 1).WillReturnRows(rows)

	// Update query with database error
	p.mock.ExpectBegin()
	p.mock.ExpectExec(regexp.QuoteMeta("UPDATE `products` SET `product_code`=? WHERE `id` = ?")).
		WithArgs(fields["product_code"], productId).
		WillReturnError(sql.ErrConnDone)
	p.mock.ExpectRollback()

	// Act
	updatedProduct, err := p.repo.PartialUpdate(productId, fields)

	// Assert
	p.Error(err)
	p.Equal(sql.ErrConnDone, err)
	p.Equal(models.Product{}, updatedProduct)
}

func (p *ProductRepositoryTestSuite) TestDelete_Success() {
	// Arrange
	productID := 1
	p.mock.ExpectBegin()
	p.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `products` WHERE `products`.`id` = ?")).
		WithArgs(productID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	p.mock.ExpectCommit()
	// Act
	err := p.repo.Delete(productID)
	// Assert
	p.NoError(err)
}

func (p *ProductRepositoryTestSuite) TestDelete_NotFound() {
	// Arrange
	productID := 999

	p.mock.ExpectBegin()
	p.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `products` WHERE `products`.`id` = ?")).
		WithArgs(productID).
		WillReturnResult(sqlmock.NewResult(1, 0)) // 0 rows affected
	p.mock.ExpectCommit()

	// Act
	err := p.repo.Delete(productID)

	// Assert
	p.Error(err)
	p.Equal(repository.ErrProductNotFound, err)
}

func (s *ProductRepositoryTestSuite) TestFindRecordsCountByProductId_Success() {
	id := 1

	expected := models.ProductReport{
		Id:           id,
		Description:  " generic description",
		RecordsCount: 2,
	}

	rows := s.mock.NewRows([]string{"id", "description", "records_count"}).
		AddRow(expected.Id, expected.Description, expected.RecordsCount)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT products.id, products.description, COUNT(product_records.id) as records_count FROM `products` inner join product_records on product_records.product_id = products.id WHERE products.id = ? GROUP BY `products`.`id`")).
		WithArgs(id).
		WillReturnRows(rows)

	report, err := s.repo.FindRecordsCountByProductId(id)
	s.NoError(err)
	s.Equal(expected, report)
}

func (s *ProductRepositoryTestSuite) TestFindRecordsCountByProductId_NotFound() {
	id := 99

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT products.id, products.description, COUNT(product_records.id) as records_count FROM `products` inner join product_records on product_records.product_id = products.id WHERE products.id = ? GROUP BY `products`.`id`")).
		WithArgs(id).
		WillReturnError(gorm.ErrRecordNotFound)

	report, err := s.repo.FindRecordsCountByProductId(id)

	s.Error(err)
	s.Equal(repository.ErrProductReportNotFound, err)
	s.Equal(models.ProductReport{}, report)
}

func (s *ProductRepositoryTestSuite) TestFindRecordsCount_Success() {

	expected := []models.ProductReport{
		{
			Id:           1,
			Description:  "generic description",
			RecordsCount: 2,
		},
		{
			Id:           2,
			Description:  "another description",
			RecordsCount: 4,
		},
	}

	rows := s.mock.NewRows([]string{"id", "description", "records_count"}).
		AddRow(expected[0].Id, expected[0].Description, expected[0].RecordsCount).
		AddRow(expected[1].Id, expected[1].Description, expected[1].RecordsCount)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT products.id, products.description, COUNT(product_records.id) as records_count FROM `products` inner join product_records on product_records.products_id = products.id GROUP BY `products`.`id`")).
		WillReturnRows(rows)

	report, err := s.repo.FindRecordsCount()

	s.NoError(err)
	s.Equal(expected, report)
}

func (s *ProductRepositoryTestSuite) TestFindRecordsCount_NotFound() {

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT products.id, products.description, COUNT(product_records.id) as records_count FROM `products` inner join product_records on product_records.products_id = products.id GROUP BY `products`.`id`")).
		WillReturnError(gorm.ErrRecordNotFound)

	report, err := s.repo.FindRecordsCount()

	s.Error(err)
	s.Equal(repository.ErrProductReportNotFound, err)
	s.Equal([]models.ProductReport{}, report)
}

// Run the test suite
func TestProductRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProductRepositoryTestSuite))
}
