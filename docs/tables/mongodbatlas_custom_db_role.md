# Table: mongodbatlas_custom_db_role

Custom roles supports a subset of MongoDB privilege actions. These are defined at the project level, for all clusters in the project.

Using custom MongoDB enables you to specify custom sets of actions which cannot be described by the built-in Atlas database user privileges.

## Examples

### List all custom db roles

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
  *
from
  mongodbatlas_custom_db_role t,
  jsonb_array_elements(t.actions) j
where
  j ->> 'action' = 'FIND'
```
