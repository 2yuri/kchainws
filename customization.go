package kchainws

//Customization is the customization for the ws
type Customization struct {
	Addresses []string `json:"addresses"`
	Types     []string `json:"subcribed_types"`
}

//Types you can choose the socket message types to receive
type Types string

const (
	UNKNOWN Types = "UNKNOWN"
	//USER_TRANSACTIONS read address transactions
	USER_TRANSACTIONS Types = "user_transaction"
	//ACCOUNTS read address account changes
	ACCOUNTS Types = "accounts"
	//TRANSACTIONS read all transactions
	TRANSACTIONS Types = "transaction"
	//BLOCKS read only blocks
	BLOCKS Types = "blocks"
)

//NewCustom creates a customization to your client
func NewCustom(types []Types, addresses ...string) *Customization {
	var parsed []string

	for _, t := range types {
		parsed = append(parsed, string(t))
	}

	return &Customization{
		Addresses: addresses,
		Types:     parsed,
	}
}

func (c *Customization) validate() error {
	if len(c.Types) == 0 {
		if len(c.Addresses) == 0 {
			c.Types = []string{string(BLOCKS)}
		} else {
			c.Types = []string{string(BLOCKS), string(TRANSACTIONS), string(ACCOUNTS), string(USER_TRANSACTIONS)}
		}
	}

	for _, t := range c.Types {
		if t != string(BLOCKS) && t != string(ACCOUNTS) && t != string(TRANSACTIONS) && t != string(USER_TRANSACTIONS) {
			return ErrInvalidType
		}
	}

	return nil
}
