package TestChain

import (
	"crypto/sha256"
	"time"
)
type Transaction = string // need to edit this variable.
type HeightOfBlockchain = uint64
type StateOfBlock = string

type Block struct {
	Index int // Current position in blockchain
	Header BlockHeader
	Hash []byte // Current block hash value
	Height HeightOfBlockchain // Blockchain' size
	TransactionList []Transaction // When calculating merkle root use this variable.
	State StateOfBlock
	counterOfTransaction int
}

type BlockHeader struct {
	version string
	Creator string
	Timestamp time.Time
	MerkleRoot string
	PreviousHash []byte
}

const (
	Generated StateOfBlock = "Generated"
	StageState StateOfBlock = "StageState"
	Published StateOfBlock = "Published"
)

func (block *Block) SetHash(hash []byte) {
	block.Hash = hash
}

func (block *Block) SetPreviousHash(previousHash []byte){
	block.Header.PreviousHash = previousHash
}

func (block *Block) SetHeight(height uint64) {
	block.Height = height
}

func (block *Block) SaveTransaction(transaction Transaction) {
	i:=0
	block.TransactionList[i] = transaction
}

func (block *Block) SetCreator(creator string){
	block.Header.Creator = creator
}

func (block *Block) SetVersion(version string) {
	block.Header.version = version
}

func (block *Block) SetTimestamp() {
	block.Header.Timestamp = time.Now()
}

func (block *Block) SetState(state StateOfBlock) {
	block.State = state
}
/*
//Convert set of byte to qword(uint64).
func byteSliceToUint64(n []byte) uint64 {
	var temp = (uint64(n[0]) << 56) | (uint64(n[1]) << 48) | (uint64(n[2]) << 40) | (uint64(n[3]) << 32) |
		(uint64(n[4]) << 24) | (uint64(n[5]) << 16) | (uint64(n[6]) << 8) | uint64(n[7])
	return temp
}*/

func (block *Block) CalculateBlockHash() {
	h := sha256.New()
	var input []byte

	input = append(input, []byte(block.Header.Creator)... )
	input = append(input, []byte(block.Header.version)...)
	input = append(input, []byte(block.Header.Timestamp.String())...)
	input = append(input, []byte(block.Header.MerkleRoot)...)
	input = append(input, block.Header.PreviousHash...)
	h.Write(input)
	block.Hash = h.Sum(nil)
}

func (block *Block) CalculateMerkleRoot() {
	h := sha256.New()
	var input []byte
	var tempHash []byte
	var temp [][]byte

	for i, j:=0, 0; i<block.counterOfTransaction; i=+2 {
		input = append(input, []byte(block.TransactionList[i])...)
		if i+1 == block.counterOfTransaction{
			h.Write(input)
			//need to save to 2-dimention array.
		}
		input = append(input, []byte(block.TransactionList[i+1])...)
		h.Write(input)
		tempHash = h.Sum(nil)
	}
}

func (block *Block) GetHash() []byte {
	return block.Hash
}

func (block *Block) GetPreviousHash() []byte {
	return block.Header.PreviousHash
}

func (block *Block) GetHeight() uint64 {
	return block.Height
}



