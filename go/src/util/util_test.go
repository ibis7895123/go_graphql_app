package util_test

import (
	"testing"

	"github.com/ibis7895123/go_graphql_app/src/util"
	"github.com/stretchr/testify/assert"
)

func Test_正常系_CreateUniqueID(t *testing.T) {
	id := util.CreateUniqueID()
	assert.Equal(t, len(id), 32)
}

func Test_正常系_NewDB(t *testing.T) {
	db := util.NewDB()
	assert.NotNil(t, db)
}
