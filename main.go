package main

import (
	"io/ioutil"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/core/asm"
)

func main() {
	contractBin, err := ioutil.ReadFile("testdata/GameStop.bin")
	if err != nil {
		log.Fatalln(err)
	}
	out, err := asm.Disassemble(contractBin)
	if err != nil {
		// Currently, the Disassemble function from go-ethereum chokes on
		// some opcodes in the GME contract bytecode. Not sure if this should
		// be expected or not, as the etherscan disassembler also chokes on some
		// opcodes. /shrug
		// Error:
		// 2021/05/26 12:12:49 incomplete push instruction at 11013
		log.Fatalln(err)
	}
	spew.Dump(out)
}
