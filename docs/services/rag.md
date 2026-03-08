# RAG Service

Namespace: `client.RAG`

## Methods
- `ListCollections(projectID string)`
- `CreateCollection(name, projectID string)`
- `DeleteCollection(collectionID, projectID string)`
- `Query(payload map[string]any, projectID string)`
- `Inference(collection, query, projectID string)`

## Endpoints
- `GET /v1/rag/collections`
- `POST /v1/rag/collections`
- `DELETE /v1/rag/collections/:id`
- `POST /v1/rag/query`
- `POST /v1/client/rag/inference/:collection`

## Example
```go
_, _ = client.RAG.CreateCollection("knowledge", "")
result, _ := client.RAG.Inference("knowledge", "what is mirrabase?", "")
_ = result
```
