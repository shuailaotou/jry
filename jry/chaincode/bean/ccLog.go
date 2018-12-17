package bean

import (
	"encoding/json"
	"fmt"
)

/**
 * 9位字符串，
 * 前3为大问题，500 服务异常，501 API异常，502  CHAINCODE
 * 中3为模块分类(101链, 102区块, 103账户)，
 * 后3为子功能问题
 *
 * @author leiming
 */

var CODE_CHAIN_DEFAULT = CODE_ERROR{"502000000", "default error"}


var CODE_502030001 = CODE_ERROR{"502030001", "register:args error"}
var CODE_502030002 = CODE_ERROR{"502030002", "register:json Unmarshal error"}
var CODE_502030003 = CODE_ERROR{"502030003", "register:signature validation error"}
var CODE_502030004 = CODE_ERROR{"502030004", "register:get user failed"}
var CODE_502030005 = CODE_ERROR{"502030005", "register:user has already exsit"}
var CODE_502030006 = CODE_ERROR{"502030006", "register:json marshal error"}

type CODE_ERROR struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func CodeRet(codeErr CODE_ERROR) string {
	jsons, _ := json.Marshal(codeErr)
	return string(jsons)
}

const  DEBUG_EN =  true
const INFO_EN   =  true

const ERROR_PREFIX  = "[Error]"
const DEBUG_PREFIX  =  "[Debug]"
const INFO_PREFIX   =  "[Info]"

func ErrorLog(a string,msg interface{})  {

	fmt.Println(ERROR_PREFIX ,a,":", msg)

}

func DebugLog(a string,msg interface{})  {

	if DEBUG_EN == true{
		fmt.Println(DEBUG_PREFIX ,a,msg)
	}

}

func InfoLog(a,msg interface{})  {

	if INFO_EN == true {
		fmt.Println(INFO_PREFIX ,a, msg)
	}

}