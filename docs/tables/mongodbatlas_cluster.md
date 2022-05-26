# Table: mongodbatlas_cluster

MongoDB Atlas cluster is a NoSQL Database-as-a-Service offering in the public cloud (available in Microsoft Azure, Google Cloud Platform, Amazon Web Services).

## Examples:

### List all MongoDB Clusters in the Project

```sql
select
    id,
    name,
    project_id
from
    mongodbatlas_cluster
```