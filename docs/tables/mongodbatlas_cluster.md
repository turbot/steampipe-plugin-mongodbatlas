# Table: mongodbatlas_cluster

MongoDB Atlas cluster is a NoSQL Database-as-a-Service offering in the public cloud (available in Microsoft Azure, Google Cloud Platform, Amazon Web Services).

## Examples

### Basic info

```sql
select
  id,
  name,
  project_id
from
  mongodbatlas_cluster
```

### Get auto scaling details of all clusters

```sql
select
  id,
  name,
  auto_scaling->>'autoIndexingEnabled' as autoscaling_autoindexingenabled,
  auto_scaling->>'diskGBEnabled' as autoscaling_diskgb_enabled,
  auto_scaling->'compute'->>'enabled' as autoscaling_compute_enabled,
  auto_scaling->'compute'->>'scaleDownEnabled' as autoscaling_compute_scaledownenabled
from
  mongodbatlas_cluster
```

### Get connection details for clusters

```sql
select
  id,
  name,
  connection_strings->>'standardSrv' as connstr_standardsrv,
  connection_strings->>'standard' as connstr_standard
from
  mongodbatlas_cluster
```

### Get all clusters which are replica sets

```sql
select
  id,
  name
from
  mongodbatlas_cluster
where
  num_shards = 1
```

### List clusters with provider backups disabled

```sql
select
  name,
  cluster_type,
  provider_backup_enabled
from
  mongodbatlas_cluster
where
  provider_backup_enabled = false
```
