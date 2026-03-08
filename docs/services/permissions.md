# Permissions Service

Namespace: `mirrabase.permissions`

## Methods
- `get(role: str, project_id: str | None = None)`
- `update(role: str, permissions: list[dict], project_id: str | None = None)`

## Endpoints
- `GET /v1/projects/:project_id/roles/:role/permissions`
- `PUT /v1/projects/:project_id/roles/:role/permissions`

## Example
```python
perms = mirrabase.permissions.get("developer")
mirrabase.permissions.update(
    "developer",
    [{"service": "db", "can_read": True, "can_write": True}],
)
```
