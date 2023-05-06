package repository

import (
	envconfig "health-care-backend/envconfig"
	"testing"

	"github.com/stretchr/testify/assert"
)

func prepareDatabaseConnection(t *testing.T) *GormDatabase {
	t.Helper()

	var env envconfig.Env
	err := envconfig.Process(&env) // intent to load config from ENV variables
	assert.NoError(t, err)

	db, err := NewGormDatabase(env.DATABASE_URL, false)
	assert.NoError(t, err)

	db.AutoMigrate()
	return db
}

func Test_NewGormDatabase(t *testing.T) {
	db := prepareDatabaseConnection(t)
	assert.NotNil(t, db)
}
