package mirrabase

import (
	"time"

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

type ClientConfig struct {
	BaseURL     string
	APIKey      string
	AccessToken string
	ProjectID   string
	Timeout     time.Duration
	Headers     map[string]string
}

type Client struct {
	config      ClientConfig
	Auth        *auth.Service
	Projects    *projects.Service
	Members     *members.Service
	Invitations *invitations.Service
	Permissions *permissions.Service
	DB          *database.Service
	Storage     *storage.Service
	RAG         *rag.Service
	APIKeys     *apikeys.Service
	Settings    *settings.Service
}

func NewClient(cfg ClientConfig) *Client {
	if cfg.Timeout == 0 {
		cfg.Timeout = 30 * time.Second
	}

	return &Client{
		config:      cfg,
		Auth:        auth.New(),
		Projects:    projects.New(),
		Members:     members.New(),
		Invitations: invitations.New(),
		Permissions: permissions.New(),
		DB:          database.New(),
		Storage:     storage.New(),
		RAG:         rag.New(),
		APIKeys:     apikeys.New(),
		Settings:    settings.New(),
	}
}

func (c *Client) SetAPIKey(key string) {
	c.config.APIKey = key
}

func (c *Client) SetAccessToken(token string) {
	c.config.AccessToken = token
}

func (c *Client) SetProjectID(projectID string) {
	c.config.ProjectID = projectID
}

func (c *Client) Config() ClientConfig {
	return c.config
}
