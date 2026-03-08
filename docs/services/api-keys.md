# API Keys Service

Namespace: `mirrabase.api_keys`

## Methods
- `list(project_id: str | None = None)`
- `create(payload: dict, project_id: str | None = None)`
- `revoke(key_id: str, project_id: str | None = None)`
- `rate_limits(project_id: str | None = None)`

## Endpoints
- `GET /v1/api-keys`
- `POST /v1/api-keys`
- `DELETE /v1/api-keys/:id`
- `GET /v1/api-keys/rate-limits`

## Example
```python
key = mirrabase.api_keys.create({"name": "frontend", "key_type": "anon"})
mirrabase.api_keys.rate_limits()
```
