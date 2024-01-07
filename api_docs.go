package oblio

import (
	"context"
	"fmt"

	"github.com/vcraescu/go-oblio-api/types"
)

func (c *Client) callDocsAPI(ctx context.Context, method, endpointSuffix string, req, resp any) error {
	if err := c.callAPI(ctx, method, "/docs", endpointSuffix, req, resp); err != nil {
		return fmt.Errorf("callAPI: %w", err)
	}

	return nil
}

type DocumentRequest struct {
	CIF        string `json:"cif,omitempty" url:"cif,omitempty"`
	SeriesName string `json:"seriesName,omitempty" url:"seriesName,omitempty"`
	Number     string `json:"number,omitempty" url:"number,omitempty"`
}

type DocumentResponse struct {
	Status

	Data types.Document `json:"data"`
}
