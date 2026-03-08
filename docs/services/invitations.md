# Invitations Service

Namespace: `mirrabase.invitations`

## Methods
- `list(project_id: str | None = None)`
- `create(payload: dict, project_id: str | None = None)`
- `revoke(invitation_id: str, project_id: str | None = None)`
- `accept(token: str)`

## Endpoints
- `GET /v1/projects/:project_id/invitations`
- `POST /v1/projects/:project_id/invitations`
- `POST /v1/projects/:project_id/invitations/revoke`
- `POST /v1/invitations/accept`

## Example
```python
mirrabase.invitations.create({"email": "dev@x.com", "role": "developer"})
mirrabase.invitations.list()
mirrabase.invitations.revoke(invitation_id)
```
