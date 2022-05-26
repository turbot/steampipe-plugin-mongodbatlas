# Table: mongodbatlas_custom_db_role

Custom roles supports a subset of MongoDB privilege actions. These are defined at the project level, for all clusters in the project.

Using custom MongoDB enables you to specify custom sets of actions which cannot be described by the built-in Atlas database user privileges.

## Examples

### Basic info

```sql
select
  role_name,
  actions
from
  mongodbatlas_custom_db_role
```

### List roles which have the 'FIND' action defined

```sql
select
  role_name
from
  mongodbatlas_custom_db_role as t,
  jsonb_array_elements(t.actions) as a
where
  a ->> 'action' = 'FIND'
```

### List roles which have at least one inhertied role

```sql
select
  role_name
from
  mongodbatlas_custom_db_role
where
  jsonb_array_length(inherited_roles) > 0
```
