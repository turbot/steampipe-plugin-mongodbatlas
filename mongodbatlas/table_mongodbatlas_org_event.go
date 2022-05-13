package mongodbatlas

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

func tableAtlasOrgEvents(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mongodbatlas_org_events",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate:       listAtlasOrgEvents,
			ParentHydrate: listAtlasProjects,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAtlasOrgEvent,
			KeyColumns: plugin.SingleColumn("alert_id"),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "Unique identifier for the event",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID").NullIfZero(),
			},
			{
				Name:        "alert_id",
				Description: "Unique identifier for the alert associated with the event",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AlertID").NullIfZero(),
			},
			{
				Name:        "alert_config_id",
				Description: "Unique identifier for the alert configuration associated to the alertId",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AlertConfigID").NullIfZero(),
			},
			{
				Name:        "api_key_id",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("APIKeyID").NullIfZero(),
			},
			{
				Name:        "collection",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Collection").NullIfZero(),
			},
			{
				Name:        "created",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Created").NullIfZero(),
			},
			{
				Name:        "current_value",
				Description: "",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("CurrentValue").NullIfZero(),
			},
			{
				Name:        "database",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Database").NullIfZero(),
			},
			{
				Name:        "event_type_name",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("EventTypeName").NullIfZero(),
			},
			{
				Name:        "project_id",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GroupID").NullIfZero(),
			},
			{
				Name:        "hostname",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Hostname").NullIfZero(),
			},
			{
				Name:        "invoice_id",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("InvoiceID").NullIfZero(),
			},
			{
				Name:        "is_global_admin",
				Description: "",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("IsGlobalAdmin").NullIfZero(),
			},
			{
				Name:        "links",
				Description: "",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Links").NullIfZero(),
			},
			{
				Name:        "metric_name",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("MetricName").NullIfZero(),
			},
			{
				Name:        "op_type",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("OpType").NullIfZero(),
			},
			{
				Name:        "org_id",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("OrgID").NullIfZero(),
			},
			{
				Name:        "payment_id",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("PaymentID").NullIfZero(),
			},
			{
				Name:        "port",
				Description: "",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Port").NullIfZero(),
			},
			{
				Name:        "public_key",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("PublicKey").NullIfZero(),
			},
			{
				Name:        "remote_address",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RemoteAddress").NullIfZero(),
			},
			{
				Name:        "replica_set_name",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ReplicaSetName").NullIfZero(),
			},
			{
				Name:        "shard_name",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ShardName").NullIfZero(),
			},
			{
				Name:        "target_public_key",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("TargetPublicKey").NullIfZero(),
			},
			{
				Name:        "target_username",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("TargetUsername").NullIfZero(),
			},
			{
				Name:        "team_id",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("TeamID").NullIfZero(),
			},
			{
				Name:        "user_id",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("UserID").NullIfZero(),
			},
			{
				Name:        "username",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Username").NullIfZero(),
			},
			{
				Name:        "whitelist_entry",
				Description: "",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("WhitelistEntry").NullIfZero(),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID").NullIfZero(),
			},
		},
	}
}

func listAtlasOrgEvents(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	project := h.Item.(*mongodbatlas.Project)
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_org_events.listAtlasOrgEvents", "connection_error", err)
		return nil, err
	}
	// Retrieve the list of incidents
	itemsPerPage := int64(500)
	// Reduce the basic request limit down if the user has only requested a small number of rows
	if d.QueryContext.Limit != nil && *d.QueryContext.Limit < itemsPerPage {
		itemsPerPage = *d.QueryContext.Limit
	}

	pageNumber := 1
	orgId := project.OrgID

	for {
		projectEvents, _, err := fetchOrgEvents(ctx, client, pageNumber, itemsPerPage, orgId)

		if err != nil {
			plugin.Logger(ctx).Error("mongodbatlas_org_events.listAtlasOrgEvents", "query_error", err)
			return nil, err
		}

		for _, projectEvent := range projectEvents.Results {
			d.StreamListItem(ctx, projectEvent)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		// find the next page
		hasNextPage := false

		for _, l := range projectEvents.Links {
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

func fetchOrgEvents(ctx context.Context, client *mongodbatlas.Client, pageNumber int, itemsPerPage int64, orgId string) (*mongodbatlas.EventResponse, *mongodbatlas.Response, error) {
	plugin.Logger(ctx).Trace("mongodbatlas_org_events.listAtlasOrgEvents", "fetchProjectEvents", orgId)
	return client.Events.ListOrganizationEvents(ctx, orgId, &mongodbatlas.EventListOptions{
		ListOptions: mongodbatlas.ListOptions{
			PageNum:      pageNumber,
			ItemsPerPage: int(itemsPerPage),
		},
	})
}

func getAtlasOrgEvent(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_org_events.getAtlasOrgEvents", "connection_error", err)
		return nil, err
	}

	projectId := *config.ProjectId

	project, _, err := client.Projects.GetOneProject(ctx, projectId)
	if err != nil {
		return nil, err
	}

	eventId := d.KeyColumnQuals["event_id"].GetStringValue()

	event, _, err := client.Events.GetOrganizationEvent(ctx, project.OrgID, eventId)

	if err != nil {
		return nil, err
	}

	return event, nil
}
