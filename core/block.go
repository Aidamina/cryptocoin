package core

type Block struct {
	Prev Hash
	Height uint64
	Data []byte
	Nonce uint64
}

func (b Block) Hash() Hash{
	return CalculateHash(b.Data)
}


func NewBlock() * Block {
	return &Block{}
}



