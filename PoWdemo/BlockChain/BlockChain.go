package BlockChain

import (
	"PoWdemo/Block"
	"fmt"
)

//通过链表的形式维护区块链的业务

type Node struct {
	//指针域
	NextNode *Node
	//数据域
	Data *Block.Block
}

var HeadNode *Node

//创建头节点，保存创世区块
func CreatHeadNode(data *Block.Block) *Node {

	var headNode *Node = new(Node)
	headNode.NextNode = nil

	headNode.Data = data

	HeadNode = headNode
	return headNode
}

//添加节点，当挖矿成功以后添加区块
func AddNode(data *Block.Block, prefNode *Node) *Node {

	var newNode *Node = new(Node)
	newNode.Data = data
	newNode.NextNode = nil
	prefNode.NextNode = newNode
	return newNode
}

func ShowBlockChain(headNode *Node) {
	var node *Node = headNode
	fmt.Print(node.Data)
	for {
		if node.NextNode != nil {
			node = node.NextNode
			fmt.Print("----", node.Data, "----")
		} else if node.NextNode == nil {
			break
		}
	}
}
