package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
	"jry/chaincode/bean"
)

type RegisterContract struct {

}

func (s *RegisterContract)Init(APIstub shim.ChaincodeStubInterface)sc.Response  {

	return shim.Success(nil)
}
func (s *RegisterContract)Invoke(APIstub shim.ChaincodeStubInterface)sc.Response  {

	function,args := APIstub.GetFunctionAndParameters()

	if function == "userRegister" {
		return s.userRegister(APIstub,args)
	}else if function == "queryUserByAddress" {
		//return s.
	}
	return  shim.Error("Invalid Smart Contract function name")
}

func main() {

	err := shim.Start(new(RegisterContract))
	if err != nil {
		bean.InfoLog("注册链码启动失败：",err)
	}
}
