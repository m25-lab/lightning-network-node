package main

import (
	"log"

	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
)

func main() {
	// partACommitment, addressA, signature := client.CreateCommitmentFromA(50, 50, "secret from A", 10, 1)
	// partBCommitment, addressB, signature := client.CreateCommitmentFromB(partACommitment, addressA, signature, "secret from B")
	// client.StoreCommitmentFromA(partBCommitment, addressB, signature)

	// ChannelID := client.OpenChannelFromA(50, 50, 1)
	// client.OpenChannelFromB(ChannelID)

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

	acc := account.NewAccount()

	AAccount, _ := acc.ImportAccount("excuse quiz oyster vendor often spray day vanish slice topic pudding crew promote floor shadow best subway slush slender good merit hollow certain repeat")
	BAccount, _ := acc.ImportAccount("claim market flip canoe wreck maid recipe bright fuel slender ladder album behind repeat come trophy come vicious frown prefer height unknown thank damp")

	log.Println("account A:", AAccount.AccAddress().String())
	log.Println("account B:", BAccount.AccAddress().String())
	log.Println("account A public key:", AAccount.PublicKey().String())

	newPKAccount := account.NewPKAccount("02E35D749E46BF716CE2C59525A162D06AC267F98E35CC56F2B8C695DF1AD16E27")
	log.Println(newPKAccount.PublicKey().String())
	log.Println(newPKAccount.AccAddress().String())
}
