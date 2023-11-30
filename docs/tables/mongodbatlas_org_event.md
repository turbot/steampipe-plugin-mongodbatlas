---
title: "Steampipe Table: mongodbatlas_org_event - Query MongoDB Atlas Organization Events using SQL"
description: "Allows users to query Organization Events in MongoDB Atlas, providing detailed insights into event activities within the organization."
---

# Table: mongodbatlas_org_event - Query MongoDB Atlas Organization Events using SQL

MongoDB Atlas Organization Events is a feature that allows you to track and monitor various activities within your organization. It provides a comprehensive view of the actions performed by users and teams, and the responses to these actions. This feature plays a crucial role in maintaining the security and integrity of your organization's data in MongoDB Atlas.

## Table Usage Guide

The `mongodbatlas_org_event` table provides insights into event activities within MongoDB Atlas. As an administrator or security professional, you can utilize this table to monitor user activities, track changes, and ensure compliance across your organization. It can be particularly useful for auditing purposes, allowing you to identify any unusual or suspicious actions.

## Examples

### Basic info
Explore which events are occurring within your MongoDB Atlas organization. This allows you to identify the associated project and target user, providing insights into user activity and project engagement.

```sql
select
  id,
  event_type_name,
  project_id,
  target_username
from
  mongodbatlas_org_event;
```

### List all events raised by a specific user
Explore which events have been triggered by a specific user within your organization. This is particularly useful for auditing user activity and understanding individual user behavior.

```sql
select
  id,
  event_type_name
from
  mongodbatlas_org_event
where
  target_username = 'billy@example.com';
```

### List all events where a user has joined a project in the last 24 hours
Determine the instances where a user has joined a project in the past day. This can be useful for tracking user activity and engagement within a specific time frame.

```sql
select
  id,
  event_type_name,
  target_username
from
  mongodbatlas_org_event
where
  event_type_name = 'JOINED_GROUP'
  and created > (now() - INTERVAL '24 hours');
```

### Check if daily bill has exceeded set threshold
Explore how often your daily bill has surpassed a predetermined limit, offering a way to monitor your expenses and manage your budget effectively. This helps in keeping track of your spending and avoiding unexpected high costs.

```sql
select
  count(id)
from
  mongodbatlas_org_event
where
  event_type_name = 'DAILY_BILL_OVER_THRESHOLD';
```