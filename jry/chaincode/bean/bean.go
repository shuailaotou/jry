package bean

//用户信息
type UserInfo struct {
	ApiVersion     string `json:"apiVersion"`     //api版本
	PhoneNoHash    string `json:"phoneNoHash"`    //电话号码hash值
	VerifyCode     string `json:"verifyCode`      //验证码
	EnterPriceCode string `json:"enterPriceCode"` //企业邀请码
	RegTime        string `json:"regTime"`        //注册时间
	Addr           string `json:"addr"`           //地址
	//Signature      string `json:"signature"`      //数据签名结果
}

//黑名单请求
type BLRequest struct {
	RequestID           string `json:"requestID"`           //请求ID
	HashOfThreeElements string `json:"hashOfThreeElements"` //hash（name，phone，idcard）
	Type                string `json:"type"`                //类型
	Time                string `json:"time"`                //时间
	//Addr                string `json:"addr"`                //地址
	//Signature           string `json:"signature"`           //以上消息签名数据
}

//黑名单请求结果报告
type BLReport struct {
	Requestid string   `json:"requestid"` //请求ID
	BaseInfo  string   `json:"baseInfo"`  //
	Result    string   `json:"result"`    //请求结果状态
	SourceArr []Source `json:"sourceArr"` //结果了来源
	Report    string   `json:"report"`    //报告内容
	GenTime   string   `json:"genTime"`   //生成时间
	//SrvAddr   string   `json:"srvAddr"`   //服务端地址
	//SrvSig    string   `json:"srvSig"`    //服务端签名
}

type Source struct {
	SourceFrom string `json:"sourceFrom"`
}

type GwPayload struct {
	Data string `json:"data"`
	Meta string `json:"meta"`
}

type Trx struct {
	Payload    string `json:"payload"`
	Signatures []Sign `json:"signatureArray"`
}

type Sign struct {
	Addr string `json:"address"`
	Sign string `json:"signature"`
}
