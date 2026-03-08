# Auth Service

Namespace: `mirrabase.auth`

## Methods
- `signup(payload: dict)`
- `verify_email(payload: dict)`
- `resend_confirmation(payload: dict)`
- `login(payload: dict)`
- `logout()`
- `get_user()`
- `login_and_set_session(payload: dict)`
- `set_session(token: str)`
- `clear_session()`
- `get_session()`

## Endpoints
- `POST /v1/client/auth/signup`
- `POST /v1/client/auth/verify-email`
- `POST /v1/client/auth/resend-confirmation`
- `POST /v1/client/auth/login`
- `POST /v1/client/auth/logout`
- `GET /v1/client/auth/user`

## Example
```python
mirrabase.auth.signup({"email": email, "password": password})
verify = mirrabase.auth.verify_email({"email": email, "token": token})
mirrabase.auth.set_session(verify["session_token"])
me = mirrabase.auth.get_user()
```
