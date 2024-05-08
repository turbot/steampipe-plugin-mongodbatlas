package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "organization_id",
			Description: "Unique identifier for the organization.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getOrganizationId,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getOrganizationIdMemoized = plugin.HydrateFunc(getOrganizationIdUncached).Memoize(memoize.WithCacheKeyFunction(getOrganizationIdCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getOrganizationId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getOrganizationIdMemoized(ctx, d, h)
}

// Build a cache key for the call to getOrganizationIdCacheKey.
func getOrganizationIdCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getOrganizationId"
	return key, nil
}

func getOrganizationIdUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	client, err := getMongoDBAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getOrganizationIdUncached", "connection_error", err)
		return nil, err
	}

	orgs, _, err := client.Organizations.List(ctx, &mongodbatlas.OrganizationsListOptions{})
	if err != nil {
		plugin.Logger(ctx).Error("getOrganizationIdUncached", "api_error", err)
		return nil, err
	}

	// We will only have the a single organization details per connection.
	// In MongoDB Atlas, the API keys (consisting of a Public Key and a Private Key) are scoped to the organization level, not the individual user level.
	// This means that each set of API keys is specific to one organization.
	return orgs.Results[0].ID, nil
}
