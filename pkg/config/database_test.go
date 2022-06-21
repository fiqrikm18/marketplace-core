package config

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	viper.Set("app_environment", "test")
}

func TestNewDbConnection_failed(t *testing.T) {
	viper.Set("db_username", "post_test")
	_, err := NewDBConnection()
	assert.Error(t, err)
	viper.Set("db_username", "postgres")
}

func TestNewDbConnection_success(t *testing.T) {
	conn, err := NewDBConnection()
	assert.NoError(t, err)
	assert.NotNil(t, conn.DB)
}
