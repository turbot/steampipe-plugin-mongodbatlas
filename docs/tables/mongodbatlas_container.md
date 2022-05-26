# Table: mongodbatlas_container

Containers in a MongoDB Atlas project allows for cloud provider backed virtual private networking - dubbed as `container network peering` in MongoDB Atlas.

## Examples:

### List all peered containers in the Project
```sql
select
    id,provider_name,atlas_cidr_block
from
    mongodbatlas_container
```

### List all peered containers in a specific cloud provider
```sql
select
    id,provider_name,atlas_cidr_block
from
    mongodbatlas_container
where
    provider_name='aws'
```