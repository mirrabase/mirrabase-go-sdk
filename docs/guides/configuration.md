# Configuration

## Client config
```python
from mirrabase_sdk.client import MirrabaseClientConfig

MirrabaseClientConfig(
    base_url="http://localhost:8000",
    api_key="...",        # optional
    access_token="...",   # optional
    project_id="...",     # optional
    timeout=30.0,
    headers={"x-custom": "1"},
)
```

## Runtime updates
```python
mirrabase.set_api_key("mb_service_xxx")
mirrabase.set_access_token("jwt-token")
mirrabase.set_project_id("00000000-0000-0000-0000-000000000001")
```

## Header strategy
- `Authorization` sent when `access_token` exists.
- `x-mirrabase-key` sent when `api_key` exists.
- If both exist, both are sent.
- `x-project-id` sent when project is resolved from argument or client config.

## Project ID resolution order
1. Method argument `project_id`
2. `client.project_id`
3. Raise `ValueError("Project ID required")` for services that require project context.
