package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

func tableMongoDBAtlasOrg(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mongodbatlas_org",
		Description: "Returns a single record containing the parent org of the project",
		List: &plugin.ListConfig{
			Hydrate:    listMongoDBAtlasOrg,
			KeyColumns: plugin.OptionalColumns([]string{"id"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "Unique identifier for the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "Name of the organization.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_deleted",
				Description: "Flag indicating if the organization is deleted.",
				Type:        proto.ColumnType_BOOL,
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

func listMongoDBAtlasOrg(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_org.listProjectParentOrg", "connection_error", err)
		return nil, err
	}

	if len(d.KeyColumnQuals["id"].GetStringValue()) != 0 {
		org, _, err := client.Organizations.Get(ctx, d.KeyColumnQuals["id"].GetStringValue())
		if err != nil {
			return nil, err
		}
		d.StreamListItem(ctx, org)
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
		orgs, response, err := client.Organizations.List(ctx, &mongodbatlas.OrganizationsListOptions{
			ListOptions: mongodbatlas.ListOptions{
				PageNum:      pageNumber,
				ItemsPerPage: int(itemsPerPage),
			},
		})
		if err != nil {
			return nil, err
		}
		for _, org := range orgs.Results {
			d.StreamListItem(ctx, org)
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
