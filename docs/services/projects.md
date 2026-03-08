# Projects Service

Namespace: `mirrabase.projects`

## Methods
- `list()`
- `create(payload: dict)`
- `get(project_id: str)`
- `update(project_id: str, payload: dict)`
- `remove(project_id: str)`
- `set_active(project_id: str)`

## Endpoints
- `GET /v1/projects`
- `POST /v1/projects`
- `GET /v1/projects/:project_id`
- `PATCH /v1/projects/:project_id`
- `DELETE /v1/projects/:project_id`

## Example
```python
projects = mirrabase.projects.list()
project = mirrabase.projects.create({"name": "My Project"})
mirrabase.projects.set_active(project["id"])
```
