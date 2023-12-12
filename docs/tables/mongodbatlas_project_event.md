---
title: "Steampipe Table: mongodbatlas_project_event - Query MongoDB Atlas Project Events using SQL"
description: "Allows users to query Project Events in MongoDB Atlas, specifically the event details, providing insights into project activities and changes."
---

# Table: mongodbatlas_project_event - Query MongoDB Atlas Project Events using SQL

MongoDB Atlas Project Events are a record of actions and changes made within a project in MongoDB Atlas. These events can include modifications to database configurations, user access changes, and general project activity. It's a crucial resource for auditing and monitoring the activities within a MongoDB Atlas project.

## Table Usage Guide

The `mongodbatlas_project_event` table provides insights into Project Events within MongoDB Atlas. As a Database Administrator or Security Analyst, explore event-specific details through this table, including event types, user details, and timestamps. Utilize it to monitor project activities, track configuration changes, and enhance auditing procedures.

## Examples

### Basic info
Explore which events are associated with specific users within a project in MongoDB Atlas. This can be useful for auditing purposes or to monitor user activities within a project.

```sql+postgres
select
  id,
  event_type_name,
  project_id,
  target_username
from
  mongodbatlas_project_event;
```

```sql+sqlite
select
  id,
  event_type_name,
  project_id,
  target_username
from
  mongodbatlas_project_event;
```

### List all events raised by a specific user
Explore all activities initiated by a specific user within a project. This helps in auditing user actions and understanding their impact on project operations.

```sql+postgres
select
  id,
  event_type_name
from
  mongodbatlas_project_event
where
  target_username = 'billy@example.com'
order by
  created;
```

```sql+sqlite
select
  id,
  event_type_name
from
  mongodbatlas_project_event
where
  target_username = 'billy@example.com'
order by
  created;
```

### Check if AWS encryption key needs rotation
Determine if your AWS encryption key requires rotation by identifying if any related events have occurred within the last 24 hours. This allows you to maintain security standards by ensuring encryption keys are updated regularly.

```sql+postgres
select
  count(id) > 0
from
  mongodbatlas_project_event
where
  event_type_name = 'AWS_ENCRYPTION_KEY_NEEDS_ROTATION'
  and created > (now() - INTERVAL '24 hours');
```

```sql+sqlite
select
  count(id) > 0
from
  mongodbatlas_project_event
where
  event_type_name = 'AWS_ENCRYPTION_KEY_NEEDS_ROTATION'
  and created > datetime('now','-24 hours');
```