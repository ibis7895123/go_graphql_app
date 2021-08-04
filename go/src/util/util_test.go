package util_test

import (
	"os"
	"testing"

	"github.com/ibis7895123/go_graphql_app/src/util"
	"github.com/stretchr/testify/assert"
)

func Test_CreateUniqueID(t *testing.T) {
	id := util.CreateUniqueID()
	assert.Equal(t, len(id), 32)
}

func Test_EnvLoad(t *testing.T) {
	// .env.sampleを呼ぶ
	util.EnvLoad("/go_graphql_app/.env.sample")
	assert.Equal(t, os.Getenv("ENV"), "local")
}

func Test_NewDB(t *testing.T) {
	db := util.NewDB()
	assert.NotNil(t, db)
}
