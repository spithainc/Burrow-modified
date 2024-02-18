// Copyright 2023 Spitha Inc.

package helpers

import (
	"context"

	"github.com/IBM/sarama"
	"golang.org/x/oauth2/clientcredentials"
)

type SOAuthTokenProvider struct {
	*clientcredentials.Config
}

func (s *SOAuthTokenProvider) Token() (*sarama.AccessToken, error) {
	ctx := context.Background()
	token, err := s.Config.Token(ctx)

	if err != nil {
		return nil, err
	}

	accessToken := &sarama.AccessToken{
		Token: token.AccessToken,
	}
	return accessToken, nil
}

func newOAuthTokenProvider(clientID, clientSecret, tokenEndpoint string) sarama.AccessTokenProvider {
	conf := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenEndpoint,
	}

	return &SOAuthTokenProvider{
		Config: conf,
	}
}
