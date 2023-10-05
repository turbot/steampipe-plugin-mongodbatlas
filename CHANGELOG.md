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
