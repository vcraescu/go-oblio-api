package testutil

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type HTTPRequestAssertionFunc func(t *testing.T, got *http.Request) bool

func NewBody(t *testing.T, v any) io.ReadCloser {
	t.Helper()

	switch a := v.(type) {
	case string:
		return io.NopCloser(strings.NewReader(a))
	case []byte:
		return io.NopCloser(bytes.NewReader(a))
	default:
		b, err := json.Marshal(v)
		require.NoError(t, err)

		return io.NopCloser(bytes.NewReader(b))
	}
}

func AssertEqualHTTPBody(t *testing.T, want io.ReadCloser, got io.ReadCloser) bool {
	t.Helper()

	if want == got {
		return true
	}

	var wantData, gotData []byte
	var err error

	if want != nil {
		wantData, err = io.ReadAll(want)
		require.NoError(t, err)
	}

	if got != nil {
		gotData, err = io.ReadAll(got)
		require.NoError(t, err)
	}

	return assert.Equal(t, string(bytes.TrimSpace(wantData)), string(bytes.TrimSpace(gotData)))
}

func AssertEqualHTTPRequest(t *testing.T, want *http.Request, got *http.Request) bool {
	return assert.Equal(t, want.URL.String(), got.URL.String()) &&
		assert.Equal(
			t, want.Header["Content-Type"], got.Header["Content-Type"]) &&
		assert.Equal(
			t, want.Header["Accept-Type"], got.Header["Accept-Type"]) &&
		assert.Equal(
			t, want.Header["Authorization"], got.Header["Authorization"]) &&
		AssertEqualHTTPBody(t, want.Body, got.Body) &&
		assert.Equal(t, want.Method, got.Method)
}

func LoadTestDataJSON(t *testing.T) []byte {
	t.Helper()

	dir := path.Join("testdata", path.Dir(t.Name()))
	filename := path.Join(dir, path.Base(t.Name())+".json")
	_ = os.MkdirAll(dir, 0755)

	if _, err := os.Stat(filename); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			require.NoError(t, err)
		}

		f, err := os.Create(filename)
		require.NoError(t, err)

		defer f.Close()

		return nil
	}

	body, err := os.ReadFile(filename)
	require.NoError(t, err)

	return body
}
