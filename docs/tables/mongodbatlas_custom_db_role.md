# Table: mongodbatlas_custom_db_role

Custom roles supports a subset of MongoDB privilege actions. These are defined at the project level, for all clusters in the project.

## Examples

### List all custom db roles
```sql
select 
    role_name,
    actions
from 
    mongodbatlas_custom_db_role
```
