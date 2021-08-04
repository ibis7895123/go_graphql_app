package util

import (
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func CreateUniqueID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

/**
 * .envファイルのローディング
 */
func EnvLoad(filePath string) {
	err := godotenv.Load(filePath)

	if err != nil {
		log.Fatal(err)
	}
}

/**
 * DBの接続
 */
func NewDB() *gorm.DB {
	// NOTE: envファイルの絶対パスを設定する
	EnvLoad("/go_graphql_app/.env")

	dataSource := os.Getenv("MYSQL_DATASOURCE")

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
