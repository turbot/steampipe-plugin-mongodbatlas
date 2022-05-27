# Table: mongodbatlas_serverless_instance

Serverless instances in MongoDB Atlas are instances which are billed on usage, rather than time like in normal clusters.

Serverless instances provide a limited feature set compared to full-blown clusters in MongoDB Atlas.

## Examples

### Basic info

```sql
select
  *
from
  mongodbatlas_serverless_instance
```

### Get connection details for serverless instances

```sql
select
  id,
  name,
  connection_strings->>'standardSrv' as connstr_standardsrv,
  connection_strings->>'standard' as connstr_standard
from
  mongodbatlas_cluster
```

### List instances with provider backups disabled

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
