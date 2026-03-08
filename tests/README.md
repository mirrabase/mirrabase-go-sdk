# Consumer Tests

Folder ini memvalidasi bahwa SDK bisa dipakai dari module Go lain menggunakan `go get`.

## Jalankan test consumer
```bash
cd tests
go test ./...
```

## Simulasi install dengan `go get`
```bash
mkdir -p /tmp/mirrabase-consumer && cd /tmp/mirrabase-consumer
go mod init example.com/mirrabase-consumer
go get github.com/mirrabase/mirrabase-go-sdk@latest
```

Catatan: di repo ini, `tests/go.mod` memakai `replace` ke `../` agar test bisa berjalan lokal sebelum publish.
