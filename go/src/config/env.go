package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// NOTE: envファイルの絶対パスを設定する
// dockerコンテナのパスを設定
const envFilePath = "/go_graphql_app/.env"

var Config ConfigStruct

type ConfigStruct struct {
	ENV              string `env:"ENV"`
	PORT             string `env:"PORT"`
	MYSQL_DATASOURCE string `env:"MYSQL_DATASOURCE"`
}

/**
 * configパッケージローディング時に呼ばれる
 */
func init() {
	EnvLoad(envFilePath)

	// envファイルの中身をセット
	err := env.Parse(&Config)

	if err != nil {
		panic(err)
	}
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
