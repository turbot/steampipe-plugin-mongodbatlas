# Table: mongodbatlas_org

An organization is the container of the users, teams, projects etc.

## Examples

### Basic info

```sql
select
  id,
  name
from
  mongodbatlas_org;
```

### List deleted organizations

```sql
select
  id,
  name
from
  mongodbatlas_org
where
  is_deleted;
```
