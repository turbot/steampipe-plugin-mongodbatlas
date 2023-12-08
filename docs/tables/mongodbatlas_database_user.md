---
title: "Steampipe Table: mongodbatlas_database_user - Query MongoDB Atlas Database Users using SQL"
description: "Allows users to query MongoDB Atlas Database Users, specifically providing details about each user's authentication, roles, and associated databases."
---

# Table: mongodbatlas_database_user - Query MongoDB Atlas Database Users using SQL

A MongoDB Atlas Database User is a unique identity recognized by MongoDB Atlas clusters, with associated roles that determine the actions the user can perform on a specific database. Database Users are separate from MongoDB Atlas Organization and Project users. They are used to authenticate applications and services to connect to MongoDB Atlas databases.

## Table Usage Guide

The `mongodbatlas_database_user` table provides insights into database users within MongoDB Atlas. As a database administrator, explore user-specific details through this table, including authentication methods, assigned roles, and the databases they have access to. Utilize it to manage and audit user access, ensuring security and compliance in your MongoDB Atlas environment.

## Examples

### Basic info
Explore which MongoDB Atlas database users are currently active, providing a quick overview of user access and potential security risks. This is useful for administrators seeking to manage user access and maintain database security.

```sql+postgres
select
  id,
  name
from
  mongodbatlas_database_user;
```

```sql+sqlite
select
  id,
  name
from
  mongodbatlas_database_user;
```

### List all scopes for each user
Explore the range of access each user has in your MongoDB Atlas database. This can assist in identifying potential security risks and ensuring appropriate access levels.

```sql+postgres
select
  username,
  jsonb_array_elements(scopes) as scopes
from
  mongodbatlas_database_user;
```

```sql+sqlite
select
  username,
  json_each.value as scopes
from
  mongodbatlas_database_user,
  json_each(scopes);
```

### List all roles for each user
Explore which roles are assigned to each user in your MongoDB Atlas database, helping you to understand user permissions and ensure appropriate access control.

```sql+postgres
select
  username,
  jsonb_array_elements(roles) as roles
from
  mongodbatlas_database_user;
```

```sql+sqlite
select
  username,
  roles.value as roles
from
  mongodbatlas_database_user,
  json_each(roles);
```

### List all database users who have 'readWriteAnyDatabase' role on the database 'admin'
Explore which database users have been granted the 'readWriteAnyDatabase' role on the 'admin' database. This can be useful in assessing user permissions and ensuring appropriate access control within your database environment.

```sql+postgres
select
  username,
  r ->> 'databaseName' as database_name
from
  mongodbatlas_database_user as u,
  jsonb_array_elements(u.roles) as r
where
  r ->> 'roleName' = 'readWriteAnyDatabase'
  AND r ->> 'databaseName' = 'admin';
```

```sql+sqlite
select
  username,
  json_extract(r.value, '$.databaseName') as database_name
from
  mongodbatlas_database_user as u,
  json_each(u.roles) as r
where
  json_extract(r.value, '$.roleName') = 'readWriteAnyDatabase'
  AND json_extract(r.value, '$.databaseName') = 'admin';
```