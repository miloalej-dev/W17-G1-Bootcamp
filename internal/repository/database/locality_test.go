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

type LocalityRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo *LocalityRepository
}

func (s *LocalityRepositoryTestSuite) SetupSuite() {
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
	s.repo = NewLocalityRepository(gormDB)
}

// Test for the find all success
func (s *LocalityRepositoryTestSuite) TestFindAllCarriers_Success() {
	// Arrange
	expectedLocalities := []models.LocalityCarrierCount{
		{
			LocalityID:		1,
			LocalityName:	"Locality_1",
			TotalCarriers:	3,
		},
		{
			LocalityID:		2,
			LocalityName:	"Locality_2",
			TotalCarriers:	0,
		},
	}

	columns := []string{
		"locality_id",
		"locality_name",
		"total_carriers",
	}

	rows := s.mock.NewRows(columns)
	for _, w := range expectedLocalities {
		rows = rows.AddRow(
			w.LocalityID,
			w.LocalityName,
			w.TotalCarriers,
		)
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT localities.id as 'locality_id', localities.locality as 'locality_name', COUNT(carriers.id) 'total_carriers' FROM `localities` LEFT JOIN carriers ON localities.id = carriers.locality_id GROUP BY `localities`.`id`",
	)).WillReturnRows(rows)

	// Act
	localities, err := s.repo.FindAllCarriers()

	// Assert
	s.NoError(err)
	s.Len(localities, 2)
	s.Equal(expectedLocalities, localities)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

// Error on retireve all
func (s *LocalityRepositoryTestSuite) TestFindAllCarriers_NoRecordsFound() {
	// Arrange
	columns := []string{
		"locality_id",
		"locality_name",
		"total_carriers",
	}

	rows := s.mock.NewRows(columns)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT localities.id as 'locality_id', localities.locality as 'locality_name', COUNT(carriers.id) 'total_carriers' FROM `localities` LEFT JOIN carriers ON localities.id = carriers.locality_id GROUP BY `localities`.`id`",
	)).WillReturnRows(rows)

	// Act
	localities, err := s.repo.FindAllCarriers()

	// Assert
	s.Error(err)
	s.Nil(localities)
	s.Equal(repository.ErrEntityNotFound, err)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *LocalityRepositoryTestSuite) TestFindAllCarriers_DatabaseError() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT localities.id as 'locality_id', localities.locality as 'locality_name', COUNT(carriers.id) 'total_carriers' FROM `localities` LEFT JOIN carriers ON localities.id = carriers.locality_id GROUP BY `localities`.`id`",
	)).WillReturnError(sql.ErrConnDone)

	// Act
	localities, err := s.repo.FindAllCarriers()

	// Assert
	s.Error(err)
	s.Nil(localities)
	s.Equal(sql.ErrConnDone, err)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *LocalityRepositoryTestSuite) TestFindAllCarriersByLocality_Success() {
	// Arrange
	id := 1
	expectedLocalities := []models.LocalityCarrierCount {
		{
			LocalityID:		1,
			LocalityName:	"Locality_1",
			TotalCarriers:	3,
		},
	}

	columns := []string{
		"locality_id",
		"locality_name",
		"total_carriers",
	}

	rows := s.mock.NewRows(columns)
	for _, w := range expectedLocalities {
		rows = rows.AddRow(
			w.LocalityID,
			w.LocalityName,
			w.TotalCarriers,
		)
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT localities.id as 'locality_id', localities.locality as 'locality_name', COUNT(carriers.id) 'total_carriers' FROM `localities` LEFT JOIN carriers ON localities.id = carriers.locality_id WHERE localities.id = ? GROUP BY `localities`.`id`",
	)).WithArgs(id).WillReturnRows(rows)

	// Act
	localities, err := s.repo.FindCarriersByLocality(id)

	// Assert
	s.NoError(err)
	s.Len(localities, 1)
	s.Equal(expectedLocalities, localities)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *LocalityRepositoryTestSuite) TestFindAllCarriersByLocality_DatabasError() {
	// Arrange
	id := 1

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT localities.id as 'locality_id', localities.locality as 'locality_name', COUNT(carriers.id) 'total_carriers' FROM `localities` LEFT JOIN carriers ON localities.id = carriers.locality_id WHERE localities.id = ? GROUP BY `localities`.`id`",
	)).WithArgs(id).WillReturnError(sql.ErrConnDone)

	// Act
	localities, err := s.repo.FindCarriersByLocality(id)

	// Assert
	s.Error(err)
	s.Nil(localities)
	s.Equal(sql.ErrConnDone, err)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *LocalityRepositoryTestSuite) TestFindAllCarriersByLocality_NoRecordsFound() {
	// Arrange
	id := 1
	columns := []string{
		"locality_id",
		"locality_name",
		"total_carriers",
	}

	rows := s.mock.NewRows(columns)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT localities.id as 'locality_id', localities.locality as 'locality_name', COUNT(carriers.id) 'total_carriers' FROM `localities` LEFT JOIN carriers ON localities.id = carriers.locality_id WHERE localities.id = ? GROUP BY `localities`.`id`",
	)).WithArgs(id).WillReturnRows(rows)

	// Act
	localities, err := s.repo.FindCarriersByLocality(id)

	// Assert
	s.Error(err)
	s.Nil(localities)
	s.Equal(repository.ErrEntityNotFound, err)
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

// Run the test suite
func TestLocalityRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(LocalityRepositoryTestSuite))
}
