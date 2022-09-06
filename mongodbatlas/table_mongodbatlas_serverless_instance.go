package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

func tableMongoDBAtlasServerlessInstance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mongodbatlas_serverless_instance",
		Description: "Serverless instances in MongoDB Atlas are instances which are billed on usage, rather than time like in normal clusters.",
		List: &plugin.ListConfig{
			Hydrate:       listMongoDBAtlasServerlessInstances,
			ParentHydrate: listMongoDBAtlasProjects,
			KeyColumns:    plugin.OptionalColumns([]string{"project_id"}),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAtlasServerlessInstance,
			KeyColumns: plugin.AllColumns([]string{"name", "project_id"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "Unique identifier of the cluster.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The name of the cluster as it appears in Atlas.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "project_id",
				Description: "Unique identifier of the project that this cluster belongs to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GroupID"),
			},
			{
				Name:        "mongodb_version",
				Description: "Version of MongoDB that the serverless instance runs, in <major version>.<minor version> format.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("MongoDBVersion"),
			},
			{
				Name:        "provider_settings",
				Description: "Configuration for the provisioned hosts on which MongoDB runs. The available options are specific to the cloud service provider.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "state_name",
				Description: "Stage of deployment of this serverless instance when the resource made its request.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "connection_strings",
				Description: "Set of connection strings that your applications use to connect to this cluster.",
				Type:        proto.ColumnType_JSON,
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

func listMongoDBAtlasServerlessInstances(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	project := h.Item.(*mongodbatlas.Project)
	client, err := getMongoDBAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_serverless_instance.listAtlasServerlessInstances", "connection_error", err)
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
		serverlessInstances, response, err := client.ServerlessInstances.List(ctx, project.ID, &mongodbatlas.ListOptions{
			PageNum:      pageNumber,
			ItemsPerPage: int(itemsPerPage),
		})
		plugin.Logger(ctx).Trace("mongodbatlas_serverless_instance.listAtlasServerlessInstances", "cluster length", len(serverlessInstances.Results))

		if err != nil {
			plugin.Logger(ctx).Error("mongodbatlas_serverless_instance.listAtlasServerlessInstances", "query_error", err)
			return nil, err
		}

		for _, cluster := range serverlessInstances.Results {
			d.StreamListItem(ctx, cluster)

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

func getAtlasServerlessInstance(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := getMongoDBAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_serverless_instance.getAtlasServerlessInstances", "connection_error", err)
		return nil, err
	}
	clusterName := d.KeyColumnQuals["name"].GetStringValue()
	projectId := d.KeyColumnQuals["project_id"].GetStringValue()

	cluster, _, err := client.ServerlessInstances.Get(ctx, projectId, clusterName)
	if err != nil {
		return nil, err
	}

	return cluster, nil
}
