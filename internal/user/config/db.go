package user

import (
	"fmt"

	userEntity "CocoaTalk_Server/internal/user/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var dbConfig = GetDBConfig()
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConfig.DBUsername,
		dbConfig.DBPassword,
		dbConfig.DBHost,
		dbConfig.DBPort,
		dbConfig.DBName,
	)

	var err error
	db, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("DB 연결 실패: %v", err))
	}
}

func checkUserTable(db *gorm.DB) error {
	if !db.Migrator().HasTable(&userEntity.User{}) {
		return fmt.Errorf("user 테이블 읽기 실패")
	}
	return nil
}

func DB() *gorm.DB {
	if db == nil {
		Init()
	}
	checkUserTable(db)

	return db
}
