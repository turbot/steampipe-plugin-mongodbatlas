# Table: mongodbatlas_database_user

A database user has access to databases in a mongodb cluster. Each user has a set of roles that provide access to all databases in the project.

## Examples

### Basic info

```sql
select
  id,
  name
from
  mongodbatlas_database_user
```

### List all database users who have 'readWriteAnyDatabase' role on the database 'admin'

```sql
select
  username,
  j ->> 'databaseName' as database_name
from
  mongodbatlas_database_user t,
  jsonb_array_elements(t.roles) j
where
  j ->> 'roleName' = 'readWriteAnyDatabase'
  AND j ->> 'databaseName' = 'admin'
```
