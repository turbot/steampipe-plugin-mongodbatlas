package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

func tableMongoDBAtlasTeam(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mongodbatlas_team",
		Description: "Teams enable you to grant project access roles to multiple users. You add any number of organization users to a team.",
		List: &plugin.ListConfig{
			Hydrate:       listMongodDBAtlasTeams,
			ParentHydrate: listMongoDBAtlasProjects,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "Unique identifier for the team.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "Name of the team.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
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

func listMongodDBAtlasTeams(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*mongodbatlas.Project)

	// Create client
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("table_mongodbatlas_team.listTeams", "connection_error", err)
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
		teams, response, err := client.Teams.List(ctx, project.OrgID, &mongodbatlas.ListOptions{
			PageNum:      pageNumber,
			ItemsPerPage: int(itemsPerPage),
		})

		if err != nil {
			plugin.Logger(ctx).Error("table_mongodbatlas_team.listTeams", "query_error", err)
			return nil, err
		}

		for _, team := range teams {
			d.StreamListItem(ctx, team)
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
