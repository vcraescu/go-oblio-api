package oblio_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vcraescu/go-oblio-api"
	"github.com/vcraescu/go-oblio-api/internal/testutil"
)

const (
	accessToken  = "test-access-token"
	clientID     = "test-client-id"
	clientSecret = "test-client-secret"
)

func StartServer(t *testing.T, respBody []byte, wantRequest testutil.HTTPRequestAssertionFunc) string {
	t.Helper()

	authHandler := NewAuthorizationHandler(t)

	srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/authorize/token" {
			authHandler(w, r)

			return
		}

		if wantRequest != nil {
			wantRequest(t, r)
		}

		if respBody != nil {
			_, err := io.Copy(w, bytes.NewReader(respBody))
			require.NoError(t, err)
		}
	}))

	srv.Start()

	t.Cleanup(func() {
		srv.Close()
	})

	return srv.URL
}

func NewAuthorizationHandler(t *testing.T) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.Helper()

		got := &oblio.GenerateAuthorizeTokenRequest{}
		err := json.NewDecoder(r.Body).Decode(got)
		require.NoError(t, err)

		require.Equal(t, got.ClientID, clientID)
		require.Equal(t, got.ClientSecret, clientSecret)

		err = json.NewEncoder(w).Encode(oblio.GenerateTokenResponse{
			AccessToken: accessToken,
			ExpiresIn:   fmt.Sprint(time.Hour.Seconds()),
			TokenType:   "Bearer",
			RequestTime: fmt.Sprint(time.Now().Second()),
		})
		require.NoError(t, err)
	}
}
