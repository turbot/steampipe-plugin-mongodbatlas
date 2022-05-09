package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

func tableAtlasProject(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mongodbatlas_project",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listAtlasProjects,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "org_id",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAtlasProject,
			KeyColumns: plugin.SingleColumn("id"),
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
	// Create client
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_project.listAtlasProjects", "connection_error", err)
		return nil, err
	}
	// Retrieve the list of incidents
	itemsPerPage := int64(100)
	// Reduce the basic request limit down if the user has only requested a small number of rows
	if d.QueryContext.Limit != nil && *d.QueryContext.Limit < itemsPerPage {
		itemsPerPage = *d.QueryContext.Limit
	}

	pageNumber := 1
	orgId := ""

	// Additional Filters
	if d.KeyColumnQuals["org_id"] != nil {
		orgId = d.KeyColumnQuals["org_id"].GetStringValue()
	}

	for {
		projects, _, err := fetchProjectsPage(ctx, client, pageNumber, itemsPerPage, orgId)

		if err != nil {
			plugin.Logger(ctx).Error("mongodbatlas_project.listAtlasProjects", "query_error", err)
			return nil, err
		}

		for _, project := range projects.Results {
			plugin.Logger(ctx).Trace("mongodbatlas_project.listAtlasProjects", "streaming item", project.ClusterCount)
			d.StreamListItem(ctx, project)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		// find the next page
		hasNextPage := false
		for _, link := range projects.Links {
			if link.Rel == "next" {
				hasNextPage = true
				break
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

func fetchProjectsPage(ctx context.Context, client *mongodbatlas.Client, pageNumber int, itemsPerPage int64, orgId string) (*mongodbatlas.Projects, *mongodbatlas.Response, error) {
	if len(orgId) != 0 {
		plugin.Logger(ctx).Trace("mongodbatlas_project.listAtlasProjects", "all_projects")
		return client.Organizations.Projects(ctx, orgId, &mongodbatlas.ListOptions{
			PageNum:      pageNumber,
			ItemsPerPage: int(itemsPerPage),
		})
	}

	plugin.Logger(ctx).Trace("mongodbatlas_project.listAtlasProjects", "org_projects", orgId)
	return client.Projects.GetAllProjects(ctx, &mongodbatlas.ListOptions{
		PageNum:      pageNumber,
		ItemsPerPage: int(itemsPerPage),
	})
}

func getAtlasProject(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_project.listAtlasProjects", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetStringValue()

	// No inputs
	if len(id) == 0 {
		return nil, nil
	}

	project, _, err := client.Projects.GetOneProject(ctx, id)
	if err != nil {
		return nil, err
	}

	return project, nil
}
