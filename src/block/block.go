package blockchain

import (
	"errors"
	_ "fmt"
	_ "math"
	"time"
)

type Transaction = string // need to edit this variable.
type BlockHeight = uint64
type BlockState = string

const SHA256_LENGTH = 32 // 32 * 8bits = 256bits

type Block struct {
	Index                int // Current position in blockchain
	Header               BlockHeader
	Hash                 []byte             // Current block hash value
	Height               BlockHeight // Blockchain' size
	TransactionList      []Transaction      // When calculating merkle root use this variable.
	State                BlockState
	counterOfTransaction int
}

type BlockHeader struct {
	version      string
	Creator      string
	Timestamp    time.Time
	MerkleRoot   string
	PreviousHash []byte
}

const (
	Created  BlockState = "Created"
	InStage BlockState = "InStage"
	Published  BlockState = "Published"
)

func (block *Block) SetHash(hash []byte) {
	block.Hash = hash
}

func (block *Block) SetPreviousHash(previousHash []byte) {
	block.Header.PreviousHash = previousHash
}

func (block *Block) SetHeight(height uint64) {
	block.Height = height
}

func (block *Block) SaveTransaction(transaction Transaction) error {
	if transaction != "" {
		block.TransactionList = append(block.TransactionList, transaction)
		block.counterOfTransaction++
		return nil
	} else {
		return nil
	}
	return errors.New("Transaction type error")

}

func (block *Block) SetCreator(creator string) {
	block.Header.Creator = creator
}

func (block *Block) SetVersion(version string) {
	block.Header.version = version
}

func (block *Block) SetTimestamp() {
	block.Header.Timestamp = time.Now()
}

func (block *Block) SetState(state BlockState) {
	block.State = state
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

func (block *Block) GetIndex() int {
	return block.Index
}

func (block *Block) GetState() BlockState {
	return block.State
}

func (block *Block) GetCountOfTransaction() int {
	return block.counterOfTransaction
}

func (block *Block) GetVersion() string {
	return block.Header.version
}

func (block *Block) GetCreator() string {
	return block.Header.Creator
}

func (block *Block) GetTimestamp() string {
	return block.Header.Timestamp.String()
}

func (block *Block) GetMerkleRoot() string {
	return block.Header.MerkleRoot
}
