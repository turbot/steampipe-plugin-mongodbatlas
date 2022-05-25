---
organization: Turbot
category: ["SaaS"]
icon_url: "/images/plugins/turbot/panos.svg"
brand_color: "#FA582D"
display_name: MongoDB Atlas
name: mongodbatlas
description: Steampipe plugin to query Mongodb Atlas.
og_description: Query MongoDB Atlas with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/panos-social-graphic.png"
---

# PAN-OS + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[MongoDB Atlas](https://www.mongodb.com/atlas) is a multi-cloud data platform powered by MongoDB.

Example query:

```sql
select 
  id, name 
from 
  mongodbatlas_project
```

```
+--------------------------+-----------+
| id                       | name      |
+--------------------------+-----------+
| 6272xxxxxxxxxxxxxxxxec00 | Project 1 |
+--------------------------+-----------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/mongodbatlas/tables)**

## Get started

### Install

Download and install the latest MongoDB Atlas Plugin

```bash
steampipe plugin install mongodbatlas
```

### Configuration

Installing the latest mongodbatlas plugin will create a config file (`~/.steampipe/config/mongodbatlas.spc`) with a single connection named `mongodbatlas`:

```hcl
connection "mongodbatlas" {
  plugin = "mongodbatlas"
  
  # Public and Private Key Pair with the necessary permissions
  # These can also be 'MONGODB_ATLAS_PUBLIC_KEY' and/or 'MONGODB_ATLAS_PRIVATE_KEY'
  # Consult https://www.mongodb.com/docs/atlas/configure-api-access/#create-an-api-key-in-an-organization on how to generate API keys
  # public_key = "public key here"
  # private_key = "private key here"
  # project_id = "project ID"
}

```

Environment variables are also available as an alternate configuration method:

- `MONGODB_ATLAS_PUBLIC_KEY`
- `MONGODB_ATLAS_PRIVATE_KEY`

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-panos
- Community: [Slack Channel](https://steampipe.io/community/join)
