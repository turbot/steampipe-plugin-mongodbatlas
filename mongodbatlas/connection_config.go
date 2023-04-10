package mongodbatlas

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type mongodbatlasConfig struct {
	PublicKey  *string `cty:"public_key"`
	PrivateKey *string `cty:"private_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"public_key": {
		Type:     schema.TypeString,
		Required: true,
	},
	"private_key": {
		Type:     schema.TypeString,
		Required: true,
	},
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
