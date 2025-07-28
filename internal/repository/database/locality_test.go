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

func (s *LocalityRepositoryTestSuite) TestCreate_Success() {
	// Arrange
	newLocality := models.Locality{
		Id:         1,
		Locality:   "Buenos Aires",
		ProvinceId: 1,
	}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `provinces` WHERE `provinces`.`id` = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocality.ProvinceId, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "province", "country_id"}).
			AddRow(1, "Buenos Aires Province", 1))

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `localities` (`locality`,`province_id`,`id`) VALUES (?,?,?)")).
		WithArgs(newLocality.Locality, newLocality.ProvinceId, newLocality.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	createdLocality, err := s.repo.Create(newLocality)

	// Asserts
	s.NoError(err)
	s.Equal(newLocality.Id, createdLocality.Id)
	s.Equal(newLocality.Locality, createdLocality.Locality)
	s.Equal(newLocality.ProvinceId, createdLocality.ProvinceId)
}

func (s *LocalityRepositoryTestSuite) TestCreate_ProvinceNotFound() {
	// Arrange
	newLocality := models.Locality{
		Id:         1,
		Locality:   "Test City",
		ProvinceId: 999, // Non-existent province ID
	}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `provinces` WHERE `provinces`.`id` = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocality.ProvinceId, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	// Act
	createdLocality, err := s.repo.Create(newLocality)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrForeignKeyViolation, err)
	s.Equal(models.Locality{}, createdLocality)
}

// Test Create - Error connected to db
func (s *LocalityRepositoryTestSuite) TestCreate_ProvinceFindDatabaseError() {
	// Arrange
	newLocality := models.Locality{
		Id:         1,
		Locality:   "Test City",
		ProvinceId: 1,
	}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `provinces` WHERE `provinces`.`id` = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocality.ProvinceId, 1).
		WillReturnError(sql.ErrConnDone)

	// Act
	createdLocality, err := s.repo.Create(newLocality)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.Locality{}, createdLocality)
}

func (s *LocalityRepositoryTestSuite) TestCreate_DuplicatedKey() {
	// Arrange
	newLocality := models.Locality{
		Id:         1,
		Locality:   "Buenos Aires",
		ProvinceId: 1,
	}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `provinces` WHERE `provinces`.`id` = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocality.ProvinceId, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "province", "country_id"}).
			AddRow(1, "Buenos Aires Province", 1))

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `localities` (`locality`,`province_id`,`id`) VALUES (?,?,?)")).
		WithArgs(newLocality.Locality, newLocality.ProvinceId, newLocality.Id).
		WillReturnError(gorm.ErrDuplicatedKey)
	s.mock.ExpectRollback()

	// Act
	createdLocality, err := s.repo.Create(newLocality)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrEntityAlreadyExists, err)
	s.Equal(models.Locality{}, createdLocality)
}

func (s *LocalityRepositoryTestSuite) TestCreate_ForeignKeyViolation() {
	// Arrange
	newLocality := models.Locality{
		Id:         1,
		Locality:   "Buenos Aires",
		ProvinceId: 1,
	}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `provinces` WHERE `provinces`.`id` = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocality.ProvinceId, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "province", "country_id"}).
			AddRow(1, "Buenos Aires Province", 1))

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `localities` (`locality`,`province_id`,`id`) VALUES (?,?,?)")).
		WithArgs(newLocality.Locality, newLocality.ProvinceId, newLocality.Id).
		WillReturnError(gorm.ErrForeignKeyViolated)
	s.mock.ExpectRollback()

	// Act
	createdLocality, err := s.repo.Create(newLocality)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrForeignKeyViolation, err)
	s.Equal(models.Locality{}, createdLocality)
}

func (s *LocalityRepositoryTestSuite) TestCreate_DatabaseError() {
	// Arrange
	newLocality := models.Locality{
		Id:         1,
		Locality:   "Test City",
		ProvinceId: 1,
	}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `provinces` WHERE `provinces`.`id` = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocality.ProvinceId, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "province", "country_id"}).
			AddRow(1, "Buenos Aires Province", 1))

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `localities` (`locality`,`province_id`,`id`) VALUES (?,?,?)")).
		WithArgs(newLocality.Locality, newLocality.ProvinceId, newLocality.Id).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	createdLocality, err := s.repo.Create(newLocality)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.Locality{}, createdLocality)
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

// Test FindLocalityBySeller - Success
func (s *LocalityRepositoryTestSuite) TestFindLocalityBySeller_Success() {
	// Arrange
	localityId := 1
	expectedLocality := models.LocalitySellerCount{
		LocalityDoc: models.LocalityDoc{
			Id:       1,
			Locality: "Buenos Aires",
			Province: "Buenos Aires",
			Country:  "Argentina",
		},
		SellerCount: func() *int { count := 5; return &count }(),
	}

	localityRows := sqlmock.NewRows([]string{"id", "locality", "province", "country", "seller_count"}).
		AddRow(expectedLocality.Id, expectedLocality.Locality, expectedLocality.Province, expectedLocality.Country, *expectedLocality.SellerCount)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT localities.id as id, localities.locality as locality, p.province as province, c.country as country, COUNT(DISTINCT s.id) as seller_count FROM `localities` JOIN provinces p ON localities.province_id = p.id JOIN countries c ON p.country_id = c.id LEFT JOIN sellers s ON localities.id = s.locality_id WHERE localities.id = ? GROUP BY localities.id, localities.locality, p.province, c.country")).
		WithArgs(localityId).
		WillReturnRows(localityRows)

	// Act
	locality, err := s.repo.FindLocalityBySeller(localityId)

	// Asserts
	s.NoError(err)
	s.Equal(expectedLocality.Id, locality.Id)
	s.Equal(expectedLocality.Locality, locality.Locality)
	s.Equal(expectedLocality.Province, locality.Province)
	s.Equal(expectedLocality.Country, locality.Country)
	s.Equal(*expectedLocality.SellerCount, *locality.SellerCount)
}

// Test FindLocalityBySeller - Entity Not Found (No Rows)
func (s *LocalityRepositoryTestSuite) TestFindLocalityBySeller_EntityNotFound() {
	// Arrange
	localityId := 999

	localityRows := sqlmock.NewRows([]string{"id", "locality", "province", "country", "seller_count"})

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT localities.id as id, localities.locality as locality, p.province as province, c.country as country, COUNT(DISTINCT s.id) as seller_count FROM `localities` JOIN provinces p ON localities.province_id = p.id JOIN countries c ON p.country_id = c.id LEFT JOIN sellers s ON localities.id = s.locality_id WHERE localities.id = ? GROUP BY localities.id, localities.locality, p.province, c.country")).
		WithArgs(localityId).
		WillReturnRows(localityRows)

	// Act
	locality, err := s.repo.FindLocalityBySeller(localityId)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal(models.LocalitySellerCount{}, locality)
}

// Test FindLocalityBySeller - Database Error
func (s *LocalityRepositoryTestSuite) TestFindLocalityBySeller_DatabaseError() {
	// Arrange
	localityId := 1

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT localities.id as id, localities.locality as locality, p.province as province, c.country as country, COUNT(DISTINCT s.id) as seller_count FROM `localities` JOIN provinces p ON localities.province_id = p.id JOIN countries c ON p.country_id = c.id LEFT JOIN sellers s ON localities.id = s.locality_id WHERE localities.id = ? GROUP BY localities.id, localities.locality, p.province, c.country")).
		WithArgs(localityId).
		WillReturnError(sql.ErrConnDone)

	// Act
	locality, err := s.repo.FindLocalityBySeller(localityId)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.LocalitySellerCount{}, locality)
}

// Test FindLocalityBySeller - GORM Record Not Found Error
func (s *LocalityRepositoryTestSuite) TestFindLocalityBySeller_GormRecordNotFound() {
	// Arrange
	localityId := 1

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT localities.id as id, localities.locality as locality, p.province as province, c.country as country, COUNT(DISTINCT s.id) as seller_count FROM `localities` JOIN provinces p ON localities.province_id = p.id JOIN countries c ON p.country_id = c.id LEFT JOIN sellers s ON localities.id = s.locality_id WHERE localities.id = ? GROUP BY localities.id, localities.locality, p.province, c.country")).
		WithArgs(localityId).
		WillReturnError(gorm.ErrRecordNotFound)

	// Act
	locality, err := s.repo.FindLocalityBySeller(localityId)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal(models.LocalitySellerCount{}, locality)
}

// Test Update - Success
func (s *LocalityRepositoryTestSuite) TestUpdate_Success() {
	// Arrange
	localityToUpdate := models.Locality{
		Id:         1,
		Locality:   "Buenos Aires Updated",
		ProvinceId: 1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `localities` SET `locality`=?,`province_id`=? WHERE `id` = ?")).
		WithArgs(localityToUpdate.Locality, localityToUpdate.ProvinceId, localityToUpdate.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	updatedLocality, err := s.repo.Update(localityToUpdate)

	// Asserts
	s.NoError(err)
	s.Equal(localityToUpdate.Id, updatedLocality.Id)
	s.Equal(localityToUpdate.Locality, updatedLocality.Locality)
	s.Equal(localityToUpdate.ProvinceId, updatedLocality.ProvinceId)
}

// Test Update - Foreign Key Violation
func (s *LocalityRepositoryTestSuite) TestUpdate_ForeignKeyViolation() {
	// Arrange
	localityToUpdate := models.Locality{
		Id:         1,
		Locality:   "Buenos Aires",
		ProvinceId: 999, // Province ID que no existe
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `localities` SET `locality`=?,`province_id`=? WHERE `id` = ?")).
		WithArgs(localityToUpdate.Locality, localityToUpdate.ProvinceId, localityToUpdate.Id).
		WillReturnError(gorm.ErrForeignKeyViolated)
	s.mock.ExpectRollback()

	// Act
	updatedLocality, err := s.repo.Update(localityToUpdate)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrForeignKeyViolation, err)
	s.Equal(models.Locality{}, updatedLocality)
}

// Test Update - Database Error
func (s *LocalityRepositoryTestSuite) TestUpdate_DatabaseError() {
	// Arrange
	localityToUpdate := models.Locality{
		Id:         1,
		Locality:   "Buenos Aires",
		ProvinceId: 1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `localities` SET `locality`=?,`province_id`=? WHERE `id` = ?")).
		WithArgs(localityToUpdate.Locality, localityToUpdate.ProvinceId, localityToUpdate.Id).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	updatedLocality, err := s.repo.Update(localityToUpdate)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.Locality{}, updatedLocality)
}

// Test PartialUpdate - Success
func (s *LocalityRepositoryTestSuite) TestPartialUpdate_Success() {
	// Arrange
	localityId := 1
	fieldsToUpdate := map[string]interface{}{
		"locality": "Buenos Aires Updated",
	}

	existingLocality := models.Locality{
		Id:         1,
		Locality:   "Buenos Aires",
		ProvinceId: 1,
	}

	localityRows := sqlmock.NewRows([]string{"id", "locality", "province_id"}).
		AddRow(existingLocality.Id, existingLocality.Locality, existingLocality.ProvinceId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `localities` WHERE `localities`.`id` = ? ORDER BY `localities`.`id` LIMIT ?")).
		WithArgs(localityId, 1).
		WillReturnRows(localityRows)

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `localities` SET `locality`=? WHERE `id` = ?")).
		WithArgs(fieldsToUpdate["locality"], localityId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	updatedLocality, err := s.repo.PartialUpdate(localityId, fieldsToUpdate)

	// Asserts
	s.NoError(err)
	s.Equal(localityId, updatedLocality.Id)
	s.Equal(fieldsToUpdate["locality"], updatedLocality.Locality)
	s.Equal(existingLocality.ProvinceId, updatedLocality.ProvinceId)
}

// Test PartialUpdate - Entity Not Found
func (s *LocalityRepositoryTestSuite) TestPartialUpdate_EntityNotFound() {
	// Arrange
	localityId := 999
	fieldsToUpdate := map[string]interface{}{
		"locality": "Non Existent",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `localities` WHERE `localities`.`id` = ? ORDER BY `localities`.`id` LIMIT ?")).
		WithArgs(localityId, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	// Act
	updatedLocality, err := s.repo.PartialUpdate(localityId, fieldsToUpdate)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal(models.Locality{}, updatedLocality)
}

// Test PartialUpdate - Find Database Error
func (s *LocalityRepositoryTestSuite) TestPartialUpdate_FindDatabaseError() {
	// Arrange
	localityId := 1
	fieldsToUpdate := map[string]interface{}{
		"locality": "Buenos Aires Updated",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `localities` WHERE `localities`.`id` = ? ORDER BY `localities`.`id` LIMIT ?")).
		WithArgs(localityId, 1).
		WillReturnError(sql.ErrConnDone)

	// Act
	updatedLocality, err := s.repo.PartialUpdate(localityId, fieldsToUpdate)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.Locality{}, updatedLocality)
}

// Test PartialUpdate - Update Foreign Key Violation
func (s *LocalityRepositoryTestSuite) TestPartialUpdate_UpdateForeignKeyViolation() {
	// Arrange
	localityId := 1
	fieldsToUpdate := map[string]interface{}{
		"province_id": 999, // Province ID que no existe
	}

	existingLocality := models.Locality{
		Id:         1,
		Locality:   "Buenos Aires",
		ProvinceId: 1,
	}

	localityRows := sqlmock.NewRows([]string{"id", "locality", "province_id"}).
		AddRow(existingLocality.Id, existingLocality.Locality, existingLocality.ProvinceId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `localities` WHERE `localities`.`id` = ? ORDER BY `localities`.`id` LIMIT ?")).
		WithArgs(localityId, 1).
		WillReturnRows(localityRows)

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `localities` SET `province_id`=? WHERE `id` = ?")).
		WithArgs(fieldsToUpdate["province_id"], localityId).
		WillReturnError(gorm.ErrForeignKeyViolated)
	s.mock.ExpectRollback()

	// Act
	updatedLocality, err := s.repo.PartialUpdate(localityId, fieldsToUpdate)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrForeignKeyViolation, err)
	s.Equal(models.Locality{}, updatedLocality)
}

// Test PartialUpdate - Update Database Error
func (s *LocalityRepositoryTestSuite) TestPartialUpdate_UpdateDatabaseError() {
	// Arrange
	localityId := 1
	fieldsToUpdate := map[string]interface{}{
		"locality": "Buenos Aires Updated",
	}

	existingLocality := models.Locality{
		Id:         1,
		Locality:   "Buenos Aires",
		ProvinceId: 1,
	}

	localityRows := sqlmock.NewRows([]string{"id", "locality", "province_id"}).
		AddRow(existingLocality.Id, existingLocality.Locality, existingLocality.ProvinceId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `localities` WHERE `localities`.`id` = ? ORDER BY `localities`.`id` LIMIT ?")).
		WithArgs(localityId, 1).
		WillReturnRows(localityRows)

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `localities` SET `locality`=? WHERE `id` = ?")).
		WithArgs(fieldsToUpdate["locality"], localityId).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	updatedLocality, err := s.repo.PartialUpdate(localityId, fieldsToUpdate)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.Locality{}, updatedLocality)
}

// Test Delete - Success
func (s *LocalityRepositoryTestSuite) TestDelete_Success() {
	// Arrange
	localityId := 1

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `localities` WHERE `localities`.`id` = ?")).
		WithArgs(localityId).
		WillReturnResult(sqlmock.NewResult(1, 1)) // 1 row affected
	s.mock.ExpectCommit()

	// Act
	err := s.repo.Delete(localityId)

	// Asserts
	s.NoError(err)
}

// Test Delete - Entity Not Found
func (s *LocalityRepositoryTestSuite) TestDelete_EntityNotFound() {
	// Arrange
	localityId := 999

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `localities` WHERE `localities`.`id` = ?")).
		WithArgs(localityId).
		WillReturnResult(sqlmock.NewResult(1, 0)) // 0 rows affected
	s.mock.ExpectCommit()

	// Act
	err := s.repo.Delete(localityId)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
}

// Test Delete - Database Error
func (s *LocalityRepositoryTestSuite) TestDelete_DatabaseError() {
	// Arrange
	localityId := 1

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `localities` WHERE `localities`.`id` = ?")).
		WithArgs(localityId).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	err := s.repo.Delete(localityId)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
}

// Test CreateWithNames - Success
func (s *LocalityRepositoryTestSuite) TestCreateWithNames_Success() {
	// Arrange
	newLocalityDoc := models.LocalityDoc{
		Id:       1,
		Locality: "Buenos Aires",
		Province: "Buenos Aires",
		Country:  "Argentina",
	}

	// Mock la búsqueda de la provincia por nombre y país
	provinceRows := sqlmock.NewRows([]string{"id", "province", "country_id"}).
		AddRow(1, "Buenos Aires", 1)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `provinces`.`id`,`provinces`.`province`,`provinces`.`country_id` FROM `provinces` INNER JOIN countries ON countries.id = provinces.country_id WHERE provinces.province = ? AND countries.country = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs("Buenos Aires", "Argentina", 1).
		WillReturnRows(provinceRows)

	// Mock la creación de la localidad
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `localities` (`locality`,`province_id`,`id`) VALUES (?,?,?)")).
		WithArgs(newLocalityDoc.Locality, 1, newLocalityDoc.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	createdLocality, err := s.repo.CreateWithNames(newLocalityDoc)

	// Asserts
	s.NoError(err)
	s.Equal(newLocalityDoc.Id, createdLocality.Id)
	s.Equal(newLocalityDoc.Locality, createdLocality.Locality)
	s.Equal(newLocalityDoc.Province, createdLocality.Province)
	s.Equal(newLocalityDoc.Country, createdLocality.Country)
}

// Test CreateWithNames - Province Not Found
func (s *LocalityRepositoryTestSuite) TestCreateWithNames_ProvinceNotFound() {
	// Arrange
	newLocalityDoc := models.LocalityDoc{
		Id:       1,
		Locality: "Unknown City",
		Province: "Unknown Province",
		Country:  "Unknown Country",
	}

	// Mock la búsqueda de la provincia por nombre y país que devuelve un error de registro no encontrado
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `provinces`.`id`,`provinces`.`province`,`provinces`.`country_id` FROM `provinces` INNER JOIN countries ON countries.id = provinces.country_id WHERE provinces.province = ? AND countries.country = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocalityDoc.Province, newLocalityDoc.Country, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	// Act
	createdLocality, err := s.repo.CreateWithNames(newLocalityDoc)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrProvinceNotFound, err)
	s.Equal(models.LocalityDoc{}, createdLocality)
}

// Test CreateWithNames - Database Error en búsqueda de provincia
func (s *LocalityRepositoryTestSuite) TestCreateWithNames_DatabaseErrorFindingProvince() {
	// Arrange
	newLocalityDoc := models.LocalityDoc{
		Id:       1,
		Locality: "Buenos Aires",
		Province: "Buenos Aires",
		Country:  "Argentina",
	}

	// Mock la búsqueda de la provincia por nombre y país que devuelve un error de base de datos
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `provinces`.`id`,`provinces`.`province`,`provinces`.`country_id` FROM `provinces` INNER JOIN countries ON countries.id = provinces.country_id WHERE provinces.province = ? AND countries.country = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocalityDoc.Province, newLocalityDoc.Country, 1).
		WillReturnError(sql.ErrConnDone)

	// Act
	createdLocality, err := s.repo.CreateWithNames(newLocalityDoc)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.LocalityDoc{}, createdLocality)
}

// Test CreateWithNames - Duplicated Key Error
func (s *LocalityRepositoryTestSuite) TestCreateWithNames_DuplicatedKeyError() {
	// Arrange
	newLocalityDoc := models.LocalityDoc{
		Id:       1,
		Locality: "Buenos Aires",
		Province: "Buenos Aires",
		Country:  "Argentina",
	}

	// Mock la búsqueda de la provincia por nombre y país
	provinceRows := sqlmock.NewRows([]string{"id", "province", "country_id"}).
		AddRow(1, "Buenos Aires", 1)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `provinces`.`id`,`provinces`.`province`,`provinces`.`country_id` FROM `provinces` INNER JOIN countries ON countries.id = provinces.country_id WHERE provinces.province = ? AND countries.country = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocalityDoc.Province, newLocalityDoc.Country, 1).
		WillReturnRows(provinceRows)

	// Mock la creación de la localidad que devuelve un error de clave duplicada
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `localities` (`locality`,`province_id`,`id`) VALUES (?,?,?)")).
		WithArgs(newLocalityDoc.Locality, 1, newLocalityDoc.Id).
		WillReturnError(gorm.ErrDuplicatedKey)
	s.mock.ExpectRollback()

	// Act
	createdLocality, err := s.repo.CreateWithNames(newLocalityDoc)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrEntityAlreadyExists, err)
	s.Equal(models.LocalityDoc{}, createdLocality)
}

// Test CreateWithNames - Foreign Key Violation
func (s *LocalityRepositoryTestSuite) TestCreateWithNames_ForeignKeyViolation() {
	// Arrange
	newLocalityDoc := models.LocalityDoc{
		Id:       1,
		Locality: "Buenos Aires",
		Province: "Buenos Aires",
		Country:  "Argentina",
	}

	// Mock la búsqueda de la provincia por nombre y país
	provinceRows := sqlmock.NewRows([]string{"id", "province", "country_id"}).
		AddRow(1, "Buenos Aires", 1)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `provinces`.`id`,`provinces`.`province`,`provinces`.`country_id` FROM `provinces` INNER JOIN countries ON countries.id = provinces.country_id WHERE provinces.province = ? AND countries.country = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocalityDoc.Province, newLocalityDoc.Country, 1).
		WillReturnRows(provinceRows)

	// Mock la creación de la localidad que devuelve un error de violación de clave foránea
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `localities` (`locality`,`province_id`,`id`) VALUES (?,?,?)")).
		WithArgs(newLocalityDoc.Locality, 1, newLocalityDoc.Id).
		WillReturnError(gorm.ErrForeignKeyViolated)
	s.mock.ExpectRollback()

	// Act
	createdLocality, err := s.repo.CreateWithNames(newLocalityDoc)

	// Asserts
	s.Error(err)
	s.Equal(repository.ErrForeignKeyViolation, err)
	s.Equal(models.LocalityDoc{}, createdLocality)
}

// Test CreateWithNames - Other Database Error
func (s *LocalityRepositoryTestSuite) TestCreateWithNames_OtherDatabaseError() {
	// Arrange
	newLocalityDoc := models.LocalityDoc{
		Id:       1,
		Locality: "Buenos Aires",
		Province: "Buenos Aires",
		Country:  "Argentina",
	}

	// Mock la búsqueda de la provincia por nombre y país
	provinceRows := sqlmock.NewRows([]string{"id", "province", "country_id"}).
		AddRow(1, "Buenos Aires", 1)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `provinces`.`id`,`provinces`.`province`,`provinces`.`country_id` FROM `provinces` INNER JOIN countries ON countries.id = provinces.country_id WHERE provinces.province = ? AND countries.country = ? ORDER BY `provinces`.`id` LIMIT ?")).
		WithArgs(newLocalityDoc.Province, newLocalityDoc.Country, 1).
		WillReturnRows(provinceRows)

	// Mock la creación de la localidad que devuelve un error de base de datos
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `localities` (`locality`,`province_id`,`id`) VALUES (?,?,?)")).
		WithArgs(newLocalityDoc.Locality, 1, newLocalityDoc.Id).
		WillReturnError(sql.ErrConnDone)
	s.mock.ExpectRollback()

	// Act
	createdLocality, err := s.repo.CreateWithNames(newLocalityDoc)

	// Asserts
	s.Error(err)
	s.Equal(sql.ErrConnDone, err)
	s.Equal(models.LocalityDoc{}, createdLocality)
}

// Test for the find all success
func (s *LocalityRepositoryTestSuite) TestFindAllCarriers_Success() {
	// Arrange
	expectedLocalities := []models.LocalityCarrierCount{
		{
			LocalityID:    1,
			LocalityName:  "Locality_1",
			TotalCarriers: 3,
		},
		{
			LocalityID:    2,
			LocalityName:  "Locality_2",
			TotalCarriers: 0,
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
	expectedLocalities := []models.LocalityCarrierCount{
		{
			LocalityID:    1,
			LocalityName:  "Locality_1",
			TotalCarriers: 3,
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

func TestLocalityRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(LocalityRepositoryTestSuite))
}
