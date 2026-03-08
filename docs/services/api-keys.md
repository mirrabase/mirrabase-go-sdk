# API Keys Service

Namespace: `client.APIKeys`

## Methods
- `List(projectID string)`
- `Create(payload map[string]any, projectID string)`
- `Revoke(keyID, projectID string)`
- `RateLimits(projectID string)`

## Endpoints
- `GET /v1/api-keys`
- `POST /v1/api-keys`
- `DELETE /v1/api-keys/:id`
- `GET /v1/api-keys/rate-limits`

## Example
```go
key, _ := client.APIKeys.Create(map[string]any{"name": "frontend", "key_type": "anon"}, "")
_, _ = client.APIKeys.RateLimits("")
_ = key
```
