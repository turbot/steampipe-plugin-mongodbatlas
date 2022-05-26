package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

func tableMongoDBAtlasCustomDBRole(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mongodbatlas_custom_db_role",
		Description: "Custom roles supports a subset of MongoDB privilege actions. These are defined at the project level, for all clusters in the project.",
		List: &plugin.ListConfig{
			Hydrate: listAtlasCustomDBRoles,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAtlasCustomDBRole,
			KeyColumns: plugin.AllColumns([]string{"role_name"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "role_name",
				Description: "The name of the role",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "actions",
				Description: "Each object in the actions array represents an individual privilege action granted by the role.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "inherited_roles",
				Description: "Each object in the inherited_roles array represents a key-value pair indicating the inherited role and the database on which the role is granted.",
				Type:        proto.ColumnType_JSON,
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RoleName"),
			},
		},
	}
}

func listAtlasCustomDBRoles(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	config := GetConfig(d.Connection)
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_custom_db_role.listAtlasCustomDBRoles", "connection_error", err)
		return nil, err
	}
	// Retrieve the list of incidents
	itemsPerPage := int64(500)
	// Reduce the basic request limit down if the user has only requested a small number of rows
	if d.QueryContext.Limit != nil && *d.QueryContext.Limit < itemsPerPage {
		itemsPerPage = *d.QueryContext.Limit
	}

	pageNumber := 1
	projectId := config.ProjectId

	for {
		databaseUsers, response, err := fetchCustomDBRolesPage(ctx, client, pageNumber, itemsPerPage, *projectId)

		if err != nil {
			plugin.Logger(ctx).Error("mongodbatlas_custom_db_role.listAtlasCustomDBRoles", "query_error", err)
			return nil, err
		}

		for _, databaseUser := range *databaseUsers {
			d.StreamListItem(ctx, databaseUser)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		// find the next page
		hasNextPage := false

		for _, l := range response.Links {
			if l.Rel == "next" {
				hasNextPage = true
			}
		}

		if hasNextPage {
			pageNumber++
			continue
		}

		break
	}

	return nil, nil
}

func fetchCustomDBRolesPage(ctx context.Context, client *mongodbatlas.Client, pageNumber int, itemsPerPage int64, projectId string) (*[]mongodbatlas.CustomDBRole, *mongodbatlas.Response, error) {
	plugin.Logger(ctx).Trace("mongodbatlas_custom_db_role.listAtlasCustomDBRoles", "fetchCustomDBRolesPage", projectId)
	return client.CustomDBRoles.List(ctx, projectId, &mongodbatlas.ListOptions{
		PageNum:      pageNumber,
		ItemsPerPage: int(itemsPerPage),
	})
}

func getAtlasCustomDBRole(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_custom_db_role.getAtlasCustomDBRole", "connection_error", err)
		return nil, err
	}
	projectId := *config.ProjectId

	roleName := d.KeyColumnQuals["role_name"].GetStringValue()

	databaseUser, _, err := client.CustomDBRoles.Get(ctx, projectId, roleName)
	if err != nil {
		return nil, err
	}

	return databaseUser, nil
}
