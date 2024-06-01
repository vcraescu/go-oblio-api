package oblio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/vcraescu/go-reqbuilder"
)

const (
	BaseURL = "https://www.oblio.eu/api"
)

type Validator interface {
	Validate() error
}

type TokenStorage interface {
	Set(ctx context.Context, value string, ttl time.Duration) error
	Get(ctx context.Context) (string, error)
}

type Client struct {
	clientID       string
	clientSecret   string
	baseURL        string
	httpClient     *http.Client
	requestBuilder reqbuilder.Builder
	tokenStorage   TokenStorage
	tokenMu        sync.Mutex
}

func NewClient(clientID, clientSecret string, opts ...Option) *Client {
	options := newOptions(opts)

	return &Client{
		clientID:       clientID,
		clientSecret:   clientSecret,
		baseURL:        options.baseURL,
		httpClient:     options.client,
		requestBuilder: reqbuilder.NewBuilder(options.baseURL),
		tokenStorage:   options.tokenStorage,
	}
}

func (c *Client) do(ctx context.Context, builder reqbuilder.Builder, out any) error {
	req, err := builder.Build(ctx)
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return UnmarshalErrorResponse(resp)
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	return nil
}

func (c *Client) doAuthorized(
	ctx context.Context, builder reqbuilder.Builder, accessToken string, resp any,
) error {
	err := c.doWithToken(ctx, builder, accessToken, resp)
	if IsUnauthorizedError(err) && accessToken != "" {
		return c.doWithToken(ctx, builder, "", resp)
	}

	return err
}

func (c *Client) doWithToken(ctx context.Context, builder reqbuilder.Builder, accessToken string, resp any) error {
	accessToken, err := c.getToken(ctx, accessToken)
	if err != nil {
		return fmt.Errorf("getToken: %w", err)
	}

	return c.do(ctx, builder.WithHeaders(reqbuilder.AuthBearerHeader(accessToken)), resp)
}

func (c *Client) getToken(ctx context.Context, accessToken string) (string, error) {
	if accessToken != "" {
		return accessToken, nil
	}

	c.tokenMu.Lock()
	defer c.tokenMu.Unlock()

	token, err := c.tokenStorage.Get(ctx)
	if err != nil {
		if token, err = c.generateToken(ctx); err != nil {
			return "", fmt.Errorf("generateToken: %w", err)
		}
	}

	return token, nil
}

func (c *Client) generateToken(ctx context.Context) (string, error) {
	resp, err := c.GenerateToken(ctx)
	if err != nil {
		return "", fmt.Errorf("generateAuthorizeToken: %w", err)
	}

	if err := c.tokenStorage.Set(ctx, resp.AccessToken, time.Duration(resp.ExpiresIn)-time.Second*10); err != nil {
		return "", fmt.Errorf("set: %w", err)
	}

	return resp.AccessToken, nil
}

func (c *Client) callAPI(ctx context.Context, method, baseURL, endpointSuffix string, req, resp any) error {
	if validator, ok := req.(Validator); ok {
		if err := validator.Validate(); err != nil {
			return err
		}
	}

	var accessToken string

	if v, ok := req.(interface{ GetAccessToken() string }); ok {
		accessToken = v.GetAccessToken()
	}

	endpoint, err := url.JoinPath(baseURL, endpointSuffix)
	if err != nil {
		return fmt.Errorf("joinPath: %w", err)
	}

	builder := c.requestBuilder.
		WithMethod(method).
		WithPath(endpoint)

	if method == http.MethodGet {
		builder = builder.WithParams(req)
	} else {
		builder = builder.WithBody(req)
	}

	if err := c.doAuthorized(ctx, builder, accessToken, resp); err != nil {
		return fmt.Errorf("doAuthorized: %w", err)
	}

	return nil
}
