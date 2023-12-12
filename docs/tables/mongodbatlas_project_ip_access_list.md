---
title: "Steampipe Table: mongodbatlas_project_ip_access_list - Query MongoDB Atlas Project IP Access Lists using SQL"
description: "Allows users to query Project IP Access Lists in MongoDB Atlas, specifically providing information about the IP addresses or CIDR blocks that can access MongoDB Atlas project resources."
---

# Table: mongodbatlas_project_ip_access_list - Query MongoDB Atlas Project IP Access Lists using SQL

A MongoDB Atlas Project IP Access List is a security feature that allows you to control which IP addresses or CIDR blocks can access your MongoDB Atlas project resources. This feature is designed to help you protect your MongoDB Atlas databases by limiting access to only trusted IP addresses or CIDR blocks. By using this feature, you can significantly reduce the attack surface of your MongoDB Atlas databases.

## Table Usage Guide

The `mongodbatlas_project_ip_access_list` table provides insights into Project IP Access Lists within MongoDB Atlas. As a database administrator or security analyst, explore details about these access lists through this table, including the allowed IP addresses or CIDR blocks, comments, and associated metadata. Utilize it to uncover information about access lists, such as those with specific IP addresses or CIDR blocks, and to ensure that only trusted sources have access to your MongoDB Atlas project resources.

## Examples

### Basic info
Explore which IP addresses have access to your MongoDB Atlas project. This can help in assessing the security and control of who can access your project.

```sql+postgres
select
  ip_address,
  cidr_block
from
  mongodbatlas_project_ip_access_list;
```

```sql+sqlite
select
  ip_address,
  cidr_block
from
  mongodbatlas_project_ip_access_list;
```

### List all IP access lists which belong to a specific `aws security group`
Identify the IP access lists linked to a certain AWS security group to gain insights into the security configurations of your MongoDB Atlas project. This could be particularly useful for reviewing access permissions and managing security measures.

```sql+postgres
select
  project_id,
  ip_address,
  cidr_block
from
  mongodbatlas_project_ip_access_list
where
  aws_security_group = 'sgr_mongodbatlas_sec_group';
```

```sql+sqlite
select
  project_id,
  ip_address,
  cidr_block
from
  mongodbatlas_project_ip_access_list
where
  aws_security_group = 'sgr_mongodbatlas_sec_group';
```

### LIST CIDR details
Gain insights into the details of the IP access list within a MongoDB Atlas project. This can be useful to understand the range of IP addresses that have been given access, which is crucial for maintaining network security and accessibility.

```sql+postgres
select
  project_id,
  cidr_block,
  host(cidr_block),
  broadcast(cidr_block),
  netmask(cidr_block),
  network(cidr_block)
from
  mongodbatlas_project_ip_access_list;
```

```sql+sqlite
Error: SQLite does not support CIDR operations.
```

### List IP access with public CIDR blocks
Identify the projects that have IP access from public CIDR blocks, excluding those from private ranges. This could be used to assess security measures and ensure that only intended networks have access.

```sql+postgres
select
  project_id,
  cidr_block
from
  mongodbatlas_project_ip_access_list
where
  not cidr_block <<= '10.0.0.0/8'
  and not cidr_block <<= '192.168.0.0/16'
  and not cidr_block <<= '172.16.0.0/12';
```

```sql+sqlite
Error: SQLite does not support CIDR operations.
```