package model

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

var bcServer chan []Block

func TcpMain() {
	bcServer = make(chan []Block)

	// start TCP and serve TCP server
	server, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, "Enter a new BPM:")
	scanner := bufio.NewScanner(conn)

	bc := Blockchain

	go func() {
		for scanner.Scan() {
			bpm, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Println("input is not a number: ", scanner.Text(), err)
				continue
			}
			newBlock, err := GenerateBlock(bc[len(bc)-1], bpm)

			if err != nil {
				log.Println(err)
				continue
			}
			if IsBlockValid(newBlock, bc[len(bc)-1]) {
				newBlockChain := append(bc, newBlock)
				ReplaceChain(newBlockChain)
			}

			bcServer <- bc
			io.WriteString(conn, "\nEnd.")

		}
	}()

	// 模拟接收广播
	go func() {
		for {
			time.Sleep(5 * time.Second)
			mutex.Lock()
			output, err := json.Marshal(bc)
			if err != nil {
				log.Fatal(err)
			}
			mutex.Unlock()
			io.WriteString(conn, "\n"+string(output))
		}
	}()

	for range bcServer {
		//spew.Dump(bc)
	}

}
