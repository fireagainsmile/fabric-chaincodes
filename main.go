package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main(){
	sc := new(SimpleContract)
	cc, err := contractapi.NewChaincode(sc)
	if err != nil {
		panic(err.Error())
	}
	if err = cc.Start(); err != nil {
		panic(err.Error())
	}
}