package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"

	"encoding/json"
	"jry/chaincode/bean"
)

func (s *RegisterContract) queryUserByAddress(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {

	}
	return shim.Error("")
}

func (s *RegisterContract) userRegister(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	para := bean.ArgsAss{args, 1, "register"}
	res := bean.ParaesAnalysis(para)
	if res.ErrStr != "" {
		return shim.Error(res.ErrStr)
	}

	userInfo := bean.UserInfo{}
	err := json.Unmarshal([]byte(res.ObjStr), &userInfo)

	if err != nil {
		bean.ErrorLog("user", bean.CodeRet(bean.CODE_502030002))
		return shim.Error(bean.CodeRet(bean.CODE_502030002))
	}

	b, _ := APIstub.GetState(bean.ARG_MARSHAL + userInfo.Addr)
	user := bean.UserInfo{}
	json.Unmarshal(b, &user)

	if len(b) > 0 && user.PhoneNoHash == userInfo.PhoneNoHash {

		bean.ErrorLog("user", bean.CodeRet(bean.CODE_502030005))
		return shim.Error(bean.CodeRet(bean.CODE_502030005))

	}

	um,err:=json.Marshal(userInfo)
	if err != nil {
		bean.ErrorLog("user", bean.CodeRet(bean.CODE_502030006))
		return shim.Error(bean.CodeRet(bean.CODE_502030006))
	}

	APIstub.PutState(bean.ARG_MARSHAL+userInfo.Addr,um)
	APIstub.PutState(bean.ARG_UNMARSHAL + userInfo.Addr,[]byte(args[0]))

	return shim.Success(nil)

}
