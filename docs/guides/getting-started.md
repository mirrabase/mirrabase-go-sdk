# Getting Started

## 1. Create client
```python
from mirrabase_sdk import MirrabaseClient, MirrabaseClientConfig

mirrabase = MirrabaseClient(
    MirrabaseClientConfig(
        base_url="http://localhost:8000",
        api_key="mirrabase-dev-public-key",
    )
)
```

## 2. Auth flow
```python
mirrabase.auth.signup({"email": email, "password": password})
verify = mirrabase.auth.verify_email({"email": email, "token": token})
mirrabase.auth.set_session(verify["session_token"])
```

## 3. Bootstrap project
```python
projects = mirrabase.projects.list()
project = projects[0] if projects else mirrabase.projects.create({"name": "My Project"})
mirrabase.projects.set_active(project["id"])
```

## 4. Use services
```python
mirrabase.db.from_table("todos").select()
mirrabase.storage.list_buckets()
mirrabase.rag.list_collections()
```
