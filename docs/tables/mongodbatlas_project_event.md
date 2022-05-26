# Table: mongodbatlas_project_event

Project Events allows you to list events for the configured project.

## Examples

### Basic Info

```sql
select
  id,
  event_type_name,
  project_id,
  target_username
from
  mongodbatlas_project_event
```

### List all events raised by a specific user

```sql
select
  id,
  event_type_name
from
  mongodbatlas_project_event
where
  target_username='billy@example.com'
```
