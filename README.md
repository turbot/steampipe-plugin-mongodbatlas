![image](https://hub.steampipe.io/images/plugins/turbot/mongodbatlas-social-graphic.png)

# MongoDB Atlas Plugin for Steampipe

Use SQL to query clusters, users, teams and more from MongoDB Atlas.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/mongodbatlas)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/mongodbatlas/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-mongodbatlas/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install mongodbatlas
```

Run a query:

```sql
select
  id,
  subject
from
  mongodbatlas_x509_authentication_database_user
where
  not_after > (now() + INTERVAL '90 days')
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-mongodbatlas.git
cd steampipe-plugin-mongodbatlas
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/mongodbatlas.spc
```

Try it!

```
steampipe query
> .inspect mongodbatlas
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-mongodbatlas/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [MongoDB Atlas Plugin](https://github.com/turbot/steampipe-plugin-mongodbatlas/labels/help%20wanted)
