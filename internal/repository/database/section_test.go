package database

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/miloalej-dev/W17-G1-Bootcamp/internal/repository"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type SectionTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	repo *SectionRepository
}

func (s *SectionTestSuite) SetupSuite() {
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
	s.repo = NewSectionRepository(gormDB)

}

func (s *SectionTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	s.Require().NoError(err)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		TranslateError: true,
	})
	s.Require().NoError(err)

	s.db = gormDB
	s.mock = mock
	s.repo = NewSectionRepository(gormDB)

}

func (s *SectionTestSuite) TestFindAllSection_Success() {

	sections := []models.Section{
		{
			Id:                 1,
			SectionNumber:      "ab12",
			CurrentTemperature: 5.5,
			MinimumTemperature: 4.5,
			CurrentCapacity:    10,
			MinimumCapacity:    5,
			MaximumCapacity:    15,
			WarehouseId:        1,
			ProductTypeId:      1,
		},
		{
			Id:                 2,
			SectionNumber:      "ab13",
			CurrentTemperature: 5.5,
			MinimumTemperature: 4.5,
			CurrentCapacity:    10,
			MinimumCapacity:    5,
			MaximumCapacity:    15,
			WarehouseId:        2,
			ProductTypeId:      2,
		},
	}
	rows := s.mock.NewRows([]string{
		"id", "section_number", "current_temperature", "minimum_temperature",
		"current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id",
	}).AddRow(
		sections[0].Id, sections[0].SectionNumber, sections[0].CurrentTemperature,
		sections[0].MinimumTemperature, sections[0].CurrentCapacity,
		sections[0].MinimumCapacity, sections[0].MaximumCapacity,
		sections[0].WarehouseId, sections[0].ProductTypeId,
	).AddRow(
		sections[1].Id, sections[1].SectionNumber, sections[1].CurrentTemperature,
		sections[1].MinimumTemperature, sections[1].CurrentCapacity,
		sections[1].MinimumCapacity, sections[1].MaximumCapacity,
		sections[1].WarehouseId, sections[1].ProductTypeId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `sections`")).WillReturnRows(rows)

	// Act
	po, err := s.repo.FindAll()

	// Assert
	s.NoError(err)
	s.Len(po, 2)
	s.Equal(sections[0].SectionNumber, po[0].SectionNumber)
	s.Equal(sections[1].SectionNumber, po[1].SectionNumber)

}

func (s *SectionTestSuite) TestFindAllSection_DatabaseError() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `sections`")).WillReturnError(sql.ErrConnDone)
	// Act
	sections, err := s.repo.FindAll()
	// Assert
	s.Error(err)
	s.Equal([]models.Section{}, sections)
	s.Equal(sql.ErrConnDone, err)

}

func (s *SectionTestSuite) TestFindByIdSection_Success() {
	sections := models.Section{
		Id:                 1,
		SectionNumber:      "ab12",
		CurrentTemperature: 5.5,
		MinimumTemperature: 4.5,
		CurrentCapacity:    10,
		MinimumCapacity:    5,
		MaximumCapacity:    15,
		WarehouseId:        1,
		ProductTypeId:      1,
	}
	rows := s.mock.NewRows([]string{
		"id", "section_number", "current_temperature", "minimum_temperature",
		"current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id",
	}).AddRow(
		sections.Id, sections.SectionNumber, sections.CurrentTemperature,
		sections.MinimumTemperature, sections.CurrentCapacity,
		sections.MinimumCapacity, sections.MaximumCapacity,
		sections.WarehouseId, sections.ProductTypeId,
	)
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `sections` WHERE `sections`.`id` = ? ORDER BY `sections`.`id` LIMIT ?")).
		WithArgs(1, 1).WillReturnRows(rows)
	// Act
	sec, err := s.repo.FindById(1)

	// Assert
	s.NoError(err)
	s.Equal(sections.Id, sec.Id)
	s.Equal(sections.SectionNumber, sec.SectionNumber)
}

func (s *SectionTestSuite) TestFindById_NotFound() {
	// Arrange
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `sections` WHERE `sections`.`id` = ? ORDER BY `sections`.`id` LIMIT ?",
	)).WithArgs(999, 1).WillReturnError(repository.ErrEntityNotFound)
	// Act
	sec, err := s.repo.FindById(999)

	// Assert
	s.Error(err)
	s.Equal(repository.ErrEntityNotFound, err)
	s.Equal(models.Section{}, sec)
}

func (s *SectionTestSuite) TestCreate_Success() {
	sections := models.Section{
		SectionNumber:      "ab12",
		CurrentTemperature: 5.5,
		MinimumTemperature: 4.5,
		CurrentCapacity:    10,
		MinimumCapacity:    5,
		MaximumCapacity:    15,
		WarehouseId:        1,
		ProductTypeId:      1,
	}
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `sections` (`section_number`,`current_temperature`,`minimum_temperature`,`current_capacity`,`minimum_capacity`,`maximum_capacity`,`warehouse_id`,`product_type_id`) VALUES (?,?,?,?,?,?,?,?)",
	)).WithArgs(
		sections.SectionNumber,
		sections.CurrentTemperature,
		sections.MinimumTemperature,
		sections.CurrentCapacity,
		sections.MinimumCapacity,
		sections.MaximumCapacity,
		sections.WarehouseId,
		sections.ProductTypeId,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	sec, err := s.repo.Create(sections)

	s.NoError(err)
	s.Equal(sections.SectionNumber, sec.SectionNumber)

}

func (s *SectionTestSuite) TestCreateSection_ForeignKeyViolated() {
	// Arrange
	sections := models.Section{
		SectionNumber:      "ab12",
		CurrentTemperature: 5.5,
		MinimumTemperature: 4.5,
		CurrentCapacity:    10,
		MinimumCapacity:    5,
		MaximumCapacity:    15,
		WarehouseId:        99,
		ProductTypeId:      1,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `sections` (`section_number`,`current_temperature`,`minimum_temperature`,`current_capacity`,`minimum_capacity`,`maximum_capacity`,`warehouse_id`,`product_type_id`) VALUES (?,?,?,?,?,?,?,?)",
	)).WithArgs(
		sections.SectionNumber,
		sections.CurrentTemperature,
		sections.MinimumTemperature,
		sections.CurrentCapacity,
		sections.MinimumCapacity,
		sections.MaximumCapacity,
		sections.WarehouseId,
		sections.ProductTypeId,
	).WillReturnError(repository.ErrForeignKeyViolation)
	s.mock.ExpectRollback()

	// Act
	sec, err := s.repo.Create(sections)

	// Assert
	s.Error(err)
	s.Equal(models.Section{}, sec)

}

func (s *SectionTestSuite) TestUpdateSection_Success() {
	section := models.Section{
		Id:                 1, // 👈 obligatorio para el Save
		SectionNumber:      "ab12",
		CurrentTemperature: 6.0,
		MinimumTemperature: 4.5,
		CurrentCapacity:    12,
		MinimumCapacity:    5,
		MaximumCapacity:    15,
		WarehouseId:        2,
		ProductTypeId:      2,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `sections` SET `section_number`=?,`current_temperature`=?,`minimum_temperature`=?,`current_capacity`=?,`minimum_capacity`=?,`maximum_capacity`=?,`warehouse_id`=?,`product_type_id`=? WHERE `id` = ?")).
		WithArgs(
			section.SectionNumber,
			section.CurrentTemperature,
			section.MinimumTemperature,
			section.CurrentCapacity,
			section.MinimumCapacity,
			section.MaximumCapacity,
			section.WarehouseId,
			section.ProductTypeId,
			section.Id,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	updated, err := s.repo.Update(section)

	s.NoError(err)
	s.Equal(section.SectionNumber, updated.SectionNumber)
	s.Equal(section.CurrentCapacity, updated.CurrentCapacity)
}

func (s *SectionTestSuite) TestUpdateSection_Error() {
	section := models.Section{
		Id:                 1,
		SectionNumber:      "ab12",
		CurrentTemperature: 6.0,
		MinimumTemperature: 4.5,
		CurrentCapacity:    12,
		MinimumCapacity:    5,
		MaximumCapacity:    15,
		WarehouseId:        2,
		ProductTypeId:      2,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `sections` SET `section_number`=?,`current_temperature`=?,`minimum_temperature`=?,`current_capacity`=?,`minimum_capacity`=?,`maximum_capacity`=?,`warehouse_id`=?,`product_type_id`=? WHERE `id` = ?")).
		WithArgs(
			section.SectionNumber,
			section.CurrentTemperature,
			section.MinimumTemperature,
			section.CurrentCapacity,
			section.MinimumCapacity,
			section.MaximumCapacity,
			section.WarehouseId,
			section.ProductTypeId,
			section.Id,
		).
		WillReturnError(errors.New("database update failed"))
	s.mock.ExpectRollback()

	updated, err := s.repo.Update(section)

	s.Error(err)
	s.EqualError(err, "database update failed")
	s.Equal(models.Section{}, updated)
}

func (s *SectionTestSuite) TestPartialUpdateSection_Success() {
	id := 1
	fields := map[string]interface{}{
		"section_number":      "XY-99",
		"current_temperature": 7.5,
		"minimum_temperature": 3.2,
		"current_capacity":    8.0,
		"minimum_capacity":    4.0,
		"maximum_capacity":    10.0,
		"warehouses_id":       1.0, // corregido nombre y tipo int
		"product_type_id":     2.0, // corregido nombre y tipo int
	}

	// Correcta expectación del SELECT con LIMIT ? (parámetro)
	s.mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `sections` WHERE `sections`.`id` = ? ORDER BY `sections`.`id` LIMIT ?"),
	).
		WithArgs(id, 1).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "section_number", "current_temperature", "minimum_temperature",
			"current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id",
		}).AddRow(
			id, "AB-01", 5.5, 3.0, 6, 3, 12, 1, 2, // valores originales, no importa mucho
		))

	// Mock para el UPDATE
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `sections` SET `section_number`=?,`current_temperature`=?,`minimum_temperature`=?,`current_capacity`=?,`minimum_capacity`=?,`maximum_capacity`=?,`warehouse_id`=?,`product_type_id`=? WHERE `id` = ?",
	)).WithArgs(
		"XY-99", 7.5, 3.2, 8, 4, 10, 1, 2, id,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	// (Opcional, para devolver la nueva sección)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `sections` WHERE `sections`.`id` = ? ORDER BY `sections`.`id` LIMIT ?",
	)).WithArgs(id, 1).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "section_number", "current_temperature", "minimum_temperature",
			"current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id",
		}).AddRow(
			id, "XY-99", 7.5, 3.2, 8, 4, 10, 1, 2,
		))

	updated, err := s.repo.PartialUpdate(id, fields)

	s.NoError(err)
	s.Equal("XY-99", updated.SectionNumber)
	s.Equal(7.5, updated.CurrentTemperature)
	s.Equal(3.2, updated.MinimumTemperature)
	s.Equal(8, updated.CurrentCapacity)
	s.Equal(4, updated.MinimumCapacity)
	s.Equal(10, updated.MaximumCapacity)
	s.Equal(1, updated.WarehouseId)
	s.Equal(2, updated.ProductTypeId)
}

func (s *SectionTestSuite) TestPartialUpdate_FailOnFirstQuery() {
	id := 99
	// DEBE ser la ÚNICA ExpectQuery activa en este test
	s.mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `sections` WHERE `sections`.`id` = ? ORDER BY `sections`.`id` LIMIT ?"),
	).
		WithArgs(sqlmock.AnyArg(), 1).
		WillReturnError(errors.New("db error on select"))

	updated, err := s.repo.PartialUpdate(id, map[string]interface{}{})

	s.Error(err)
	s.EqualError(err, "db error on select")
	s.Equal(models.Section{}, updated)

	// Verifica que no queden expectations colgadas (buena práctica)
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *SectionTestSuite) TestPartialUpdate_NoRowsFound() {
	id := 42
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `sections` WHERE `sections`.`id` = ? ORDER BY `sections`.`id` LIMIT 1",
	)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{})) // sin filas

	updated, err := s.repo.PartialUpdate(id, map[string]interface{}{})

	s.Error(err)
	s.Equal(models.Section{}, updated)
}

func (s *SectionTestSuite) TestPartialUpdate_FailOnSave() {
	id := 1
	fields := map[string]interface{}{
		"section_number": "FAIL-01",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `sections` WHERE `sections`.`id` = ? ORDER BY `sections`.`id` LIMIT ?",
	)).
		WithArgs(id, 1). // ← Corregido aquí
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "section_number", "current_temperature", "minimum_temperature",
			"current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id",
		}).AddRow(id, "A", 1.1, 1.1, 1, 1, 1, 1, 1))

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `sections` SET `section_number`=?,`current_temperature`=?,`minimum_temperature`=?,`current_capacity`=?,`minimum_capacity`=?,`maximum_capacity`=?,`warehouse_id`=?,`product_type_id`=? WHERE `id` = ?",
	)).WithArgs(
		"FAIL-01", 1.1, 1.1, 1, 1, 1, 1, 1, id,
	).WillReturnError(errors.New("save failed"))
	s.mock.ExpectRollback()

	updated, err := s.repo.PartialUpdate(id, fields)

	s.Error(err)
	s.EqualError(err, "save failed")
	s.Equal(models.Section{}, updated)
}

func (s *SectionTestSuite) TestDelete_Success() {
	id := 1

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		"DELETE FROM `sections` WHERE `sections`.`id` = ?",
	)).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.repo.Delete(id)
	s.NoError(err)
}

func (s *SectionTestSuite) TestFindSectionReport_SectionNotFound() {
	id := 5
	s.mock.ExpectQuery(regexp.QuoteMeta( // lo que hace GORM en First(&section, id)
		"SELECT * FROM `sections` WHERE `sections`.`id` = ? ORDER BY `sections`.`id` LIMIT ?"),
	).WithArgs(id, 1).
		WillReturnError(errors.New("some db error"))

	_, err := s.repo.FindSectionReport(id)
	s.ErrorIs(err, repository.ErrSectionNotFound)
}

func (s *SectionTestSuite) TestFindSectionReport_QueryError() {
	id := 5
	// Mock First ok (existe la sección)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `sections` WHERE `sections`.`id` = ? ORDER BY `sections`.`id` LIMIT ?"),
	).WithArgs(id, 1).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "section_number", "current_temperature", "minimum_temperature",
			"current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id",
		}).AddRow(
			id, "SEC123", 4.0, 3.0, 8, 2, 10, 1, 1,
		))

	// Mock error en el Scan de la segunda query (JOIN, GROUP BY, etc.)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT s.id as section_id, s.section_number, COUNT(p.id) as products_count FROM sections as s INNER JOIN product_batches as p ON s.id = p.section_id WHERE s.id = ? GROUP BY s.id, s.section_number")).
		WithArgs(id).
		WillReturnError(errors.New("query error"))

	_, err := s.repo.FindSectionReport(id)
	s.Error(err)
	s.Contains(err.Error(), "query error")
}

func (s *SectionTestSuite) TestFindSectionReport_Success() {
	id := 5
	// Mock Find sección
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `sections` WHERE `sections`.`id` = ? ORDER BY `sections`.`id` LIMIT ?"),
	).WithArgs(id, 1).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "section_number", "current_temperature", "minimum_temperature",
			"current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id",
		}).AddRow(
			id, "SEC123", 4.0, 3.0, 8, 2, 10, 1, 1,
		))
	// Mock JOIN query exitosa
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT s.id as section_id, s.section_number, COUNT(p.id) as products_count FROM sections as s INNER JOIN product_batches as p ON s.id = p.section_id WHERE s.id = ? GROUP BY s.id, s.section_number")).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"section_id", "section_number", "products_count"}).
			AddRow(id, "SEC123", 3), // <--- products_count 3!
		)

	rep, err := s.repo.FindSectionReport(id)
	s.NoError(err)
	s.Equal(id, rep.SectionId)
	s.Equal("SEC123", rep.SectionNumber)
	s.Equal(3, rep.ProductsCount)
}

func (s *SectionTestSuite) TestFindAllSectionReports_Success() {
	// Simula dos resultados para la consulta
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT s.id as section_id, s.section_number, COUNT(p.id) as products_count FROM sections as s INNER JOIN product_batches as p ON s.id = p.section_id GROUP BY s.id, s.section_number")).
		WillReturnRows(sqlmock.NewRows([]string{"section_id", "section_number", "products_count"}).
			AddRow(1, "S-001", 4).
			AddRow(2, "S-002", 2),
		)

	result, err := s.repo.FindAllSectionReports()

	s.NoError(err)
	s.Len(result, 2)
	s.Equal(1, result[0].SectionId)
	s.Equal("S-001", result[0].SectionNumber)
	s.Equal(4, result[0].ProductsCount)
	s.Equal(2, result[1].SectionId)
	s.Equal("S-002", result[1].SectionNumber)
	s.Equal(2, result[1].ProductsCount)
}

func (s *SectionTestSuite) TestFindAllSectionReports_DBError() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT s.id as section_id, s.section_number, COUNT(p.id) as products_count FROM sections as s INNER JOIN product_batches as p ON s.id = p.section_id GROUP BY s.id, s.section_number")).
		WillReturnError(errors.New("db explosion"))

	result, err := s.repo.FindAllSectionReports()

	s.Error(err)
	s.EqualError(err, "db explosion")
	s.Nil(result)
}

func (s *SectionTestSuite) TestFindAllSectionReports_Empty() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT s.id as section_id, s.section_number, COUNT(p.id) as products_count FROM sections as s INNER JOIN product_batches as p ON s.id = p.section_id GROUP BY s.id, s.section_number")).
		WillReturnRows(sqlmock.NewRows([]string{"section_id", "section_number", "products_count"}))

	result, err := s.repo.FindAllSectionReports()

	s.NoError(err)
	s.Equal(0, len(result))
}

// Run the test suite
func TestSectionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(SectionTestSuite))
}
