package oblio

import (
	"context"
	"net/http"

	"github.com/vcraescu/go-oblio-api/types"
)

type CreateProformaRequest struct {
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
	NoticeNumber       string              `json:"noticeNumber,omitempty"`
	InternalNote       string              `json:"internalNote,omitempty"`
	DeputyName         string              `json:"deputyName,omitempty"`
	DeputyIdentityCard string              `json:"deputyIdentityCard,omitempty"`
	DeputyAuto         string              `json:"deputyAuto,omitempty"`
	SalesAgent         string              `json:"selesAgent,omitempty"`
	Mentions           string              `json:"mentions,omitempty"`
	WorkStation        string              `json:"workStation,omitempty"`
	SendEmail          bool                `json:"sendEmail,omitempty"`
}

type CreateProformaResponse struct {
	Status

	Data types.Document `json:"data"`
}

func (c *Client) CreateProforma(ctx context.Context, req *CreateProformaRequest) (*CreateProformaResponse, error) {
	resp := &CreateProformaResponse{}

	if err := c.callDocsAPI(ctx, http.MethodPost, "/proforma", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetProforma(ctx context.Context, req *DocumentRequest) (*DocumentResponse, error) {
	resp := &DocumentResponse{}

	if err := c.callDocsAPI(ctx, http.MethodGet, "/proforma", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) CancelProforma(ctx context.Context, req *DocumentRequest) (*DocumentResponse, error) {
	resp := &DocumentResponse{}

	if err := c.callDocsAPI(ctx, http.MethodPut, "/proforma/cancel", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) RestoreProforma(ctx context.Context, req *DocumentRequest) (*DocumentResponse, error) {
	resp := &DocumentResponse{}

	if err := c.callDocsAPI(ctx, http.MethodPut, "/proforma/restore", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) DeleteProforma(ctx context.Context, req *DocumentRequest) (*DocumentResponse, error) {
	resp := &DocumentResponse{}

	if err := c.callDocsAPI(ctx, http.MethodDelete, "/proforma", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
