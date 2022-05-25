# Table: mongodbatlas_x509_authentication_database_user

[Database Users](/plugins/turbot/pagerduty/tables/database_user) can authenticate against databases using X.509 certificates. Certificates can be managed by Atlas or can be [self-managed](https://www.mongodb.com/docs/atlas/security-self-managed-x509/#set-up-self-managed-x.509-authentication)

## Examples

### List all X.509 Certificates for a user
```sql
select 
    id,subject
from 
    mongodbatlas_x509_auth_database_user
where
    username='billy'
```

### List all X.509 certificates expiring in 15 days
```sql
select 
    id,subject
from 
    mongodbatlas_x509_auth_database_user
where
    not_after < (now() + INTERVAL '15 days')
```