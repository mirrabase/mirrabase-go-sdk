# Testing

## Setup
```bash
python3.10 -m venv env
source env/bin/activate
pip install -e .[dev]
```

## E2E (real backend in Docker)
```bash
pytest tests/e2e -v
```

## E2E env
- `MIRRABASE_BASE_URL` default `http://localhost:8000`
- `MIRRABASE_PUBLIC_KEY` default `mirrabase-dev-public-key`
- `MIRRABASE_SERVICE_KEY` default `mirrabase-dev-security-key`
- `MIRRABASE_PROJECT_ID` default `00000000-0000-0000-0000-000000000001`
