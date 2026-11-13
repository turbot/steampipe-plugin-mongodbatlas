package mongodbatlas

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type mongodbatlasConfig struct {
	PublicKey  *string `hcl:"public_key"`
	PrivateKey *string `hcl:"private_key"`
}

func ConfigInstance() interface{} {
	return &mongodbatlasConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) mongodbatlasConfig {
	if connection == nil || connection.GetConfig() == nil {
		return mongodbatlasConfig{}
	}
	config, _ := connection.GetConfig().(mongodbatlasConfig)
	return config
}
