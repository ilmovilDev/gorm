package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func Init() error {
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger := NewLogger(p)
	return logger
}
