# Getting Started

## 1. Create client
```go
client := mirrabase.NewClient(mirrabase.ClientConfig{
    BaseURL: "http://localhost:8000",
    APIKey:  "mirrabase-dev-public-key",
})
```

## 2. Auth flow
```go
_, _ = client.Auth.Signup(map[string]any{"email": email, "password": password})
verify, _ := client.Auth.VerifyEmail(map[string]any{"email": email, "token": token})
_ = client.Auth.SetSession(verify["session_token"].(string))
```

## 3. Bootstrap project
```go
projects, _ := client.Projects.List()
if len(projects) > 0 {
    client.SetProjectID(projects[0]["id"].(string))
}
```

## 4. Use services
```go
_, _ = client.DB.FromTable("todos").Select("")
_, _ = client.Storage.ListBuckets("")
_, _ = client.RAG.ListCollections("")
```
