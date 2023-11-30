---
title: "Steampipe Table: mongodbatlas_team - Query MongoDB Atlas Teams using SQL"
description: "Allows users to query Teams in MongoDB Atlas, specifically the team details including ID, name, and associated users, providing insights into team structures and user access within MongoDB Atlas."
---

# Table: mongodbatlas_team - Query MongoDB Atlas Teams using SQL

MongoDB Atlas Teams is a feature within MongoDB Atlas that allows you to manage user access and permissions. It provides a centralized way to group users and assign specific roles and privileges to them. MongoDB Atlas Teams helps you maintain security and governance by controlling user access to your MongoDB Atlas resources.

## Table Usage Guide

The `mongodbatlas_team` table provides insights into Teams within MongoDB Atlas. As a Database Administrator, explore team-specific details through this table, including team IDs, names, and associated users. Utilize it to uncover information about teams, such as user access levels, the distribution of roles among team members, and the overall structure of teams within your MongoDB Atlas environment.

## Examples

### Basic info
Explore which team IDs and names are available in your MongoDB Atlas environment. This can be useful for managing access and permissions within your database system.

```sql
select
  id,
  name
from
  mongodbatlas_team;
```

### List users in all teams
Identify all users across various teams to understand team composition and collaboration dynamics in your MongoDB Atlas environment.

```sql
select
  id,
  name,
  jsonb_array_elements(users) as user
from
  mongodbatlas_team;
```