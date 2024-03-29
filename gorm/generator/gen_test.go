package generator

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func connectMysql(dsn string, options ...Option) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestGenerateTableStruct(t *testing.T) {
	db, err := connectMysql("root:123456@tcp(192.168.119.128:3306)/poker")
	if err != nil {
		t.Fatal(err)
	}

	GenerateTableStruct(db,
		WithOutPath("./gen"),
		GenTables("cfg_currency"),
	)
}
