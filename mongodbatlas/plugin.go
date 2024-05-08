package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

const pluginName = "steampipe-plugin-mongodbatlas"

// Plugin creates this (mongodbatlas) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{},
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "organization_id",
				Hydrate: getOrganizationId,
			},
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		TableMap: map[string]*plugin.Table{
			"mongodbatlas_cluster":                           tableMongoDBAtlasCluster(ctx),
			"mongodbatlas_container":                         tableMongoDBAtlasContainer(ctx),
			"mongodbatlas_custom_db_role":                    tableMongoDBAtlasCustomDBRole(ctx),
			"mongodbatlas_database_user":                     tableMongoDBAtlasDatabaseUser(ctx),
			"mongodbatlas_org":                               tableMongoDBAtlasOrg(ctx),
			"mongodbatlas_org_event":                         tableMongoDBAtlasOrgEvents(ctx),
			"mongodbatlas_project":                           tableMongoDBAtlasProject(ctx),
			"mongodbatlas_project_event":                     tableMongoDBAtlasProjectEvents(ctx),
			"mongodbatlas_project_ip_access_list":            tableMongoDBAtlasProjectIpAccessList(ctx),
			"mongodbatlas_serverless_instance":               tableMongoDBAtlasServerlessInstance(ctx),
			"mongodbatlas_team":                              tableMongoDBAtlasTeam(ctx),
			"mongodbatlas_x509_authentication_database_user": tableMongoDBAtlasX509AuthenticationDatabaseUser(ctx),
		},
	}

	return p
}

func hasNextPage(r *mongodbatlas.Response) bool {
	for _, l := range r.Links {
		if l.Rel == "next" {
			return true
		}
	}
	return false
}