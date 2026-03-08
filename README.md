# `mirrabase-sdk`
Official Python SDK for Mirrabase Core.

## Install
```bash
pip install mirrabase-sdk
```

## Local setup (venv)
```bash
python3.10 -m venv env
source env/bin/activate
pip install -e .[dev]
```

## Quick Start
```python
from mirrabase_sdk import MirrabaseClient, MirrabaseClientConfig

mirrabase = MirrabaseClient(
    MirrabaseClientConfig(
        base_url="http://localhost:8000",
        api_key="mirrabase-dev-public-key",
    )
)

projects = mirrabase.projects.list()
project = projects[0] if projects else mirrabase.projects.create({"name": "My Project"})
mirrabase.set_project_id(project["id"])

rows = mirrabase.db.from_table("todos").select()
print(rows)
```

## Documentation Index
- [Getting Started](docs/guides/getting-started.md)
- [Configuration](docs/guides/configuration.md)
- [Error Handling](docs/guides/error-handling.md)
- [Testing](docs/guides/testing.md)
- [Auth Service](docs/services/auth.md)
- [Projects Service](docs/services/projects.md)
- [Members Service](docs/services/members.md)
- [Invitations Service](docs/services/invitations.md)
- [Permissions Service](docs/services/permissions.md)
- [Database Service](docs/services/database.md)
- [Storage Service](docs/services/storage.md)
- [RAG Service](docs/services/rag.md)
- [API Keys Service](docs/services/api-keys.md)
- [Settings Service](docs/services/settings.md)
