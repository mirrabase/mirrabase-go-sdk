# Storage Service

Namespace: `client.Storage`

## Methods
- `ListBuckets(projectID string)`
- `CreateBucket(payload map[string]any, projectID string)`
- `DeleteBucket(bucketID, projectID string)`
- `ListFiles(params map[string]string)`
- `UploadObject(bucket, path, filePath, projectID string)`
- `DownloadObject(bucket, path, projectID string)`

## Endpoints
- `GET /v1/storage/buckets`
- `POST /v1/storage/buckets`
- `DELETE /v1/storage/buckets/:id`
- `GET /v1/storage/files`
- `POST /v1/client/storage/object/:bucket/:path`
- `GET /v1/client/storage/object/:bucket/:path`

## Example
```go
_, _ = client.Storage.CreateBucket(map[string]any{"name": "assets", "is_public": true}, "")
_, _ = client.Storage.UploadObject("assets", "docs/a.txt", "./a.txt", "")
_, _ = client.Storage.DownloadObject("assets", "docs/a.txt", "")
```
