package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

type rowContainer struct {
	*mongodbatlas.Container
	ProjectId string
}

func tableMongoDBAtlasContainer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mongodbatlas_container",
		Description: "Containers in a project allows for cloud provider backed virtual private networking - dubbed as container network peering in MongoDB Atlas.",
		List: &plugin.ListConfig{
			Hydrate:       listMongoDBAtlasContainers,
			ParentHydrate: listMongoDBAtlasProjects,
			KeyColumns:    plugin.OptionalColumns([]string{"provider_name", "project_id"}),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getContainer,
			KeyColumns: plugin.AllColumns([]string{"id", "project_id"}),
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "Unique identifier for the container.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "project_id",
				Description: "Unique identifier for the project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "provider_name",
				Description: "Cloud provider for this Network Peering connection.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "atlas_cidr_block",
				Description: "CIDR block that Atlas uses for your clusters.",
				Type:        proto.ColumnType_CIDR,
				Transform:   transform.FromField("AtlasCIDRBlock"),
			},
			{
				Name:        "azure_subscription_id",
				Description: "Unique identifer of the Azure subscription in which the VNet resides.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AzureSubscriptionID").NullIfZero(),
			},
			{
				Name:        "gcp_project_id",
				Description: "Unique identifier of the Google Cloud project in which the network peer resides. Returns null until a peering connection is created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GCPProjectID").NullIfZero(),
			},
			{
				Name:        "network_name",
				Description: "Unique identifier of the Network Peering connection in the Atlas project. Returns null until a peering connection is created.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "provisioned",
				Description: "Flag that indicates if the project has clusters deployed in the Network Peering container or Azure VNet.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "region",
				Description: "AWS region where the VCP resides or Azure region where the VNet resides.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "vnet_name",
				Description: "Unique identifier of your Azure VNet. The value is null if there are no network peering connections in the container.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "vpc_id",
				Description: "Unique identifier of the project's Network Peering container.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VPCID"),
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("NetworkName"),
			},
		}),
	}
}

func listMongoDBAtlasContainers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*mongodbatlas.Project)
	// Create client
	client, err := getMongoDBAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("table_mongodbatlas_container.listContainers", "connection_error", err)
		return nil, err
	}
	// Retrieve the list of incidents
	itemsPerPage := int64(500)
	// Reduce the basic request limit down if the user has only requested a small number of rows
	if d.QueryContext.Limit != nil && *d.QueryContext.Limit < itemsPerPage {
		itemsPerPage = *d.QueryContext.Limit
	}

	pageNumber := 1
	providerName := d.EqualsQuals["provider_name"].GetStringValue()

	for {
		listOptions := &mongodbatlas.ContainersListOptions{
			ListOptions: mongodbatlas.ListOptions{
				PageNum:      pageNumber,
				ItemsPerPage: int(itemsPerPage),
			},
		}
		if len(providerName) > 0 {
			listOptions.ProviderName = providerName
		}

		containers, response, err := client.Containers.List(ctx, project.ID, listOptions)

		if err != nil {
			plugin.Logger(ctx).Error("table_mongodbatlas_container.listContainers", "query_error", err)
			return nil, err
		}

		for _, container := range containers {
			c := rowContainer{
				Container: &container,
				ProjectId: project.ID,
			}
			d.StreamListItem(ctx, c)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
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

func getContainer(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := getMongoDBAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_project_events.getAtlasProjectEvents", "connection_error", err)
		return nil, err
	}

	projectId := d.EqualsQuals["project_id"].GetStringValue()
	containerId := d.EqualsQuals["event_id"].GetStringValue()

	event, _, err := client.Containers.Get(ctx, projectId, containerId)

	if err != nil {
		return nil, err
	}

	return event, nil
}
