# Table: mongodbatlas_project_ip_access_list

Atlas only allows client connections to the database deployment from entries in the project's IP access list. Each entry is either a single IP address or a CIDR-notated range of addresses. For AWS clusters with one or more VPC Peering connections to the same AWS region, you can specify a Security Group associated with a peered VPC.

The IP access list applies to all database deployments in the project and can have up to 200 IP access list entries, with the following exception: projects with an existing sharded cluster created before August 25, 2017 can have up to 100 IP access list entries.

Atlas supports creating temporary IP access list entries that expire within a user-configurable 7-day period.

## Examples

### Basic info

```sql
select
  ip_address,
  cidr_block
from
  mongodbatlas_project_ip_access_list
```

### List all IP access lists which belong to a specific `aws security group`

```sql
select
  project_id,
  ip_address,
  cidr_block
from
  mongodbatlas_project_ip_access_list
where
  aws_security_group = 'sgr_mongodbatlas_sec_group'
```

### Show CIDR details

```sql
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

### Show IP accesses with public CIDR blocks

```sql
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
