package config

import (
	configureConts "github.com/fiqrikm18/markerplace_core/pkg/const"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewAppConfig_failed(t *testing.T) {
	var appConfig *AppConfig
	var err error
	assert.Nil(t, appConfig)

	appConfig, err = NewAppConfig("test_failed")
	assert.Error(t, err)
}

func TestNewAppConfig_success(t *testing.T) {
	var appConfig *AppConfig
	var err error
	assert.Nil(t, appConfig)

	err = os.Setenv("MARKETPLACE_CORE_CONFIG", "/Users/e180/Projects/personal/marketplace_core")
	assert.Nil(t, err)

	appConfig, err = NewAppConfig(configureConts.ConfigurationFileName)
	assert.Nil(t, err)

	assert.NotNil(t, appConfig)
	assert.NotEmpty(t, appConfig.BaseUrl)
}
