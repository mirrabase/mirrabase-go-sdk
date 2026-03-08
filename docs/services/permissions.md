# Permissions Service

Namespace: `client.Permissions`

## Methods
- `Get(role, projectID string)`
- `Update(role string, permissions []map[string]any, projectID string)`

## Endpoints
- `GET /v1/projects/:project_id/roles/:role/permissions`
- `PUT /v1/projects/:project_id/roles/:role/permissions`

## Example
```go
perms, _ := client.Permissions.Get("developer", "")
_, _ = client.Permissions.Update("developer", []map[string]any{{"service": "db", "can_read": true}}, "")
_ = perms
```
