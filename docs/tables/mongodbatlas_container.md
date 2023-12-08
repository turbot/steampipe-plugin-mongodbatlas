---
title: "Steampipe Table: mongodbatlas_container - Query MongoDB Atlas Containers using SQL"
description: "Allows users to query Containers in MongoDB Atlas, providing detailed insights into the configuration and status of containers."
---

# Table: mongodbatlas_container - Query MongoDB Atlas Containers using SQL

MongoDB Atlas Containers are isolated environments that hold your data and are the base building blocks of MongoDB Atlas clusters. Containers are typically used to store data for a specific project or team. They provide isolation and security for your data, while also allowing for easy scalability and performance tuning.

## Table Usage Guide

The `mongodbatlas_container` table provides insights into the containers within MongoDB Atlas. As a database administrator, you can explore container-specific details through this table, including its configuration, status, and associated metadata. Utilize it to uncover information about containers, such as their capacity, region, and current usage, aiding in efficient management and optimization of your MongoDB Atlas resources.

## Examples

### Basic info
Explore the provider's name and associated network block information in your MongoDB Atlas resources. This can help you manage and organize your resources more effectively.

```sql+postgres
select
  id,
  provider_name,
  atlas_cidr_block
from
  mongodbatlas_container;
```

```sql+sqlite
select
  id,
  provider_name,
  atlas_cidr_block
from
  mongodbatlas_container;
```

### List all peered containers in a specific cloud provider
Explore which containers in a specific cloud provider are peered. This is particularly useful for understanding your cloud network's configuration and identifying any potential security risks.

```sql+postgres
select
  id,
  provider_name,
  atlas_cidr_block
from
  mongodbatlas_container
where
  provider_name = 'aws';
```

```sql+sqlite
select
  id,
  provider_name,
  atlas_cidr_block
from
  mongodbatlas_container
where
  provider_name = 'aws';
```