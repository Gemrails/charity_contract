package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type CharityNote struct {
	Direction string `json:"direction"`
	CostMoney int32 `json:"costMoney"`
	DonationName string `json:"donationName"`
	ALLMoney  int32	`json:"allMoney"`
	LeftMoney	int32	`json:"leftMoney"`
}

func (s *SmartContract) Init(api shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(api shim.ChaincodeStubInterface) peer.Response {
	// Extract the function and args from the transaction proposal
	function, args := stub.GetFunctionAndParameters()
	if function == "donation"{
		return s.donation(api, args)
	}else if function == "queryDealOnce"{
		return s.queryDealOnce(api, args)
	}else if function == "queryDealALL"{
		return s.queryDealALL(api)
	}
	return shim.Error("Invalid function name.")
}

func (s *SmartContract)donation(api shim.ChaincodeStubInterface, args[]string) peer.Response{
	if len(args) != 2 {
		return "", fmt.Errorf("need your name and money acount")
	}
	cn := &charityNote{
		UserName: args[1]
		
	}
	return shim.Success(nil)
}

func (s *SmartContract)queryDealOnce(api shim.ChaincodeStubInterface, args[]string) peer.Response{

	return shim.Success(nil)
}

func (s *SmartContract)queryDealALL(api shim.ChaincodeStubInterface) peer.Response{

	return shim.Success(nil)
}

func main() {
	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
