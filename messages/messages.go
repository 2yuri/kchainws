package messages

import (
	"encoding/json"
	"github.com/hyperyuri/kchainws"
)

type Response struct {
	Type string `json:"type"`
	Data []byte `json:"data"`
}

type AcceptableTypes interface {
	Block | Account | Transaction | UserTransaction
}

//ParseMessage parse message from ws
func ParseMessage[T AcceptableTypes](data []byte, input *T) error {
	var message Response
	err := json.Unmarshal(data, &message)
	if err != nil {
		return err
	}

	err = json.Unmarshal(message.Data, input)
	if err != nil {
		return err
	}

	return nil
}

//GetType get the message type
func GetType(data []byte) (kchainws.Types, error) {
	var message Response
	err := json.Unmarshal(data, &message)
	if err != nil {
		return kchainws.UNKNOWN, err
	}
	return kchainws.Types(message.Type), nil
}
