# Table: mongodbatlas_database_user

A database user has access to databases in a mongodb cluster. Each user has a set of roles that provide access to all databases in the project.

## Examples

### Basic info

```sql
select
  id,
  name
from
  mongodbatlas_database_user;
```

### List all scopes for each user

```sql
select
  username,
  jsonb_array_elements(scopes) as scopes
from
  mongodbatlas_database_user;
```

### List all roles for each user

```sql
select
  username,
  jsonb_array_elements(roles) as roles
from
  mongodbatlas_database_user;
```

### List all database users who have 'readWriteAnyDatabase' role on the database 'admin'

```sql
select
  username,
  r ->> 'databaseName' as database_name
from
  mongodbatlas_database_user u,
  jsonb_array_elements(u.roles) r
where
  r ->> 'roleName' = 'readWriteAnyDatabase'
  AND r ->> 'databaseName' = 'admin';
```
