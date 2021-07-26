package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

const DOC_COM_TYPE = "comObj"

func PutCom(stub shim.ChaincodeStubInterface, com Commodity) ([]byte, bool) {

	com.ObjectType = DOC_COM_TYPE

	b, err := json.Marshal(com)
	if err != nil {
		fmt.Println("序列化信息时发生错误")
		return nil, false
	}

	err = stub.PutState(com.EntityID, b)
	if err != nil {
		fmt.Println("putstate时发生错误")
		return nil, false
	}

	return b, true
}

func (t *EducationChaincode) addCom(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2{
		return shim.Error("给定的参数个数不符合要求")
	}

	var com Commodity
	err := json.Unmarshal([]byte(args[0]), &com)
	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}

	_, bl := PutCom(stub, com)
	if !bl {
		return shim.Error("PutCom时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息添加成功"))
}
