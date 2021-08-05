package config_test

import (
	"os"
	"testing"

	"github.com/ibis7895123/go_graphql_app/src/config"
	"github.com/stretchr/testify/assert"
)

func Test_正常系_init(t *testing.T) {
	env := config.Config.ENV
	assert.Equal(t, env, "local")
}

func Test_正常系_EnvLoad(t *testing.T) {
	// .env.sampleを呼ぶ
	config.EnvLoad("/go_graphql_app/.env.sample")
	assert.Equal(t, os.Getenv("ENV"), "local")
}
