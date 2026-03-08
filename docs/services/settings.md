# Settings Service

Namespace: `mirrabase.settings`

## Methods
- `get_general(project_id: str | None = None)`
- `update_general(payload: dict, project_id: str | None = None)`
- `list_domains(project_id: str | None = None)`
- `pause(project_id: str | None = None)`
- `delete(project_id: str | None = None)`

## Endpoints
- `GET /v1/settings/general`
- `PUT /v1/settings/general`
- `GET /v1/settings/domains`
- `POST /v1/settings/danger-zone/pause`
- `DELETE /v1/settings/danger-zone/delete`

## Example
```python
settings = mirrabase.settings.get_general()
mirrabase.settings.update_general({"app_name": "My App"})
```
