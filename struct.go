package goton

const (
	// TONInputEncodingText value string text
	TONInputEncodingText = "text"

	// TONInputEncodingHex value string hex
	TONInputEncodingHex = "hex"

	// TONInputEncodingBase64 value string base64
	TONInputEncodingBase64 = "base64"

	// TONOutputEncodingText value string Text
	TONOutputEncodingText = "Text"

	// TONOutputEncodingHex value string Hex
	TONOutputEncodingHex = "Hex"

	// TONOutputEncodingHexUppercase value string HexUppercase
	TONOutputEncodingHexUppercase = "HexUppercase"

	// TONOutputEncodingBase64 value string Base64
	TONOutputEncodingBase64 = "Base64"
)

var (
	chains        = getChains()
	LensMnemonic  = getlengthCounWordsInMnemonic()
	SortDirection = getSortDirection()
	// TONMnemonicDictionary map with dictionary
	TONMnemonicDictionary = getTONMnemonicDictionary()
)

// TONKey struct with keys
type TONKey struct {
	Public string `json:"public"`
	Secret string `json:"secret"`
}

// InputMessage struct for generate message
type InputMessage struct {
	Text   string `json:"text,omitempty"`
	Hex    string `json:"hex,omitempty"`
	Base64 string `json:"base64,omitempty"`
}

// MaxFactorizeResult method crypto MaxFactorizeResult
type MaxFactorizeResult struct {
	A string `json:"a"`
	B string `json:"b"`
}

// ModularPowerRequest method crypto ModularPowerRequest
type ModularPowerRequest struct {
	Base     string `json:"base"`
	Exponent string `json:"exponent"`
	Modulus  string `json:"modulus"`
}

// RandomGenerateBytesRequest method crypto RandomGenerateBytesRequest
type RandomGenerateBytesRequest struct {
	Length         int    `json:"length"`
	OutputEncoding string `json:"outputEncoding,omitempty"`
}

// MnemonicStructRequest method crypto MnemonicStructRequest
type MnemonicStructRequest struct {
	*InputMessage `json:"entropy,omitempty"`
	Dictionary    int    `json:"dictionary,omitempty"`
	WordCount     int    `json:"wordCount,omitempty"`
	Phrase        string `json:"phrase,omitempty"`
}

// MessageInputMessage generate message->input message struct
type MessageInputMessage struct {
	*InputMessage `json:"message"`
}

// ScryptDate method crypto ScryptDate
type ScryptDate struct {
	Password       *InputMessage `json:"password"`
	Salt           *InputMessage `json:"salt"`
	LogN           int           `json:"logN"`
	R              int           `json:"r"`
	P              int           `json:"p"`
	DkLen          int           `json:"dkLen"`
	OutputEncoding string        `json:"outputEncoding,omitempty"`
}

// NaclBoxIn for method crypto nalc box
type NaclBoxIn struct {
	*InputMessage  `json:"message"`
	Nonce          string `json:"nonce"`
	TheirPublicKey string `json:"theirPublicKey"`
	SecretKey      string `json:"secretKey"`
	OutputEncoding string `json:"outputEncoding,omitempty"`
}

// NaclSecretBox for method crypto nalc box secret
type NaclSecretBox struct {
	*InputMessage  `json:"message"`
	Nonce          string `json:"nonce"`
	Key            string `json:"key"`
	OutputEncoding string `json:"outputEncoding,omitempty"`
}

// NaclSign for method crypto NaclSign
type NaclSign struct {
	*InputMessage  `json:"message"`
	Key            string `json:"key"`
	OutputEncoding string `json:"outputEncoding,omitempty"`
}

// HDSerialized for method crypto HDKeys
type HDSerialized struct {
	Serialized string `json:"serialized"`
}

// HDDerivery for method HDKeysDerivery
type HDDerivery struct {
	Serialized string `json:"serialized"`
	Index      int    `json:"index"`
	Hardened   bool   `json:"hardened"`
	Compliant  bool   `json:"compliant"`
}

// HDPathDerivery for methods HDKeysDeriveryPath
type HDPathDerivery struct {
	Serialized string `json:"serialized"`
	Path       string `json:"path"`
	Compliant  bool   `json:"compliant"`
}

// TomlConfig struct with config data
type TomlConfig struct {
	BaseURL                            string   `toml:"base_URL" json:"baseURL"`
	Servers                            []string `toml:"servers" json:"servers"`
	MessageRetriesCount                int      `toml:"message_retries_count" json:"messageRetriesCount"`
	MessageExpirationTimeout           int      `toml:"message_expiration_timeout" json:"messageExpirationTimeout"`
	MessageExpirationTimeoutGrowFactor float32  `toml:"message_expiration_timeout_grow_factor" json:"messageExpirationTimeoutGrowFactor"`
	MessageProcessingTimeout           int      `toml:"message_processing_timeout" json:"messageProcessingTimeout"`
	WaitForTimeout                     int      `toml:"wait_for_timeout" json:"waitForTimeout"`
	AccessKey                          string   `toml:"access_key" json:"accessKey"`
	OutOfSyncThreshold                 int      `toml:"out_of_sync_threshold" json:"outOfSyncThreshold"`
}

// NewConfig create new config for connect client
// chanID 0-devnet, 1-mainnet,
func NewConfig(chanID int) *TomlConfig {
	config := TomlConfig{}
	config.BaseURL = chains[chanID]
	config.Servers = append(config.Servers, config.BaseURL)
	config.MessageRetriesCount = 10
	config.MessageExpirationTimeout = 10000 //ms
	config.MessageExpirationTimeoutGrowFactor = 1.5
	config.MessageProcessingTimeout = 40000 //ms
	config.WaitForTimeout = 40000           //ms
	config.OutOfSyncThreshold = 15000       //ms

	return &config
}

func getChains() map[int]string {
	return map[int]string{0: "net.ton.dev", 1: "main.ton.dev"}
}

func getlengthCounWordsInMnemonic() map[int]string {
	return map[int]string{12: "12", 15: "15", 18: "18", 21: "21", 24: "24"}
}

func getTONMnemonicDictionary() map[string]int {
	return map[string]int{
		"TON":                 0,
		"ENGLISH":             1,
		"CHINESE_SIMPLIFIED":  2,
		"CHINESE_TRADITIONAL": 3,
		"FRENCH":              4,
		"ITALIAN":             5,
		"JAPANESE":            6,
		"KOREAN":              7,
		"SPANISH":             8,
	}
}

func getSortDirection() map[int]string {
	return map[int]string{
		0: "ASC",
		1: "DESC",
	}
}

// FixInputMessage make InputMessage struct
func FixInputMessage(value string, format string) *InputMessage {
	inpMess := &InputMessage{}
	switch format {
	case "text":
		inpMess.Text = value
	case "hex":
		inpMess.Hex = value
	case "base64":
		inpMess.Base64 = value
	}

	return inpMess
}

// Queries structs

// ParamsOfQuery ...
type ParamsOfQuery struct {
	Table   string  `json:"table"`
	Filter  string  `json:"filter"`
	Result  string  `json:"result"`
	OrderBy OrderBy `json:"order_by,omitempty"`
	Limit   int     `json:"limit,omitempty"` //max:50
}

// ParamsOfSubscribe ...
type ParamsOfSubscribe struct {
	Table  string `json:"table"`
	Filter string `json:"filter"`
	Result string `json:"result"`
}

// ParamsOfWaitFor ...
type ParamsOfWaitFor struct {
	Table   string `json:"table"`
	Filter  string `json:"filter"`
	Result  string `json:"result"`
	Timeout int    `json:"timeout"`
}

// OrderBy ...
type OrderBy struct {
	Path      string `json:"path"`
	Direction string `json:"direction"`
}

// HandleStruct ...
type HandleStruct struct {
	Handle int `json:"handle"`
}
