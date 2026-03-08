# Database Service

Namespace: `client.DB`

## Management methods
- `ListTables(projectID string)`
- `DropTable(name, projectID string)`
- `Query(sql, projectID string)`
- `ListBackups(projectID string)`
- `CreateBackup(projectID string)`

## Client CRUD builder
- `FromTable(table).Select(projectID)`
- `FromTable(table).SelectByID(id, projectID)`
- `FromTable(table).Insert(payload, projectID)`
- `FromTable(table).Update(id, payload, projectID)`
- `FromTable(table).Delete(id, projectID)`

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
```go
_, _ = client.DB.Query("select * from todos", "")
rows, _ := client.DB.FromTable("todos").Select("")
_ = rows
```
