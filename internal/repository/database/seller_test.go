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

type SellerRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo *SellerRepository
}

func (s *SellerRepositoryTestSuite) SetupSuite() {
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
	s.repo = NewSellerRepository(gormDB)
}

func (s *SellerRepositoryTestSuite) TestFindAll_Success() {
	// Arrange
	expectedSellers := []models.Seller{
		{
			Id:         1,
			Name:       "Company A",
			Address:    "123 Main St",
			Telephone:  "555-0001",
			LocalityId: 1,
		},
		{
			Id:         2,
			Name:       "Company B",
			Address:    "456 Oak Ave",
			Telephone:  "555-0002",
			LocalityId: 2,
		},
	}

	rows := s.mock.NewRows([]string{"id", "name", "address", "telephone", "locality_id"}).
		AddRow(expectedSellers[0].Id, expectedSellers[0].Name, expectedSellers[0].Address,
			expectedSellers[0].Telephone, expectedSellers[0].LocalityId).
		AddRow(expectedSellers[1].Id, expectedSellers[1].Name, expectedSellers[1].Address,
			expectedSellers[1].Telephone, expectedSellers[1].LocalityId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `sellers`")).WillReturnRows(rows)

	// Act
	sellers, err := s.repo.FindAll()

	// Assert
	s.NoError(err)
	s.Len(sellers, 2)
	s.Equal(expectedSellers[0].Name, sellers[0].Name)
	s.Equal(expectedSellers[1].Name, sellers[1].Name)
}

func (s *SellerRepositoryTestSuite) TestFindAll_DatabaseError() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `sellers`")).WillReturnError(sql.ErrConnDone)

	// Act
	sellers, err := s.repo.FindAll()

	// Assert
	s.Error(err)
	s.Nil(sellers)
	s.Equal(sql.ErrConnDone, err)
}

func (s *SellerRepositoryTestSuite) TestFindById_Success() {
	// Arrange
	expectedSeller := models.Seller{
		Id:         1,
		Name:       "Company A",
		Address:    "123 Main St",
		Telephone:  "555-0001",
		LocalityId: 1,
	}

	rows := s.mock.NewRows([]string{"id", "name", "address", "telephone", "locality_id"}).
		AddRow(expectedSeller.Id, expectedSeller.Name, expectedSeller.Address,
			expectedSeller.Telephone, expectedSeller.LocalityId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `sellers` WHERE `sellers`.`id` = ? ORDER BY `sellers`.`id` LIMIT ?")).
		WithArgs(1, 1).WillReturnRows(rows)

	// Act
	seller, err := s.repo.FindById(1)

	// Assert
	s.NoError(err)
	s.Equal(expectedSeller.Id, seller.Id)
	s.Equal(expectedSeller.Name, seller.Name)
}

func (s *SellerRepositoryTestSuite) TestFindById_NotFound() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `sellers` WHERE `sellers`.`id` = ? ORDER BY `sellers`.`id` LIMIT ?")).
		WithArgs(999, 1).WillReturnError(repository.ErrEntityNotFound)

	// Act
	seller, err := s.repo.FindById(999)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal(models.Seller{}, seller)
}

func (s *SellerRepositoryTestSuite) TestFindById_DatabaseError() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `sellers` WHERE `sellers`.`id` = ? ORDER BY `sellers`.`id` LIMIT ?")).
		WithArgs(1, 1).WillReturnError(sql.ErrConnDone)

	// Act
	seller, err := s.repo.FindById(1)

	// Assert
	s.Error(err)
	s.Equal(models.Seller{}, seller)
	s.Equal(sql.ErrConnDone, err)
}

func (s *SellerRepositoryTestSuite) TestCreate_Success() {
	// Arrange
	newSeller := models.Seller{
		Name:       "New Company",
		Address:    "789 Pine St",
		Telephone:  "555-0003",
		LocalityId: 3,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `sellers` (`name`,`address`,`telephone`,`locality_id`) VALUES (?,?,?,?)")).
		WithArgs(newSeller.Name, newSeller.Address, newSeller.Telephone, newSeller.LocalityId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	createdSeller, err := s.repo.Create(newSeller)

	// Assert
	s.NoError(err)
	s.Equal(newSeller.Name, createdSeller.Name)
	s.Equal(1, createdSeller.Id)
}

func (s *SellerRepositoryTestSuite) TestCreate_ForeignKeyViolated() {
	// Arrange
	newSeller := models.Seller{
		Name:       "Duplicate Company",
		Address:    "789 Pine St",
		Telephone:  "555-0003",
		LocalityId: 3,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `sellers` (`name`,`address`,`telephone`,`locality_id`) VALUES (?,?,?,?)")).
		WithArgs(newSeller.Name, newSeller.Address, newSeller.Telephone, newSeller.LocalityId).
		WillReturnError(repository.ErrForeignKeyViolation)
	s.mock.ExpectRollback()

	// Act
	createdSeller, err := s.repo.Create(newSeller)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrForeignKeyViolation, err)
	s.Equal(models.Seller{}, createdSeller)
}

func (s *SellerRepositoryTestSuite) TestCreate_DatabaseError() {
	// Arrange
	newSeller := models.Seller{
		Name:       "New Company",
		Address:    "789 Pine St",
		Telephone:  "555-0003",
		LocalityId: 3,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `sellers` (`name`,`address`,`telephone`,`locality_id`) VALUES (?,?,?,?)")).
		WithArgs(newSeller.Name, newSeller.Address, newSeller.Telephone, newSeller.LocalityId).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	createdSeller, err := s.repo.Create(newSeller)

	// Assert
	s.Error(err)
	s.Equal(models.Seller{}, createdSeller)
	s.Equal(sql.ErrConnDone, err)
}

func (s *SellerRepositoryTestSuite) TestUpdate_Success() {
	// Arrange
	existingSeller := models.Seller{
		Id:         1,
		Name:       "Updated Company",
		Address:    "Updated Address",
		Telephone:  "555-9999",
		LocalityId: 1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `sellers` SET `name`=?,`address`=?,`telephone`=?,`locality_id`=? WHERE `id` = ?")).
		WithArgs(existingSeller.Name, existingSeller.Address, existingSeller.Telephone, existingSeller.LocalityId, existingSeller.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	updatedSeller, err := s.repo.Update(existingSeller)

	// Assert
	s.NoError(err)
	s.Equal(existingSeller.Name, updatedSeller.Name)
	s.Equal(existingSeller.Id, updatedSeller.Id)
}

func (s *SellerRepositoryTestSuite) TestUpdate_ForeignKeyViolated() {
	// Arrange
	existingSeller := models.Seller{
		Id:         1,
		Name:       "Updated Company",
		Address:    "Updated Address",
		Telephone:  "555-9999",
		LocalityId: 999, // Invalid locality ID
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `sellers` SET `name`=?,`address`=?,`telephone`=?,`locality_id`=? WHERE `id` = ?")).
		WithArgs(existingSeller.Name, existingSeller.Address, existingSeller.Telephone, existingSeller.LocalityId, existingSeller.Id).
		WillReturnError(repository.ErrForeignKeyViolation)
	s.mock.ExpectRollback()

	// Act
	updatedSeller, err := s.repo.Update(existingSeller)

	// Assert
	s.Error(err)
	s.Equal(models.Seller{}, updatedSeller)
	s.Equal(repository.ErrForeignKeyViolation, err)
}

func (s *SellerRepositoryTestSuite) TestUpdate_DatabaseError() {
	// Arrange
	existingSeller := models.Seller{
		Id:         1,
		Name:       "Updated Company",
		Address:    "Updated Address",
		Telephone:  "555-9999",
		LocalityId: 1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `sellers` SET `name`=?,`address`=?,`telephone`=?,`locality_id`=? WHERE `id` = ?")).
		WithArgs(existingSeller.Name, existingSeller.Address, existingSeller.Telephone, existingSeller.LocalityId, existingSeller.Id).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	updatedSeller, err := s.repo.Update(existingSeller)

	// Assert
	s.Error(err)
	s.Equal(models.Seller{}, updatedSeller)
	s.Equal(sql.ErrConnDone, err)
}

func (s *SellerRepositoryTestSuite) TestPartialUpdate_Success() {
	// Arrange
	sellerID := 1
	fields := map[string]interface{}{
		"name":    "Partially Updated Company",
		"address": "New Address",
	}

	expectedSeller := models.Seller{
		Id:         sellerID,
		Name:       "Partially Updated Company",
		Address:    "New Address",
		Telephone:  "555-0001",
		LocalityId: 1,
	}

	// First query to find the seller
	rows := s.mock.NewRows([]string{"id", "name", "address", "telephone", "locality_id"}).
		AddRow(expectedSeller.Id, "Old Company", "Old Address", expectedSeller.Telephone, expectedSeller.LocalityId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `sellers` WHERE `sellers`.`id` = ? ORDER BY `sellers`.`id` LIMIT ?")).
		WithArgs(sellerID, 1).WillReturnRows(rows)

	// Update query
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `sellers` SET `address`=?,`name`=? WHERE `id` = ?")).
		WithArgs(fields["address"], fields["name"], sellerID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	updatedSeller, err := s.repo.PartialUpdate(sellerID, fields)

	// Assert
	s.NoError(err)
	s.Equal(sellerID, updatedSeller.Id)
}

func (s *SellerRepositoryTestSuite) TestPartialUpdate_NotFound() {
	// Arrange
	sellerID := 999
	fields := map[string]interface{}{
		"name": "Updated Name",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `sellers` WHERE `sellers`.`id` = ? ORDER BY `sellers`.`id` LIMIT ?")).
		WithArgs(sellerID, 1).WillReturnError(gorm.ErrRecordNotFound)

	// Act
	updatedSeller, err := s.repo.PartialUpdate(sellerID, fields)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal(models.Seller{}, updatedSeller)
}

func (s *SellerRepositoryTestSuite) TestPartialUpdate_ForeignKeyViolated() {
	// Arrange
	sellerID := 1
	fields := map[string]interface{}{
		"locality_id": 999, // Invalid locality ID
	}

	expectedSeller := models.Seller{
		Id:         sellerID,
		Name:       "Company A",
		Address:    "123 Main St",
		Telephone:  "555-0001",
		LocalityId: 1,
	}

	// First query to find the seller
	rows := s.mock.NewRows([]string{"id", "name", "address", "telephone", "locality_id"}).
		AddRow(expectedSeller.Id, expectedSeller.Name, expectedSeller.Address, expectedSeller.Telephone, expectedSeller.LocalityId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `sellers` WHERE `sellers`.`id` = ? ORDER BY `sellers`.`id` LIMIT ?")).
		WithArgs(sellerID, 1).WillReturnRows(rows)

	// Update query with foreign key violation
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `sellers` SET `locality_id`=? WHERE `id` = ?")).
		WithArgs(fields["locality_id"], sellerID).
		WillReturnError(repository.ErrForeignKeyViolation)
	s.mock.ExpectRollback()

	// Act
	updatedSeller, err := s.repo.PartialUpdate(sellerID, fields)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrForeignKeyViolation, err)
	s.Equal(models.Seller{}, updatedSeller)
}

func (s *SellerRepositoryTestSuite) TestPartialUpdate_DatabaseError() {
	// Arrange
	sellerID := 1
	fields := map[string]interface{}{
		"name": "Updated Name",
	}

	expectedSeller := models.Seller{
		Id:         sellerID,
		Name:       "Company A",
		Address:    "123 Main St",
		Telephone:  "555-0001",
		LocalityId: 1,
	}

	// First query to find the seller
	rows := s.mock.NewRows([]string{"id", "name", "address", "telephone", "locality_id"}).
		AddRow(expectedSeller.Id, expectedSeller.Name, expectedSeller.Address, expectedSeller.Telephone, expectedSeller.LocalityId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `sellers` WHERE `sellers`.`id` = ? ORDER BY `sellers`.`id` LIMIT ?")).
		WithArgs(sellerID, 1).WillReturnRows(rows)

	// Update query with database error
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `sellers` SET `name`=? WHERE `id` = ?")).
		WithArgs(fields["name"], sellerID).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	updatedSeller, err := s.repo.PartialUpdate(sellerID, fields)

	// Assert
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.Seller{}, updatedSeller)
}
func (s *SellerRepositoryTestSuite) TestDelete_Success() {
	// Arrange
	sellerID := 1
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `sellers` WHERE `sellers`.`id` = ?")).
		WithArgs(sellerID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	// Act
	err := s.repo.Delete(sellerID)
	// Assert
	s.NoError(err)
}

func (s *SellerRepositoryTestSuite) TestDelete_NotFound() {
	// Arrange
	sellerID := 999

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `sellers` WHERE `sellers`.`id` = ?")).
		WithArgs(sellerID).
		WillReturnResult(sqlmock.NewResult(1, 0)) // 0 rows affected
	s.mock.ExpectCommit()

	// Act
	err := s.repo.Delete(sellerID)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
}

// Run the test suite
func TestSellerRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(SellerRepositoryTestSuite))
}
