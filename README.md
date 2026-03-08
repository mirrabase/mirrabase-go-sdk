# mirrabase-go-sdk

Official Go SDK for Mirrabase Core.

## Install
```bash
go get github.com/mirrabase/mirrabase-go-sdk@latest
```

## Local setup
```bash
go mod tidy
go test ./...
```

## Quick Start
```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mirrabase/mirrabase-go-sdk"
)

func main() {
	client := mirrabase.NewClient(mirrabase.ClientConfig{
		BaseURL: "http://localhost:8000",
		APIKey:  "mirrabase-dev-public-key",
		Timeout: 30 * time.Second,
	})

	projects, err := client.Projects.List()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(projects)
}
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
