package util

import (
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/ibis7895123/go_graphql_app/src/config"
	"github.com/jinzhu/gorm"
)

func CreateUniqueID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

/**
 * DBの接続
 */
func NewDB() *gorm.DB {
	dataSource := config.Config.MYSQL_DATASOURCE

	if dataSource == "" {
		panic("DataSource is empty. ")
	}

	// DB接続
	db, err := gorm.Open("mysql", dataSource)

	// DB起動エラー返す
	if err != nil || db == nil {
		panic(err)
	}

	// ログ出力をON
	db.LogMode(true)

	return db
}
