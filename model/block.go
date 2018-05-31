package model

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
	"log"
	"github.com/davecgh/go-spew/spew"
	"sync"
)

type Block struct {
	Index int
	Timestamp string
	BPM int
	Hash string
	Prehash string
}

var Blockchain []Block

var mutex = &sync.Mutex{}


func CalcualteHash(block Block) string {
	record:=string(block.Index)+block.Timestamp+string(block.BPM)+block.Prehash
	h:=sha256.New()
	b2:=[]byte(record)
	h.Write(b2)
	hashed:=h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func GenerateBlock(oldBlock Block,BPM int) (Block,error) {
	log.Println("generateblock, oldblock:"+spew.Sdump(oldBlock))
	var newBlock Block
	t:=time.Now()
	newBlock.Index=oldBlock.Index+1
	newBlock.Timestamp=t.String()
	newBlock.BPM=BPM
	newBlock.Prehash=oldBlock.Hash
	newBlock.Hash=CalcualteHash(newBlock)
	log.Println("generateblock, newblock:"+spew.Sdump(newBlock))
	return newBlock,nil
}

func IsBlockValid(newBlock ,oldBlock Block) bool {
	if oldBlock.Index+1!=newBlock.Index {
		return false
	}
	if oldBlock.Hash!=newBlock.Prehash {
		return false
	}
	if CalcualteHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func ReplaceChain(newBlocks []Block) {
	mutex.Lock()
	if len(newBlocks)>len(Blockchain) {
		Blockchain = newBlocks
	}
	mutex.Unlock()
}

