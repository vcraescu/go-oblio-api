package oblio

import (
	"context"
	"net/http"

	"github.com/vcraescu/go-oblio-api/types"
)

type CreateNoticeRequest struct {
	CIF                string              `json:"cif,omitempty"`
	Client             types.Client        `json:"client,omitempty"`
	IssueDate          types.Date          `json:"issueDate,omitempty"`
	DueDate            types.Date          `json:"dueDate,omitempty"`
	SeriesName         string              `json:"seriesName,omitempty"`
	Language           string              `json:"language,omitempty"`
	Precision          types.Int           `json:"precision,omitempty"`
	Currency           string              `json:"currency,omitempty"`
	ExchangeRate       int                 `json:"exchangeRate,omitempty"`
	Products           []types.DocumentRow `json:"products,omitempty"`
	IssuerName         string              `json:"issuerName,omitempty"`
	IssuerID           int64               `json:"issuerId,omitempty"`
	InternalNote       string              `json:"internalNote,omitempty"`
	DeputyName         string              `json:"deputyName,omitempty"`
	DeputyIdentityCard string              `json:"deputyIdentityCard,omitempty"`
	DeputyAuto         string              `json:"deputyAuto,omitempty"`
	SalesAgent         string              `json:"selesAgent,omitempty"`
	Mentions           string              `json:"mentions,omitempty"`
	WorkStation        string              `json:"workStation,omitempty"`
	SendEmail          types.Bool          `json:"sendEmail,omitempty"`
	UseStock           types.Bool          `json:"useStock"`
}

type CreateNoticeResponse struct {
	Status

	Data types.Document `json:"data"`
}

func (c *Client) CreateNotice(ctx context.Context, req *CreateNoticeRequest) (*CreateNoticeResponse, error) {
	resp := &CreateNoticeResponse{}

	if err := c.callDocsAPI(ctx, http.MethodPost, "/notice", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetNotice(ctx context.Context, req *DocumentRequest) (*DocumentResponse, error) {
	resp := &DocumentResponse{}

	if err := c.callDocsAPI(ctx, http.MethodGet, "/notice", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) CancelNotice(ctx context.Context, req *DocumentRequest) (*DocumentResponse, error) {
	resp := &DocumentResponse{}

	if err := c.callDocsAPI(ctx, http.MethodPut, "/notice/cancel", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) RestoreNotice(ctx context.Context, req *DocumentRequest) (*DocumentResponse, error) {
	resp := &DocumentResponse{}

	if err := c.callDocsAPI(ctx, http.MethodPut, "/notice/restore", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) DeleteNotice(ctx context.Context, req *DocumentRequest) (*DocumentResponse, error) {
	resp := &DocumentResponse{}

	if err := c.callDocsAPI(ctx, http.MethodDelete, "/notice", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
