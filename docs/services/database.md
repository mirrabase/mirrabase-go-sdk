# Database Service

Namespace: `mirrabase.db`

## Management methods
- `list_tables(project_id: str | None = None)`
- `drop_table(name: str, project_id: str | None = None)`
- `query(sql: str, project_id: str | None = None)`
- `list_backups(project_id: str | None = None)`
- `create_backup(project_id: str | None = None)`

## Client CRUD builder
- `from_table(table).select(project_id=None)`
- `from_table(table).select_by_id(id, project_id=None)`
- `from_table(table).insert(payload, project_id=None)`
- `from_table(table).update(id, payload, project_id=None)`
- `from_table(table).delete(id, project_id=None)`

## Endpoints
- `GET /v1/database/tables`
- `DELETE /v1/database/tables/:name`
- `POST /v1/database/query`
- `GET /v1/database/backups`
- `POST /v1/database/backups`
- `GET /v1/client/db/:table`
- `GET /v1/client/db/:table/:id`
- `POST /v1/client/db/:table`
- `PUT /v1/client/db/:table/:id`
- `DELETE /v1/client/db/:table/:id`

## Example
```python
mirrabase.db.query("select * from todos")
rows = mirrabase.db.from_table("todos").select()
```
