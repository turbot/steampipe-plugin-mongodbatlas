package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

func tableAtlasX509Auth(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "table_x509_authentication_database_user",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate:       listDatabaseUserX509Auth,
			ParentHydrate: listAtlasDatabaseUsers,
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
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "created_at",
				Description: "Time when Atlas created this X.509 certificate.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreatedAt"),
			},
			{
				Name:        "not_after",
				Description: "Time when this certificate expires.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("NotAfter"),
			},
			{
				Name:        "months_until_expiration",
				Description: "A number of months that the created certificate is valid for before expiry, up to 24 months.default 3.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("MonthsUntilExpiration"),
			},
			{
				Name:        "certificate",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Certificate"),
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
		plugin.Logger(ctx).Error("x509_authentication_database_user.listAtlasProjectIpAccessList", "connection_error", err)
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

	x509Stuff, _, err := fetchX509AuthUser(ctx, client, pageNumber, itemsPerPage, dbUser.Username, *projectId)
	if err != nil {
		plugin.Logger(ctx).Error("x509_authentication_database_user.listAtlasProjectIpAccessList", "query_error", err)
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

func fetchX509AuthUser(ctx context.Context, client *mongodbatlas.Client, pageNumber int, itemsPerPage int64, username, projectId string) ([]mongodbatlas.UserCertificate, *mongodbatlas.Response, error) {
	plugin.Logger(ctx).Trace("x509_authentication_database_user.listDatabaseUserX509Auth", "fetchX509AuthUser", projectId)
	return client.X509AuthDBUsers.GetUserCertificates(ctx, projectId, username)
}
