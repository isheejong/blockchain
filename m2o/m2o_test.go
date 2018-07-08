package main

import (
	"testing"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

/*
func TestUpdateExchangeRate(t *testing.T) {
	stub := shim.NewMockStub("mockChaincodeStub",new(M2OChaincode))
	if stub == nil{
		t.Fatalf("MockStub creation failed")
	}
	args := [][]byte{
		[]byte("updateExchangeRate"), []byte("국민카드"),
		[]byte("10.123")}

	fmt.Println("- start Test")

	invokeResult := stub.MockInvoke("1111", args)
	if invokeResult.Status != 200 {
		t.Errorf("updateExchangeRate returned non-OK status, got: %d, want: %d.", invokeResult.Status, 200)
	}
}
*/
func TestConvertToM2OToken(t *testing.T){
	stub := shim.NewMockStub("mockChaincodeStub",new(M2OChaincode))
	if stub == nil{
		t.Fatalf("MockStub creation failed")
	}


	dummy := [][]byte{
		[]byte("updateExchangeRate"), []byte("국민카드"),
		[]byte("10.123")}

	invokeResult1 := stub.MockInvoke("1111", dummy)
	if invokeResult1.Status != 200 {
		t.Errorf("updateExchangeRate returned non-OK status, got: %d, want: %d.", invokeResult1.Status, 200)
	}

	args1 := [][]byte{
		[]byte("convertToM2OToken"),            // chaincode function
		[]byte("isheejong"),                    // M2OUID
		[]byte("국민카드"), 				 	 // partnerName
		[]byte("1111-2222-3333-4444-5555"),     // partnerUID
		[]byte("2000"),                         // FromPoint
		[]byte("2018-07-06 12:57:11"),			// Date
	}


	invokeResult2 := stub.MockInvoke("2222", args1)
	if invokeResult2.Status != 200 {
		t.Errorf("updateExchangeRate returned non-OK status, got: %d, want: %d.", invokeResult2.Status, 200)
	}

	args2 := [][]byte{[]byte("queryUserTotal"), []byte("isheejong"),}

	invokeResult3 := stub.MockInvoke("3333", args2)
	if invokeResult3.Status != 200 {
		t.Errorf("updateExchangeRate returned non-OK status, got: %d, want: %d.", invokeResult3.Status, 200)
	}

	fmt.Println("----------------")
	fmt.Println("Final result")
	fmt.Println(string(invokeResult3.GetPayload()))
}
/*
func TestQueryUserTransactionsForAll(t *testing.T) {
	stub := shim.NewMockStub("mockChaincodeStub",new(M2OChaincode))
	if stub == nil{
		t.Fatalf("MockStub creation failed")
	}


	dummy1 := [][]byte{
		[]byte("convertToM2OToken"),           // chaincode function
		[]byte("isheejong"),                   // M2OUID
		[]byte("국민카드"), 					// partnerName
		[]byte("1111-2222-3333-4444-5555"),    // partnerUID
		[]byte("2000"),                        // FromPoint
	}

	dummy2 := [][]byte{
		[]byte("convertToM2OToken"),           // chaincode function
		[]byte("isheejong"),                   // M2OUID
		[]byte("현대카드"), 					// partnerName
		[]byte("2222-1111-3333-1111-4444"),    // partnerUID
		[]byte("3000"),                        // FromPoint
	}

	stub.MockInvoke("4444", dummy1)
	stub.MockInvoke("5555", dummy2)

	fmt.Println("- start Test")
	args := [][]byte{
		[]byte("queryUserTransactionsForAll"), []byte("isheejong"),}

	invokeResult := stub.MockInvoke("1111", args)
	if invokeResult.Status != 200 {
		t.Errorf("queryUserTransactionsForAll returned non-OK status, got: %d, want: %d.", invokeResult.Status, 200)
	}
}
*/