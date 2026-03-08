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

	// Auth
	_, _ = client.Auth.Signup(map[string]any{"email": "dev@example.com", "password": "secret"})

	// Projects
	projects, _ := client.Projects.List()
	if len(projects) > 0 {
		if id, ok := projects[0]["id"].(string); ok {
			client.SetProjectID(id)
		}
	}

	// Members / Invitations / Permissions
	_, _ = client.Members.List("")
	_, _ = client.Invitations.List("")
	_, _ = client.Permissions.Get("developer", "")

	// Database
	_, _ = client.DB.Query("select * from todos", "")
	_, _ = client.DB.FromTable("todos").Select("")

	// Storage
	_, _ = client.Storage.ListBuckets("")

	// RAG
	_, _ = client.RAG.ListCollections("")

	// API Keys
	_, _ = client.APIKeys.List("")

	// Settings
	_, _ = client.Settings.GetGeneral("")

	cfg := client.Config()
	fmt.Println("Mirrabase SDK initialized for", cfg.BaseURL)
	log.Println("example finished")
}
