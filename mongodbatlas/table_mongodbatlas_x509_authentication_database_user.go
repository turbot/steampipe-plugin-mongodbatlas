package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

func tableMongoDBAtlasX509AuthenticationDatabaseUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mongodbatlas_x509_authentication_database_user",
		Description: "Database Users can authenticate against databases using X.509 certificates. Certificates can be managed by Atlas or can be self-managed",
		List: &plugin.ListConfig{
			Hydrate:       listDatabaseUserX509Auth,
			ParentHydrate: listMongoDBAtlasDatabaseUsers,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "Serial number of this certificate.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "subject",
				Description: "Fully distinguished name of the database user to which this certificate belongs. To learn more, see RFC 2253.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "Time when Atlas created this X.509 certificate.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "not_after",
				Description: "Time when this certificate expires.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "months_until_expiration",
				Description: "A number of months that the created certificate is valid for before expiry, up to 24 months.default 3.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "certificate",
				Description: "Certificate Data",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "project_id",
				Description: "Unique identifier of the Atlas project to which this certificate belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GroupID"),
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

func listDatabaseUserX509Auth(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	dbUser := h.Item.(mongodbatlas.DatabaseUser)
	// Create client
	config := GetConfig(d.Connection)
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("x509_authentication_database_user.listDatabaseUserX509Auth", "connection_error", err)
		return nil, err
	}

	projectId := config.ProjectId

	x509Stuff, _, err := client.X509AuthDBUsers.GetUserCertificates(ctx, *projectId, dbUser.Username)
	if err != nil {
		plugin.Logger(ctx).Error("x509_authentication_database_user.listDatabaseUserX509Auth", "query_error", err)
		return nil, err
	}
	for _, uc := range x509Stuff {
		d.StreamListItem(ctx, uc)
		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.QueryStatus.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
