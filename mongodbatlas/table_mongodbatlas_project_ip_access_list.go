package mongodbatlas

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
	"go.mongodb.org/atlas/mongodbatlas"
)

func tableMongoDBAtlasProjectIpAccessList(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "mongodbatlas_project_ip_access_list",
		Description: "Atlas only allows client connections to the database deployment from entries in the project's IP access list. Each entry is either a single IP address or a CIDR-notated range of addresses. For AWS clusters with one or more VPC Peering connections to the same AWS region, you can specify a Security Group associated with a peered VPC. The IP access list applies to all database deployments in the project and can have up to 200 IP access list entries, with the following exception: projects with an existing sharded cluster created before August 25, 2017 can have up to 100 IP access list entries.",
		List: &plugin.ListConfig{
			Hydrate:       listMongoDBAtlasProjectIpAccessList,
			ParentHydrate: listMongoDBAtlasProjects,
			KeyColumns:    plugin.OptionalColumns([]string{"project_id"}),
		},
		Get: &plugin.GetConfig{
			Hydrate: getAtlasProjectIpAccessList,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "project_id",
					Require: plugin.Required,
				},
				{
					Name:    "aws_security_group",
					Require: plugin.Optional,
				},
				{
					Name:    "cidr_block",
					Require: plugin.Optional,
				},
				{
					Name:    "ip_address",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "aws_security_group",
				Description: "Unique identifier of AWS security group in this access list entry.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "cidr_block",
				Description: "Range of IP addresses in CIDR notation in this access list entry.",
				Type:        proto.ColumnType_INET,
				Transform:   transform.FromField("CIDRBlock"),
			},
			{
				Name:        "comment",
				Description: "Comment associated with this access list entry.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "delete_after_date",
				Description: "Timestamp in ISO 8601 date and time format in UTC after which Atlas deletes the temporary access list entry. Atlas returns this field if you specified an expiration date when creating this access list entry.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromGo().NullIfZero(),
			},
			{
				Name:        "ip_address",
				Description: "Entry using an IP address in this access list entry.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("IPAddress"),
			},
			{
				Name:        "project_id",
				Description: "Unique identifier of the project to which this access list entry applies.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("GroupID"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("CIDRBlock"),
			},
		},
	}
}

func listMongoDBAtlasProjectIpAccessList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	project := h.Item.(*mongodbatlas.Project)
	client, err := getMongoDBAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_project_ip_access_list.listAtlasProjectIpAccessList", "connection_error", err)
		return nil, err
	}
	// Retrieve the list of incidents
	itemsPerPage := int64(500)
	// Reduce the basic request limit down if the user has only requested a small number of rows
	if d.QueryContext.Limit != nil && *d.QueryContext.Limit < itemsPerPage {
		itemsPerPage = *d.QueryContext.Limit
	}

	pageNumber := 1

	for {
		projectIpAccessLists, response, err := client.ProjectIPAccessList.List(ctx, project.ID, &mongodbatlas.ListOptions{
			PageNum:      pageNumber,
			ItemsPerPage: int(itemsPerPage),
		})

		plugin.Logger(ctx).Trace("mongodbatlas_project_ip_access_list.listAtlasProjectIpAccessList", "cluster length", len(projectIpAccessLists.Results))

		if err != nil {
			plugin.Logger(ctx).Error("mongodbatlas_project_ip_access_list.listAtlasProjectIpAccessList", "query_error", err)
			return nil, err
		}

		for _, projectIPAccessList := range projectIpAccessLists.Results {
			d.StreamListItem(ctx, projectIPAccessList)
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

func getAtlasProjectIpAccessList(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	client, err := getMongoDBAtlasClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("mongodbatlas_project_ip_access_list.getAtlasProjectIpAccessList", "connection_error", err)
		return nil, err
	}
	projectId := d.KeyColumnQuals["project_id"].GetStringValue()

	listName := ""

	if len(d.KeyColumnQuals["aws_security_group"].GetStringValue()) != 0 {
		listName = d.KeyColumnQuals["aws_security_group"].GetStringValue()
	} else if len(d.KeyColumnQuals["cidr_block"].GetInetValue().GetCidr()) != 0 {
		listName = d.KeyColumnQuals["cidr_block"].GetInetValue().GetCidr()
	} else if len(d.KeyColumnQuals["ip_address"].GetInetValue().GetAddr()) != 0 {
		listName = d.KeyColumnQuals["ip_address"].GetInetValue().GetAddr()
	} else {
		return nil, fmt.Errorf("one of 'aws_security_group', 'cidr_block' or 'ip_address' is required")
	}

	ipAccess, _, err := client.ProjectIPAccessList.Get(ctx, projectId, listName)
	if err != nil {
		return nil, err
	}

	return ipAccess, nil
}
