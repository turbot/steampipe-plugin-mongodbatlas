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