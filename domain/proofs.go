package domain

import "encoding/json"

// ProofsErrorCode ...
var ProofsErrorCode map[string]int

type (

	// ParamsOfProofBlockData ...
	ParamsOfProofBlockData struct {
		Block json.RawMessage `json:"block"`
	}

	// ParamsOfProofTransactionData ...
	ParamsOfProofTransactionData struct {
		Transaction json.RawMessage `json:"transaction"`
	}

	// ProofsUseCase ...
	ProofsUseCase interface {
		ProofBlockData(*ParamsOfProofBlockData) error
		ProofTransactionData(*ParamsOfProofTransactionData) error
	}
)

func init() {
	AbiErrorCode = map[string]int{
		"InvalidData":           901,
		"ProofCheckFailed":      902,
		"InternalError":         903,
		"DataDiffersFromProven": 904,
	}
}
