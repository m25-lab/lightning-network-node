package core

type Commitment struct {
	channel         *Channel
	commitAmount    []uint64
	csvLock         uint32
	sender          Address
	receiver        Address
	senderSignature Signature
}
