# Table: mongodbatlas_org_event

Org Events allows you to list events for the parent organization of the configured project.

## Examples

### Basic info

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
  target_username = 'billy@example.com'
```

### List all events where a user has joined a project

```sql
select
  id,
  event_type_name,
  target_username
from
  mongodbatlas_org_event
where
  event_type_name = 'JOINED_GROUP'
```

### Check if daily bill has exceeded set threshold

```sql
select
  count(id)
from
  mongodbatlas_org_event
where
  event_type_name = 'DAILY_BILL_OVER_THRESHOLD'
```
