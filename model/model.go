package model

type Request struct {
	Text       string `json:"text"`
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
}
type Response struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

type DeepLRequest struct {
	Jsonrpc string    `json:"jsonrpc"`
	Method  string    `json:"method"`
	Params  ReqParams `json:"params"`
	Id      int64     `json:"id"`
}

type ReqParams struct {
	Texts           []ReqParamsTexts         `json:"texts"`
	Splitting       string                   `json:"splitting"`
	Lang            ReqParamsLang            `json:"lang"`
	Timestamp       int64                    `json:"timestamp"`
	CommonJobParams ReqParamsCommonJobParams `json:"commonJobParams"`
}

type ReqParamsTexts struct {
	Text                string `json:"text"`
	RequestAlternatives int    `json:"requestAlternatives"`
}
type ReqParamsLang struct {
	TargetLang             string `json:"target_lang"`
	SourceLangUserSelected string `json:"source_lang_user_selected"`
	//SourceLangUserSelected string `json:"source_lang_user_selected,omitempty"`
}
type ReqParamsCommonJobParams struct {
	//RegionalVariant string `json:"regionalVariant,omitempty"`
	WasSpoken    bool   `json:"wasSpoken"`
	TranscribeAS string `json:"transcribe_as"`
}

type DeepLResponse struct {
	Jsonrpc string     `json:"jsonrpc"`
	Id      int        `json:"id"`
	Result  RespResult `json:"result"`
	Error   RespError  `json:"error"`
}

type RespResult struct {
	Texts             []RespResultText            `json:"texts"`
	Lang              string                      `json:"lang"`
	LangIsConfident   bool                        `json:"lang_is_confident"`
	DetectedLanguages RespResultDetectedLanguages `json:"detectedLanguages"`
}

type RespResultText struct {
	Alternatives []interface{} `json:"alternatives"`
	Text         string        `json:"text"`
}

type RespResultDetectedLanguages struct {
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

type RespError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//
//- `DE`: 德语
//- `EN`: 英语
//- `ES`: 西班牙语
//- `FR`: 法语
//- `IT`: 意大利语
//- `JA`: 日语
//- `NL`: 荷兰语
//- `PL`: 波兰语
//- `PT`: 葡萄牙语
//- `RU`: 俄语
//- `ZH`: 中文
//- `BG`: 保加利亚语
//- `CS`: 捷克语
//- `DA`: 丹麦语
//- `EL`: 希腊语
//- `ET`: 爱沙尼亚语
//- `FI`: 芬兰语
//- `HU`: 匈牙利语
//- `LT`: 立陶宛语
//- `LV`: 拉脱维亚语
//- `RO`: 罗马尼亚语
//- `SK`: 斯洛伐克语
//- `SL`: 斯洛文尼亚语
//- `SV`: 瑞典语
