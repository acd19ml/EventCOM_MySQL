package conf_test

import (
	"testing"

	"github.com/acd19ml/EventCOM_MySQL/conf"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		should.Equal("EventCOM", conf.C().App.Name)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	should := assert.New(t)
	// os.Setenv("MYSQL_DATABASE", "unit_test")
	err := conf.LoadConfigFromEnv()
	if should.NoError(err) {
		should.Equal("EventCOM", conf.C().MySQL.Database)
	}
}

func TestGetDB(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		conf.C().MySQL.GetDB()
	}
}
