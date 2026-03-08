# RAG Service

Namespace: `mirrabase.rag`

## Methods
- `list_collections(project_id: str | None = None)`
- `create_collection(name: str, project_id: str | None = None)`
- `delete_collection(collection_id: str, project_id: str | None = None)`
- `query(payload: dict, project_id: str | None = None)`
- `inference(collection: str, query: str, project_id: str | None = None)`

## Endpoints
- `GET /v1/rag/collections`
- `POST /v1/rag/collections`
- `DELETE /v1/rag/collections/:id`
- `POST /v1/rag/query`
- `POST /v1/client/rag/inference/:collection`

## Example
```python
mirrabase.rag.create_collection("knowledge")
result = mirrabase.rag.inference("knowledge", "what is mirrabase?")
```
