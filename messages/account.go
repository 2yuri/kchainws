package messages

type Account []struct {
	BalanceChange   bool   `json:"BalanceChange"`
	IsSender        bool   `json:"IsSender"`
	IsKDAOperation  bool   `json:"IsKDAOperation"`
	IsNFTOperation  bool   `json:"IsNFTOperation"`
	TokenIdentifier string `json:"TokenIdentifier"`
	NFTNonce        string `json:"NFTNonce"`
	Type            string `json:"Type"`
	MarketplaceID   string `json:"MarketplaceID"`
	OrderID         string `json:"OrderID"`
}
