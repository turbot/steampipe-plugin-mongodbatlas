# Table: mongodbatlas_advanced_cluster

MongoDB Atlas Cluster is a NoSQL Database-as-a-Service offering in the public cloud (available in Microsoft Azure, Google Cloud Platform, Amazon Web Services).

## Examples

### List all MongoDB advanced clusters

```sql
select
  id,
  name,
  project_id
from
  mongodbatlas_advanced_cluster
```

### Get connection details for all advanced clusters

```sql
select
  id,
  name,
  connection_strings['standardSrv'] as connstr_standardsrv,
  connection_strings['standard'] as connstr_standard
from
    mongodbatlas_cluster
```

### Get all advanced clusters which are replica sets

```sql
select
  id,
  name
from
  mongodbatlas_cluster
where
  num_shards = 1
```
