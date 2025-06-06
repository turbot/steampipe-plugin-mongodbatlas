## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#36](https://github.com/turbot/steampipe-plugin-mongodbatlas/pull/36))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#36](https://github.com/turbot/steampipe-plugin-mongodbatlas/pull/36))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#34](https://github.com/turbot/steampipe-plugin-mongodbatlas/pull/34))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#34](https://github.com/turbot/steampipe-plugin-mongodbatlas/pull/34))

## v0.4.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#30](https://github.com/turbot/steampipe-plugin-mongodbatlas/pull/30))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#30](https://github.com/turbot/steampipe-plugin-mongodbatlas/pull/30))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-mongodbatlas/blob/main/docs/LICENSE). ([#30](https://github.com/turbot/steampipe-plugin-mongodbatlas/pull/30))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#29](https://github.com/turbot/steampipe-plugin-mongodbatlas/pull/29))

## v0.3.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#18](https://github.com/turbot/steampipe-plugin-mongodbatlas/pull/18))

## v0.3.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#16](https://github.com/turbot/steampipe-plugin-mongodbatlas/pull/16))
- Recompiled plugin with Go version `1.21`. ([#16](https://github.com/turbot/steampipe-plugin-mongodbatlas/pull/16))

## v0.2.0 [2023-04-10]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#10](https://github.com/turbot/steampipe-plugin-mongodbatlas/pull/10))

## v0.1.1 [2022-09-27]

_Bug fixes_

- Fixed inconsistent table names in `mongodbatlas_org_event` and `mongodbatlas_project_event` tables. (([#7](https://github.com/turbot/steampipe-plugin-mongodbatlas/pull/7))

## v0.1.0 [2022-09-27]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. (([#5](https://github.com/turbot/steampipe-plugin-mongodbatlas/pull/5))
- Recompiled plugin with Go version `1.19`. ([#5](https://github.com/turbot/steampipe-plugin-mongodbatlas/pull/5))

## v0.0.1 [2022-06-02]

_What's new?_

- New tables added

  - [mongodbatlas_cluster](https://hub.steampipe.io/plugins/turbot/mongodbatlas/tables/mongodbatlas_cluster)
  - [mongodbatlas_container](https://hub.steampipe.io/plugins/turbot/mongodbatlas/tables/mongodbatlas_container)
  - [mongodbatlas_custom_db_role](https://hub.steampipe.io/plugins/turbot/mongodbatlas/tables/mongodbatlas_custom_db_role)
  - [mongodbatlas_database_user](https://hub.steampipe.io/plugins/turbot/mongodbatlas/tables/mongodbatlas_database_user)
  - [mongodbatlas_org](https://hub.steampipe.io/plugins/turbot/mongodbatlas/tables/mongodbatlas_org)
  - [mongodbatlas_org_event](https://hub.steampipe.io/plugins/turbot/mongodbatlas/tables/mongodbatlas_org_event)
  - [mongodbatlas_project](https://hub.steampipe.io/plugins/turbot/mongodbatlas/tables/mongodbatlas_project)
  - [mongodbatlas_project_event](https://hub.steampipe.io/plugins/turbot/mongodbatlas/tables/mongodbatlas_project_event)
  - [mongodbatlas_project_ip_access_list](https://hub.steampipe.io/plugins/turbot/mongodbatlas/tables/mongodbatlas_project_ip_access_list)
  - [mongodbatlas_serverless_instance](https://hub.steampipe.io/plugins/turbot/mongodbatlas/tables/mongodbatlas_serverless_instance)
  - [mongodbatlas_team](https://hub.steampipe.io/plugins/turbot/mongodbatlas/tables/mongodbatlas_team)
  - [mongodbatlas_x509_authentication_database_user](https://hub.steampipe.io/plugins/turbot/mongodbatlas/tables/mongodbatlas_x509_authentication_database_user)
