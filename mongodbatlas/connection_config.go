package mongodbatlas

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
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
	if connection == nil || connection.Config == nil {
		return mongodbatlasConfig{}
	}
	config, _ := connection.Config.(mongodbatlasConfig)
	return config
}
