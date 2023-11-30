---
title: "Steampipe Table: mongodbatlas_cluster - Query MongoDB Atlas Clusters using SQL"
description: "Allows users to query MongoDB Atlas Clusters, specifically detailed information about each cluster, including its configuration, status, and statistics."
---

# Table: mongodbatlas_cluster - Query MongoDB Atlas Clusters using SQL

MongoDB Atlas is a fully managed cloud database service for modern applications. It provides an easy way to deploy, operate, and scale MongoDB in the cloud, removing the operational burden from your shoulders and making it easy to focus on building your applications. MongoDB Atlas Clusters are the heart of your MongoDB deployments, where your data is stored and from where it is served.

## Table Usage Guide

The `mongodbatlas_cluster` table provides insights into MongoDB Atlas Clusters. As a Database Administrator or Developer, you can explore cluster-specific details through this table, including configuration, status, and statistics. Utilize it to uncover information about clusters, such as their current operational status, the configuration settings in use, and performance statistics, which can be useful for monitoring and troubleshooting.

## Examples

### Basic info
Explore which MongoDB Atlas clusters are associated with specific project IDs to streamline project management and resource allocation.

```sql
select
  id,
  name,
  project_id
from
  mongodbatlas_cluster;
```

### Get auto-scaling details of all clusters
Explore the auto-scaling configurations of your clusters to understand if automatic indexing, disk space adjustments, and scale-down options are enabled. This can be useful in managing resources and optimizing performance in your database environment.

```sql
select
  id,
  name,
  auto_scaling ->> 'autoIndexingEnabled' as auto_scaling_auto_indexing_enabled,
  auto_scaling ->> 'diskGBEnabled' as auto_scaling_diskgb_enabled,
  auto_scaling -> 'compute' ->> 'enabled' as auto_scaling_compute_enabled,
  auto_scaling -> 'compute' ->> 'scaleDownEnabled' as autos_caling_compute_scale_down_enabled
from
  mongodbatlas_cluster;
```

### Get connection details for clusters
Discover the connection details for clusters to better manage and troubleshoot your MongoDB Atlas setup. This helps in simplifying the process of connecting to your clusters.

```sql
select
  id,
  name,
  connection_strings ->> 'standardSrv' as conn_str_standard_srv,
  connection_strings ->> 'standard' as conn_str_standard
from
  mongodbatlas_cluster;
```

### Get all clusters which are replica sets
Determine the areas in which MongoDB Atlas clusters are functioning as replica sets. This can help in managing resource allocation and understanding the distribution of your database workloads.

```sql
select
  id,
  name
from
  mongodbatlas_cluster
where
  num_shards = 1;
```

### List clusters with provider backups disabled
Discover the segments that have provider backups disabled in MongoDB Atlas clusters. This can help identify potential risk areas where data might not be recoverable in the event of a system failure.

```sql
select
  name,
  cluster_type,
  provider_backup_enabled
from
  mongodbatlas_cluster
where
  provider_backup_enabled = false;
```