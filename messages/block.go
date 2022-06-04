package messages

type Block struct {
	Hash               string   `json:"hash"`
	Nonce              int      `json:"nonce"`
	ParentHash         string   `json:"parentHash"`
	Timestamp          int64    `json:"timestamp"`
	Slot               int      `json:"slot"`
	Epoch              int      `json:"epoch"`
	IsEpochStart       bool     `json:"isEpochStart"`
	Size               int      `json:"size"`
	SizeTxs            int      `json:"sizeTxs"`
	TxRootHash         string   `json:"txRootHash"`
	TrieRoot           string   `json:"trieRoot"`
	ValidatorsTrieRoot string   `json:"validatorsTrieRoot"`
	KappsTrieRoot      string   `json:"kappsTrieRoot"`
	ProducerSignature  string   `json:"producerSignature"`
	Signature          string   `json:"signature"`
	PrevRandSeed       string   `json:"prevRandSeed"`
	RandSeed           string   `json:"randSeed"`
	TxCount            int      `json:"txCount"`
	TxFees             int      `json:"txFees"`
	TxBurnedFees       int      `json:"txBurnedFees"`
	BlockRewards       int      `json:"blockRewards"`
	StakingRewards     int      `json:"stakingRewards"`
	TxHashes           []string `json:"txHashes"`
	Validators         []string `json:"validators"`
	SoftwareVersion    string   `json:"softwareVersion"`
	ChainID            string   `json:"chainID"`
	Reserved           string   `json:"reserved"`
}
