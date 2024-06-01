package oblio_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vcraescu/go-oblio-api"
	"github.com/vcraescu/go-oblio-api/internal/testutil"
	"github.com/vcraescu/go-oblio-api/types"
)

func TestClient_GetInvoices(t *testing.T) {
	t.Parallel()

	type args struct {
		req *oblio.GetInvoicesRequest
	}

	tests := []struct {
		name    string
		args    args
		want    *oblio.GetInvoicesResponse
		wantReq testutil.HTTPRequestAssertionFunc
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				req: &oblio.GetInvoicesRequest{
					CIF:        "1234567",
					SeriesName: "SC",
					Number:     "001",
					Draft:      true,
					Client: oblio.ClientFilter{
						CIF: "client-cif",
					},
					Canceled:           true,
					IssuedAfter:        types.NewDate(2024, 1, 1),
					IssuedBefore:       types.NewDate(2024, 1, 1),
					WithProducts:       true,
					WithEInvoiceStatus: true,
					OrderBy:            oblio.IDOrderBy,
					OrderDir:           oblio.DescOrderDir,
					LimitPerPage:       10,
					Offset:             5,
				},
			},
			wantReq: func(t *testing.T, got *http.Request) bool {
				return assert.Equal(t, url.Values{
					"canceled":           []string{"1"},
					"cif":                []string{"1234567"},
					"client[cif]":        []string{"client-cif"},
					"draft":              []string{"1"},
					"issuedAfter":        []string{"2024-01-01"},
					"issuedBefore":       []string{"2024-01-01"},
					"limitPerPage":       []string{"10"},
					"number":             []string{"001"},
					"offset":             []string{"5"},
					"orderBy":            []string{"id"},
					"orderDir":           []string{"DESC"},
					"seriesName":         []string{"SC"},
					"withEInvoiceStatus": []string{"1"},
					"withProducts":       []string{"1"},
				}, got.URL.Query())
			},
			want: &oblio.GetInvoicesResponse{
				Status: oblio.Status{
					Status:        http.StatusOK,
					StatusMessage: "Success",
				},
				Data: []types.Invoice{
					{
						ID:                 "10000",
						Draft:              true,
						Collected:          true,
						SeriesName:         "SC",
						Number:             "0001",
						IssueDate:          types.NewDate(2023, 1, 31),
						DueDate:            types.NewDate(2023, 2, 15),
						Precision:          2,
						Currency:           "RON",
						ExchangeRate:       "0.000000",
						Total:              "21000.0000",
						IssuerName:         "Ion Popescu",
						IssuerID:           "1820913326911",
						DeputyIdentityCard: "MZ123475",
						UseStock:           true,
						Type:               "Factura",
						Link:               "https://www.oblio.eu/utils/show_file/?ic=11111&id=11111&it=1111&preload=1",
						EInvoice:           "https://www.oblio.eu/utils/show_file/?ic=111111&id=111111&it=11111&einvoice=1&api=1",
						Client: types.Client{
							ClientID: "1111111",
							CIF:      "49834438",
							Name:     "FOOBAR DEVELOPMENT SRL",
							RC:       "J12/1811/2019",
							Code:     "FOO",
							Address:  "BLD. Vasile Lupu,  NR.5",
							State:    "CLUJ",
							City:     "MUN. CLUJ-NAPOCA",
							Country:  "ROMANIA",
							IBAN:     "RO66INGB0000999900000000",
							Bank:     "ING BANK NV",
							Email:    "ion.popescu@foobar.com",
							Contact:  "ion.popescu@foobar.com",
							VATPayer: true,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			RunTest(t, Test{
				wantReq: tt.wantReq,
				want:    tt.want,
				wantErr: tt.wantErr,
				method:  "GetInvoices",
				arg:     tt.args.req,
			})
		})
	}
}
