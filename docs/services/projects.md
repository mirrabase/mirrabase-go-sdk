# Projects Service

Namespace: `client.Projects`

## Methods
- `List()`
- `Create(payload map[string]any)`
- `Get(projectID string)`
- `Update(projectID string, payload map[string]any)`
- `Remove(projectID string)`
- `SetActive(projectID string)`

## Endpoints
- `GET /v1/projects`
- `POST /v1/projects`
- `GET /v1/projects/:project_id`
- `PATCH /v1/projects/:project_id`
- `DELETE /v1/projects/:project_id`

## Example
```go
projects, _ := client.Projects.List()
project, _ := client.Projects.Create(map[string]any{"name": "My Project"})
client.Projects.SetActive(project["id"].(string))
_ = projects
```
