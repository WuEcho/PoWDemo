package main

import (
	"PoWdemo/Block"
	"PoWdemo/BlockChain"
	//"fmt"
)

func main() {

	var headBlock = Block.GenerateFirstBlock("创世区块")

	var nextBlock = Block.GenerateNextBlock("第二区块", headBlock)

	var blockChain = BlockChain.CreatHeadNode(&headBlock)

	BlockChain.AddNode(&nextBlock, blockChain)

	BlockChain.ShowBlockChain(blockChain)
}
