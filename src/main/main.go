package main

import (
	"block"
)

func main() {
	blockTest := new(block.Block)

	blockTest.SaveTransaction("hello")
	blockTest.SaveTransaction("world")
	blockTest.SaveTransaction("block")

	blockTest.CalculateMerkleRoot()

	/*
	blockTest.SetCreator("test")
	blockTest.SetVersion("1")
	blockTest.SetTimestamp()
	blockTest.Header.MerkleRoot = "hello"
	blockTest.Header.PreviousHash = []byte{1,2}

	blockTest.CalculateBlockHash()
	fmt.Print(blockTest.Hash)
	*/


}
