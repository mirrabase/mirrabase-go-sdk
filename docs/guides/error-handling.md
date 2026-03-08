# Error Handling

SDK raises `MirrabaseError` for API and network failures.

## Shape
```python
class MirrabaseError(Exception):
    status_code: int
    request_id: str | None
    rate_limit: RateLimitInfo | None
    body: object
```

## Mapping
- `400`: validation/input
- `401`: unauthenticated
- `403`: forbidden
- `404`: not found
- `429`: rate limited + `rate_limit`
- `5xx`: server error
- timeout: `408`
- network error: `0`

## Example
```python
from mirrabase_sdk import MirrabaseError

try:
    mirrabase.projects.list()
except MirrabaseError as error:
    print(error.status_code, str(error))
```
