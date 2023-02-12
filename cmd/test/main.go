package main

import (
	"encoding/hex"
	"fmt"

	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
)

func main() {
	// partACommitment, addressA, signature := client.CreateCommitmentFromA(50, 50, "secret from A", 10, 1)
	// partBCommitment, addressB, signature := client.CreateCommitmentFromB(partACommitment, addressA, signature, "secret from B")
	// client.StoreCommitmentFromA(partBCommitment, addressB, signature)

	// channelId := client.OpenChannelFromA(50, 50, 1)
	// client.OpenChannelFromB(channelId)

	// partACommitment, addressA, signature = client.CreateCommitmentFromA(40, 60, "secret from AA", 10, 1)
	// partBCommitment, addressB, signature = client.CreateCommitmentFromB(partACommitment, addressA, signature, "secret from BB")
	// client.StoreCommitmentFromA(partBCommitment, addressB, signature)

	// partACommitment, addressA, signature = client.CreateCommitmentFromA(70, 30, "secret from AA", 10, 1)
	// partBCommitment, addressB, signature = client.CreateCommitmentFromB(partACommitment, addressA, signature, "secret from BB")
	// client.StoreCommitmentFromA(partBCommitment, addressB, signature)

	// partAFundCommitment, addressA, signature := client.FundFromA(10, "secret from AA", 10, 1)
	// partBAcceptFundCommitment, addressB, signature := client.AcceptFundFromB(partAFundCommitment, addressA, signature, "secret from BB")
	// client.StoreAcceptFundCommitmentFromA(partBAcceptFundCommitment, addressB, signature)

	// closeChannel, addressA, signature := client.CloseChannel(10, 10, 1)
	// client.AcceptCloseChannel(closeChannel, addressA, signature)

	//get public ke

	acc := account.NewAccount(60)

	AAccount, _ := acc.ImportAccount("excuse quiz oyster vendor often spray day vanish slice topic pudding crew promote floor shadow best subway slush slender good merit hollow certain repeat")
	BAccount, _ := acc.ImportAccount("claim market flip canoe wreck maid recipe bright fuel slender ladder album behind repeat come trophy come vicious frown prefer height unknown thank damp")

	fmt.Println("account A:", AAccount.AccAddress().String())
	fmt.Println("account B:", BAccount.AccAddress().String())
	fmt.Println("account A public key:", AAccount.PublicKey().String())

	var pubkey secp256k1.PubKey
	data, err := hex.DecodeString("02E35D749E46BF716CE2C59525A162D06AC267F98E35CC56F2B8C695DF1AD16E27")
	if err != nil {
		panic(err)
	}
	pubkey.Key = data
	fmt.Println(pubkey.String())
}

func GetPublicKeyFromAddress(address string) {
	accAddress, err := types.AccAddressFromBech32(address)
	if err != nil {
		fmt.Println(err)
	}

	publicKey := accAddress.Bytes()
	fmt.Println(publicKey)

}
