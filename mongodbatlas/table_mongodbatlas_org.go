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
			Hydrate:       listProjectParentOrg,
			ParentHydrate: listAtlasProjects,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "Unique identifier for the organization.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
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

func listProjectParentOrg(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*mongodbatlas.Project)

	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_org.listProjectParentOrg", "connection_error", err)
		return nil, err
	}

	org, _, err := client.Organizations.Get(ctx, project.OrgID)
	if err != nil {
		return nil, err
	}
	d.StreamListItem(ctx, org)

	return nil, nil
}
