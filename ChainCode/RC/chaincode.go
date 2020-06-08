/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Chaincode is the definition of the chaincode structure.
type Chaincode struct {
}

type lookUpTable struct {
	MethodName string `json:"MethodName"`
	Subject    string `json:"Subject"`
	Object     string `json:"Object"`
	ScName     string `json:"ScName"`
	Creator    string `json:"Creator"`
	ScAddress  string `json:"ScAddress"`
	ABI        string `json:"ABI"`
}

// Init is called when the chaincode is instantiated by the blockchain network.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Init()", fcn, params)
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()

	switch fcn {
	case "methodRegister":
		return cc.methodRegister(stub, params)
	case "methodUpdate":
		return cc.methodUpdate(stub, params)
	case "methodDelete":
		return cc.methodDelete(stub, params)
	case "getContract":
		return cc.getContract(stub, params)
	default:
		return shim.Error("error")
	}
}

func (cc *Chaincode) methodRegister(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	if len(params) != 7 {
		return shim.Error("Incorrect number of parameter")
	}

	methodName, subject, object, scName, creator, scAddress, abi := params[0], params[1], params[2], params[3], params[4], params[5], params[6]

	lookUpTableData := &lookUpTable{MethodName: methodName, Subject: subject, Object: object, ScName: scName, Creator: creator, ScAddress: scAddress, ABI: abi}

	lookUpTableDataBytes, err := json.Marshal(lookUpTableData)
	fmt.Println(string(lookUpTableDataBytes))

	if err != nil {
		return shim.Error("failed to Marshal lookUpTableData, error : " + err.Error())
	}

	err = stub.PutState(methodName, lookUpTableDataBytes)

	if err != nil {
		return shim.Error("failed to PutState lookUpTableData, error : " + err.Error())
	}

	return shim.Success(nil)
}

func (cc *Chaincode) methodUpdate(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	if len(params) != 7 {
		return shim.Error("Incorrect number of parameter")
	}

	methodName, subject, object, scName, creator, scAddress, abi := params[0], params[1], params[2], params[3], params[4], params[5], params[6]

	lookUpTableData := lookUpTable{}
	lookUpTableDataBytes, err := stub.GetState(methodName)

	if err != nil {
		return shim.Error("failed to GetState lookUpTableData, error : " + err.Error())
	}

	err = json.Unmarshal(lookUpTableDataBytes, &lookUpTableData)

	if err != nil {
		return shim.Error("failed to UnMarshal lookUpTableData, error : " + err.Error())
	}

	lookUpTableData.Subject = subject
	lookUpTableData.Object = object
	lookUpTableData.ScName = scName
	lookUpTableData.Creator = creator
	lookUpTableData.ScAddress = scAddress
	lookUpTableData.ABI = abi

	lookUpTableDataBytes, err = json.Marshal(lookUpTableData)

	if err != nil {
		return shim.Error("failed to Marshal lookUpTableData, error : " + err.Error())
	}

	err = stub.PutState(methodName, lookUpTableDataBytes)

	if err != nil {
		return shim.Error("failed to PutState lookUpTableData, error : " + err.Error())
	}

	return shim.Success(nil)
}

func (cc *Chaincode) methodDelete(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	if len(params) != 1 {
		return shim.Error("Incorrect number of parameter")
	}

	methodName := params[0]

	err := stub.DelState(methodName)

	if err != nil {
		return shim.Error("failed to DelState methodName, error : " + err.Error())
	}

	return shim.Success(nil)
}

func (cc *Chaincode) getContract(stub shim.ChaincodeStubInterface, params []string) sc.Response {
	if len(params) != 1 {
		return shim.Error("Incorrect number of parameter")
	}

	methodName := params[0]

	lookUpTableData := lookUpTable{}
	lookUpTableDataBytes, err := stub.GetState(methodName)

	if err != nil {
		return shim.Error("failed to GetState lookUpTableData, error : " + err.Error())
	}

	err = json.Unmarshal(lookUpTableDataBytes, &lookUpTableData)

	if err != nil {
		return shim.Error("failed to UnMarshal lookUpTableData, error : " + err.Error())
	}

	fmt.Println(string(lookUpTableDataBytes))

	response := shim.Success([]byte("success"))
	response.Message = string('반환할 데이터들')
	return response
}
