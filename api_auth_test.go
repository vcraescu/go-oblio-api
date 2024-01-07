package oblio_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vcraescu/go-oblio-api"
)

func TestClient_GenerateAuthorizeToken(t *testing.T) {
	t.Parallel()

	type fields struct {
		clientID     string
		clientSecret string
	}

	tests := []struct {
		name    string
		fields  fields
		want    *oblio.GenerateTokenResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				clientID:     clientID,
				clientSecret: clientSecret,
			},
			want: &oblio.GenerateTokenResponse{
				AccessToken: accessToken,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var (
				baseURL  = StartServer(t, nil, nil)
				client   = oblio.NewClient(tt.fields.clientID, tt.fields.clientSecret, oblio.WithBaseURL(baseURL))
				got, err = client.GenerateToken(context.Background())
			)

			if tt.wantErr != nil {
				tt.wantErr(t, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want.AccessToken, got.AccessToken)
		})
	}
}
