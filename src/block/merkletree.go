package block

type merkletree struct {
	root *transactionNode
	merkleRoot string
	nodes []*transactionNode
}

type transactionNode struct {
	ptrTree *merkletree
	ptrParent *transactionNode
	ptrLeft *transactionNode
	ptrLight *transactionNode
	isLeaf bool
	hash []byte
	transaction Transaction
}

func Create() *merkletree {
	mt := new(merkletree)

	return mt
}

func (mt *merkletree) IsEmpty(){

}

func (mt *merkletree) MakeMT() {

}

func (mt *merkletree) GetMerklePath() {

}

func (mt *merkletree) GetMerkleRoot() {

}

func (mt *merkletree) CalculateHash() {

}