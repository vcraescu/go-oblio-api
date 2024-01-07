package oblio

import (
	"context"
	"net/http"

	"github.com/vcraescu/go-oblio-api/types"
)

type CreateInvoiceRequest struct {
	CIF                string                  `json:"cif,omitempty"`
	Client             types.Client            `json:"client,omitempty"`
	IssueDate          types.Date              `json:"issueDate,omitempty"`
	DueDate            types.Date              `json:"dueDate,omitempty"`
	DeliveryDate       types.Date              `json:"deliveryDate,omitempty"`
	CollectDate        types.Date              `json:"collectDate,omitempty"`
	SeriesName         string                  `json:"seriesName,omitempty"`
	Language           string                  `json:"language,omitempty"`
	Precision          types.Int               `json:"precision,omitempty"`
	Currency           string                  `json:"currency,omitempty"`
	ExchangeRate       int                     `json:"exchangeRate,omitempty"`
	Products           []types.DocumentRow     `json:"products,omitempty"`
	IssuerName         string                  `json:"issuerName,omitempty"`
	IssuerID           string                  `json:"issuerId,omitempty"`
	NoticeNumber       string                  `json:"noticeNumber,omitempty"`
	InternalNote       string                  `json:"internalNote,omitempty"`
	DeputyName         string                  `json:"deputyName,omitempty"`
	DeputyIdentityCard string                  `json:"deputyIdentityCard,omitempty"`
	DeputyAuto         string                  `json:"deputyAuto,omitempty"`
	SalesAgent         string                  `json:"selesAgent,omitempty"`
	Mentions           string                  `json:"mentions,omitempty"`
	WorkStation        string                  `json:"workStation,omitempty"`
	Collect            types.Collect           `json:"collect,omitempty"`
	ReferenceDocument  types.ReferenceDocument `json:"referenceDocument,omitempty"`
	SendEmail          types.Bool              `json:"sendEmail,omitempty"`
	UseStock           types.Bool              `json:"useStock,omitempty"`
}

type CreateInvoiceResponse struct {
	Status

	Data types.Document `json:"data"`
}

func (c *Client) CreateInvoice(ctx context.Context, req *CreateInvoiceRequest) (*CreateInvoiceResponse, error) {
	resp := &CreateInvoiceResponse{}

	if err := c.callDocsAPI(ctx, http.MethodPost, "/invoice", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

type CollectRequest struct {
	CIF        string          `json:"cif,omitempty"`
	SeriesName string          `json:"seriesName,omitempty"`
	Number     string          `json:"number,omitempty"`
	Collects   []types.Collect `json:"collects,omitempty"`
}

type CollectResponse struct {
	Status

	Data types.Document `json:"data"`
}

func (c *Client) Collect(ctx context.Context, req *CollectRequest) (*CollectResponse, error) {
	resp := &CollectResponse{}

	if err := c.callDocsAPI(ctx, http.MethodPut, "/invoice/collect", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetInvoice(ctx context.Context, req *DocumentRequest) (*DocumentResponse, error) {
	resp := &DocumentResponse{}

	if err := c.callDocsAPI(ctx, http.MethodGet, "/invoice", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) CancelInvoice(ctx context.Context, req *DocumentRequest) (*DocumentResponse, error) {
	resp := &DocumentResponse{}

	if err := c.callDocsAPI(ctx, http.MethodPut, "/invoice/cancel", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) RestoreInvoice(ctx context.Context, req *DocumentRequest) (*DocumentResponse, error) {
	resp := &DocumentResponse{}

	if err := c.callDocsAPI(ctx, http.MethodPut, "/invoice/restore", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) DeleteInvoice(ctx context.Context, req *DocumentRequest) (*DocumentResponse, error) {
	resp := &DocumentResponse{}

	if err := c.callDocsAPI(ctx, http.MethodDelete, "/invoice", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

type ClientFilter struct {
	CIF   string `json:"CIF,omitempty" url:"cif,omitempty"`
	Email string `json:"email,omitempty" url:"email,omitempty"`
	Phone string `json:"phone,omitempty" url:"phone,omitempty"`
	Code  string `json:"code,omitempty" url:"code,omitempty"`
}

type OrderDir string

const (
	AscOrderDir  = "ASC"
	DescOrderDir = "DESC"
)

type OrderBy string

const (
	IDOrderBy        = "id"
	IssueDateOrderBy = "issueDate"
	NumberOrderBy    = "number"
)

type GetInvoicesRequest struct {
	CIF                string       `json:"cif,omitempty" url:"cif,omitempty"`
	SeriesName         string       `json:"seriesName,omitempty" url:"seriesName,omitempty"`
	Number             string       `json:"number,omitempty" url:"number,omitempty"`
	Draft              types.Bool   `json:"draft,omitempty" url:"draft,omitempty"`
	Client             ClientFilter `json:"client,omitempty" url:"client,omitempty"`
	Canceled           types.Bool   `json:"canceled,omitempty" url:"canceled,omitempty"`
	IssuedAfter        types.Date   `json:"issuedAfter" url:"issuedAfter,omitempty"`
	IssuedBefore       types.Date   `json:"issuedBefore" url:"issuedBefore,omitempty"`
	WithProducts       types.Bool   `json:"withProducts,omitempty" url:"withProducts,omitempty"`
	WithEInvoiceStatus types.Bool   `json:"withEInvoiceStatus,omitempty" url:"withEInvoiceStatus,omitempty"`
	OrderBy            OrderBy      `json:"orderBy,omitempty" url:"orderBy,omitempty"`
	OrderDir           OrderDir     `json:"orderDir,omitempty" url:"orderDir,omitempty"`
	LimitPerPage       int          `json:"limitPerPage,omitempty" url:"limitPerPage,omitempty"`
	Offset             int          `json:"offset,omitempty" url:"offset,omitempty"`
}

type GetInvoicesResponse struct {
	Status

	Data []types.Invoice `json:"data,omitempty"`
}

func (c *Client) GetInvoices(ctx context.Context, req *GetInvoicesRequest) (*GetInvoicesResponse, error) {
	resp := &GetInvoicesResponse{}

	if err := c.callDocsAPI(ctx, http.MethodGet, "/invoice/list", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
