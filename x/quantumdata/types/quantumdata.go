package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// QuantumDataFragment представляет фрагмент данных в системе
type QuantumDataFragment struct {
	ID            string         `json:"id"`
	Data          string         `json:"data"`
	Expiry        int64          `json:"expiry"`
	AccessLevel   string         `json:"access_level"`
	Owner         sdk.AccAddress `json:"owner"`
	EncryptionKey string         `json:"encryption_key"`
	Status        string         `json:"status"`
	LastAccessed  int64          `json:"last_accessed"`
	AccessCount   int            `json:"access_count"`
	Conditions    []string       `json:"conditions"`
	FragmentSize  int            `json:"fragment_size"`
}
