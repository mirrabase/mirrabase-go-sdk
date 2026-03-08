package mirrabase_test

import (
	"testing"
	"time"

	"github.com/mirrabase/mirrabase-go-sdk"
)

func TestNewClient(t *testing.T) {
	client := mirrabase.NewClient(mirrabase.ClientConfig{
		BaseURL: "http://localhost:8000",
		APIKey:  "mirrabase-dev-public-key",
		Timeout: 10 * time.Second,
	})

	if client == nil {
		t.Fatal("client must not be nil")
	}
}
