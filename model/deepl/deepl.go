package deepl

type Response struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  Result `json:"result"`
	Error   Error  `json:"error"`
}

type Result struct {
	Texts             []Texts           `json:"texts"`
	Lang              string            `json:"lang"`
	LangIsConfident   bool              `json:"lang_is_confident"`
	DetectedLanguages DetectedLanguages `json:"detectedLanguages"`
}

type Texts struct {
	Alternatives []interface{} `json:"alternatives"`
	Text         string        `json:"text"`
}

type DetectedLanguages struct {
	EN          float64 `json:"EN"`
	DE          float64 `json:"DE"`
	FR          float64 `json:"FR"`
	ES          float64 `json:"ES"`
	PT          float64 `json:"PT"`
	IT          float64 `json:"IT"`
	NL          float64 `json:"NL"`
	PL          float64 `json:"PL"`
	RU          float64 `json:"RU"`
	ZH          float64 `json:"ZH"`
	JA          float64 `json:"JA"`
	BG          float64 `json:"BG"`
	CS          float64 `json:"CS"`
	DA          float64 `json:"DA"`
	EL          float64 `json:"EL"`
	ET          float64 `json:"ET"`
	FI          float64 `json:"FI"`
	HU          float64 `json:"HU"`
	LT          float64 `json:"LT"`
	LV          float64 `json:"LV"`
	RO          float64 `json:"RO"`
	SK          float64 `json:"SK"`
	SL          float64 `json:"SL"`
	SV          float64 `json:"SV"`
	TR          float64 `json:"TR"`
	ID          float64 `json:"ID"`
	Unsupported float64 `json:"unsupported"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
