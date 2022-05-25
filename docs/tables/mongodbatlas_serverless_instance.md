# Table: mongodbatlas_serverless_instance

Serverless instances in MongoDB Atlas are instances which are billed on usage, rather than time like in normal clusters. 

Serverless instances provide a limited feature set compared to full-blown clusters in MongoDB Atlas.

## Examples

### List all serverless instances in the project
```sql
select
    *
from
    mongodbatlas_serverless_instance
```
