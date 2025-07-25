package database

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
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
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = gormDB
	s.repo = NewLocalityRepository(s.db)
}

// Test Create - Caso exitoso: crea una locality
func (s *LocalityRepositoryTestSuite) TestCreate_Success() {
	// Arrange
	newLocality := models.LocalityDoc{
		Id:       1,
		Locality: "Buenos Aires",
		Province: "Buenos Aires",
		Country:  "Argentina",
	}

	expectedProvince := models.Province{
		Id:        1,
		Province:  "Buenos Aires",
		CountryId: 1,
	}

	// Mock para buscar la provincia
	provinceRows := sqlmock.NewRows([]string{"id", "province", "country_id"}).
		AddRow(expectedProvince.Id, expectedProvince.Province, expectedProvince.CountryId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `provinces`.`id`,`provinces`.`province`,`provinces`.`country_id` FROM `provinces` INNER JOIN countries ON countries.id = provinces.country_id WHERE provinces.province = ? AND countries.country = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocality.Province, newLocality.Country, 1).
		WillReturnRows(provinceRows)

	// Mock para crear la locality
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `localities` (`id`,`locality`,`province_id`) VALUES (?,?,?)")).
		WithArgs(newLocality.Id, newLocality.Locality, expectedProvince.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	createdLocality, err := s.repo.Create(newLocality)

	// Asserts
	s.NoError(err)
	s.Equal(newLocality.Id, createdLocality.Id)
	s.Equal(newLocality.Locality, createdLocality.Locality)
	s.Equal(newLocality.Province, createdLocality.Province)
	s.Equal(newLocality.Country, createdLocality.Country)
}

// Test Create - Province no encontrada
func (s *LocalityRepositoryTestSuite) TestCreate_ProvinceNotFound() {
	// Arrange
	newLocality := models.LocalityDoc{
		Id:       1,
		Locality: "Test City",
		Province: "NonExistent Province",
		Country:  "NonExistent Country",
	}

	// Mock para buscar provincia que no existe
	provinceRows := sqlmock.NewRows([]string{"id", "province", "country_id"})
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `provinces`.`id`,`provinces`.`province`,`provinces`.`country_id` FROM `provinces` INNER JOIN countries ON countries.id = provinces.country_id WHERE provinces.province = ? AND countries.country = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocality.Province, newLocality.Country, 1).
		WillReturnRows(provinceRows)

	// Act
	createdLocality, err := s.repo.Create(newLocality)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrProvinceNotFound, err)
	s.Equal(models.LocalityDoc{}, createdLocality)
}

// Test Create - Error de base de datos al buscar provincia
func (s *LocalityRepositoryTestSuite) TestCreate_ProvinceFindDatabaseError() {
	// Arrange
	newLocality := models.LocalityDoc{
		Id:       1,
		Locality: "Test City",
		Province: "Buenos Aires",
		Country:  "Argentina",
	}

	// Mock para error de conexión al buscar provincia
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `provinces`.`id`,`provinces`.`province`,`provinces`.`country_id` FROM `provinces` INNER JOIN countries ON countries.id = provinces.country_id WHERE provinces.province = ? AND countries.country = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocality.Province, newLocality.Country, 1).
		WillReturnError(sql.ErrConnDone)

	// Act
	createdLocality, err := s.repo.Create(newLocality)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.LocalityDoc{}, createdLocality)
}

// Test Create - Locality ya existe (Duplicate Key)
func (s *LocalityRepositoryTestSuite) TestCreate_DuplicatedKey() {
	// Arrange
	newLocality := models.LocalityDoc{
		Id:       1,
		Locality: "Buenos Aires",
		Province: "Buenos Aires",
		Country:  "Argentina",
	}

	expectedProvince := models.Province{
		Id:        1,
		Province:  "Buenos Aires",
		CountryId: 1,
	}

	// Mock para buscar la provincia exitosamente
	provinceRows := sqlmock.NewRows([]string{"id", "province", "country_id"}).
		AddRow(expectedProvince.Id, expectedProvince.Province, expectedProvince.CountryId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `provinces`.`id`,`provinces`.`province`,`provinces`.`country_id` FROM `provinces` INNER JOIN countries ON countries.id = provinces.country_id WHERE provinces.province = ? AND countries.country = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocality.Province, newLocality.Country, 1).
		WillReturnRows(provinceRows)

	// Mock para crear locality que ya existe
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `localities` (`id`,`locality`,`province_id`) VALUES (?,?,?)")).
		WithArgs(newLocality.Id, newLocality.Locality, expectedProvince.Id).
		WillReturnError(gorm.ErrDuplicatedKey)
	s.mock.ExpectRollback()

	// Act
	createdLocality, err := s.repo.Create(newLocality)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrEntityAlreadyExists, err)
	s.Equal(models.LocalityDoc{}, createdLocality)
}

// Test Create - Foreign Key Violation
func (s *LocalityRepositoryTestSuite) TestCreate_ForeignKeyViolation() {
	// Arrange
	newLocality := models.LocalityDoc{
		Id:       1,
		Locality: "Test City",
		Province: "Buenos Aires",
		Country:  "Argentina",
	}

	expectedProvince := models.Province{
		Id:        999, // ID que causará violación de FK
		Province:  "Buenos Aires",
		CountryId: 999,
	}

	// Mock para buscar la provincia exitosamente
	provinceRows := sqlmock.NewRows([]string{"id", "province", "country_id"}).
		AddRow(expectedProvince.Id, expectedProvince.Province, expectedProvince.CountryId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `provinces`.`id`,`provinces`.`province`,`provinces`.`country_id` FROM `provinces` INNER JOIN countries ON countries.id = provinces.country_id WHERE provinces.province = ? AND countries.country = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocality.Province, newLocality.Country, 1).
		WillReturnRows(provinceRows)

	// Mock para crear locality con violación de FK
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `localities` (`id`,`locality`,`province_id`) VALUES (?,?,?)")).
		WithArgs(newLocality.Id, newLocality.Locality, expectedProvince.Id).
		WillReturnError(gorm.ErrForeignKeyViolated)
	s.mock.ExpectRollback()

	// Act
	createdLocality, err := s.repo.Create(newLocality)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrForeignKeyViolation, err)
	s.Equal(models.LocalityDoc{}, createdLocality)
}

// Test Create - Error genérico de base de datos durante creación
func (s *LocalityRepositoryTestSuite) TestCreate_DatabaseError() {
	// Arrange
	newLocality := models.LocalityDoc{
		Id:       1,
		Locality: "Test City",
		Province: "Buenos Aires",
		Country:  "Argentina",
	}

	expectedProvince := models.Province{
		Id:        1,
		Province:  "Buenos Aires",
		CountryId: 1,
	}

	// Mock para buscar la provincia exitosamente
	provinceRows := sqlmock.NewRows([]string{"id", "province", "country_id"}).
		AddRow(expectedProvince.Id, expectedProvince.Province, expectedProvince.CountryId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `provinces`.`id`,`provinces`.`province`,`provinces`.`country_id` FROM `provinces` INNER JOIN countries ON countries.id = provinces.country_id WHERE provinces.province = ? AND countries.country = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocality.Province, newLocality.Country, 1).
		WillReturnRows(provinceRows)

	// Mock para error genérico durante creación
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `localities` (`id`,`locality`,`province_id`) VALUES (?,?,?)")).
		WithArgs(newLocality.Id, newLocality.Locality, expectedProvince.Id).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	createdLocality, err := s.repo.Create(newLocality)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.LocalityDoc{}, createdLocality)
}

func TestLocalityRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(LocalityRepositoryTestSuite))
}
