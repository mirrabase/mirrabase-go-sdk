# Testing

## Setup
```bash
go mod tidy
```

## Unit/smoke test
```bash
go test ./...
```

## Consumer import test
```bash
cd tests
go test ./...
```

## E2E env (untuk tahap lanjut)
- `MIRRABASE_BASE_URL` default `http://localhost:8000`
- `MIRRABASE_PUBLIC_KEY` default `mirrabase-dev-public-key`
- `MIRRABASE_SERVICE_KEY` default `mirrabase-dev-security-key`
- `MIRRABASE_PROJECT_ID` default `00000000-0000-0000-0000-000000000001`
