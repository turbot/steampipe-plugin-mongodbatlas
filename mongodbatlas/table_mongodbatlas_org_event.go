package mongodbatlas

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

func tableMongoDBAtlasOrgEvents(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "mongodbatlas_org_events",
		Description: "Org Events allows you to list events for the parent organization of the configured project.",
		List: &plugin.ListConfig{
			Hydrate:       listMongoDBAtlasOrgEvents,
			ParentHydrate: listMongoDBAtlasOrg,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:      "created",
					Require:   plugin.Optional,
					Operators: []string{">", ">=", "=", "<", "<="},
				},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAtlasOrgEvent,
			KeyColumns: plugin.AllColumns([]string{"alert_id", "org_id"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "Unique identifier for the event",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
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
			},
			{
				Name:        "api_key_id",
				Description: "Unique identifier for the API Key that triggered the event. If this field is present in the response, Atlas does not return the userId field.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("APIKeyID").NullIfZero(),
			},
			{
				Name:        "collection",
				Description: "Name of the collection on which the event occurred. This field can be present when the eventTypeName is either DATA_EXPLORER or DATA_EXPLORER_CRUD.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created",
				Description: "UTC date when the event occurred.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "current_value",
				Description: "Describes the value of the metricName at the time of the event.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "database",
				Description: "Name of the database on which the event occurred. This field can be present when the eventTypeName is either DATA_EXPLORER or DATA_EXPLORER_CRUD.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "event_type_name",
				Description: "Human-readable label that indicates the type of event.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "project_id",
				Description: "The unique identifier for the project in which the event occurred.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GroupID").NullIfZero(),
			},
			{
				Name:        "hostname",
				Description: "The hostname of the Atlas host machine associated to the event.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "invoice_id",
				Description: "The unique identifier of the invoice associated to the event.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_global_admin",
				Description: "Indicates whether the user who triggered the event is a MongoDB employee.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "links",
				Description: "One or more uniform resource locators that link to sub-resources and/or related resources. The Web Linking Specification explains the relation-types between URLs.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "metric_name",
				Description: "The name of the metric associated to the alertId.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "op_type",
				Description: "Type of operation that occurred. This field is present when the eventTypeName is either DATA_EXPLORER or DATA_EXPLORER_CRUD.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "org_id",
				Description: "The unique identifier for the organization in which the event occurred.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "payment_id",
				Description: "The unique identifier of the invoice payment associated to the event.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "port",
				Description: "The port on which the mongod or mongos listens.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "public_key",
				Description: "Public key associated with the API Key that triggered the event. If this field is present in the response, Atlas does not return the 'username' field.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "remote_address",
				Description: "IP address of the userId Atlas user who triggered the event.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "replica_set_name",
				Description: "The name of the replica set associated to the event.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "shard_name",
				Description: "The name of the shard associated to the event.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "target_public_key",
				Description: "The public key of the API Key targeted by the event.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "target_username",
				Description: "The username for the Atlas user targeted by the event.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "team_id",
				Description: "The unique identifier for the Atlas team associated to the event.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("TeamID").NullIfZero(),
			},
			{
				Name:        "user_id",
				Description: "The unique identifier for the Atlas user who triggered the event. If this field is present in the response, Atlas does not return the apiKeyId field.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("UserID").NullIfZero(),
			},
			{
				Name:        "username",
				Description: "The username for the Atlas user who triggered the event. If this field is present in the response, Atlas does not return the publicKey field.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "whitelist_entry",
				Description: "The white list entry of the API Key targeted by the event.",
				Type:        proto.ColumnType_STRING,
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

func listMongoDBAtlasOrgEvents(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	org := h.Item.(*mongodbatlas.Organization)
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
	orgId := org.ID

	for {
		listOptions := &mongodbatlas.EventListOptions{
			ListOptions: mongodbatlas.ListOptions{
				PageNum:      pageNumber,
				ItemsPerPage: int(itemsPerPage),
			},
		}

		if d.Quals["created"] != nil {
			for _, q := range d.Quals["created"].Quals {
				givenTime := q.Value.GetTimestampValue().AsTime()
				switch q.Operator {
				case ">":
				case ">=":
					listOptions.MinDate = givenTime.Format(time.RFC3339)
				case "=":
					listOptions.MinDate = givenTime.Format(time.RFC3339)
					listOptions.MaxDate = givenTime.Format(time.RFC3339)
				case "<=":
				case "<":
					listOptions.MaxDate = givenTime.Format(time.RFC3339)
				}
			}
		}

		orgEvents, response, err := client.Events.ListOrganizationEvents(ctx, orgId, listOptions)

		if err != nil {
			plugin.Logger(ctx).Error("mongodbatlas_org_events.listAtlasOrgEvents", "query_error", err)
			return nil, err
		}

		for _, projectEvent := range orgEvents.Results {
			d.StreamListItem(ctx, projectEvent)
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

func getAtlasOrgEvent(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := getMongodbAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_org_events.getAtlasOrgEvents", "connection_error", err)
		return nil, err
	}

	eventId := d.KeyColumnQuals["event_id"].GetStringValue()
	orgId := d.KeyColumnQuals["org_id"].GetStringValue()

	event, _, err := client.Events.GetOrganizationEvent(ctx, orgId, eventId)

	if err != nil {
		return nil, err
	}

	return event, nil
}
