# Configuration

## Client config
```go
cfg := mirrabase.ClientConfig{
    BaseURL:     "http://localhost:8000",
    APIKey:      "...", // optional
    AccessToken: "...", // optional
    ProjectID:   "...", // optional
    Timeout:     30 * time.Second,
    Headers:     map[string]string{"x-custom": "1"},
}
client := mirrabase.NewClient(cfg)
```

## Runtime updates
```go
client.SetAPIKey("mb_service_xxx")
client.SetAccessToken("jwt-token")
client.SetProjectID("00000000-0000-0000-0000-000000000001")
```

## Header strategy
- `Authorization` dikirim saat `AccessToken` tersedia.
- `x-mirrabase-key` dikirim saat `APIKey` tersedia.
- Jika keduanya ada, keduanya dikirim.
- `x-project-id` dikirim saat project context tersedia.

## Project ID resolution order
1. Argumen method `projectID`
2. `ClientConfig.ProjectID`
3. Return error `project id required` untuk endpoint yang wajib project context.
