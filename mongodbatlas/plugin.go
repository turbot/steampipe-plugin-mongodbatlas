package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

const pluginName = "steampipe-plugin-mongodbatlas"

// Plugin creates this (mongodbatlas) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().Transform(transform.NullIfZeroValue),
		DefaultGetConfig: &plugin.GetConfig{},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"mongodbatlas_project":                tableAtlasProject(ctx),
			"mongodbatlas_org":                    tableAtlasOrg(ctx),
			"mongodbatlas_cluster":                tableAtlasCluster(ctx),
			"mongodbatlas_advanced_cluster":       tableAtlasAdvancedCluster(ctx),
			"mongodbatlas_serverless_instance":    tableAtlasServerlessInstance(ctx),
			"mongodbatlas_project_ip_access_list": tableAtlasProjectIpAccessList(ctx),
		},
	}

	return p
}
