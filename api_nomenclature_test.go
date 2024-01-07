package oblio_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vcraescu/go-oblio-api"
	"github.com/vcraescu/go-oblio-api/internal/testutil"
	"github.com/vcraescu/go-oblio-api/types"
	"github.com/vcraescu/go-reqbuilder"
)

func TestClient_GetCompanies(t *testing.T) {
	t.Parallel()

	type args struct {
		req *oblio.GetCompaniesRequest
	}

	reqBuilder := reqbuilder.
		NewBuilder("/").
		WithPath("/nomenclature/companies").
		WithMethod(http.MethodGet)

	tests := []struct {
		name    string
		args    args
		want    *oblio.GetCompaniesResponse
		wantReq testutil.HTTPRequestAssertionFunc
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "with access token",
			args: args{
				req: &oblio.GetCompaniesRequest{
					Authorized: oblio.Authorized{
						AccessToken: "access-token",
					},
				},
			},
			wantReq: func(t *testing.T, got *http.Request) bool {
				want, err := reqBuilder.
					WithHeaders(reqbuilder.AuthBearerHeader("access-token")).
					Build(context.Background())
				require.NoError(t, err)

				return testutil.AssertEqualHTTPRequest(t, want, got)
			},
			want: &oblio.GetCompaniesResponse{
				Status: oblio.Status{
					Status:        http.StatusOK,
					StatusMessage: "Success",
				},
				Data: []types.Company{
					{
						CIF:            "123456",
						Company:        "FOOBAR S.R.L.",
						UserTypeAccess: "Administrator",
						UseStock:       false,
					},
				},
			},
		},
		{
			name: "with generated access token",
			args: args{
				req: &oblio.GetCompaniesRequest{},
			},
			wantReq: func(t *testing.T, got *http.Request) bool {
				want, err := reqBuilder.
					WithHeaders(reqbuilder.AuthBearerHeader(accessToken)).
					Build(context.Background())
				require.NoError(t, err)

				return testutil.AssertEqualHTTPRequest(t, want, got)
			},
			want: &oblio.GetCompaniesResponse{
				Status: oblio.Status{
					Status:        http.StatusOK,
					StatusMessage: "Success",
				},
				Data: []types.Company{
					{
						CIF:            "RO37311090",
						Company:        "OBLIO SOFTWARE SRL",
						UserTypeAccess: "admin",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			runSubTest(t, test{
				wantReq: tt.wantReq,
				want:    tt.want,
				wantErr: tt.wantErr,
				method:  "GetCompanies",
				arg:     tt.args.req,
			})
		})
	}
}

func TestClient_GetVATRates(t *testing.T) {
	t.Parallel()

	type args struct {
		in *oblio.GetVATRatesRequest
	}

	reqBuilder := reqbuilder.
		NewBuilder("/").
		WithPath("/nomenclature/vat_rates").
		WithMethod(http.MethodGet)

	tests := []struct {
		name    string
		args    args
		want    *oblio.GetVATRatesResponse
		wantReq testutil.HTTPRequestAssertionFunc
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				in: &oblio.GetVATRatesRequest{
					CIF: "123",
				},
			},
			wantReq: func(t *testing.T, got *http.Request) bool {
				want, err := reqBuilder.
					WithHeaders(reqbuilder.AuthBearerHeader(accessToken)).
					WithParams(oblio.GetVATRatesRequest{CIF: "123"}).
					Build(context.Background())
				require.NoError(t, err)

				return testutil.AssertEqualHTTPRequest(t, want, got)
			},
			want: &oblio.GetVATRatesResponse{
				Status: oblio.Status{
					Status:        http.StatusOK,
					StatusMessage: "Success",
				},
				Data: []types.VATRate{
					{
						Name:    "Normala",
						Percent: 19,
						Default: true,
					},
					{
						Name:    "Redusa",
						Percent: 9,
						Default: false,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			runSubTest(t, test{
				wantReq: tt.wantReq,
				want:    tt.want,
				wantErr: tt.wantErr,
				method:  "GetVATRates",
				arg:     tt.args.in,
			})
		})
	}
}

func TestClient_GetClients(t *testing.T) {
	t.Parallel()

	type args struct {
		in *oblio.GetClientsRequest
	}

	reqBuilder := reqbuilder.
		NewBuilder("/").
		WithPath("/nomenclature/clients").
		WithMethod(http.MethodGet)

	tests := []struct {
		name    string
		args    args
		want    *oblio.GetClientsResponse
		wantReq testutil.HTTPRequestAssertionFunc
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				in: &oblio.GetClientsRequest{
					CIF: "123",
				},
			},
			wantReq: func(t *testing.T, got *http.Request) bool {
				want, err := reqBuilder.
					WithHeaders(reqbuilder.AuthBearerHeader(accessToken)).
					WithParams(oblio.GetClientsRequest{CIF: "123"}).
					Build(context.Background())
				require.NoError(t, err)

				return testutil.AssertEqualHTTPRequest(t, want, got)
			},
			want: &oblio.GetClientsResponse{
				Status: oblio.Status{
					Status:        http.StatusOK,
					StatusMessage: "Success",
				},
				Data: []types.Client{
					{
						CIF:      "RO37311090",
						Name:     "OBLIO SOFTWARE SRL",
						RC:       "J13/887/2017",
						State:    "Constanta",
						City:     "Constanta",
						VATPayer: true,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			runSubTest(t, test{
				wantReq: tt.wantReq,
				want:    tt.want,
				wantErr: tt.wantErr,
				method:  "GetClients",
				arg:     tt.args.in,
			})
		})
	}
}

func TestClient_GetProducts(t *testing.T) {
	t.Parallel()

	type args struct {
		in *oblio.GetProductsRequest
	}

	reqBuilder := reqbuilder.
		NewBuilder("/").
		WithPath("/nomenclature/products").
		WithMethod(http.MethodGet)

	tests := []struct {
		name    string
		args    args
		want    *oblio.GetProductsResponse
		wantReq testutil.HTTPRequestAssertionFunc
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				in: &oblio.GetProductsRequest{
					CIF: "123",
				},
			},
			wantReq: func(t *testing.T, got *http.Request) bool {
				want, err := reqBuilder.
					WithHeaders(reqbuilder.AuthBearerHeader(accessToken)).
					WithParams(oblio.GetProductsRequest{CIF: "123"}).
					Build(context.Background())
				require.NoError(t, err)

				return testutil.AssertEqualHTTPRequest(t, want, got)
			},
			want: &oblio.GetProductsResponse{
				Status: oblio.Status{
					Status:        http.StatusOK,
					StatusMessage: "Success",
				},
				Data: []types.Product{
					{
						Name:          "Montare",
						MeasuringUnit: "buc",
						ProductType:   "Serviciu",
						Price:         "119.00",
						Currency:      "RON",
						VATName:       "Normala",
						VATPercentage: 19,
						VATIncluded:   true,
					},
					{
						Name:          "Birou",
						MeasuringUnit: "buc",
						ProductType:   "Marfa",
						Active:        true,
						Stock: []types.Stock{
							{
								WorkStation:   "Sediu",
								Management:    "Magazin",
								Quantity:      2,
								Price:         "200.00",
								Currency:      "RON",
								VATName:       "Normala",
								VATPercentage: 19,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			runSubTest(t, test{
				wantReq: tt.wantReq,
				want:    tt.want,
				wantErr: tt.wantErr,
				method:  "GetProducts",
				arg:     tt.args.in,
			})
		})
	}
}

func TestClient_GetSeries(t *testing.T) {
	t.Parallel()

	type args struct {
		in *oblio.GetSeriesRequest
	}

	reqBuilder := reqbuilder.
		NewBuilder("/").
		WithPath("/nomenclature/series").
		WithMethod(http.MethodGet)

	tests := []struct {
		name    string
		args    args
		want    *oblio.GetSeriesResponse
		wantReq testutil.HTTPRequestAssertionFunc
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				in: &oblio.GetSeriesRequest{
					CIF: "123",
				},
			},
			wantReq: func(t *testing.T, got *http.Request) bool {
				want, err := reqBuilder.
					WithHeaders(reqbuilder.AuthBearerHeader(accessToken)).
					WithParams(oblio.GetSeriesRequest{CIF: "123"}).
					Build(context.Background())
				require.NoError(t, err)

				return testutil.AssertEqualHTTPRequest(t, want, got)
			},
			want: &oblio.GetSeriesResponse{
				Status: oblio.Status{
					Status:        http.StatusOK,
					StatusMessage: "Success",
				},
				Data: []types.Series{
					{
						Type:    "Factura",
						Name:    "FCT",
						Start:   "0001",
						Next:    "0051",
						Default: true,
					},
					{
						Type:    "Proforma",
						Name:    "PR",
						Start:   "0001",
						Next:    "0008",
						Default: true,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			runSubTest(t, test{
				wantReq: tt.wantReq,
				want:    tt.want,
				wantErr: tt.wantErr,
				method:  "GetSeries",
				arg:     tt.args.in,
			})
		})
	}
}

func TestClient_GetLanguages(t *testing.T) {
	t.Parallel()

	type args struct {
		in *oblio.GetLanguagesRequest
	}

	reqBuilder := reqbuilder.
		NewBuilder("/").
		WithPath("/nomenclature/languages").
		WithMethod(http.MethodGet)

	tests := []struct {
		name    string
		args    args
		want    *oblio.GetLanguagesResponse
		wantReq testutil.HTTPRequestAssertionFunc
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				in: &oblio.GetLanguagesRequest{
					CIF: "123",
				},
			},
			wantReq: func(t *testing.T, got *http.Request) bool {
				want, err := reqBuilder.
					WithHeaders(reqbuilder.AuthBearerHeader(accessToken)).
					WithParams(oblio.GetLanguagesRequest{CIF: "123"}).
					Build(context.Background())
				require.NoError(t, err)

				return testutil.AssertEqualHTTPRequest(t, want, got)
			},
			want: &oblio.GetLanguagesResponse{
				Status: oblio.Status{
					Status:        http.StatusOK,
					StatusMessage: "Success",
				},
				Data: []types.Language{
					{
						Code: "EN",
						Name: "Engleza",
					},
					{
						Code: "FR",
						Name: "Franceza",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			runSubTest(t, test{
				wantReq: tt.wantReq,
				want:    tt.want,
				wantErr: tt.wantErr,
				method:  "GetLanguages",
				arg:     tt.args.in,
			})
		})
	}
}

func TestClient_GetManagement(t *testing.T) {
	t.Parallel()

	type args struct {
		in *oblio.GetManagementRequest
	}

	reqBuilder := reqbuilder.
		NewBuilder("/").
		WithPath("/nomenclature/management").
		WithMethod(http.MethodGet)

	tests := []struct {
		name    string
		args    args
		want    *oblio.GetManagementResponse
		wantReq testutil.HTTPRequestAssertionFunc
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				in: &oblio.GetManagementRequest{
					CIF: "123",
				},
			},
			wantReq: func(t *testing.T, got *http.Request) bool {
				want, err := reqBuilder.
					WithHeaders(reqbuilder.AuthBearerHeader(accessToken)).
					WithParams(oblio.GetManagementRequest{CIF: "123"}).
					Build(context.Background())
				require.NoError(t, err)

				return testutil.AssertEqualHTTPRequest(t, want, got)
			},
			want: &oblio.GetManagementResponse{
				Status: oblio.Status{
					Status:        http.StatusOK,
					StatusMessage: "Success",
				},
				Data: []types.Management{
					{
						Management:     "Magazin",
						WorkStation:    "Sediu",
						ManagementType: "Cantitativ Valorica",
					},
					{
						Management:     "Mobila",
						WorkStation:    "Depozit",
						ManagementType: "Cantitativ Valorica",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			runSubTest(t, test{
				wantReq: tt.wantReq,
				want:    tt.want,
				wantErr: tt.wantErr,
				method:  "GetManagement",
				arg:     tt.args.in,
			})
		})
	}
}

type test struct {
	wantReq testutil.HTTPRequestAssertionFunc
	want    any
	wantErr assert.ErrorAssertionFunc
	method  string
	arg     any
}

func runSubTest(t *testing.T, test test) {
	t.Parallel()

	var (
		respBody = testutil.LoadTestDataJSON(t)
		baseURL  = StartServer(t, respBody, test.wantReq)
		client   = oblio.NewClient(clientID, clientSecret, oblio.WithBaseURL(baseURL))
	)

	out := reflect.
		ValueOf(client).
		MethodByName(test.method).
		Call([]reflect.Value{reflect.ValueOf(context.Background()), reflect.ValueOf(test.arg)})

	got := out[0].Interface()
	err, _ := out[1].Interface().(error)

	if test.wantErr != nil {
		test.wantErr(t, err)

		return
	}

	require.NoError(t, err)
	require.Equal(t, test.want, got)
}
