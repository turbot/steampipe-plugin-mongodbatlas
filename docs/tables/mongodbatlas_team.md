# Table: mongodbatlas_team

Teams enable you to grant project access roles to multiple users. You add any number of organization users to a team. You grant a team roles for specific projects. All members of a team share the same project access.

Needs `Organization Owner` access in the provided key pair.

## Examples

### Basic info

```sql
select
  id,
  name
from
  mongodbatlas_team
```

### List users in all teams

```sql
select
  id,
  name,
  jsonb_array_elements(users) as user
from
  mongodbatlas_team
```
