package Block

import (
	"time"
	"strconv"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

//通过代码，完成PoW挖矿

//声明区块
type Block struct {
	//区块高度
	 Index int
	 //时间戳
     TimeStamp string
     //难度
     Diff  int
     //上一个区块的哈希
     PreHash  string
     //当前区块哈希
     HashCode  string
     //随机数
     Nonce  int
     //交易信息
     Data string
}

//创建创世区块
func GenerateFirstBlock(data string) Block {
	var firstBlock Block
	firstBlock.Index = 1
	firstBlock.TimeStamp = time.Now().String()
	firstBlock.Diff = 4
	firstBlock.Nonce = 0
	firstBlock.Data = data
	firstBlock.HashCode = GenerateBlockHashValue(firstBlock)
	return firstBlock
}


//生产区块的哈希值
func GenerateBlockHashValue(block Block) string {

	var hashData = strconv.Itoa(block.Index) + block.TimeStamp + strconv.Itoa(block.Diff) + strconv.Itoa(block.Nonce) +
		block.Data
    //hash算法
    var hash = sha256.New()
    hash.Write([]byte(hashData))
    hashed := hash.Sum(nil)
    //将字节装换成字符串
    return hex.EncodeToString(hashed)
}

//产生下一区块
func GenerateNextBlock(data string,oldBlock Block) Block {

	  var newBlock Block
	  newBlock.Data = data
	  newBlock.TimeStamp = time.Now().String()
	  //应该由矿工操作处理 暂时设置为0
	  newBlock.Nonce = 0
	  newBlock.Diff = 4
	  newBlock.PreHash = oldBlock.HashCode
      newBlock.Index = 2
      //应该填写PoW挖矿成功后的值
      newBlock.HashCode = pow(newBlock.Diff,&newBlock)

      return newBlock
}

//PoW工作量证明
func pow(diff int, block *Block) string {
	for {
		hash := GenerateBlockHashValue(*block)

		if strings.HasPrefix(hash,strings.Repeat("0",diff)) {
			//挖矿成功
			return hash
		}else {
			block.Nonce ++
		}
	}
	return ""
}

