package database

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"regexp"
	"testing"
)

type CarrierRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo *CarrierDB
}

func (s *CarrierRepositoryTestSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	if err != nil {
		s.T().Fatal(err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:						db,
		SkipInitializeWithVersion:	true,
	}), &gorm.Config{
		TranslateError: true,
		Logger: logger.Discard,
	})

	if err != nil {
		s.T().Fatal(err)
	}

	s.db = gormDB
	s.repo = NewCarrierDB(gormDB)
}

// Test for the find all success
func (s *CarrierRepositoryTestSuite) TestFindAll_Success() {
	// Arrange
	expectedCarriers := []models.Carrier{
		{
			ID:                 1,
			CId:				"CID#01",
			CompanyName:        "Meli",
			Address:            "Boulevard",
			Telephone:          "123-456789",
			LocalityId:         1,
		},
		{
			ID:                 1,
			CId:				"CID#01",
			CompanyName:        "Meli",
			Address:            "Boulevard",
			Telephone:          "123-456789",
			LocalityId:         1,
		},
	}

	columns := []string{
		"id",
		"cid",
		"name",
		"address",
		"telephone",
		"locality_id",
	}

	rows := s.mock.NewRows(columns)
	for _, w := range expectedCarriers {
		rows = rows.AddRow(
			w.ID,
			w.CId,
			w.CompanyName,
			w.Address,
			w.Telephone,
			w.LocalityId,
		)
	}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `carriers`")).WillReturnRows(rows)

	// Act
	carriers, err := s.repo.FindAll()

	// Assert
	s.NoError(err)
	s.Len(carriers, 2)
	s.Equal(expectedCarriers, carriers)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

// Error on retireve all
func (s *CarrierRepositoryTestSuite) TestFindAll_DatabaseError() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `carriers`")).WillReturnError(sql.ErrConnDone)

	// Act
	carriers, err := s.repo.FindAll()

	// Assert
	s.Error(err)
	s.Nil(carriers)
	s.Equal(sql.ErrConnDone, err)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

// Find by Id returns the specified carrier
func (s *CarrierRepositoryTestSuite) TestFindById_Success() {
	// Arrange
	expectedCarrier := models.Carrier{
		ID:                 1,
		CId:				"CID#01",
		CompanyName:        "Meli",
		Address:            "Boulevard",
		Telephone:          "123-456789",
		LocalityId:         1,
	}

	columns := []string{
		"id",
		"cid",
		"name",
		"address",
		"telephone",
		"locality_id",
	}

	rows := s.mock.NewRows(columns).
		AddRow(
			expectedCarrier.ID,
			expectedCarrier.CId,
			expectedCarrier.CompanyName,
			expectedCarrier.Address,
			expectedCarrier.Telephone,
			expectedCarrier.LocalityId,
		)


	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `carriers` WHERE `carriers`.`id` = ? ORDER BY `carriers`.`id` LIMIT ?",
	)).WithArgs(1, 1).WillReturnRows(rows)

	// Act
	carrier, err := s.repo.FindById(1)

	// Assert
	s.NoError(err)
	s.Equal(expectedCarrier, carrier)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

// Find by a non existent Id returns nothing
func (s *CarrierRepositoryTestSuite) TestFindById_NotFound() {
	id := 999

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `carriers` WHERE `carriers`.`id` = ? ORDER BY `carriers`.`id` LIMIT ?",
	)).WithArgs(id, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	result, err := s.repo.FindById(id)

	s.Error(err)
	s.Equal(gorm.ErrRecordNotFound, err)
	s.Equal(models.Carrier{}, result)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *CarrierRepositoryTestSuite) TestCreate_Success() {

	// Arrange
	inputCarrier := models.Carrier{
		ID:           0, // Will be auto-incremented
		CId:          "CID-XYZ",
		CompanyName:  "Test Carrier",
		Address:      "123 Test Lane",
		Telephone:    "+1234567890",
		LocalityId:   456,
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT 1 FROM `carriers` WHERE cid = ? ORDER BY `carriers`.`id` LIMIT ?",
	)).WithArgs(inputCarrier.CId, 1).WillReturnError(gorm.ErrRecordNotFound)

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `carriers` (`cid`,`name`,`address`,`telephone`,`locality_id`) VALUES (?,?,?,?,?)",
	)).
		WithArgs(
			inputCarrier.CId,
			inputCarrier.CompanyName,
			inputCarrier.Address,
			inputCarrier.Telephone,
			inputCarrier.LocalityId,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	carrier, err := s.repo.Create(inputCarrier)

	// Assert
	s.NoError(err)
	s.Equal("CID-XYZ", carrier.CId)
	s.Equal("Test Carrier", carrier.CompanyName)
	s.Equal("123 Test Lane", carrier.Address)
	s.Equal("+1234567890", carrier.Telephone)
	s.Equal(456, carrier.LocalityId)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *CarrierRepositoryTestSuite) TestCreate_FailErrorOnCIdValidation() {

	// Arrange
	inputCarrier := models.Carrier{
		ID:           0, // Will be auto-incremented
		CId:          "CID-XYZ",
		CompanyName:  "Test Carrier",
		Address:      "123 Test Lane",
		Telephone:    "+1234567890",
		LocalityId:   456,
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT 1 FROM `carriers` WHERE cid = ? ORDER BY `carriers`.`id` LIMIT ?",
	)).WithArgs(inputCarrier.CId, 1).WillReturnError(gorm.ErrInvalidDB)

	// Act
	carrier, err := s.repo.Create(inputCarrier)

	// Assert
	s.Error(err)
	s.Equal(models.Carrier{}, carrier)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *CarrierRepositoryTestSuite) TestCreate_FailCidAlreadyExists() {

	// Arrange
	inputCarrier := models.Carrier{
		ID:           0,
		CId:          "CID-XYZ",
		CompanyName:  "Test Carrier",
		Address:      "123 Test Lane",
		Telephone:    "+1234567890",
		LocalityId:   456,
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT 1 FROM `carriers` WHERE cid = ? ORDER BY `carriers`.`id` LIMIT ?",
	)).WithArgs(inputCarrier.CId, 1).
		WillReturnRows(s.mock.NewRows([]string{"1"}).AddRow(1))

	// Act
	carrier, err := s.repo.Create(inputCarrier)

	// Assert
	s.Error(err)
	s.Equal(models.Carrier{}, carrier)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *CarrierRepositoryTestSuite) TestCreate_FailLocalityNotFound() {
	// Arrange
	inputCarrier := models.Carrier{
		ID:           0, // Will be auto-incremented
		CId:          "CID-XYZ",
		CompanyName:  "Test Carrier",
		Address:      "123 Test Lane",
		Telephone:    "+1234567890",
		LocalityId:   456,
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT 1 FROM `carriers` WHERE cid = ? ORDER BY `carriers`.`id` LIMIT ?",
	)).WithArgs(inputCarrier.CId, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `carriers` (`cid`,`name`,`address`,`telephone`,`locality_id`) VALUES (?,?,?,?,?)",
	)).WithArgs(
		inputCarrier.CId,
		inputCarrier.CompanyName,
		inputCarrier.Address,
		inputCarrier.Telephone,
		inputCarrier.LocalityId,
	).WillReturnError(gorm.ErrForeignKeyViolated)
	s.mock.ExpectRollback()

	// Act
	carrier, err := s.repo.Create(inputCarrier)

	// Assert
	s.Error(err)
	s.ErrorIs(err, repository.ErrLocalityNotFound)
	s.Equal(models.Carrier{}, carrier)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *CarrierRepositoryTestSuite) TestUpdate_Success() {
	// Arrange
	existingCarrier := models.Carrier{
		ID:                 1,
		CId:				"CID#01",
		CompanyName:        "Meli",
		Address:            "Boulevard",
		Telephone:          "123-456789",
		LocalityId:         1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `carriers` SET `cid`=?,`name`=?,`address`=?,`telephone`=?,`locality_id`=? WHERE `id` = ?",
	)).
		WithArgs(
			existingCarrier.CId,
			existingCarrier.CompanyName,
			existingCarrier.Address,
			existingCarrier.Telephone,
			existingCarrier.LocalityId,
			existingCarrier.ID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	updatedCarrier, err := s.repo.Update(existingCarrier)

	// Assert
	s.NoError(err)
	s.Equal(existingCarrier, updatedCarrier)

	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *CarrierRepositoryTestSuite) TestUpdate_ForeignKeyViolation() {
	// Arrange
	existingCarrier := models.Carrier{
		ID:                 1,
		CId:				"CID#01",
		CompanyName:        "Meli",
		Address:            "Boulevard",
		Telephone:          "123-456789",
		LocalityId:         1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `carriers` SET `cid`=?,`name`=?,`address`=?,`telephone`=?,`locality_id`=? WHERE `id` = ?",
	)).
		WithArgs(
			existingCarrier.CId,
			existingCarrier.CompanyName,
			existingCarrier.Address,
			existingCarrier.Telephone,
			existingCarrier.LocalityId,
			existingCarrier.ID,
		).
		WillReturnError(gorm.ErrForeignKeyViolated)
	s.mock.ExpectRollback()

	// Act
	updatedCarrier, err := s.repo.Update(existingCarrier)

	// Assert
	s.Error(err)
	s.Equal(models.Carrier{}, updatedCarrier)

	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *CarrierRepositoryTestSuite) TestPartialUpdate_Success() {
	// Arrange
	id := 1
	originalCarrier := models.Carrier{
		ID:                 id,
		CId:				"CID#01",
		CompanyName:        "Meli",
		Address:            "Boulevard",
		Telephone:          "123-456789",
		LocalityId:         1,
	}

	fields := map[string]interface{}{
		"cid":			"NEW-CODE",
		"company_name":	"LibreMercado",
		"address":		"New Address",
		"telephone":	"123-321",
		"locality_id":	float64(2),
	}

	columns := []string{
		"id",
		"cid",
		"name",
		"address",
		"telephone",
		"locality_id",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `carriers` WHERE `carriers`.`id` = ? ORDER BY `carriers`.`id` LIMIT ?",
	)).WithArgs(id, 1).
		WillReturnRows(s.mock.NewRows(columns).
			AddRow(
				originalCarrier.ID,
				originalCarrier.CId,
				originalCarrier.CompanyName,
				originalCarrier.Address,
				originalCarrier.Telephone,
				originalCarrier.LocalityId,
			),
		)

	expectedUpdatedCarrier := models.Carrier{
		ID:				id,
		CId:			"NEW-CODE",
		CompanyName:	"LibreMercado",
		Address:		"New Address",
		Telephone:		"123-321",
		LocalityId:		2,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `carriers` SET `cid`=?,`name`=?,`address`=?,`telephone`=?,`locality_id`=? WHERE `id` = ?",
	)).
		WithArgs(
			expectedUpdatedCarrier.CId,
			expectedUpdatedCarrier.CompanyName,
			expectedUpdatedCarrier.Address,
			expectedUpdatedCarrier.Telephone,
			expectedUpdatedCarrier.LocalityId,
			expectedUpdatedCarrier.ID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	result, err := s.repo.PartialUpdate(id, fields)

	// Assert
	s.NoError(err)
	s.Equal(expectedUpdatedCarrier, result)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *CarrierRepositoryTestSuite) TestPartialUpdate_DatabaseErrorOnFirst() {
	// Arrange
	id := 1
	fields := map[string]interface{}{
		"cid":			"NEW-CODE",
		"company_name":	"LibreMercado",
		"address":		"New Address",
		"telephone":	"123-321",
		"locality_id":	float64(2),
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `carriers` WHERE `carriers`.`id` = ? ORDER BY `carriers`.`id` LIMIT ?",
	)).WithArgs(id, 1).
		WillReturnError(gorm.ErrInvalidDB)

	// Act
	updatedCarrier, err := s.repo.PartialUpdate(id, fields)

	// Assert
	s.Error(err)
	s.ErrorIs(err, gorm.ErrInvalidDB)
	s.Equal(models.Carrier{}, updatedCarrier)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *CarrierRepositoryTestSuite) TestPartialUpdate_ErrorNotFound() {
	// Arrange
	id := 1111
	fields := map[string]interface{}{
		"cid":			"NEW-CODE",
		"company_name":	"LibreMercado",
		"address":		"New Address",
		"telephone":	"123-321",
		"locality_id":	float64(2),
	}

	columns := []string{
		"id",
		"cid",
		"name",
		"address",
		"telephone",
		"locality_id",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `carriers` WHERE `carriers`.`id` = ? ORDER BY `carriers`.`id` LIMIT ?",
	)).WithArgs(id, 1).WillReturnRows(s.mock.NewRows(columns))

	// Act
	updatedCarrier, err := s.repo.PartialUpdate(id, fields)

	// Assert
	s.Error(err)
	s.ErrorIs(err, repository.ErrEntityNotFound)
	s.Equal(models.Carrier{}, updatedCarrier)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *CarrierRepositoryTestSuite) TestPartialUpdate_ErrorOnSave() {
	// Arrange
	id := 1
	originalCarrier := models.Carrier{
		ID:				id,
		CId:			"CID#01",
		CompanyName:	"Meli",
		Address:		"Boulevard",
		Telephone:		"123-456789",
		LocalityId:		1,
	}

	fields := map[string]interface{}{
		"cid":			"NEW-CODE",
		"company_name":	"LibreMercado",
		"address":		"New Address",
		"telephone":	"123-321",
		"locality_id":	float64(2),
	}

	columns := []string{
		"id",
		"cid",
		"name",
		"address",
		"telephone",
		"locality_id",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `carriers` WHERE `carriers`.`id` = ? ORDER BY `carriers`.`id` LIMIT ?",
	)).
		WithArgs(id, 1).
		WillReturnRows(s.mock.NewRows(columns).
			AddRow(
				originalCarrier.ID,
				originalCarrier.CId,
				originalCarrier.CompanyName,
				originalCarrier.Address,
				originalCarrier.Telephone,
				originalCarrier.LocalityId,
			),
		)

	expectedUpdatedCarrier := models.Carrier{
		ID:				id,
		CId:			"NEW-CODE",
		CompanyName:	"LibreMercado",
		Address:		"New Address",
		Telephone:		"123-321",
		LocalityId:		2,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `carriers` SET `cid`=?,`name`=?,`address`=?,`telephone`=?,`locality_id`=? WHERE `id` = ?",
	)).
		WithArgs(
			expectedUpdatedCarrier.CId,
			expectedUpdatedCarrier.CompanyName,
			expectedUpdatedCarrier.Address,
			expectedUpdatedCarrier.Telephone,
			expectedUpdatedCarrier.LocalityId,
			expectedUpdatedCarrier.ID,
		).
		WillReturnError(gorm.ErrInvalidValue)
	s.mock.ExpectRollback()

	// Act
	result, err := s.repo.PartialUpdate(id, fields)

	// Assert
	s.Error(err)
	s.Equal(models.Carrier{}, result)
	s.NoError(s.mock.ExpectationsWereMet())
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *CarrierRepositoryTestSuite) TestDelete_Success() {
	// Arrange
	carrierID := 1
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"DELETE FROM `carriers` WHERE `carriers`.`id` = ?",
	)).WithArgs(carrierID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	err := s.repo.Delete(carrierID)

	// Assert
	s.NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *CarrierRepositoryTestSuite) TestDelete_NotFound() {
	// Arrange
	carrierID := 999

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"DELETE FROM `carriers` WHERE `carriers`.`id` = ?",
	)).WithArgs(carrierID).
		WillReturnResult(sqlmock.NewResult(1, 0))
	s.mock.ExpectCommit()

	// Act
	err := s.repo.Delete(carrierID)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

// Run the test suite
func TestCarrierRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CarrierRepositoryTestSuite))
}
