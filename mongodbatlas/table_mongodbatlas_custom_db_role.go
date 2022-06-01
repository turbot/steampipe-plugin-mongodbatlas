package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

type rowCustomDBRole struct {
	mongodbatlas.CustomDBRole
	ProjectID string
}

func tableMongoDBAtlasCustomDBRole(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mongodbatlas_custom_db_role",
		Description: "Custom roles supports a subset of MongoDB privilege actions. These are defined at the project level, for all clusters in the project.",
		List: &plugin.ListConfig{
			Hydrate:       listMongoDBAtlasCustomDBRoles,
			ParentHydrate: listMongoDBAtlasProjects,
			KeyColumns:    plugin.OptionalColumns([]string{"project_id"}),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAtlasCustomDBRole,
			KeyColumns: plugin.AllColumns([]string{"role_name", "project_id"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "role_name",
				Description: "The name of the role.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "project_id",
				Description: "The unique identifier of the project for this role.",
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

func listMongoDBAtlasCustomDBRoles(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*mongodbatlas.Project)
	// Create client
	client, err := getMongoDBAtlasClient(ctx, d)
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

	for {
		roles, response, err := client.CustomDBRoles.List(ctx, project.ID, &mongodbatlas.ListOptions{
			PageNum:      pageNumber,
			ItemsPerPage: int(itemsPerPage),
		})

		if err != nil {
			plugin.Logger(ctx).Error("mongodbatlas_custom_db_role.listAtlasCustomDBRoles", "query_error", err)
			return nil, err
		}

		for _, role := range *roles {
			r := rowCustomDBRole{
				CustomDBRole: role,
				ProjectID:    project.ID,
			}
			d.StreamListItem(ctx, r)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if hasNextPage(response) {
			pageNumber++
			continue
		}

		break
	}

	return nil, nil
}

func getAtlasCustomDBRole(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := getMongoDBAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_custom_db_role.getAtlasCustomDBRole", "connection_error", err)
		return nil, err
	}

	roleName := d.KeyColumnQuals["role_name"].GetStringValue()
	projectId := d.KeyColumnQuals["project_id"].GetStringValue()

	role, _, err := client.CustomDBRoles.Get(ctx, projectId, roleName)
	if err != nil {
		return nil, err
	}

	return role, nil
}
