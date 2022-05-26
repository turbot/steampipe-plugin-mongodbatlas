# Table: mongodbatlas_org_event

Org Events allows you to list events for the parent organization of the configured project.

## Examples

### List all events for the project

```sql
select
  id,
  event_type_name,
  project_id,
  target_username
from
  mongodbatlas_org_event
```

### List all events raised by a specific user

```sql
select
  id,
  event_type_name
from
  mongodbatlas_org_event
where
  target_username='billy@example.com'
```
