package pgsql

import (
	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

var options = &pg.Options{
	Addr:     "db-test:5432",
	User:     "root",
	Password: "root",
	Database: "testdb",
}

func TestNewPgDb_WhenOptionsOk_ThenCreateClient(t *testing.T) {
	got := NewPgDb(options)
	defer func(got *pg.DB) {
		_ = got.Close()
	}(got)

	assert.NotNil(t, got)
	assert.Equal(t, options.Addr, got.Options().Addr)
	assert.Equal(t, options.User, got.Options().User)
	assert.Equal(t, options.Password, got.Options().Password)
	assert.Equal(t, options.Database, got.Options().Database)
}
