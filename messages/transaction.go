package messages

import "time"

type UserTransaction struct {
	Hash         string                   `json:"hash"`
	BlockNum     uint64                   `json:"blockNum,omitempty"`
	Sender       string                   `json:"sender"`
	Nonce        uint64                   `json:"nonce"`
	PermissionID int32                    `json:"permissionID,omitempty"`
	Data         []string                 `json:"data,omitempty"`
	Timestamp    time.Duration            `json:"timestamp,omitempty"`
	KAppFee      int64                    `json:"kAppFee"`
	BandwidthFee int64                    `json:"bandwidthFee"`
	Status       string                   `json:"status"`
	ResultCode   string                   `json:"resultCode,omitempty"`
	Version      uint32                   `json:"version,omitempty"`
	ChainID      string                   `json:"chainID,omitempty"`
	Signature    []string                 `json:"signature,omitempty"`
	SearchOrder  uint32                   `json:"searchOrder"`
	Receipts     []map[string]interface{} `json:"receipts"`
	Contracts    []struct {
		Type       int                    `json:"type"`
		TypeString string                 `json:"typeString"`
		Parameter  map[string]interface{} `json:"parameter,omitempty"`
	} `json:"contract"`
}

type Transaction []struct {
	Hash         string                   `json:"hash"`
	BlockNum     uint64                   `json:"blockNum,omitempty"`
	Sender       string                   `json:"sender"`
	Nonce        uint64                   `json:"nonce"`
	PermissionID int32                    `json:"permissionID,omitempty"`
	Data         []string                 `json:"data,omitempty"`
	Timestamp    time.Duration            `json:"timestamp,omitempty"`
	KAppFee      int64                    `json:"kAppFee"`
	BandwidthFee int64                    `json:"bandwidthFee"`
	Status       string                   `json:"status"`
	ResultCode   string                   `json:"resultCode,omitempty"`
	Version      uint32                   `json:"version,omitempty"`
	ChainID      string                   `json:"chainID,omitempty"`
	Signature    []string                 `json:"signature,omitempty"`
	SearchOrder  uint32                   `json:"searchOrder"`
	Receipts     []map[string]interface{} `json:"receipts"`
	Contracts    []struct {
		Type       int                    `json:"type"`
		TypeString string                 `json:"typeString"`
		Parameter  map[string]interface{} `json:"parameter,omitempty"`
	} `json:"contract"`
}

type foo struct {
}
