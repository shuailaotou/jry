package bean

import (
	"encoding/hex"
	"github.com/haltingstate/secp256k1-go"

	"crypto/sha256"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"strings"
	"time"
)

func SignatureValid(msg, address, sig string) bool {

	var messageByte = []byte(msg)

	if len(sig) != 130 {
		return false
	}

	signature, err := hex.DecodeString(sig)
	if err != nil {
		return false
	}
	recoverPubKey := secp256k1.RecoverPubkey(Keccak256(messageByte), signature)
	uncompressedRecoverPubkey := secp256k1.UncompressPubkey(recoverPubKey)
	recoverAddress := PubkeyToAddress(uncompressedRecoverPubkey)
	if address == hex.EncodeToString(recoverAddress) {
		return true
	}
	return false

}

func SignatureValidData(info *Trx) bool {

	if strings.Count(info.Payload, "") == 1 {
		return false
	}
	if strings.Count(info.Signatures[0].Sign, "") == 1 {

		return false
	}
	if strings.Count(info.Signatures[0].Addr, "") == 1 {
		return false
	}

	return SignatureValid(info.Payload, info.Signatures[0].Addr, info.Signatures[0].Sign)

}

func Keccak256(data ...[]byte) []byte {

	d := sha3.NewKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

func PubkeyToAddress(publicKey []byte) []byte {

	return Keccak256(publicKey[1:])[12:]

}

func CheckTimestamp(tranTimeStamp string) bool {

	the_time, err := time.Parse("2006-01-02T15:04:05Z", tranTimeStamp)

	if err != nil {
		return false
	}
	tranTimeStampUnix := the_time.Unix()
	timeStampSys := time.Now().UTC().Unix()

	if tranTimeStampUnix < (timeStampSys-10*60) || tranTimeStampUnix > (timeStampSys+10*60) {
		return false
	}
	return true
}

func CheckTxid(msg, txid string) bool {

	h := sha256.New()
	h.Write([]byte(msg))
	txidH := h.Sum(nil)
	if txid == hex.EncodeToString(txidH) {
		return true
	}
	return false
}
func GenMsgHash(msg string) string {

	h := sha256.New()
	h.Write([]byte(msg))
	txidH := h.Sum(nil)

	return hex.EncodeToString(txidH)

}
