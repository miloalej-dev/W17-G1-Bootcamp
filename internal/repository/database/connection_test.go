package database

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"testing"
)

type ConnectionTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
}

func (s *ConnectionTestSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		s.T().Fatal(err)
	}

	// Expect the ping that GORM will perform during connection initialization
	s.mock.ExpectPing()

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
}

// Test configure function with valid environment variables
func (s *ConnectionTestSuite) TestConfigure_Success() {
	// Arrange
	user := "testuser"
	password := "testpass"
	host := "testhost"
	port := "33060"
	database := "testdb"

	// Set up environment variables
	s.Require().NoError(os.Setenv("DB_USER", user))
	s.Require().NoError(os.Setenv("DB_PASSWORD", password))
	s.Require().NoError(os.Setenv("DB_HOST", host))
	s.Require().NoError(os.Setenv("DB_PORT", port))
	s.Require().NoError(os.Setenv("DB_NAME", database))

	defer func() {
		_ = os.Unsetenv("DB_USER")
		_ = os.Unsetenv("DB_PASSWORD")
		_ = os.Unsetenv("DB_HOST")
		_ = os.Unsetenv("DB_PORT")
		_ = os.Unsetenv("DB_NAME")
	}()

	config, err := configure()

	s.NoError(err)
	s.NotNil(config)
	s.Equal(user, config.user)
	s.Equal(password, config.password)
	s.Equal(host, config.host)
	s.Equal(port, config.port)
	s.Equal(database, config.database)
}

// Test configure function with default values
func (s *ConnectionTestSuite) TestConfigure_WithDefaults() {
	// Arrange
	user := "testuser"
	password := "testpass"

	// Set only required environment variables
	s.Require().NoError(os.Setenv("DB_USER", user))
	s.Require().NoError(os.Setenv("DB_PASSWORD", password))

	defer func() {
		_ = os.Unsetenv("DB_USER")
		_ = os.Unsetenv("DB_PASSWORD")
	}()

	config, err := configure()

	s.Require().NoError(err)
	s.NotNil(config)
	s.Equal(user, config.user)
	s.Equal(password, config.password)
	s.Equal("localhost", config.host)   // default value
	s.Equal("3306", config.port)        // default value
	s.Equal("frescos", config.database) // default value
}

// Test configure function with missing user
func (s *ConnectionTestSuite) TestConfigure_MissingUser() {
	s.Require().NoError(os.Setenv("DB_PASSWORD", "testpass"))
	defer func() { _ = os.Unsetenv("DB_PASSWORD") }()

	config, err := configure()

	s.Error(err)
	s.Nil(config)
	s.Contains(err.Error(), "user environment variable is required")
}

// Test configure function with missing password
func (s *ConnectionTestSuite) TestConfigure_MissingPassword() {
	s.Require().NoError(os.Setenv("DB_USER", "testuser"))
	defer func() { _ = os.Unsetenv("DB_USER") }()

	config, err := configure()

	s.Error(err)
	s.Nil(config)
	s.Contains(err.Error(), "password environment variable is required")
}

// Test buildDSN function
func (s *ConnectionTestSuite) TestBuildDSN() {
	config := &configuration{
		user:     "testuser",
		password: "testpass",
		host:     "testhost",
		port:     "3307",
		database: "testdb",
	}

	expectedDSN := "testuser:testpass@tcp(testhost:3307)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	actualDSN := buildDSN(config)

	s.Equal(expectedDSN, actualDSN)
}

// Test configurePool function
func (s *ConnectionTestSuite) TestConfigurePool_Success() {
	err := configurePool(s.db)
	s.NoError(err)
	// Verify pool settings
	sqlDB, err := s.db.DB()
	s.NoError(err)

	stats := sqlDB.Stats()
	s.Equal(100, stats.MaxOpenConnections)
}

// Test NewConnection function with mocked environment
func (s *ConnectionTestSuite) TestNewConnection_WithValidEnv() {
	// Skip this test in CI or when actual DB connection is not available
	if testing.Short() {
		s.T().Skip("Skipping integration test in short mode")
	}

	// Set up environment variables for the test
	s.Require().NoError(os.Setenv("DB_USER", "root"))
	s.Require().NoError(os.Setenv("DB_PASSWORD", "password"))
	s.Require().NoError(os.Setenv("DB_HOST", "localhost"))
	s.Require().NoError(os.Setenv("DB_PORT", "3306"))
	s.Require().NoError(os.Setenv("DB_NAME", "test_db"))

	defer func() {
		_ = os.Unsetenv("DB_USER")
		_ = os.Unsetenv("DB_PASSWORD")
		_ = os.Unsetenv("DB_HOST")
		_ = os.Unsetenv("DB_PORT")
		_ = os.Unsetenv("DB_NAME")
	}()

	// This test will fail if there's no actual database running
	// It's mainly to test the function logic, not the actual connection
	_, err := NewConnection()

	// We expect an error here since we don't have a real database running
	// But we can verify that the error is related to connection, not configuration
	if err != nil {
		s.Contains(err.Error(), "failed to connect to database")
	}
}

// Test NewConnection with invalid configuration
func (s *ConnectionTestSuite) TestNewConnection_InvalidConfig() {
	// Clear environment variables to force configuration error
	_ = os.Unsetenv("DB_USER")
	_ = os.Unsetenv("DB_PASSWORD")

	db, err := NewConnection()

	s.Error(err)
	s.Nil(db)
	s.Contains(err.Error(), "user environment variable is required")
}

// Test configuration struct creation
func (s *ConnectionTestSuite) TestConfigurationStruct() {
	config := &configuration{
		user:     "testuser",
		password: "testpass",
		host:     "testhost",
		port:     "3307",
		database: "testdb",
	}

	s.Equal("testuser", config.user)
	s.Equal("testpass", config.password)
	s.Equal("testhost", config.host)
	s.Equal("3307", config.port)
	s.Equal("testdb", config.database)
}

// Test DSN building with special characters in password
func (s *ConnectionTestSuite) TestBuildDSN_SpecialCharacters() {
	config := &configuration{
		user:     "user@domain",
		password: "pass@word!#$",
		host:     "db.example.com",
		port:     "3306",
		database: "my_database",
	}

	dsn := buildDSN(config)
	expected := "user@domain:pass@word!#$@tcp(db.example.com:3306)/my_database?charset=utf8mb4&parseTime=True&loc=Local"

	s.Equal(expected, dsn)
}

// Test that environment variables are properly cleaned between tests
func (s *ConnectionTestSuite) TearDownTest() {
	// Clean up environment variables after each test
	envVars := []string{"DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME"}
	for _, env := range envVars {
		_ = os.Unsetenv(env)
	}
}

// Run the test suite
func TestConnectionTestSuite(t *testing.T) {
	suite.Run(t, new(ConnectionTestSuite))
}
