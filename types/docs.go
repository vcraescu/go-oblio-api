package types

import (
	"fmt"
	"net/url"
)

type DocumentRow interface {
	isDocumentRow()
}

type DiscountType string

const (
	PercentageDiscountType DiscountType = "procentual"
	FlatDiscountType       DiscountType = "valoric"
)

var _ DocumentRow = (*Discount)(nil)

type Discount struct {
	RefItem          string       `json:"refItem,omitempty"`
	Name             string       `json:"name,omitempty"`
	Discount         int          `json:"discount,omitempty"`
	DiscountType     DiscountType `json:"discountType,omitempty"`
	DiscountAllAbove Bool         `json:"discountAllAbove,omitempty"`
}

func (d *Discount) isDocumentRow() {}

type CollectType string

const (
	ReceiptCollectType        CollectType = "Chitanta"
	TaxReceiptCollectType     CollectType = "Bon fiscal"
	CashCollectType           CollectType = "Alta incasare numerar"
	PaymentOrderCollectType   CollectType = "Ordin de plata"
	PostalOrderCollectType    CollectType = "Mandat postal"
	CardCollectType           CollectType = "Card"
	CheckCollectType          CollectType = "CEC"
	PromissoryNoteCollectType CollectType = "Bilet ordin"
	BankCollectType           CollectType = "Alta incasare banca"
)

type Collect struct {
	Type           CollectType `json:"type,omitempty"`
	SeriesName     string      `json:"seriesName,omitempty"`
	DocumentNumber string      `json:"documentNumber,omitempty"`
	Value          string      `json:"value,omitempty"`
	IssueDate      Date        `json:"issueDate,omitempty"`
	Mentions       string      `json:"mentions,omitempty"`
}

type ReferenceDocument struct {
	Type       string `json:"type,omitempty"`
	SeriesName string `json:"seriesName,omitempty"`
	Number     int    `json:"number,omitempty"`
}

type Document struct {
	DocumentType string    `json:"documentType,omitempty"`
	SeriesName   string    `json:"seriesName,omitempty"`
	Number       string    `json:"number,omitempty"`
	Link         string    `json:"link,omitempty"`
	EInvoice     string    `json:"einvoice,omitempty"`
	Total        string    `json:"total,omitempty"`
	Collects     []Collect `json:"collects,omitempty"`
}

func (d *Document) GetID() (string, error) {
	if d.Link == "" {
		return "", nil
	}

	rawURL, err := url.ParseRequestURI(d.Link)
	if err != nil {
		return "", fmt.Errorf("parseRequestURI: %w", err)
	}

	q := rawURL.Query()

	return q.Get("id"), nil
}

type Invoice struct {
	ID                 string    `json:"id,omitempty"`
	Draft              Bool      `json:"draft,omitempty"`
	Canceled           Bool      `json:"canceled,omitempty"`
	Collected          Bool      `json:"collected,omitempty"`
	SeriesName         string    `json:"seriesName,omitempty"`
	Number             string    `json:"number,omitempty"`
	IssueDate          Date      `json:"issueDate,omitempty"`
	DueDate            Date      `json:"dueDate,omitempty"`
	Precision          Int       `json:"precision,omitempty"`
	Currency           string    `json:"currency,omitempty"`
	ExchangeRate       string    `json:"exchangeRate,omitempty"`
	Total              string    `json:"total,omitempty"`
	IssuerName         string    `json:"issuerName,omitempty"`
	IssuerID           string    `json:"issuerId,omitempty"`
	NoticeNumber       string    `json:"noticeNumber,omitempty"`
	DeputyName         string    `json:"deputyName,omitempty"`
	DeputyIdentityCard string    `json:"deputyIdentityCard,omitempty"`
	DeputyAuto         string    `json:"deputyAuto,omitempty"`
	Mentions           string    `json:"mentions,omitempty"`
	UseStock           Bool      `json:"useStock,omitempty"`
	Type               string    `json:"type,omitempty"`
	Link               string    `json:"link,omitempty"`
	EInvoice           string    `json:"einvoice,omitempty"`
	Client             Client    `json:"client,omitempty"`
	Products           []Product `json:"products,omitempty"`
}

const (
	SimplePrecision = 2
	DoublePrecision = 4
)

var _ DocumentRow = (*ProductRow)(nil)

type ProductRow struct {
	Product

	Item       string `json:"item,omitempty"`
	Management string `json:"management,omitempty"`
	Quantity   int    `json:"quantity,omitempty"`
	Save       Bool   `json:"save,omitempty"`
}

func (p *ProductRow) isDocumentRow() {}
