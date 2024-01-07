package oblio

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vcraescu/go-oblio-api/types"
)

func (c *Client) callNomenclatureAPI(ctx context.Context, endpointSuffix string, req, resp any) error {
	if err := c.callAPI(ctx, http.MethodGet, "/nomenclature", endpointSuffix, req, resp); err != nil {
		return fmt.Errorf("callAPI: %w", err)
	}

	return nil
}

type GetCompaniesRequest struct {
	Authorized
}

type GetCompaniesResponse struct {
	Status

	Data []types.Company `json:"data"`
}

func (c *Client) GetCompanies(ctx context.Context, req *GetCompaniesRequest) (*GetCompaniesResponse, error) {
	resp := &GetCompaniesResponse{}

	if err := c.callNomenclatureAPI(ctx, "companies", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

type GetVATRatesRequest struct {
	Authorized

	CIF string `json:"cif,omitempty" url:"cif,omitempty"`
}

type GetVATRatesResponse struct {
	Status

	Data []types.VATRate `json:"data"`
}

func (c *Client) GetVATRates(ctx context.Context, req *GetVATRatesRequest) (*GetVATRatesResponse, error) {
	resp := &GetVATRatesResponse{}

	if err := c.callNomenclatureAPI(ctx, "vat_rates", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

type GetClientsRequest struct {
	Authorized

	CIF       string `json:"cif,omitempty" url:"cif,omitempty"`
	Name      string `json:"name,omitempty" url:"name,omitempty"`
	ClientCIF string `json:"clientCif,omitempty" url:"clientCif,omitempty"`
	Offset    int    `json:"offset,omitempty" url:"offset,omitempty"`
}

type GetClientsResponse struct {
	Status

	Data []types.Client `json:"data"`
}

func (c *Client) GetClients(ctx context.Context, req *GetClientsRequest) (*GetClientsResponse, error) {
	resp := &GetClientsResponse{}

	if err := c.callNomenclatureAPI(ctx, "clients", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

type GetProductsRequest struct {
	Authorized

	CIF         string `json:"cif,omitempty" url:"cif,omitempty"`
	Name        string `json:"name,omitempty" url:"name,omitempty"`
	Code        string `json:"code,omitempty" url:"code,omitempty"`
	Management  string `json:"management,omitempty" url:"management,omitempty"`
	WorkStation string `json:"workStation,omitempty" url:"workStation,omitempty"`
	Offset      int    `json:"offset,omitempty" url:"offset,omitempty"`
}

type GetProductsResponse struct {
	Status

	Data []types.Product `json:"data"`
}

func (c *Client) GetProducts(ctx context.Context, req *GetProductsRequest) (*GetProductsResponse, error) {
	resp := &GetProductsResponse{}

	if err := c.callNomenclatureAPI(ctx, "products", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

type GetSeriesRequest struct {
	Authorized

	CIF string `json:"cif,omitempty" url:"cif,omitempty"`
}

type GetSeriesResponse struct {
	Status

	Data []types.Series `json:"data"`
}

func (c *Client) GetSeries(ctx context.Context, req *GetSeriesRequest) (*GetSeriesResponse, error) {
	resp := &GetSeriesResponse{}

	if err := c.callNomenclatureAPI(ctx, "series", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

type GetLanguagesRequest struct {
	Authorized

	CIF string `json:"cif,omitempty" url:"cif,omitempty"`
}

type GetLanguagesResponse struct {
	Status

	Data []types.Language `json:"data"`
}

func (c *Client) GetLanguages(ctx context.Context, req *GetLanguagesRequest) (*GetLanguagesResponse, error) {
	resp := &GetLanguagesResponse{}

	if err := c.callNomenclatureAPI(ctx, "languages", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

type GetManagementRequest struct {
	Authorized

	CIF string `json:"cif,omitempty" url:"cif,omitempty"`
}

type GetManagementResponse struct {
	Status

	Data []types.Management `json:"data"`
}

func (c *Client) GetManagement(ctx context.Context, req *GetManagementRequest) (*GetManagementResponse, error) {
	resp := &GetManagementResponse{}

	if err := c.callNomenclatureAPI(ctx, "management", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
