# Table: mongodbatlas_project

Within an organization, Projects are used to:
* Isolate different environments (for instance, development/qa/prod environments) from each other.
* Associate different users or teams with different environments, or give different permissions to users in different environments.
* Maintain separate cluster security configurations. For example:
    * Create/manage different sets of database user credentials for each project.
    * Isolate networks in different VPCs.
* Create different alert settings. For example, configure alerts for Production environments differently than Development environments.

This table lists a single entry which contains the details of the `project_id` configured.

## Examples

### List configured project

```sql
select
  id,name
from
  mongodbatlas_project;
```
