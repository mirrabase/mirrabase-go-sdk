# Members Service

Namespace: `client.Members`

## Methods
- `List(projectID string)`
- `UpdateRole(userID, role, projectID string)`
- `Remove(userID, projectID string)`

## Endpoints
- `GET /v1/projects/:project_id/members`
- `PATCH /v1/projects/:project_id/members/:user_id`
- `DELETE /v1/projects/:project_id/members/:user_id`

## Example
```go
_, _ = client.Members.List("")
_, _ = client.Members.UpdateRole(userID, "admin", "")
_ = client.Members.Remove(userID, "")
```
