package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

type rowTeam struct {
	mongodbatlas.Team
	OrgId string
}

func tableMongoDBAtlasTeam(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mongodbatlas_team",
		Description: "Teams enable you to grant project access roles to multiple users. You add any number of organization users to a team.",
		List: &plugin.ListConfig{
			Hydrate:       listMongodDBAtlasTeams,
			ParentHydrate: listMongoDBAtlasOrg,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "Unique identifier for the team.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "Name of the team.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "users",
				Description: "Users assigned to the team.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromValue(),
				Hydrate:     getTeamUsers,
			},
			{
				Name:        "org_id",
				Description: "Unique identifier of the organization for this team.",
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

func listMongodDBAtlasTeams(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	org := h.Item.(*mongodbatlas.Organization)

	// Create client
	client, err := getMongoDBAtlasClient(ctx, d)
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
		teams, response, err := client.Teams.List(ctx, org.ID, &mongodbatlas.ListOptions{
			PageNum:      pageNumber,
			ItemsPerPage: int(itemsPerPage),
		})

		if err != nil {
			plugin.Logger(ctx).Error("table_mongodbatlas_team.listTeams", "query_error", err)
			return nil, err
		}

		for _, team := range teams {
			rTeam := rowTeam{
				Team:  team,
				OrgId: org.ID,
			}
			d.StreamListItem(ctx, rTeam)
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

func getTeamUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	data := h.Item.(rowTeam)

	client, err := getMongoDBAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("table_mongodbatlas_team.listTeams", "connection_error", err)
		return nil, err
	}

	users, _, err := client.Teams.GetTeamUsersAssigned(ctx, data.OrgId, data.ID)
	plugin.Logger(ctx).Trace("error", err)
	plugin.Logger(ctx).Trace("users", users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
