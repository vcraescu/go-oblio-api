package oblio

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vcraescu/go-reqbuilder"
)

type generateTokenRequest struct {
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}

func (r *generateTokenRequest) Validate() error {
	if r.ClientID == "" {
		return fmt.Errorf("clientID is empty: %w", ErrInvalidArgument)
	}

	if r.ClientSecret == "" {
		return fmt.Errorf("clientSecret is empty: %w", ErrInvalidArgument)
	}

	return nil
}

type GenerateTokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   string `json:"expires_in,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
	Scope       string `json:"scope,omitempty"`
	RequestTime string `json:"request_time,omitempty"`
}

func (c *Client) GenerateToken(ctx context.Context) (*GenerateTokenResponse, error) {
	req := generateTokenRequest{
		ClientID:     c.clientID,
		ClientSecret: c.clientSecret,
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	builder := c.requestBuilder.
		WithMethod(http.MethodPost).
		WithPath("/authorize/token").
		WithHeaders(reqbuilder.JSONContentHeader).
		WithBody(req)
	resp := &GenerateTokenResponse{}

	if err := c.do(ctx, builder, resp); err != nil {
		return resp, fmt.Errorf("do: %w", err)
	}

	return resp, nil
}
