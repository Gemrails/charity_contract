package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type CharityNote struct {
	Direction    string `json:"direction"`
	CostMoney    int32  `json:"costMoney"`
	DonationName string `json:"donationName"`
}

type CharityUser struct {
	DonationName string `json:"donationName"`
	ALLMoney     string `json:"allMoney"`
	LeftMoney    int32  `json:"leftMoney"`
}

func (s *SmartContract) Init(api shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(api shim.ChaincodeStubInterface) peer.Response {

	function, args := stub.GetFunctionAndParameters()

	switch function {
	case "donation":
		return s.donation(api, args)
	case "queryDealOnce":
		return s.queryDealOnce(api, args)
	case "queryDealALL":
		return s.queryDealALL(api)
	}

	return shim.Error("Invalid function name.")
}

func (s *SmartContract) donation(api shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return "", fmt.Errorf("need your name and money acount")
	}
	moneyCount, err := strconv.Atoi(args[2])
	if err != nil {
		return "", fmt.Errorf("strconv money error.")
	}
	cn := &CharityNote{
		Direction:    "origin",
		CostMoney:    0,
		DonationName: args[1],
	}
	cu := &CharityUser{
		DonationName: args[1],
		ALLMoney:     moneyCount,
		LeftMoney:    moneyCount,
	}

	return shim.Success(nil)
}

func (s *SmartContract) queryDealOnce(api shim.ChaincodeStubInterface, args []string) peer.Response {

	return shim.Success(nil)
}

func (s *SmartContract) queryDealALL(api shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)
}

func main() {
	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
