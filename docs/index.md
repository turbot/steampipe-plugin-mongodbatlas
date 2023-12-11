---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/mongodbatlas.svg"
brand_color: "#00ED64"
display_name: MongoDB Atlas
name: mongodbatlas
description: Steampipe plugin for querying clusters, users, teams and more from MongoDB Atlas.
og_description: Query MongoDB Atlas with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/mongodbatlas-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# MongoDB Atlas + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[MongoDB Atlas](https://www.mongodb.com/atlas) is a multi-cloud data platform powered by MongoDB.

Example query:

```sql
select
  id,
  name
from
  mongodbatlas_project;
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

Download and install the latest MongoDB Atlas plugin

```bash
steampipe plugin install mongodbatlas
```

### Configuration

Installing the latest mongodbatlas plugin will create a config file (`~/.steampipe/config/mongodbatlas.spc`) with a single connection named `mongodbatlas`:

```hcl
connection "mongodbatlas" {
  plugin = "mongodbatlas"

  # See https://www.mongodb.com/docs/atlas/configure-api-access/#create-an-api-key-in-an-organization
  # for information on how to generate API keys.

  # Public key of the API key.
  # Can also be set with the MONGODB_ATLAS_PUBLIC_API_KEY environment variable.
  # public_key = "hnxxxxxo"

  # Private key of the API key.
  # Can also be set with the MONGODB_ATLAS_PRIVATE_API_KEY environment variable.
  # private_key = "xxxxxxxx-xxxx-4xxx-axxx-dxxxxxxxd9fc"
}
```

- `public_key` - (optional) The API public key from the MongoDB Atlas console. Can also be set with the `MONGODB_ATLAS_PUBLIC_API_KEY` environment variable.
- `private_key` - (optional) The API private key from the MongoDB Atlas console. Can also be set with the `MONGODB_ATLAS_PRIVATE_API_KEY` environment variable.


