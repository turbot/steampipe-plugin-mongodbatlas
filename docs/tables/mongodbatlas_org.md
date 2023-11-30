---
title: "Steampipe Table: mongodbatlas_org - Query MongoDB Atlas Organizations using SQL"
description: "Allows users to query MongoDB Atlas Organizations, providing insights into the configuration, roles, and users associated with each organization."
---

# Table: mongodbatlas_org - Query MongoDB Atlas Organizations using SQL

MongoDB Atlas is a cloud database service for applications that work with data in a variety of formats, including structured, semi-structured, and polymorphic. It provides a scalable and secure way for businesses to manage their data. MongoDB Atlas Organizations are higher-level entities that allow for the management of multiple projects and their associated members, teams, and roles.

## Table Usage Guide

The `mongodbatlas_org` table provides insights into the organizations within MongoDB Atlas. As a Database Administrator, explore organization-specific details through this table, including roles, users, and configuration settings. Utilize it to manage and monitor your organizations, such as understanding user permissions, role assignments, and the configuration of each organization.

## Examples

### Basic info
Explore the basic information of your MongoDB Atlas organizations by identifying their unique identifiers and names. This can be beneficial for managing and tracking your organizations in a more organized manner.

```sql
select
  id,
  name
from
  mongodbatlas_org;
```

### List deleted organizations
Discover the segments that have been removed from your MongoDB Atlas organizations. This is useful for auditing and understanding changes in your organizational structure.

```sql
select
  id,
  name
from
  mongodbatlas_org
where
  is_deleted;
```