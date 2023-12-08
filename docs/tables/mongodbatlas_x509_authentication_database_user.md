---
title: "Steampipe Table: mongodbatlas_x509_authentication_database_user - Query MongoDB Atlas X509 Authentication Database Users using SQL"
description: "Allows users to query X509 Authentication Database Users in MongoDB Atlas, specifically the database user details, providing insights into user authentication and access control."
---

# Table: mongodbatlas_x509_authentication_database_user - Query MongoDB Atlas X509 Authentication Database Users using SQL

X509 Authentication Database User is a resource within MongoDB Atlas that provides a way to manage user authentication via X509 certificates. It helps in maintaining the security of the database by controlling the access of users based on the X509 certificates. This resource is particularly useful in environments where security and access control are critical.

## Table Usage Guide

The `mongodbatlas_x509_authentication_database_user` table provides insights into X509 Authentication Database Users within MongoDB Atlas. As a Database Administrator or Security Analyst, explore user-specific details through this table, including user roles, databases, and associated metadata. Utilize it to uncover information about users, such as their authentication mechanism, the roles assigned to them, and the databases they have access to.

## Examples

### Basic info
Gain insights into the X.509 authentication database user by identifying instances where a specific username is used. This is useful for understanding user activity and managing access controls.

```sql+postgres
select
  id,
  subject
from
  mongodbatlas_x509_authentication_database_user
where
  username = 'billy';
```

```sql+sqlite
select
  id,
  subject
from
  mongodbatlas_x509_authentication_database_user
where
  username = 'billy';
```

### List all X.509 certificates expiring in 15 days
Discover the segments that have MongoDB Atlas X.509 certificates nearing expiration in the next 15 days. This helps in proactive certificate management, preventing potential access issues due to expired certificates.

```sql+postgres
select
  id,
  subject
from
  mongodbatlas_x509_authentication_database_user
where
  not_after < (now() + INTERVAL '15 days');
```

```sql+sqlite
select
  id,
  subject
from
  mongodbatlas_x509_authentication_database_user
where
  not_after < datetime('now', '+15 days');
```

### List all X.509 certificates expiring after 90 days
Discover the segments that hold X.509 certificates with an extended validity period. This can be useful in identifying the certificates that won't require immediate renewal, allowing you to better manage your certificate renewal timelines.

```sql+postgres
select
  id,
  subject
from
  mongodbatlas_x509_authentication_database_user
where
  not_after > (now() + INTERVAL '90 days');
```

```sql+sqlite
select
  id,
  subject
from
  mongodbatlas_x509_authentication_database_user
where
  not_after > datetime('now', '+90 days');
```

### List all X.509 certificates expiring within 2 months
Explore which X.509 certificates are on the verge of expiration. This is crucial to avoid service interruptions due to expired certificates.

```sql+postgres
select
  id,
  subject
from
  mongodbatlas_x509_authentication_database_user
where
  months_until_expiration <= 2;
```

```sql+sqlite
select
  id,
  subject
from
  mongodbatlas_x509_authentication_database_user
where
  months_until_expiration <= 2;
```