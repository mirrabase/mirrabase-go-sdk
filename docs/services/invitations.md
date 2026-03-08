# Invitations Service

Namespace: `client.Invitations`

## Methods
- `List(projectID string)`
- `Create(payload map[string]any, projectID string)`
- `Revoke(invitationID, projectID string)`
- `Accept(token string)`

## Endpoints
- `GET /v1/projects/:project_id/invitations`
- `POST /v1/projects/:project_id/invitations`
- `POST /v1/projects/:project_id/invitations/revoke`
- `POST /v1/invitations/accept`

## Example
```go
_, _ = client.Invitations.Create(map[string]any{"email": "dev@x.com", "role": "developer"}, "")
_, _ = client.Invitations.List("")
_ = client.Invitations.Revoke(invitationID, "")
```
