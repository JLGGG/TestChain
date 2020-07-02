package TestChain

import (
	"bytes"
	"encoding/json"
	"reflect"
	"time"
)
type Block struct {
	PreviousHash []byte
	Hash []byte
	Height uint64
	TransactionList []Transaction
	//addition need..
	Timestamp time.Time
	Creator string
	State StateOfBlock
}

type HeightOfBlockchain = uint64
type StateOfBlock = string

const (
	Generated StateOfBlock = "Generated"
	StageState StateOfBlock = "StageState"
	Published StateOfBlock = "Published"
)
