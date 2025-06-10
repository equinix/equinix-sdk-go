package equinixoauth2

import (
	"sync"

	"github.com/equinix/equinix-sdk-go/services/accesstokenv1"
)

const (
	defTokenTimeout = 3600
)

// Config describes oauth2 client credentials flow
type Config struct {
	// ClientID is the application's ID.
	ClientID string
	// ClientSecret is the application's secret.
	ClientSecret string
	// BaseURL is the base endpoint of a server that  token endpoint
	BaseURL string
}

// Error describes oauth2 err
type Error struct {
	Code    string
	Message string
}

// TokenSource returns a TokenSource that returns t until t expires,
// automatically refreshing it as necessary using the provided context and the
// client ID and client secret.
func (c *Config) TokenSource() *ContextAwareTokenSource {
	config := accesstokenv1.NewConfiguration()
	config.Servers = accesstokenv1.ServerConfigurations{
		accesstokenv1.ServerConfiguration{
			URL: c.BaseURL,
		},
	}
	restClient := accesstokenv1.NewAPIClient(config)
	source := ContextAwareTokenSource{
		c,
		restClient,
		sync.Mutex{},
		nil,
	}
	return &source
}
