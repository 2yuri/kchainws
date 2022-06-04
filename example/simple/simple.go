package main

import (
	"github.com/hyperyuri/kchainws"
	"github.com/hyperyuri/kchainws/messages"
	"log"
)

func main() {
	//create a new ws client
	c, err := kchainws.NewClient(
		//setup the node URL
		kchainws.WithURL("localhost:8080"),
		//setup the customization - here im only watching for blocks
		//if i select ACCOUNTS or USER_TRANSACTION i need to fill with an address
		//if not provide any type it will fill automatic
		//if provide wrong types it will fail
		kchainws.WithCustom(kchainws.NewCustom([]kchainws.Types{kchainws.BLOCKS, kchainws.TRANSACTIONS})))

	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	// get a new reader chan
	reader := c.GetReader()

	// start conn to receive messages
	err = c.StarConnection()
	if err != nil {
		log.Fatalln(err)
	}

	//messages handling
	for {
		select {
		case <-c.Context().Done():
			log.Println("connection is over")
			return
		case m := <-reader:
			//you can check the message type with messages.GetType()
			if t, _ := messages.GetType(m); t == kchainws.BLOCKS {
				var b messages.Block
				err := messages.ParseMessage(m, &b)
				if err != nil {
					log.Println(err)
					continue
				}
				log.Printf("%+v\n", b)
				continue
			}

			var txs messages.Transaction
			err := messages.ParseMessage(m, &txs)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Printf("%+v\n", txs)
		}
	}
}
