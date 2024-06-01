package types

type Company struct {
	CIF            string `json:"cif,omitempty"`
	Company        string `json:"company,omitempty"`
	UserTypeAccess string `json:"userTypeAccess,omitempty"`
	UseStock       Bool   `json:"useStock,omitempty"`
}

type VATRate struct {
	Name    string `json:"name,omitempty"`
	Percent int    `json:"percent,omitempty"`
	Default bool   `json:"default,omitempty"`
}

type Client struct {
	ClientID     string `json:"clientId,omitempty"`
	CIF          string `json:"cif,omitempty"`
	Name         string `json:"name,omitempty"`
	RC           string `json:"rc,omitempty"`
	Code         string `json:"code,omitempty"`
	Address      string `json:"address,omitempty"`
	State        string `json:"state,omitempty"`
	City         string `json:"city,omitempty"`
	Country      string `json:"country,omitempty"`
	IBAN         string `json:"iban,omitempty"`
	Bank         string `json:"bank,omitempty"`
	Email        string `json:"email,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Contact      string `json:"contact,omitempty"`
	VATPayer     Bool   `json:"vatPayer,omitempty"`
	Save         Bool   `json:"save,omitempty"`
	Autocomplete Bool   `json:"autocomplete,omitempty"`
}

type Stock struct {
	WorkStation   string `json:"workStation,omitempty"`
	Management    string `json:"management,omitempty"`
	Quantity      int    `json:"quantity,omitempty"`
	Price         string `json:"price,omitempty"`
	Currency      string `json:"currency,omitempty"`
	VATName       string `json:"vatName,omitempty"`
	VATPercentage int    `json:"vatPercentage,omitempty"`
	VATIncluded   bool   `json:"vatIncluded,omitempty"`
}

const (
	MerchandiseProductType  ProductType = "Marfa"
	ServiceProductType      ProductType = "Serviciu"
	RawProductType          ProductType = "Materii prime"
	ConsumableProductType   ProductType = "Materiale consumabile"
	SemiProductType         ProductType = "Semifabricate"
	FinishedProductType     ProductType = "Produs finit"
	WasteProductType        ProductType = "Produs rezidual"
	AgriculturalProductType ProductType = "Produse agricole"
	LivestockProductType    ProductType = "Animale si pasari"
	PackingProductType      ProductType = "Ambalaje"
	InventoryProductType    ProductType = "Obiecte de inventar"
	NoneProductType         ProductType = "-"
)

type ProductType string

type Product struct {
	Name          string      `json:"name,omitempty"`
	Code          string      `json:"code,omitempty"`
	Description   string      `json:"description,omitempty"`
	MeasuringUnit string      `json:"measuringUnit,omitempty"`
	ProductType   ProductType `json:"productType,omitempty"`
	Stock         []Stock     `json:"stock,omitempty"`
	Price         string      `json:"price,omitempty"`
	Currency      string      `json:"currency,omitempty"`
	VATName       string      `json:"vatName,omitempty"`
	VATPercentage Int         `json:"vatPercentage,omitempty"`
	VATIncluded   Bool        `json:"vatIncluded,omitempty"`
	Active        Bool        `json:"active,omitempty"`
	Image         string      `json:"image,omitempty"`
}

type Series struct {
	Type    string `json:"type,omitempty"`
	Name    string `json:"name,omitempty"`
	Start   string `json:"start,omitempty"`
	Next    string `json:"next,omitempty"`
	Default bool   `json:"default,omitempty"`
}

type Language struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type Management struct {
	Management     string `json:"management,omitempty"`
	WorkStation    string `json:"workStation,omitempty"`
	ManagementType string `json:"managementType,omitempty"`
}
