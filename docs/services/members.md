# Members Service

Namespace: `mirrabase.members`

## Methods
- `list(project_id: str | None = None)`
- `update_role(user_id: str, role: str, project_id: str | None = None)`
- `remove(user_id: str, project_id: str | None = None)`

## Endpoints
- `GET /v1/projects/:project_id/members`
- `PATCH /v1/projects/:project_id/members/:user_id`
- `DELETE /v1/projects/:project_id/members/:user_id`

## Example
```python
mirrabase.members.list()
mirrabase.members.update_role(user_id, "admin")
mirrabase.members.remove(user_id)
```
