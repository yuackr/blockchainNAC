/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type Chaincode struct {
}

type SubjectOfRecord struct {
	Object       string
	Misbehavior  string
	Time         string
	Penalty      int
	LimitPenalty string
}

type MisshaviorRecordsJC struct {
	Subject          string            `json:"Subject"`
	SubjectOfRecords []SubjectOfRecord `json:"SubjectOfRecords"`
}

func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	fmt.Println("Init()", fcn, params)
	return shim.Success(nil)
}

// Invoke is called as a result of an application request to run the chaincode.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()

	switch fcn {
	case "misbehaviorJudge":
		return cc.misbehaviorJudge(stub, params)
	default:
		return shim.Error("error")
	}
}

func (cc *Chaincode) misbehaviorJudge(stub shim.ChaincodeStubInterface, params []string) sc.Response {

	//subject, msb := params[0], params[1]

	result := shim.Success([]byte("asdf"))
	result.Message = "2m"
	return result

	//  valueJSONAsBytes, err := stub.GetState(subject)

	//  var accessTime string = time.Now().Format("1996-07-28 00:00:00")
	//  // if not exist MisshaviorRecordsJC
	//  if err != nil {
	// 	 misshaviorRecordsJC := MisshaviorRecordsJC{
	// 		 Subject:          subject,
	// 		 SubjectOfRecords: nil,
	// 	 }
	// 	 valueJSONAsBytes, err := json.Marshal(misshaviorRecordsJC)

	// 	 if err != nil {
	// 		 return shim.Error(err.Error()), false, -1
	// 	 }

	// 	 err = stub.PutState(subject, valueJSONAsBytes)

	// 	 if err != nil {
	// 		 return shim.Error(err.Error()), false, -1
	// 	 }

	// 	 valueJSONAsBytes, err = stub.GetState(subject)
	// 	 if err != nil {
	// 		 return shim.Error(err.Error()), false, -1
	// 	 }

	//  }

	//  if valueJSONAsBytes.SubjectOfRecords != nil {
	// 	 var index int
	// 	 index = -1

	// 	 for i := 0; len(valueJSONAsBytes.SubjectOfRecords); i++ {
	// 		 if valueJSONAsBytes.SubjectOfRecords.Object == object {
	// 			 index = i
	// 			 break
	// 		 }
	// 	 }

	// 	 if index != -1{// if exist panalty
	// 		 if valueJSONAsBytes.SubjectOfRecords[index].LimitPenalty>= accessTime
	// 		 return shim.Error(err.Error()), false, 0
	// 	 }

	//  }
	//  var num1 int
	//  num1,err = strconv.Atoi(accessTime)
	//  var num2 int
	//  num2, err = strconv.Atoi(toLR)
	//  var duration int
	//  duration = num1 - num2
	//  var tempPenalty string
	//  tempPenalty, err = strconv.Itoi(num1+100)

	//  if duration <= 1{
	// 	 if valueJSONAsBytes.SubjectOfRecords != nil {
	// 		 var index int
	// 		 index = -1

	// 		 for i := 0; len(valueJSONAsBytes.SubjectOfRecords); i++ {
	// 			 if valueJSONAsBytes.SubjectOfRecords.Object == object {
	// 				 index = i
	// 				 break
	// 			 }
	// 		 }

	// 		 if index != -1{// if exist panalty
	// 			 valueJSONAsBytes.SubjectOfRecords[index].Time = accessTime
	// 			 valueJSONAsBytes.SubjectOfRecords[index].Penalty = 1
	// 			 valueJSONAsBytes.SubjectOfRecords[index].LimitPenalty = tempPenalty
	// 		 }

	// 	 }else{
	// 		 tempSOR := SubjectOfRecord{
	// 			 Object: object,
	// 			 Missbehavior: "too Prequently Access",
	// 			 Time: accessTime,
	// 			 Penalty: 1,
	// 			 LimitPenalty: tempPenalty,
	// 		 }
	// 		 valueJSONAsBytes.SubjectOfRecords.append(valueJSONAsBytes.SubjectOfRecords,tempSOR)
	// 	 }

	// 	 err = stub.PutState(subject, valueJSONAsBytes)

	// 	 if err != nil {
	// 		 return shim.Error(err.Error()), false, -1
	// 	 }
	//  }

	//  return shim.Success(nil), true, 0
}
