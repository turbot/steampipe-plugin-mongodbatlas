---
title: "Steampipe Table: mongodbatlas_serverless_instance - Query MongoDB Atlas Serverless Instances using SQL"
description: "Allows users to query MongoDB Atlas Serverless Instances, providing insights into the configuration, status, and performance of these resources."
---

# Table: mongodbatlas_serverless_instance - Query MongoDB Atlas Serverless Instances using SQL

MongoDB Atlas Serverless Instances are a part of MongoDB's fully managed cloud database service. This service allows you to deploy, operate, and scale MongoDB databases in the cloud. Serverless instances provide on-demand, pay-per-use access to your MongoDB data, scaling automatically to meet your application's read and write workload needs.

## Table Usage Guide

The `mongodbatlas_serverless_instance` table provides insights into the configuration and performance of MongoDB Atlas Serverless Instances. As a Database Administrator or Developer, you can use this table to examine details about these instances, including their configuration, status, and associated metrics. This can help you optimize resource usage, troubleshoot performance issues, and ensure the smooth operation of your MongoDB databases in the cloud.

## Examples

### Basic info
Explore which serverless instances are in use in your MongoDB Atlas setup. This can help you monitor and manage your resources more effectively.

```sql+postgres
select
  id,
  name
from
  mongodbatlas_serverless_instance;
```

```sql+sqlite
select
  id,
  name
from
  mongodbatlas_serverless_instance;
```

### Get connection details for serverless instances
Explore the connection details for your serverless instances to understand how to connect to them in different scenarios. This can be particularly useful when setting up new applications or troubleshooting connectivity issues.

```sql+postgres
select
  id,
  name,
  connection_strings ->> 'standardSrv' as conn_str_standard_srv,
  connection_strings ->> 'standard' as conn_str_standard
from
  mongodbatlas_serverless_instance;
```

```sql+sqlite
select
  id,
  name,
  json_extract(connection_strings, '$.standardSrv') as conn_str_standard_srv,
  json_extract(connection_strings, '$.standard') as conn_str_standard
from
  mongodbatlas_serverless_instance;
```

### List instances with provider backups disabled
Explore which serverless instances have provider backups disabled to identify potential data loss risks and prioritize areas for improved data security.

```sql+postgres
select
  name,
  cluster_type,
  provider_backup_enabled
from
  mongodbatlas_serverless_instance
where
  provider_backup_enabled = false;
```

```sql+sqlite
select
  name,
  cluster_type,
  provider_backup_enabled
from
  mongodbatlas_serverless_instance
where
  provider_backup_enabled = 0;
```