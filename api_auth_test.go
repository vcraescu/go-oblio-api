package oblio_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vcraescu/go-oblio-api"
	"github.com/vcraescu/go-oblio-api/internal/testutil"
	"github.com/vcraescu/go-oblio-api/types"
	"github.com/vcraescu/go-reqbuilder"
)

func TestClient_GenerateAuthorizeToken(t *testing.T) {
	t.Parallel()

	type fields struct {
		clientID     string
		clientSecret string
	}

	reqBuilder := reqbuilder.
		NewBuilder("/").
		WithPath("/authorize/token").
		WithMethod(http.MethodPost)

	tests := []struct {
		name    string
		fields  fields
		want    *oblio.GenerateTokenResponse
		wantReq testutil.HTTPRequestAssertionFunc
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				clientID:     clientID,
				clientSecret: clientSecret,
			},
			wantReq: func(t *testing.T, got *http.Request) bool {
				want, err := reqBuilder.
					WithHeaders(reqbuilder.JSONContentHeader).
					WithBody(oblio.GenerateAuthorizeTokenRequest{
						ClientID:     clientID,
						ClientSecret: clientSecret,
					}).
					Build(context.Background())
				require.NoError(t, err)

				return testutil.AssertEqualHTTPRequest(t, want, got)
			},
			want: &oblio.GenerateTokenResponse{
				AccessToken: accessToken,
				ExpiresIn:   types.Int(3600),
				TokenType:   "Bearer",
				RequestTime: types.Timestamp(now),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RunTest(t, Test{
				wantReq:    tt.wantReq,
				want:       tt.want,
				wantErr:    tt.wantErr,
				method:     "GenerateToken",
				noFixtures: true,
			})
		})
	}
}
