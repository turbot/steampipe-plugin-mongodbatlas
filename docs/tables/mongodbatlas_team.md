# Table: mongodbatlas_team

Teams enable you to grant project access roles to multiple users. You add any number of organization users to a team. You grant a team roles for specific projects. All members of a team share the same project access.

Needs `Organization Owner` access in the provided key pair.

## Example

### List all teams in the parent org

```sql
select 
    id, name
from 
    mongodbatlas_team
```