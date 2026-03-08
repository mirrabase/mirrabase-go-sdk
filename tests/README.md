# Consumer Tests

Folder ini memvalidasi import SDK dari sisi consumer code yang berada di package terpisah (`tests`).

## Jalankan test dari root
```bash
go test ./...
```

## Jalankan test langsung dari folder tests
```bash
cd tests
go test ./...
```

## Simulasi install dengan `go get` (project baru)
```bash
mkdir -p /tmp/mirrabase-consumer && cd /tmp/mirrabase-consumer
go mod init example.com/mirrabase-consumer
go get github.com/mirrabase/mirrabase-go-sdk@latest
```
