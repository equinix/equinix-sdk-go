package equinixoauth2

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/equinix/equinix-sdk-go/services/accesstokenv1"
	"golang.org/x/oauth2"
)

type ContextAwareTokenSource struct {
	conf   *Config
	client *accesstokenv1.APIClient
	mu     sync.Mutex
	token  *oauth2.Token
}

func (s *ContextAwareTokenSource) TokenWithContext(ctx context.Context) (*oauth2.Token, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.token.Valid() {
		req := accesstokenv1.Oauth2TokenRequest{
			GrantType:    accesstokenv1.PtrString("client_credentials"),
			ClientId:     s.conf.ClientID,
			ClientSecret: s.conf.ClientSecret,
		}
		result, _, err := s.client.OAuth2TokenApi.GetOAuth2AccessToken(ctx).Payload(req).Execute()
		if err != nil {
			return nil, fmt.Errorf("oauth2: failed to fetch token: %s", err)
		}

		token := oauth2.Token{
			AccessToken:  result.AccessToken,
			TokenType:    "Bearer",
			RefreshToken: result.GetRefreshToken(),
		}

		timeout, err := strconv.Atoi(result.TokenTimeout)
		if err != nil {
			timeout = defTokenTimeout
		}
		if timeout != 0 {
			token.Expiry = time.Now().Add(time.Duration(timeout) * time.Second)
		}
		if token.AccessToken == "" {
			return nil, fmt.Errorf("oauth2: server response missing access token")
		}
		s.token = &token
	}

	return s.token, nil
}
