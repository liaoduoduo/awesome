package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type EducationChaincode struct {

}

func (t *EducationChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response{

	return shim.Success(nil)
}

func (t *EducationChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// 获取用户意图
	fun, args := stub.GetFunctionAndParameters()

	if fun == "addCom" {
		return t.addCom(stub, args)
	}
	return shim.Error("没找到函数")
}

func main(){
	err := shim.Start(new(EducationChaincode))
	if err != nil{
		fmt.Printf("启动EducationChaincode时发生错误: %s", err)
	}
}