package types

type Header struct {
	// Version of the block.  This is not the same as the protocol version.
	Version int32

	// time when block was created
	TimeStamp int64

	// Height    uint64 // block height

	// hash of the previous block
	PrevBlock []byte

	// hash of the block state
	MerkelRoot []byte

	// state of the block : 0->commit ; 1->valid ; 2->invalid
	state byte
}

type Body struct {
	Transactions []*Transaction
}

type Block struct {
	Header *Header
	Body   *Body
}

func NewBlock() *Block {
	return &Block{
		Header: &Header{},
		Body:   &Body{},
	}
}
