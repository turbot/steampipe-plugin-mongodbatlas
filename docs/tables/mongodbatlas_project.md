---
title: "Steampipe Table: mongodbatlas_project - Query MongoDB Atlas Projects using SQL"
description: "Allows users to query MongoDB Atlas Projects, providing detailed information about each project's configuration, status, and associated resources."
---

# Table: mongodbatlas_project - Query MongoDB Atlas Projects using SQL

A MongoDB Atlas Project is a logical grouping of resources that are deployed in MongoDB Atlas, MongoDB's fully managed cloud database service. Projects in MongoDB Atlas serve as a means to manage access and security for the databases and clusters within them. It provides a centralized way to manage database configurations, user roles, IP whitelists, and more.

## Table Usage Guide

The `mongodbatlas_project` table provides insights into projects within MongoDB Atlas. As a database administrator or DevOps engineer, explore project-specific details through this table, including project ID, cluster count, and associated metadata. Utilize it to uncover information about projects, such as their associated clusters, user roles, and IP whitelists.

## Examples

### Basic info
Explore the basic details of your MongoDB Atlas projects by identifying each project's unique identifier and name. This can help you keep track of your projects and maintain an organized database.

```sql
select
  id,
  name
from
  mongodbatlas_project;
```

### List projects with at least 1 cluster
Discover the projects that have one or more clusters associated with them, allowing you to identify areas of resource allocation and usage. This can be beneficial in understanding project resource utilization and managing resources efficiently.

```sql
select
  id,
  name,
  cluster_count
from
  mongodbatlas_project
where
  cluster_count > 0;
```