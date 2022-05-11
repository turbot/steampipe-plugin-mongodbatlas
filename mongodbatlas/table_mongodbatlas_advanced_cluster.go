package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

func tableAtlasAdvancedCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mongodbatlas_advanced_cluster",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listAtlasClusters,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAtlasCluster,
			KeyColumns: plugin.AllColumns([]string{"name"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "Unique identifier of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the cluster as it appears in Atlas.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "project_id",
				Description: "Unique identifier of the project that this cluster belongs to",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GroupID"),
			},
			{
				Name:        "auto_scaling",
				Description: "TDB",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("AutoScaling"),
			},
			{
				Name:        "bi_connector_config",
				Description: "TDB",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("BiConnector"),
			},
			{
				Name:        "cluster_type",
				Description: "TDB",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ClusterType"),
			},
			{
				Name:        "disk_size_gb",
				Description: "TDB",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("DiskSizeGB"),
			},
			{
				Name:        "encryption_at_rest_provider",
				Description: "TDB",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("EncryptionAtRestProvider"),
			},
			{
				Name:        "labels",
				Description: "TDB",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Labels"),
			},
			{
				Name:        "mongo_db_version",
				Description: "TDB",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("MongoDBVersion"),
			},
			{
				Name:        "mongodb_major_version",
				Description: "TDB",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("MongoDBMajorVersion"),
			},
			{
				Name:        "mongo_uri",
				Description: "TDB",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("MongoURI"),
			},
			{
				Name:        "mongo_uri_updated",
				Description: "TDB",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("MongoURIUpdated"),
			},
			{
				Name:        "mongo_uri_with_options",
				Description: "TDB",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("MongoURIWithOptions"),
			},
			{
				Name:        "num_shards",
				Description: "TDB",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("NumShards"),
			},
			{
				Name:        "paused",
				Description: "TDB",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Paused"),
			},
			{
				Name:        "pit_enabled",
				Description: "TDB",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("PitEnabled"),
			},
			{
				Name:        "provider_backup_enabled",
				Description: "TDB",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("ProviderBackupEnabled"),
			},
			{
				Name:        "provider_settings",
				Description: "TDB",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ProviderSettings"),
			},
			{
				Name:        "replication_factor",
				Description: "Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ReplicationFactor"),
			},
			{
				Name:        "replication_spec",
				Description: "Configuration of each region in the cluster. Each element in this object represents a region where Atlas deploys your cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ReplicationSpec"),
			},
			{
				Name:        "replication_specs",
				Description: "Configuration for each zone in a Global Cluster. Each object in this array represents a zone where Atlas deploys nodes for your Global Cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ReplicationSpecs"),
			},
			{
				Name:        "srv_address",
				Description: "TDB",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SrvAddress"),
			},
			{
				Name:        "state_name",
				Description: "TDB",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("StateName"),
			},
			{
				Name:        "connection_strings",
				Description: "Set of connection strings that your applications use to connect to this cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ConnectionStrings"),
			},
			{
				Name:        "version_release_system",
				Description: "TDB",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("VersionReleaseSystem"),
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

func listAtlasAdvancedClusters(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	config := GetConfig(d.Connection)
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_cluster.listAtlasAdvancedClusters", "connection_error", err)
		return nil, err
	}
	// Retrieve the list of incidents
	itemsPerPage := int64(500)
	// Reduce the basic request limit down if the user has only requested a small number of rows
	if d.QueryContext.Limit != nil && *d.QueryContext.Limit < itemsPerPage {
		itemsPerPage = *d.QueryContext.Limit
	}

	pageNumber := 1
	projectId := config.ProjectId

	for {
		advancedClustersResponse, _, err := fetchAdvancedClustersPage(ctx, client, pageNumber, itemsPerPage, *projectId)
		plugin.Logger(ctx).Trace("mongodbatlas_cluster.listAtlasCluster", "cluster length", len(advancedClustersResponse.Results))

		if err != nil {
			plugin.Logger(ctx).Error("mongodbatlas_cluster.listAtlasClusters", "query_error", err)
			return nil, err
		}

		for _, cluster := range advancedClustersResponse.Results {
			d.StreamListItem(ctx, cluster)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		// find the next page
		hasNextPage := false

		for _, l := range advancedClustersResponse.Links {
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

func fetchAdvancedClustersPage(ctx context.Context, client *mongodbatlas.Client, pageNumber int, itemsPerPage int64, projectId string) (*mongodbatlas.AdvancedClustersResponse, *mongodbatlas.Response, error) {
	plugin.Logger(ctx).Trace("mongodbatlas_advanced_cluster.listAtlasCluster", "project_clusters", projectId)
	return client.AdvancedClusters.List(ctx, projectId, &mongodbatlas.ListOptions{
		PageNum:      pageNumber,
		ItemsPerPage: int(itemsPerPage),
	})
}

func getAtlasAdvancedCluster(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_advanced_cluster.getAtlasAdvancedCluster", "connection_error", err)
		return nil, err
	}
	projectId := config.ProjectId
	clusterName := d.KeyColumnQuals["name"].GetStringValue()

	cluster, _, err := client.AdvancedClusters.Get(ctx, *projectId, clusterName)
	if err != nil {
		return nil, err
	}

	return cluster, nil
}
