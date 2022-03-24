package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func NewPostgresConn(config Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  GetDSN(config),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// Create the connection pool

	sqlDB, err := db.DB()

	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, err
}
func CloseDBConnection(conn *gorm.DB) error {
	sqlDB, err := conn.DB()
	if err != nil {
		return err
	}
	defer sqlDB.Close()
	return nil
}
