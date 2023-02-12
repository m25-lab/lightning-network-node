package message

import (
	"encoding/json"

	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
)

type AddWhitelist struct {
	Publickey string
}

func ValidateAddWhitelist(req *pb.SendMessageRequest) bool {
	var addWhitelist AddWhitelist

	if err := json.Unmarshal([]byte(req.Data), &addWhitelist); err != nil {
		return false
	}

	//import pubkey
	var pubkey *secp256k1.PubKey
	pubkey.Key = []byte(addWhitelist.Publickey)

	//check pubkey
	if pubkey == nil {
		return false
	}

	if addWhitelist.Publickey == "" {
		return false
	}
}
