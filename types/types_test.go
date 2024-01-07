package types_test

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vcraescu/go-oblio-api/types"
)

func TestDate_MarshalJSON(t *testing.T) {
	t.Parallel()

	type args struct {
		Date types.Date `json:"date"`
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				Date: types.NewDate(2024, 1, 1),
			},
			want: `{"date":"2024-01-01"}`,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := json.Marshal(tt.args)

			if tt.wantErr != nil {
				tt.wantErr(t, err)

				return
			}

			require.NoError(t, err)
			require.JSONEq(t, tt.want, string(got))
		})
	}
}

func TestDate_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	type args struct {
		js string
	}

	type Got struct {
		Date types.Date
	}

	tests := []struct {
		name    string
		args    args
		want    types.Date
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "empty",
			args: args{
				js: `{"date":""}`,
			},
			want: types.NewDate(0, 0, 0),
		},
		{
			name: "null",
			args: args{
				js: `{"date":null}`,
			},
			want: types.NewDate(0, 0, 0),
		},
		{
			name: "success",
			args: args{
				js: `{"date":"2024-01-01"}`,
			},
			want: types.NewDate(2024, 1, 1),
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Got{}
			err := json.Unmarshal([]byte(tt.args.js), &got)

			if tt.wantErr != nil {
				tt.wantErr(t, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got.Date)
		})
	}
}

func TestDate_EncodeValues(t *testing.T) {
	t.Parallel()

	type args struct {
		Date types.Date `json:"date" url:"date,omitempty"`
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				Date: types.NewDate(2024, 1, 1),
			},
			want: `date=2024-01-01`,
		},
		{
			name: "empty",
			args: args{},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := query.Values(tt.args)

			if tt.wantErr != nil {
				tt.wantErr(t, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got.Encode())
		})
	}
}

func TestBool_MarshalJSON(t *testing.T) {
	t.Parallel()

	type args struct {
		Value types.Bool `json:"value"`
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "true",
			args: args{
				Value: true,
			},
			want: `{"value":"1"}`,
		},
		{
			name: "false",
			args: args{
				Value: false,
			},
			want: `{"value":"0"}`,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := json.Marshal(tt.args)

			if tt.wantErr != nil {
				tt.wantErr(t, err)

				return
			}

			require.NoError(t, err)
			require.JSONEq(t, tt.want, string(got))
		})
	}
}

func TestBool_EncodeValues(t *testing.T) {
	t.Parallel()

	type args struct {
		Value types.Bool `json:"value" url:"value"`
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "true",
			args: args{
				Value: true,
			},
			want: "value=1",
		},
		{
			name: "false",
			args: args{
				Value: false,
			},
			want: "value=0",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := query.Values(tt.args)

			if tt.wantErr != nil {
				tt.wantErr(t, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got.Encode())
		})
	}
}

func TestBool_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	type args struct {
		js string
	}

	type Got struct {
		Value types.Bool `json:"value"`
	}

	tests := []struct {
		name    string
		args    args
		want    types.Bool
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "true",
			args: args{
				js: `{"value":"1"}`,
			},
			want: true,
		},
		{
			name: "false",
			args: args{
				js: `{"value":"0"}`,
			},
		},
		{
			name: "empty",
			args: args{
				js: `{"value":""}`,
			},
			wantErr: func(t assert.TestingT, err error, _ ...any) bool {
				return assert.Error(t, err)
			},
		},
		{
			name: "invalid value",
			args: args{
				js: `{"value":"foobar"}`,
			},
			wantErr: func(t assert.TestingT, err error, _ ...any) bool {
				return assert.Error(t, err)
			},
		},
		{
			name: "missing",
			args: args{
				js: `{"foo":"bar"}`,
			},
		},
		{
			name: "int",
			args: args{
				js: `{"value":0}`,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Got{}
			err := json.Unmarshal([]byte(tt.args.js), &got)

			if tt.wantErr != nil {
				tt.wantErr(t, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got.Value)
		})
	}
}

func TestInt_MarshalJSON(t *testing.T) {
	t.Parallel()

	type args struct {
		Value types.Int `json:"value"`
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "non empty",
			args: args{
				Value: 4,
			},
			want: `{"value":"4"}`,
		},
		{
			name: "zero",
			args: args{
				Value: 0,
			},
			want: `{"value":"0"}`,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := json.Marshal(tt.args)

			if tt.wantErr != nil {
				tt.wantErr(t, err)

				return
			}

			require.NoError(t, err)
			require.JSONEq(t, tt.want, string(got))
		})
	}
}

func TestInt_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	type Got struct {
		Value types.Int `json:"value"`
	}

	type args struct {
		js string
	}

	tests := []struct {
		name    string
		args    args
		want    types.Int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "non empty",
			args: args{
				js: `{"value":"2"}`,
			},
			want: 2,
		},
		{
			name: "empty",
			args: args{
				js: `{"value":""}`,
			},
		},
		{
			name: "invalid value",
			args: args{
				js: `{"value":"foobar"}`,
			},
			wantErr: func(t assert.TestingT, err error, _ ...any) bool {
				return assert.Error(t, err)
			},
		},
		{
			name: "int",
			args: args{
				js: `{"value":2}`,
			},
			want: 2,
		},
		{
			name: "float",
			args: args{
				js: `{"value":2.5}`,
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Got{}
			err := json.Unmarshal([]byte(tt.args.js), &got)

			if tt.wantErr != nil {
				tt.wantErr(t, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got.Value)
		})
	}
}
