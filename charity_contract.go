package main

import (
	"charity_contract/tools"
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
	DealNumbers  int    `json:"dealNumbers"`
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
	case "queryUserInfo":
		return s.queryUserInfo(api, args)
	case "donationRules":
		return s.donationRules(api, args)
	}

	return shim.Error("Invalid function name.")
}

func (s *SmartContract) set(api shim.ChaincodeStubInterface, key string, value string) error {
	err := stub.PutState(args[0], []byte(value))
	if err != nil {
		return fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return nil
}

func (s *SmartContract) get(api shim.ChaincodeStubInterface, key string) (string, error) {
	value, err := stub.GetState(key)
	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}
	if value == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}
	return value, nil
}

func (s *SmartContract)getRange(api shim.ChaincodeStubInterface, keyStart string, keyEnd string) (shim.StateQueryIteratorInterface, error){
	resultsIterator, err := api.GetStateByRange(keyStart, keyEnd)
	if err != nil {
		return shim.StateQueryIteratorInterface, fmt.Errorf("get range error, %s", err)
	}
	defer resultsIterator.Close()
	return resultsIterator, nil
}

func (s *SmartContract) donation(api shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("need your name and deal numbers")
	}
	moneyCount, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("strconv money error.")
	}
	cNote := &CharityNote{
		Direction:    tools.D0,
		CostMoney:    0,
		DonationName: args[1],
	}
	cUser := &CharityUser{
		DonationName: args[1],
		ALLMoney:     moneyCount,
		LeftMoney:    moneyCount,
		DealNumbers:  0
	}
	cNoteKey := tools.Skey(cNote.DonationName, cUser.DealNumbers)
	cNoteBytes, _ := json.Marshal(cNote)
	err := s.set(api, cNoteKey, cNoteBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	cUserKey := cNote.DonationName
	cUserBytes, _ := json.Marshal(cUser)
	s.set(api, cUserKey, cUserBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (s *SmartContract) queryDealOnce(api shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("need your name and deal numbers")
	}
	if args[2] == 0 {
		return shim.Error("cant query 0 nums deal.")
	}
	cNoteKey := tools.Skey(args[1], args[2])
	cNoteValue, err := s.get(api, cNoteKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(cNoteValue)
}

func (s *SmartContract) queryDealALL(api shim.ChaincodeStubInterface) peer.Response {
	if len(args) != 1{
		return shim.Error("need your name")	
	}
	cUserValue := s.get(api, args[1])
	cUser := &CharityUser{}
	json.Unmarshal(cUserValue, &cUser)
	totalDealNums, err := strconv.Atoi(cUser.DealNumbers)
	if err != nil {
		return shim.Error("strconv totalDealNums error.")
	}
	cNoteKeyEnd := tools.Skey(args[1], totalDealNums)
	cNoteKeyStart := tools.Skey(args[1], 1)
	resultsIter, err := api.GetStateByRange(cNoteKeyStart, cNoteKeyEnd)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()	

	var buffer bytes.Buffer
	buffer.WriteString("[[[")
	bArrayMemberAlreadyWritten := false
	for resultsIter.HasNext(){
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten == true	
	}
	buffer.WriteString("]]]")	

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract)queryUserInfo(api shim.ChaincodeStubInterface, name string)peer.Response{

	return shim.Success(nil)
}

func (s *SmartContract)donationRules(api shim.ChaincodeStubInterface, model string)peer.Response{
	
	return shim.Success(nil)
}

func main() {
	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
