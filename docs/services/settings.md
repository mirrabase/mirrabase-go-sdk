# Settings Service

Namespace: `client.Settings`

## Methods
- `GetGeneral(projectID string)`
- `UpdateGeneral(payload map[string]any, projectID string)`
- `ListDomains(projectID string)`
- `Pause(projectID string)`
- `Delete(projectID string)`

## Endpoints
- `GET /v1/settings/general`
- `PUT /v1/settings/general`
- `GET /v1/settings/domains`
- `POST /v1/settings/danger-zone/pause`
- `DELETE /v1/settings/danger-zone/delete`

## Example
```go
settings, _ := client.Settings.GetGeneral("")
_, _ = client.Settings.UpdateGeneral(map[string]any{"app_name": "My App"}, "")
_ = settings
```
