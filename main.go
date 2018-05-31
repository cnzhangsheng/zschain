package main

import (
	"github.com/cnzhangsheng/zschain/model"
	"log"
	"time"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	log.SetPrefix("[blockchain]")
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)

	createGenesisBlock()

	tcpMain()
	//httpMain()

}

func createGenesisBlock() {
	t := time.Now()
	genesisBlock := model.Block{0, t.String(), 0, "", ""}
	model.Blockchain = append(model.Blockchain, genesisBlock)

	spew.Dump(genesisBlock)
}

func httpMain() {
	model.MakeMuxRouter()
}

func tcpMain() {
	model.TcpMain()
}
