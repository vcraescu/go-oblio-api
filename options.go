package oblio

import (
	"github.com/vcraescu/go-oblio-api/token"
	"net/http"
)

type options struct {
	client       *http.Client
	baseURL      string
	tokenStorage TokenStorage
}

type Option interface {
	apply(opts *options)
}

var _ Option = optionFunc(nil)

type optionFunc func(opts *options)

func (fn optionFunc) apply(opts *options) {
	fn(opts)
}

func WithClient(client *http.Client) Option {
	return optionFunc(func(opts *options) {
		opts.client = client
	})
}

func WithBaseURL(baseURL string) Option {
	return optionFunc(func(opts *options) {
		opts.baseURL = baseURL
	})
}

func WithTokenStorage(tokenStorage TokenStorage) Option {
	return optionFunc(func(opts *options) {
		opts.tokenStorage = tokenStorage
	})
}

func newOptions(opts []Option) *options {
	options := &options{
		baseURL:      BaseURL,
		client:       http.DefaultClient,
		tokenStorage: token.NewInMemStorage(),
	}

	for _, opt := range opts {
		opt.apply(options)
	}

	return options
}
