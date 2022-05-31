package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

func tableMongoDBAtlasProject(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mongodbatlas_project",
		Description: "Returns details of the project configured in the connection config",
		List: &plugin.ListConfig{
			Hydrate:    listMongoDBAtlasProjects,
			KeyColumns: plugin.OptionalColumns([]string{"id"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "A unique identifier of the project.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cluster_count",
				Description: "The number of Atlas clusters deployed in the project.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromGo(), // 0 is possible
			},
			{
				Name:        "org_id",
				Description: "The unique identifier of the Atlas organization to which the project belongs.",
				Type:        proto.ColumnType_STRING,
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

func listMongoDBAtlasProjects(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_project.listAtlasProjects", "connection_error", err)
		return nil, err
	}

	if len(d.KeyColumnQuals["id"].GetStringValue()) != 0 {
		project, _, err := client.Projects.GetOneProject(ctx, d.KeyColumnQuals["id"].GetStringValue())
		if err != nil {
			return nil, err
		}
		d.StreamListItem(ctx, project)
		return nil, nil
	}

	// Retrieve the list of incidents
	itemsPerPage := int64(500)
	// Reduce the basic request limit down if the user has only requested a small number of rows
	if d.QueryContext.Limit != nil && *d.QueryContext.Limit < itemsPerPage {
		itemsPerPage = *d.QueryContext.Limit
	}

	pageNumber := 1

	for {
		projects, response, err := client.Projects.GetAllProjects(ctx, &mongodbatlas.ListOptions{
			PageNum:      pageNumber,
			ItemsPerPage: int(itemsPerPage),
		})
		if err != nil {
			return nil, err
		}

		for _, project := range projects.Results {
			d.StreamListItem(ctx, project)
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
