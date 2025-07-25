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

// Test Create - Success
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

	provinceRows := sqlmock.NewRows([]string{"id", "province", "country_id"}).
		AddRow(expectedProvince.Id, expectedProvince.Province, expectedProvince.CountryId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `provinces`.`id`,`provinces`.`province`,`provinces`.`country_id` FROM `provinces` INNER JOIN countries ON countries.id = provinces.country_id WHERE provinces.province = ? AND countries.country = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocality.Province, newLocality.Country, 1).
		WillReturnRows(provinceRows)

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `localities` (`locality`,`province_id`,`id`) VALUES (?,?,?)")).
		WithArgs(newLocality.Locality, expectedProvince.Id, newLocality.Id).
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

// Test Create - Province notFound
func (s *LocalityRepositoryTestSuite) TestCreate_ProvinceNotFound() {
	// Arrange
	newLocality := models.LocalityDoc{
		Id:       1,
		Locality: "Test City",
		Province: "NonExistent Province",
		Country:  "NonExistent Country",
	}

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

// Test Create - Error connected to db
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

// Test Create - Locality al ready exist (Duplicate Key)
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
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `localities` (`locality`,`province_id`,`id`) VALUES (?,?,?)")).
		WithArgs(newLocality.Locality, expectedProvince.Id, newLocality.Id).
		WillReturnError(gorm.ErrDuplicatedKey)
	s.mock.ExpectRollback()

	// Act
	createdLocality, err := s.repo.Create(newLocality)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrEntityAlreadyExists, err)
	s.Equal(models.LocalityDoc{}, createdLocality)
}

// Test Create - Foreign key violation
func (s *LocalityRepositoryTestSuite) TestCreate_ForeignKeyViolation() {
	// Arrange
	newLocality := models.LocalityDoc{
		Id:       1,
		Locality: "Test City",
		Province: "Buenos Aires",
		Country:  "Argentina",
	}

	expectedProvince := models.Province{
		Id:        999,
		Province:  "Buenos Aires",
		CountryId: 999,
	}

	provinceRows := sqlmock.NewRows([]string{"id", "province", "country_id"}).
		AddRow(expectedProvince.Id, expectedProvince.Province, expectedProvince.CountryId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `provinces`.`id`,`provinces`.`province`,`provinces`.`country_id` FROM `provinces` INNER JOIN countries ON countries.id = provinces.country_id WHERE provinces.province = ? AND countries.country = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocality.Province, newLocality.Country, 1).
		WillReturnRows(provinceRows)

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `localities` (`locality`,`province_id`,`id`) VALUES (?,?,?)")).
		WithArgs(newLocality.Locality, expectedProvince.Id, newLocality.Id).
		WillReturnError(gorm.ErrForeignKeyViolated)
	s.mock.ExpectRollback()

	// Act
	createdLocality, err := s.repo.Create(newLocality)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrForeignKeyViolation, err)
	s.Equal(models.LocalityDoc{}, createdLocality)
}

// Test Create - Error generic database
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

	provinceRows := sqlmock.NewRows([]string{"id", "province", "country_id"}).
		AddRow(expectedProvince.Id, expectedProvince.Province, expectedProvince.CountryId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `provinces`.`id`,`provinces`.`province`,`provinces`.`country_id` FROM `provinces` INNER JOIN countries ON countries.id = provinces.country_id WHERE provinces.province = ? AND countries.country = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocality.Province, newLocality.Country, 1).
		WillReturnRows(provinceRows)

	// Mock para error genérico durante creación
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `localities` (`locality`,`province_id`,`id`) VALUES (?,?,?)")).
		WithArgs(newLocality.Locality, expectedProvince.Id, newLocality.Id).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	createdLocality, err := s.repo.Create(newLocality)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.LocalityDoc{}, createdLocality)
}

// Test FindById - Success
func (s *LocalityRepositoryTestSuite) TestFindById_Success() {
	// Arrange
	expectedLocality := models.Locality{
		Id:         1,
		Locality:   "Buenos Aires",
		ProvinceId: 1,
	}

	localityRows := sqlmock.NewRows([]string{"id", "locality", "province_id"}).
		AddRow(expectedLocality.Id, expectedLocality.Locality, expectedLocality.ProvinceId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `localities` WHERE `localities`.`id` = ? ORDER BY `localities`.`id` LIMIT ?")).
		WithArgs(1, 1).
		WillReturnRows(localityRows)

	// Act
	locality, err := s.repo.FindById(1)

	// Asserts
	s.NoError(err)
	s.Equal(expectedLocality.Id, locality.Id)
	s.Equal(expectedLocality.Locality, locality.Locality)
	s.Equal(expectedLocality.ProvinceId, locality.ProvinceId)
}

// Test FindById - Entity Not Found
func (s *LocalityRepositoryTestSuite) TestFindById_EntityNotFound() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `localities` WHERE `localities`.`id` = ? ORDER BY `localities`.`id` LIMIT ?")).
		WithArgs(999, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	// Act
	locality, err := s.repo.FindById(999)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal(models.Locality{}, locality)
}

// Test FindById - Database Error
func (s *LocalityRepositoryTestSuite) TestFindById_DatabaseError() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `localities` WHERE `localities`.`id` = ? ORDER BY `localities`.`id` LIMIT ?")).
		WithArgs(1, 1).
		WillReturnError(sql.ErrConnDone)

	// Act
	locality, err := s.repo.FindById(1)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.Locality{}, locality)
}

// Test FindAll - Success
func (s *LocalityRepositoryTestSuite) TestFindAll_Success() {
	// Arrange
	expectedLocalities := []models.Locality{
		{
			Id:         1,
			Locality:   "Buenos Aires",
			ProvinceId: 1,
		},
		{
			Id:         2,
			Locality:   "Córdoba",
			ProvinceId: 2,
		},
	}

	localityRows := sqlmock.NewRows([]string{"id", "locality", "province_id"}).
		AddRow(expectedLocalities[0].Id, expectedLocalities[0].Locality, expectedLocalities[0].ProvinceId).
		AddRow(expectedLocalities[1].Id, expectedLocalities[1].Locality, expectedLocalities[1].ProvinceId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `localities`")).
		WillReturnRows(localityRows)

	// Act
	localities, err := s.repo.FindAll()

	// Asserts
	s.NoError(err)
	s.Len(localities, 2)
	s.Equal(expectedLocalities[0].Id, localities[0].Id)
	s.Equal(expectedLocalities[0].Locality, localities[0].Locality)
	s.Equal(expectedLocalities[1].Id, localities[1].Id)
	s.Equal(expectedLocalities[1].Locality, localities[1].Locality)
}

// Test FindAll - Success Empty Result
func (s *LocalityRepositoryTestSuite) TestFindAll_SuccessEmptyResult() {
	// Arrange
	localityRows := sqlmock.NewRows([]string{"id", "locality", "province_id"})

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `localities`")).
		WillReturnRows(localityRows)

	// Act
	localities, err := s.repo.FindAll()

	// Asserts
	s.NoError(err)
	s.Empty(localities)
	s.Len(localities, 0)
}

// Test FindAll - Database Error
func (s *LocalityRepositoryTestSuite) TestFindAll_DatabaseError() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `localities`")).
		WillReturnError(sql.ErrConnDone)

	// Act
	localities, err := s.repo.FindAll()

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Nil(localities)
}

// Test FindAllLocality - Success
func (s *LocalityRepositoryTestSuite) TestFindAllLocality_Success() {
	// Arrange
	expectedLocalities := []models.LocalitySellerCount{
		{
			LocalityDoc: models.LocalityDoc{
				Id:       1,
				Locality: "Buenos Aires",
				Province: "Buenos Aires",
				Country:  "Argentina",
			},
			SellerCount: func() *int { count := 5; return &count }(),
		},
		{
			LocalityDoc: models.LocalityDoc{
				Id:       2,
				Locality: "Córdoba",
				Province: "Córdoba",
				Country:  "Argentina",
			},
			SellerCount: func() *int { count := 3; return &count }(),
		},
	}

	localityRows := sqlmock.NewRows([]string{"id", "locality", "province", "country", "seller_count"}).
		AddRow(expectedLocalities[0].Id, expectedLocalities[0].Locality, expectedLocalities[0].Province, expectedLocalities[0].Country, *expectedLocalities[0].SellerCount).
		AddRow(expectedLocalities[1].Id, expectedLocalities[1].Locality, expectedLocalities[1].Province, expectedLocalities[1].Country, *expectedLocalities[1].SellerCount)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT localities.id as id, localities.locality as locality, p.province as province, c.country as country, COUNT(DISTINCT s.id) as seller_count FROM `localities` JOIN provinces p ON localities.province_id = p.id JOIN countries c ON p.country_id = c.id LEFT JOIN sellers s ON localities.id = s.locality_id GROUP BY localities.id, localities.locality, p.province, c.country")).
		WillReturnRows(localityRows)

	// Act
	localities, err := s.repo.FindAllLocality()

	// Asserts
	s.NoError(err)
	s.Len(localities, 2)
	s.Equal(expectedLocalities[0].Id, localities[0].Id)
	s.Equal(expectedLocalities[0].Locality, localities[0].Locality)
	s.Equal(expectedLocalities[0].Province, localities[0].Province)
	s.Equal(expectedLocalities[0].Country, localities[0].Country)
	s.Equal(*expectedLocalities[0].SellerCount, *localities[0].SellerCount)
	s.Equal(expectedLocalities[1].Id, localities[1].Id)
	s.Equal(expectedLocalities[1].Locality, localities[1].Locality)
}

// Test FindAllLocality - Success Empty Result
func (s *LocalityRepositoryTestSuite) TestFindAllLocality_SuccessEmptyResult() {
	// Arrange
	localityRows := sqlmock.NewRows([]string{"id", "locality", "province", "country", "seller_count"})

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT localities.id as id, localities.locality as locality, p.province as province, c.country as country, COUNT(DISTINCT s.id) as seller_count FROM `localities` JOIN provinces p ON localities.province_id = p.id JOIN countries c ON p.country_id = c.id LEFT JOIN sellers s ON localities.id = s.locality_id GROUP BY localities.id, localities.locality, p.province, c.country")).
		WillReturnRows(localityRows)

	// Act
	localities, err := s.repo.FindAllLocality()

	// Asserts
	s.NoError(err)
	s.Empty(localities)
	s.Len(localities, 0)
}

// Test FindAllLocality - Database Error
func (s *LocalityRepositoryTestSuite) TestFindAllLocality_DatabaseError() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT localities.id as id, localities.locality as locality, p.province as province, c.country as country, COUNT(DISTINCT s.id) as seller_count FROM `localities` JOIN provinces p ON localities.province_id = p.id JOIN countries c ON p.country_id = c.id LEFT JOIN sellers s ON localities.id = s.locality_id GROUP BY localities.id, localities.locality, p.province, c.country")).
		WillReturnError(sql.ErrConnDone)

	// Act
	localities, err := s.repo.FindAllLocality()

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Nil(localities)
}

func TestLocalityRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(LocalityRepositoryTestSuite))
}
