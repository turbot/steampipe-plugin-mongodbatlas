package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableAtlasProject(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mongodbatlas_project",
		Description: "Returns details of the project configured in the connection config",
		List: &plugin.ListConfig{
			Hydrate: listAtlasProjects,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "An unique identifier of the project.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the project.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "cluster_count",
				Description: "The number of Atlas clusters deployed in the project.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ClusterCount"),
			},
			{
				Name:        "org_id",
				Description: "The unique identifier of the Atlas organization to which the project belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("OrgID"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listAtlasProjects(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_project.listAtlasProjects", "connection_error", err)
		return nil, err
	}
	config := GetConfig(d.Connection)
	id := *config.ProjectId

	project, _, err := client.Projects.GetOneProject(ctx, id)
	if err != nil {
		return nil, err
	}

	d.StreamListItem(ctx, project)

	return nil, nil
}
