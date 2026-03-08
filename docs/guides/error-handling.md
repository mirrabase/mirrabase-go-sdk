# Error Handling

Untuk fase bootstrap ini, method service mengembalikan `error` standar Go.

## Shape (target)
```go
type Error struct {
    StatusCode int
    RequestID  string
    Body       any
}
```

## Mapping (target)
- `400`: validation/input
- `401`: unauthenticated
- `403`: forbidden
- `404`: not found
- `429`: rate limited
- `5xx`: server error
- timeout: `408`
- network error: `0`

## Example
```go
projects, err := client.Projects.List()
if err != nil {
    log.Printf("projects list failed: %v", err)
    return
}
fmt.Println(projects)
```
