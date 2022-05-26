package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

func tableMongoDBAtlasDatabaseUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mongodbatlas_database_user",
		Description: "A database user has access to databases in a mongodb cluster.",
		List: &plugin.ListConfig{
			Hydrate: listAtlasDatabaseUsers,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAtlasDatabaseUser,
			KeyColumns: plugin.AllColumns([]string{"username", "database_name"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "database_name",
				Description: "Database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "delete_after_date",
				Description: "Timestamp in ISO 8601 date and time format in UTC after which Atlas deletes the temporary access list entry. Atlas returns this field if you specified an expiration date when creating this access list entry.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "project_id",
				Description: "Unique identifier of the project to which this access list entry applies.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GroupID"),
			},
			{
				Name:        "labels",
				Description: "List that contains key-value pairs that tag and categorize the database user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "username",
				Description: "Username needed to authenticate to the MongoDB database or collection.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "roles",
				Description: "List that contains key-value pairs that tag and categorize the database user.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "scopes",
				Description: "List that contains key-value pairs that tag and categorize the database user.",
				Type:        proto.ColumnType_JSON,
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DatabaseName"),
			},
		},
	}
}

func listAtlasDatabaseUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	config := GetConfig(d.Connection)
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_database_user.listAtlasDatabaseUsers", "connection_error", err)
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
		databaseUsers, response, err := fetchDatabaseUsersPage(ctx, client, pageNumber, itemsPerPage, *projectId)

		if err != nil {
			plugin.Logger(ctx).Error("mongodbatlas_database_user.listAtlasDatabaseUsers", "query_error", err)
			return nil, err
		}

		for _, databaseUser := range databaseUsers {
			d.StreamListItem(ctx, databaseUser)
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

func fetchDatabaseUsersPage(ctx context.Context, client *mongodbatlas.Client, pageNumber int, itemsPerPage int64, projectId string) ([]mongodbatlas.DatabaseUser, *mongodbatlas.Response, error) {
	plugin.Logger(ctx).Trace("mongodbatlas_database_user.listAtlasDatabaseUsers", "fetchDatabaseUsersPage", projectId)
	return client.DatabaseUsers.List(ctx, projectId, &mongodbatlas.ListOptions{
		PageNum:      pageNumber,
		ItemsPerPage: int(itemsPerPage),
	})
}

func getAtlasDatabaseUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_database_user.getAtlasDatabaseUser", "connection_error", err)
		return nil, err
	}
	projectId := *config.ProjectId

	username := d.KeyColumnQuals["username"].GetStringValue()
	databaseName := d.KeyColumnQuals["database_name"].GetStringValue()

	databaseUser, _, err := client.DatabaseUsers.Get(ctx, databaseName, projectId, username)
	if err != nil {
		return nil, err
	}

	return databaseUser, nil
}
