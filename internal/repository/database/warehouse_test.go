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

type WarehouseRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo *WarehouseDB
}

func (s *WarehouseRepositoryTestSuite) SetupSuite() {
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
	s.repo = NewWarehouseDB(gormDB)
}

// Test for the find all success
func (s *WarehouseRepositoryTestSuite) TestFindAll_Success() {
	// Arrange
	expectedWarehouses := []models.Warehouse{
		{
			Id:                 1,
			WarehouseCode:      "CID#01",
			Address:            "Boulevard",
			Telephone:          "123-456789",
			MinimumCapacity:    100,
			MinimumTemperature: 19,
			LocalityId:         1,
		},
		{
			Id:                 2,
			WarehouseCode:      "CID#02",
			Address:            "Plaza",
			Telephone:          "222-456789",
			MinimumCapacity:    200,
			MinimumTemperature: 5,
			LocalityId:         1,
		},
	}

	columns := []string{
		"id",
		"warehouse_code",
		"address",
		"telephone",
		"minimum_capacity",
		"minimum_temperature",
		"locality_id",
	}

	rows := s.mock.NewRows(columns)
	for _, w := range expectedWarehouses {
		rows = rows.AddRow(
			w.Id,
			w.WarehouseCode,
			w.Address,
			w.Telephone,
			w.MinimumCapacity,
			w.MinimumTemperature,
			w.LocalityId,
		)
	}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `warehouses`")).WillReturnRows(rows)

	// Act
	warehouses, err := s.repo.FindAll()

	// Assert
	s.NoError(err)
	s.Len(warehouses, 2)
	s.Equal(expectedWarehouses, warehouses)
	s.mock.ExpectationsWereMet()
}

// Error on retireve all
func (s *WarehouseRepositoryTestSuite) TestFindAll_DatabaseError() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `warehouses`")).WillReturnError(sql.ErrConnDone)

	// Act
	warehouses, err := s.repo.FindAll()

	// Assert
	s.Error(err)
	s.Nil(warehouses)
	s.Equal(sql.ErrConnDone, err)
	s.mock.ExpectationsWereMet()
}

// Find by Id returns the specified warehouse
func (s *WarehouseRepositoryTestSuite) TestFindById_Success() {
	// Arrange
	expectedWarehouse := models.Warehouse{
		Id:                 1,
		WarehouseCode:      "CID#01",
		Address:            "Boulevard",
		Telephone:          "123-456789",
		MinimumCapacity:    100,
		MinimumTemperature: 19,
		LocalityId:         1,
	}

	fields := []string{
		"id",
		"warehouse_code",
		"address",
		"telephone",
		"minimum_capacity",
		"minimum_temperature",
		"locality_id",
	}

	rows := s.mock.NewRows(fields).
		AddRow(
			expectedWarehouse.Id,
			expectedWarehouse.WarehouseCode,
			expectedWarehouse.Address,
			expectedWarehouse.Telephone,
			expectedWarehouse.MinimumCapacity,
			expectedWarehouse.MinimumTemperature,
			expectedWarehouse.LocalityId,
		)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `warehouses` WHERE `warehouses`.`id` = ? ORDER BY `warehouses`.`id` LIMIT ?")).
		WithArgs(1, 1).WillReturnRows(rows)

	// Act
	warehouse, err := s.repo.FindById(1)

	// Assert
	s.NoError(err)
	s.Equal(expectedWarehouse, warehouse)
	s.mock.ExpectationsWereMet()
}

// Find by a non existent Id returns nothing
func (s *WarehouseRepositoryTestSuite) TestFindById_NotFound() {
	id := 42 // Some ID not in the database

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `warehouses` WHERE `warehouses`.`id` = ? ORDER BY `warehouses`.`id` LIMIT ?",
	)).
		WithArgs(id, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	result, err := s.repo.FindById(id)

	s.Error(err)
	s.Equal(gorm.ErrRecordNotFound, err)
	s.Equal(models.Warehouse{}, result)
}

func (s *WarehouseRepositoryTestSuite) TestCreate_Success() {
	// Arrange
	newWarehouse := models.Warehouse{
		WarehouseCode:      "CID#01",
		Address:            "Boulevard",
		Telephone:          "123-456789",
		MinimumCapacity:    100,
		MinimumTemperature: 19,
		LocalityId:         1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `warehouses` (`warehouse_code`,`address`,`telephone`,`minimum_capacity`,`minimum_temperature`,`locality_id`) VALUES (?,?,?,?,?,?)")).
		WithArgs(newWarehouse.WarehouseCode, newWarehouse.Address, newWarehouse.Telephone, newWarehouse.MinimumCapacity, newWarehouse.MinimumTemperature, newWarehouse.LocalityId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	createdWarehouse, err := s.repo.Create(newWarehouse)

	// Assert
	s.NoError(err)
	s.Equal(newWarehouse.WarehouseCode, createdWarehouse.WarehouseCode)
	s.Equal(newWarehouse.Address, createdWarehouse.Address)
	s.Equal(newWarehouse.Telephone, createdWarehouse.Telephone)
	s.Equal(newWarehouse.MinimumCapacity, createdWarehouse.MinimumCapacity)
	s.Equal(newWarehouse.MinimumTemperature, createdWarehouse.MinimumTemperature)
	s.Equal(newWarehouse.LocalityId, createdWarehouse.LocalityId)
	s.mock.ExpectationsWereMet()
}

func (s *WarehouseRepositoryTestSuite) TestCreate_ForeignKeyViolated() {
	// Arrange
	newWarehouse := models.Warehouse{
		//Id:					1,
		WarehouseCode:      "CID#01",
		Address:            "Boulevard",
		Telephone:          "123-456789",
		MinimumCapacity:    100,
		MinimumTemperature: 19,
		LocalityId:         111,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `warehouses` (`warehouse_code`,`address`,`telephone`,`minimum_capacity`,`minimum_temperature`,`locality_id`) VALUES (?,?,?,?,?,?)")).
		WithArgs(newWarehouse.WarehouseCode, newWarehouse.Address, newWarehouse.Telephone, newWarehouse.MinimumCapacity, newWarehouse.MinimumTemperature, newWarehouse.LocalityId).
		WillReturnError(gorm.ErrForeignKeyViolated)
	s.mock.ExpectRollback()

	// Act
	createdWarehouse, err := s.repo.Create(newWarehouse)

	// Assert
	s.Error(err)
	s.Equal(err, repository.ErrLocalityNotFound)
	s.Equal(models.Warehouse{}, createdWarehouse)
	s.mock.ExpectationsWereMet()
}

func (s *WarehouseRepositoryTestSuite) TestCreate_AnotherGormError() {
	// Arrange
	newWarehouse := models.Warehouse{
		//Id:					1,
		WarehouseCode:      "CID#01",
		Address:            "Boulevard",
		Telephone:          "123-456789",
		MinimumCapacity:    100,
		MinimumTemperature: 19,
		LocalityId:         111,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `warehouses` (`warehouse_code`,`address`,`telephone`,`minimum_capacity`,`minimum_temperature`,`locality_id`) VALUES (?,?,?,?,?,?)")).
		WithArgs(newWarehouse.WarehouseCode, newWarehouse.Address, newWarehouse.Telephone, newWarehouse.MinimumCapacity, newWarehouse.MinimumTemperature, newWarehouse.LocalityId).
		WillReturnError(gorm.ErrInvalidValue)
	s.mock.ExpectRollback()

	// Act
	createdWarehouse, err := s.repo.Create(newWarehouse)

	// Assert
	s.Error(err)
	s.Equal(err, gorm.ErrInvalidValue)
	s.Equal(models.Warehouse{}, createdWarehouse)
	s.mock.ExpectationsWereMet()
}

func (s *WarehouseRepositoryTestSuite) TestUpdate_Success() {
	// Arrange
	existingWarehouse := models.Warehouse{
		Id:                 1,
		WarehouseCode:      "CID#01",
		Address:            "Boulevard",
		Telephone:          "123-456789",
		MinimumCapacity:    100,
		MinimumTemperature: 19,
		LocalityId:         1,
	}

	// Match the UPDATE SQL GORM will use
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `warehouses` SET `warehouse_code`=?,`address`=?,`telephone`=?,`minimum_capacity`=?,`minimum_temperature`=?,`locality_id`=? WHERE `id` = ?",
	)).
		WithArgs(
			existingWarehouse.WarehouseCode,
			existingWarehouse.Address,
			existingWarehouse.Telephone,
			existingWarehouse.MinimumCapacity,
			existingWarehouse.MinimumTemperature,
			existingWarehouse.LocalityId,
			existingWarehouse.Id,
		).
		WillReturnResult(sqlmock.NewResult(1, 1)) // result: row affected

	s.mock.ExpectCommit()

	// Act
	updatedWarehouse, err := s.repo.Update(existingWarehouse)

	// Assert
	s.NoError(err)
	s.Equal(existingWarehouse, updatedWarehouse)

	// Verify all expectations were met
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *WarehouseRepositoryTestSuite) TestUpdate_ForeignKeyViolation() {
	// Arrange
	existingWarehouse := models.Warehouse{
		Id:                 1,
		WarehouseCode:      "CID#01",
		Address:            "Boulevard",
		Telephone:          "123-456789",
		MinimumCapacity:    100,
		MinimumTemperature: 19,
		LocalityId:         1,
	}

	// Match the UPDATE SQL GORM will use
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `warehouses` SET `warehouse_code`=?,`address`=?,`telephone`=?,`minimum_capacity`=?,`minimum_temperature`=?,`locality_id`=? WHERE `id` = ?",
	)).
		WithArgs(
			existingWarehouse.WarehouseCode,
			existingWarehouse.Address,
			existingWarehouse.Telephone,
			existingWarehouse.MinimumCapacity,
			existingWarehouse.MinimumTemperature,
			existingWarehouse.LocalityId,
			existingWarehouse.Id,
		).
		WillReturnError(gorm.ErrForeignKeyViolated)
	s.mock.ExpectRollback()

	// Act
	updatedWarehouse, err := s.repo.Update(existingWarehouse)

	// Assert
	s.Error(err)
	s.Equal(models.Warehouse{}, updatedWarehouse)

	// Verify all expectations were met
	err = s.mock.ExpectationsWereMet()
	s.NoError(err)
}

func (s *WarehouseRepositoryTestSuite) TestPartialUpdate_Success() {
	// Arrange
	id := 1
	originalWarehouse := models.Warehouse{
		Id:                 id,
		WarehouseCode:      "OLD-CODE",
		Address:            "Old Address",
		Telephone:          "555-1111",
		MinimumCapacity:    10,
		MinimumTemperature: 20,
		LocalityId:         1,
	}

	// These are the fields we'll update
	fields := map[string]interface{}{
		"code":               "NEW-CODE",
		"address":            "New Address",
		"telephone":		  "123-321",
		"minimum_capacity":   float64(99), // float64 is common from JSON/mapstructure
		"minimum_temperature": float64(5),
		"locality_id":			float64(2),
	}

	columns := []string{
		"id", "warehouse_code", "address", "telephone",
		"minimum_capacity", "minimum_temperature", "locality_id",
	}

	// Mock the SELECT FROM warehouses WHERE id=?
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `warehouses` WHERE `warehouses`.`id` = ? ORDER BY `warehouses`.`id` LIMIT ?",
	)).WithArgs(id, 1).
		WillReturnRows(s.mock.NewRows(columns).
			AddRow(
				originalWarehouse.Id,
				originalWarehouse.WarehouseCode,
				originalWarehouse.Address,
				originalWarehouse.Telephone,
				originalWarehouse.MinimumCapacity,
				originalWarehouse.MinimumTemperature,
				originalWarehouse.LocalityId,
			),
		)

	// The updated struct as expected in DB after field change
	expectedUpdatedWarehouse := models.Warehouse{
		Id:                 id,
		WarehouseCode:      "NEW-CODE",
		Address:            "New Address",
		Telephone:          "123-321",
		MinimumCapacity:    99,
		MinimumTemperature: 5,
		LocalityId:         2,
	}

	// Mock the UPDATE. GORM uses UPDATE with SET ... WHERE id = ?
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `warehouses` SET `warehouse_code`=?,`address`=?,`telephone`=?,`minimum_capacity`=?,`minimum_temperature`=?,`locality_id`=? WHERE `id` = ?")).
		WithArgs(
			expectedUpdatedWarehouse.WarehouseCode,
			expectedUpdatedWarehouse.Address,
			expectedUpdatedWarehouse.Telephone,
			expectedUpdatedWarehouse.MinimumCapacity,
			expectedUpdatedWarehouse.MinimumTemperature,
			expectedUpdatedWarehouse.LocalityId,
			expectedUpdatedWarehouse.Id,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// Act
	result, err := s.repo.PartialUpdate(id, fields)

	// Assert
	s.NoError(err)
	s.Equal(expectedUpdatedWarehouse, result)
}

func (s *WarehouseRepositoryTestSuite) TestPartialUpdate_DatabaseErrorOnFirst() {
	// Arrange
	id := 1
	fields := map[string]any{
		"code":     		  "CID#01",
		"address":            "Boulevard plaza",
		"telephone":          "123-456789",
		"minimum_capacity":    float64(200),
		"minimum_temperature": float64(15),
		"locality_id":         float64(1),
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `warehouses` WHERE `warehouses`.`id` = ? ORDER BY `warehouses`.`id` LIMIT ?",
	)).WithArgs(id, 1).
		WillReturnError(gorm.ErrInvalidDB)

	// Act
	updatedWarehouse, err := s.repo.PartialUpdate(id, fields)

	// Assert
	s.Error(err)
	s.ErrorIs(err, gorm.ErrInvalidDB)
	s.Equal(models.Warehouse{}, updatedWarehouse)
	s.mock.ExpectationsWereMet()
}

func (s *WarehouseRepositoryTestSuite) TestPartialUpdate_ErrorNotFound() {
	// Arrange
	id := 1111
	fields := map[string]any{
		"code":     		  "CID#01",
		"address":            "Boulevard plaza",
		"telephone":          "123-456789",
		"minimum_capacity":    float64(200),
		"minimum_temperature": float64(15),
		"locality_id":         float64(1),
	}

	columns := []string{
		"id", "warehouse_code", "address", "telephone",
		"minimum_capacity", "minimum_temperature", "locality_id",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `warehouses` WHERE `warehouses`.`id` = ? ORDER BY `warehouses`.`id` LIMIT ?",
	)).WithArgs(id, 1).WillReturnRows(s.mock.NewRows(columns))

	// Act
	updatedWarehouse, err := s.repo.PartialUpdate(id, fields)

	// Assert
	s.Error(err)
	s.ErrorIs(err, repository.ErrEntityNotFound)
	s.Equal(models.Warehouse{}, updatedWarehouse)
	s.mock.ExpectationsWereMet()
}

func (s *WarehouseRepositoryTestSuite) TestPartialUpdate_ErrorOnSave() {
	// Arrange
	id := 1
	originalWarehouse := models.Warehouse{
		Id:                 id,
		WarehouseCode:      "OLD-CODE",
		Address:            "Old Address",
		Telephone:          "555-1111",
		MinimumCapacity:    10,
		MinimumTemperature: 20,
		LocalityId:         1,
	}

	// These are the fields we'll update
	fields := map[string]interface{}{
		"code":               "NEW-CODE",
		"address":            "New Address",
		"telephone":		  "123-321",
		"minimum_capacity":   float64(99), // float64 is common from JSON/mapstructure
		"minimum_temperature": float64(5),
		"locality_id":			float64(2),
	}

	columns := []string{
		"id", "warehouse_code", "address", "telephone",
		"minimum_capacity", "minimum_temperature", "locality_id",
	}

	// Mock the SELECT FROM warehouses WHERE id=?
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `warehouses` WHERE `warehouses`.`id` = ? ORDER BY `warehouses`.`id` LIMIT ?",
	)).WithArgs(id, 1).
		WillReturnRows(s.mock.NewRows(columns).
			AddRow(
				originalWarehouse.Id,
				originalWarehouse.WarehouseCode,
				originalWarehouse.Address,
				originalWarehouse.Telephone,
				originalWarehouse.MinimumCapacity,
				originalWarehouse.MinimumTemperature,
				originalWarehouse.LocalityId,
			),
		)

	// The updated struct as expected in DB after field change
	expectedUpdatedWarehouse := models.Warehouse{
		Id:                 id,
		WarehouseCode:      "NEW-CODE",
		Address:            "New Address",
		Telephone:          "123-321",
		MinimumCapacity:    99,
		MinimumTemperature: 5,
		LocalityId:         2,
	}

	// Mock the UPDATE. GORM uses UPDATE with SET ... WHERE id = ?
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `warehouses` SET `warehouse_code`=?,`address`=?,`telephone`=?,`minimum_capacity`=?,`minimum_temperature`=?,`locality_id`=? WHERE `id` = ?")).
		WithArgs(
			expectedUpdatedWarehouse.WarehouseCode,
			expectedUpdatedWarehouse.Address,
			expectedUpdatedWarehouse.Telephone,
			expectedUpdatedWarehouse.MinimumCapacity,
			expectedUpdatedWarehouse.MinimumTemperature,
			expectedUpdatedWarehouse.LocalityId,
			expectedUpdatedWarehouse.Id,
		).
		WillReturnError(gorm.ErrInvalidValue)
	s.mock.ExpectRollback()

	// Act
	result, err := s.repo.PartialUpdate(id, fields)

	// Assert
	s.Error(err)
	s.Equal(models.Warehouse{}, result)
}

func (s *WarehouseRepositoryTestSuite) TestDelete_Success() {
	// Arrange
	warehouseID := 1
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `warehouses` WHERE `warehouses`.`id` = ?")).
		WithArgs(warehouseID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	// Act
	err := s.repo.Delete(warehouseID)
	// Assert
	s.NoError(err)
	s.mock.ExpectationsWereMet()
}

func (s *WarehouseRepositoryTestSuite) TestDelete_NotFound() {
	// Arrange
	warehouseID := 999

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `warehouses` WHERE `warehouses`.`id` = ?")).
		WithArgs(warehouseID).
		WillReturnResult(sqlmock.NewResult(1, 0)) // 0 rows affected
	s.mock.ExpectCommit()

	// Act
	err := s.repo.Delete(warehouseID)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.mock.ExpectationsWereMet()
}

// Run the test suite
func TestWarehouseRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(WarehouseRepositoryTestSuite))
}
