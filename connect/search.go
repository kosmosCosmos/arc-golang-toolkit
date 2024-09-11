package connect

import (
	"fmt"
	"github.com/meilisearch/meilisearch-go"
)

// MeilisearchClientConfig holds the configuration for a Meilisearch client
type MeilisearchClientConfig struct {
	Host   string
	APIKey string
}

// NewMeilisearchClient creates and returns a new Meilisearch client
func NewMeilisearchClient(config MeilisearchClientConfig) (*meilisearch.ServiceManager, error) {
	client := meilisearch.New(config.Host, meilisearch.WithAPIKey(config.APIKey))

	// 测试连接
	_, err := client.Health()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Meilisearch: %w", err)
	}

	return &client, nil
}
