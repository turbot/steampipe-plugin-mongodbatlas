---
title: "Steampipe Table: mongodbatlas_custom_db_role - Query MongoDB Atlas Custom Database Roles using SQL"
description: "Allows users to query Custom Database Roles in MongoDB Atlas, specifically the role name, database name, and associated actions, providing insights into the role-based access control."
---

# Table: mongodbatlas_custom_db_role - Query MongoDB Atlas Custom Database Roles using SQL

MongoDB Atlas Custom Database Roles represent a collection of permissions that you can assign to users. These roles can be used to grant specific privileges to the users on a specific database. The privileges determine the operations that the users can perform on the database.

## Table Usage Guide

The `mongodbatlas_custom_db_role` table provides insights into custom database roles within MongoDB Atlas. As a database administrator, explore specific details about these roles, including database name, role name, and associated actions. Utilize it to manage and control access to your databases, ensuring that users have the appropriate permissions for their roles.

## Examples

### Basic info
Explore which custom database roles have been assigned in your MongoDB Atlas and identify the associated actions. This can help in understanding user permissions and improve the security management of your database.

```sql+postgres
select
  role_name,
  actions
from
  mongodbatlas_custom_db_role;
```

```sql+sqlite
select
  role_name,
  actions
from
  mongodbatlas_custom_db_role;
```

### List roles which have the 'FIND' action defined
Explore which roles have the 'FIND' action defined to understand the distribution of permissions within your database, which can help in maintaining security and access controls.

```sql+postgres
select
  role_name
from
  mongodbatlas_custom_db_role as r,
  jsonb_array_elements(t.actions) as a
where
  a ->> 'action' = 'FIND';
```

```sql+sqlite
select
  role_name
from
  mongodbatlas_custom_db_role as r,
  json_each(r.actions) as a
where
  json_extract(a.value, '$.action') = 'FIND';
```

### List roles which have at least one inherited role
Discover which roles in your MongoDB Atlas database have inherited roles, allowing you to better understand role hierarchies and permissions in your database system. This can be particularly useful in larger systems where role management may become complex.

```sql+postgres
select
  role_name
from
  mongodbatlas_custom_db_role
where
  jsonb_array_length(inherited_roles) > 0;
```

```sql+sqlite
select
  role_name
from
  mongodbatlas_custom_db_role
where
  json_array_length(inherited_roles) > 0;
```