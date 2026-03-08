# Storage Service

Namespace: `mirrabase.storage`

## Methods
- `list_buckets(project_id: str | None = None)`
- `create_bucket(payload: dict, project_id: str | None = None)`
- `delete_bucket(bucket_id: str, project_id: str | None = None)`
- `list_files(params: dict | None = None)`
- `upload_object(bucket: str, path: str, file_path: str, project_id: str | None = None)`
- `download_object(bucket: str, path: str, project_id: str | None = None)`

## Endpoints
- `GET /v1/storage/buckets`
- `POST /v1/storage/buckets`
- `DELETE /v1/storage/buckets/:id`
- `GET /v1/storage/files`
- `POST /v1/client/storage/object/:bucket/:path`
- `GET /v1/client/storage/object/:bucket/:path`

## Example
```python
mirrabase.storage.create_bucket({"name": "assets", "is_public": True})
mirrabase.storage.upload_object("assets", "docs/a.txt", "./a.txt")
data = mirrabase.storage.download_object("assets", "docs/a.txt")
```
