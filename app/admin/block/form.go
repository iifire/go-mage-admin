package block

// FieldsetType 列类型
type FieldsetType struct {
	Label  string   `json:"label"`
	Class  string   `json:"class"`
	Style  string   `json:"style"`
	Layout []string `json:"layout"`
	Items  []TrType `json:"items"`
}

// TrType 表格行类型
type TrType struct {
	Class string   `json:"class"`
	Items []TdType `json:"items"`
}

// TdType 表格列类型
type TdType struct {
	Label       string `json:"label"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Values      string `json:"values"`
	Required    string `json:"required"`
	Disabled    string `json:"disabled"`
	Readonly    string `json:"readonly"`
	AfterHtml   string `json:"afterHtml"`
	Placeholder string `json:"placeholder"`
	Style       string `json:"style"`
	Class       string `json:"class"`
}
