package bean

import (
	"encoding/json"
)

const (
	ARG_UNMARSHAL = "arg_unmarshal"
	ARG_MARSHAL = "arg_marshal"
)

type ArgsAss struct {
	Args   []string //参数数组
	ArgNum int      //期望参数count
	Object string   //对象名称
}
type PaResponse struct {
	ErrStr string
	ObjStr string
}

func ParaesAnalysis(als ArgsAss) PaResponse {

	pr := PaResponse{}

	if len(als.Args) != als.ArgNum {

		ErrorLog(als.Object, "args error")
		pr.ErrStr = als.Object + "args error"

		return pr

	}

	trx := &Trx{}
	err := json.Unmarshal([]byte(als.Args[0]), trx)

	if err != nil {

		ErrorLog(als.Object, "json Unmarshal error")
		pr.ErrStr = als.Object + "json Unmarshal error"
		return pr
	}

	b := SignatureValidData(trx)

	if !b {

		ErrorLog(als.Object, "signature validation error")
		pr.ErrStr = als.Object + "signature validation error"
		return pr
	}

	payloadDataString := trx.Payload
	payloadData := &GwPayload{}
	err = json.Unmarshal([]byte(payloadDataString), payloadData)

	if err != nil {
		ErrorLog(als.Object, "json Unmarshal error")
		pr.ErrStr = als.Object + "json Unmarshal error"
		return pr

	}
	if payloadData.Data == "" {

		ErrorLog(als.Object, "GW signature validation error")
		pr.ErrStr = als.Object + "GW signature validation error"
		return pr
	}

	objectPayloadString := payloadData.Data

	trxClient := &Trx{}
	err = json.Unmarshal([]byte(objectPayloadString), &trxClient)

	if err != nil {

		ErrorLog(als.Object, "json Unmarshal error")
		pr.ErrStr = als.Object + "json Unmarshal error"
		return pr

	}

	b = SignatureValidData(trx)
	if !b {

		ErrorLog(als.Object, "Cli signature validation error")
		pr.ErrStr = als.Object + "Cli signature validation error"
		return pr

	}
	objectString := trxClient.Payload
	pr.ObjStr = objectString

	return pr
}
