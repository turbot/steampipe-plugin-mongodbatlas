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
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"mongodbatlas_project":                           tableMongoDBAtlasProject(ctx),
			"mongodbatlas_org":                               tableMongoDBAtlasOrg(ctx),
			"mongodbatlas_cluster":                           tableMongoDBAtlasCluster(ctx),
			"mongodbatlas_advanced_cluster":                  tableMongoDBAtlasAdvancedCluster(ctx),
			"mongodbatlas_serverless_instance":               tableMongoDBAtlasServerlessInstance(ctx),
			"mongodbatlas_project_ip_access_list":            tableMongoDBAtlasProjectIpAccessList(ctx),
			"mongodbatlas_database_user":                     tableMongoDBAtlasDatabaseUser(ctx),
			"mongodbatlas_x509_authentication_database_user": tableMongoDBAtlasX509Auth(ctx),
			"mongodbatlas_custom_db_role":                    tableMongoDBAtlasCustomDBRole(ctx),
			"mongodbatlas_project_event":                     tableMongoDBAtlasProjectEvents(ctx),
			"mongodbatlas_org_event":                         tableMongoDBAtlasOrgEvents(ctx),
			"mongodbatlas_team":                              tableMongoDBAtlasTeam(ctx),
			"mongodbatlas_container":                         tableMongoDBAtlasContainer(ctx),
		},
	}

	return p
}
