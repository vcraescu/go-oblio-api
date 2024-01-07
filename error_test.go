package oblio_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vcraescu/go-oblio-api"
	"github.com/vcraescu/go-oblio-api/internal/testutil"
)

func TestUnmarshalErrorResponse(t *testing.T) {
	t.Parallel()

	type args struct {
		resp *http.Response
	}

	tests := []struct {
		name string
		args args
		want *oblio.ErrorResponse
	}{
		{
			name: "valid response",
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusUnauthorized,
					Body: testutil.NewBody(t, oblio.Status{
						Status:        http.StatusBadRequest,
						StatusMessage: "something wrong",
					}),
				},
			},
			want: &oblio.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "something wrong",
			},
		},
		{
			name: "text response",
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusBadRequest,
					Body:       testutil.NewBody(t, "foobar"),
				},
			},
			want: &oblio.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: `foobar`,
			},
		},
		{
			name: "unknown response",
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusBadRequest,
					Body: testutil.NewBody(t, map[string]string{
						"foo": "bar",
					}),
				},
			},
			want: &oblio.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: `{"foo":"bar"}`,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := oblio.UnmarshalErrorResponse(tt.args.resp)
			require.Equal(t, tt.want, got)
		})
	}
}
