package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

func tableMongoDBAtlasContainer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "table_mongodbatlas_container",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate:    listContainers,
			KeyColumns: plugin.OptionalColumns([]string{"provider_name"}),
		},
		Get: &plugin.GetConfig{
			Hydrate:    getContainer,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "Unique identifier for the team.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "provider_name",
				Description: "Cloud provider for this Network Peering connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ProviderName").NullIfZero(),
			},
			{
				Name:        "atlas_cidr_block",
				Description: "CIDR block that Atlas uses for your clusters.",
				Type:        proto.ColumnType_CIDR,
				Transform:   transform.FromField("AtlasCIDRBlock").NullIfZero(),
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
				Transform:   transform.FromField("NetworkName").NullIfZero(),
			},
			{
				Name:        "provisioned",
				Description: "Flag that indicates if the project has clusters deployed in the Network Peering container or Azure VNet.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Provisioned"),
			},
			{
				Name:        "region",
				Description: "AWS region where the VCP resides or Azure region where the VNet resides.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Region").NullIfZero(),
			},
			{
				Name:        "vnet_name",
				Description: "Unique identifier of your Azure VNet. The value is null if there are no network peering connections in the container.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("NetworkName").NullIfZero(),
			},
			{
				Name:        "vpc_id",
				Description: "Unique identifier of the project's Network Peering container.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VPCID").NullIfZero(),
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AtlasCIDRBlock"),
			},
		},
	}
}

func listContainers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getMongodbAtlasClient(ctx, d)
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
	projectId := GetConfig(d.Connection).ProjectId
	providerName := d.KeyColumnQuals["provider_name"].GetStringValue()

	for {
		containers, response, err := fetchContainers(ctx, client, pageNumber, itemsPerPage, *projectId, providerName)

		if err != nil {
			plugin.Logger(ctx).Error("table_mongodbatlas_container.listContainers", "query_error", err)
			return nil, err
		}

		for _, container := range containers {
			d.StreamListItem(ctx, container)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		// find the next page
		hasNextPage := false

		for _, l := range response.Links {
			if l.Rel == "next" {
				hasNextPage = true
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

func fetchContainers(ctx context.Context, client *mongodbatlas.Client, pageNumber int, itemsPerPage int64, projectId string, providerName string) ([]mongodbatlas.Container, *mongodbatlas.Response, error) {
	plugin.Logger(ctx).Trace("table_mongodbatlas_container.listContainers", "fetchContainers", projectId)

	if len(providerName) != 0 {
		return client.Containers.List(ctx, projectId, &mongodbatlas.ContainersListOptions{
			ProviderName: providerName,
			ListOptions: mongodbatlas.ListOptions{
				PageNum:      pageNumber,
				ItemsPerPage: int(itemsPerPage),
			},
		})
	} else {
		return client.Containers.ListAll(ctx, projectId, &mongodbatlas.ListOptions{
			PageNum:      pageNumber,
			ItemsPerPage: int(itemsPerPage),
		})
	}

}

func getContainer(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_project_events.getAtlasProjectEvents", "connection_error", err)
		return nil, err
	}

	projectId := *config.ProjectId
	containerId := d.KeyColumnQuals["event_id"].GetStringValue()

	event, _, err := client.Containers.Get(ctx, projectId, containerId)

	if err != nil {
		return nil, err
	}

	return event, nil
}
