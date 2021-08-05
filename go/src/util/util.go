package util

import (
	"strings"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/ibis7895123/go_graphql_app/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateUniqueID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

/**
 * DBの接続
 * driverNameは基本 "mysql" 固定だが、テスト用はトランザクションごとに名前をつける
 */
func NewDB(driverName string) *gorm.DB {
	dataSource := config.Config.MYSQL_DATASOURCE

	if dataSource == "" {
		panic("DataSource is empty. ")
	}

	dialector := mysql.New(
		mysql.Config{
			DriverName: driverName,
			DSN:        dataSource,
		},
	)

	// DB接続
	db, err := gorm.Open(
		dialector,
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)

	// DB起動エラー返す
	if err != nil || db == nil {
		panic(err)
	}

	return db
}

/**
 * テストDBの接続
 */
func NewTestDB() *gorm.DB {
	dataSource := config.Config.MYSQL_DATASOURCE

	if dataSource == "" {
		panic("DataSource is empty. ")
	}

	// ドライバの名前をランダムで生成
	driverName := uuid.New().String()

	// txdbに登録(トランザクションを独立させる)
	txdb.Register(driverName, "mysql", dataSource)

	// DB生成
	db := NewDB(driverName)

	return db
}
