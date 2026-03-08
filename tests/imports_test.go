package tests

import (
	"testing"
	"time"

	"github.com/mirrabase/mirrabase-go-sdk"
	"github.com/mirrabase/mirrabase-go-sdk/pkg/apikeys"
	"github.com/mirrabase/mirrabase-go-sdk/pkg/auth"
	"github.com/mirrabase/mirrabase-go-sdk/pkg/database"
	"github.com/mirrabase/mirrabase-go-sdk/pkg/invitations"
	"github.com/mirrabase/mirrabase-go-sdk/pkg/members"
	"github.com/mirrabase/mirrabase-go-sdk/pkg/permissions"
	"github.com/mirrabase/mirrabase-go-sdk/pkg/projects"
	"github.com/mirrabase/mirrabase-go-sdk/pkg/rag"
	"github.com/mirrabase/mirrabase-go-sdk/pkg/settings"
	"github.com/mirrabase/mirrabase-go-sdk/pkg/storage"
)

func TestImportsAndConstructors(t *testing.T) {
	_ = auth.New()
	_ = projects.New()
	_ = members.New()
	_ = invitations.New()
	_ = permissions.New()
	_ = database.New()
	_ = storage.New()
	_ = rag.New()
	_ = apikeys.New()
	_ = settings.New()

	c := mirrabase.NewClient(mirrabase.ClientConfig{
		BaseURL: "http://localhost:8000",
		APIKey:  "mirrabase-dev-public-key",
		Timeout: 10 * time.Second,
	})

	if c.Auth == nil || c.Projects == nil || c.DB == nil {
		t.Fatal("client services should be initialized")
	}
}
