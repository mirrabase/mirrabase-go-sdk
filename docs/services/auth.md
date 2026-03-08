# Auth Service

Namespace: `client.Auth`

## Methods
- `Signup(payload map[string]any)`
- `VerifyEmail(payload map[string]any)`
- `ResendConfirmation(payload map[string]any)`
- `Login(payload map[string]any)`
- `Logout()`
- `GetUser()`
- `LoginAndSetSession(payload map[string]any)`
- `SetSession(token string)`
- `ClearSession()`
- `GetSession()`

## Endpoints
- `POST /v1/client/auth/signup`
- `POST /v1/client/auth/verify-email`
- `POST /v1/client/auth/resend-confirmation`
- `POST /v1/client/auth/login`
- `POST /v1/client/auth/logout`
- `GET /v1/client/auth/user`

## Example
```go
_, _ = client.Auth.Signup(map[string]any{"email": email, "password": password})
verify, _ := client.Auth.VerifyEmail(map[string]any{"email": email, "token": token})
_ = client.Auth.SetSession(verify["session_token"].(string))
me, _ := client.Auth.GetUser()
_ = me
```
