package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

func tableMongoDBAtlasCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mongodbatlas_cluster",
		Description: "MongoDB Atlas Cluster is a NoSQL Database-as-a-Service offering in the public cloud (available in Microsoft Azure, Google Cloud Platform, Amazon Web Services).",
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
				Description: "Unique 24-hexadecimal digit string that identifies the cluster.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The name of the cluster as it appears in Atlas.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "project_id",
				Description: "Unique identifier of the project that this cluster belongs to",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GroupID"),
			},
			{
				Name:        "auto_scaling",
				Description: "Collection of settings that configures auto-scaling information for the cluster.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "bi_connector_config",
				Description: "Configuration settings applied to BI Connector for Atlas on this cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("BiConnector"),
			},
			{
				Name:        "cluster_type",
				Description: "Type of the cluster",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "disk_size_gb",
				Description: "Capacity, in gigabytes, of the host's root volume. Increase this number to add capacity, up to a maximum possible value of 4096 (i.e., 4 TB). This value must be a positive number.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("DiskSizeGB"),
			},
			{
				Name:        "encryption_at_rest_provider",
				Description: "Cloud service provider that offers Encryption at Rest.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "labels",
				Description: "Collection of key-value pairs that tag and categorize the cluster. Each key and value has a maximum length of 255 characters.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "mongodb_version",
				Description: "Version of MongoDB that the cluster is running, in X.Y.Z format.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("MongoDBVersion"),
			},
			{
				Name:        "mongodb_major_version",
				Description: "MongoDB Version of the cluster.",
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
				Description: "Positive integer that specifies the number of shards for a sharded cluster. If this is set to 1, the cluster is a replica set. If this is set to 2 or higher, the cluster is a sharded cluster with the number of shards specified.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "paused",
				Description: "Flag that indicates whether the cluster has been paused.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "pit_enabled",
				Description: "Flag that indicates whether the cluster uses continuous cloud backups. More information is available at https://www.mongodb.com/docs/atlas/backup/cloud-backup/overview/#continuous-cloud-backups",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "provider_backup_enabled",
				Description: "Flag that indicates if the cluster uses Back Up Your Database Deployment for backups.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "provider_settings",
				Description: "Configuration for the provisioned hosts on which MongoDB runs. The available options are specific to the cloud service provider.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "replication_factor",
				Description: "Number of replica set members. Each member keeps a copy of your databases, providing high availability and data redundancy.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "replication_spec",
				Description: "Configuration of each region in the cluster. Each element in this object represents a region where Atlas deploys your cluster.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "replication_specs",
				Description: "Configuration for each zone in a Global Cluster. Each object in this array represents a zone where Atlas deploys nodes for your Global Cluster.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "srv_address",
				Description: "Connection string for connecting to the Atlas cluster. The +srv modifier forces the connection to use TLS. The mongoURI parameter lists additional options.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "state_name",
				Description: "Condition in which the API resource finds the cluster when you called the resource. The resource returns one of the following states: IDLE, CREATING, UPDATING, DELETING, DELETED, REPAIRING.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "connection_strings",
				Description: "Set of connection strings that your applications use to connect to this cluster.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "version_release_system",
				Description: `Release cadence that Atlas uses for this cluster.`,
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

func listAtlasClusters(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	config := GetConfig(d.Connection)
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_cluster.listAtlasClusters", "connection_error", err)
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
		clusters, _, err := fetchClustersPage(ctx, client, pageNumber, itemsPerPage, *projectId)
		plugin.Logger(ctx).Trace("mongodbatlas_cluster.listAtlasCluster", "cluster length", len(clusters))

		if err != nil {
			plugin.Logger(ctx).Error("mongodbatlas_cluster.listAtlasClusters", "query_error", err)
			return nil, err
		}

		hasNextPage := false

		for _, cluster := range clusters {
			d.StreamListItem(ctx, cluster)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}

			// find the next page
			// for some odd reason, the SDK puts this in the array element
			if !hasNextPage {
				for _, link := range cluster.Links {
					if link.Rel == "next" {
						hasNextPage = true
						break
					}
				}
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

func fetchClustersPage(ctx context.Context, client *mongodbatlas.Client, pageNumber int, itemsPerPage int64, projectId string) ([]mongodbatlas.Cluster, *mongodbatlas.Response, error) {
	plugin.Logger(ctx).Trace("mongodbatlas_cluster.listAtlasCluster", "project_clusters", projectId)
	return client.Clusters.List(ctx, projectId, &mongodbatlas.ListOptions{
		PageNum:      pageNumber,
		ItemsPerPage: int(itemsPerPage),
	})
}

func getAtlasCluster(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_cluster.getAtlasCluster", "connection_error", err)
		return nil, err
	}
	projectId := config.ProjectId
	clusterName := d.KeyColumnQuals["name"].GetStringValue()

	cluster, _, err := client.Clusters.Get(ctx, *projectId, clusterName)
	if err != nil {
		return nil, err
	}

	return cluster, nil
}
