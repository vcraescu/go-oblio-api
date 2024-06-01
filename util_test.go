package oblio_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vcraescu/go-oblio-api"
	"github.com/vcraescu/go-oblio-api/internal/testutil"
	"github.com/vcraescu/go-oblio-api/types"
)

const (
	accessToken  = "test-access-token"
	clientID     = "test-client-id"
	clientSecret = "test-client-secret"
)

var now = time.Now()

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
			ExpiresIn:   types.Int(3600),
			TokenType:   "Bearer",
			RequestTime: types.Timestamp(now),
		})
		require.NoError(t, err)
	}
}

type Test struct {
	wantReq    testutil.HTTPRequestAssertionFunc
	want       any
	wantErr    assert.ErrorAssertionFunc
	method     string
	arg        any
	noFixtures bool
}

func RunTest(t *testing.T, test Test) {
	t.Parallel()

	var respBody []byte

	if !test.noFixtures {
		respBody = testutil.LoadTestDataJSON(t)
	}

	var (
		baseURL = StartServer(t, respBody, test.wantReq)
		client  = oblio.NewClient(clientID, clientSecret, oblio.WithBaseURL(baseURL))
		args    = []reflect.Value{reflect.ValueOf(context.Background())}
	)

	if test.arg != nil {
		args = append(args, reflect.ValueOf(test.arg))
	}

	out := reflect.
		ValueOf(client).
		MethodByName(test.method).
		Call(args)

	got := out[0].Interface()
	err, _ := out[1].Interface().(error)

	if test.wantErr != nil {
		test.wantErr(t, err)

		return
	}

	require.NoError(t, err)
	require.EqualExportedValues(t, test.want, got)
}
