package main

import (
	"io/ioutil"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/core/asm"
)

func main() {
	dat, err := ioutil.ReadFile("testdata/GameStop.bin")
	if err != nil {
		log.Fatalln(err)
	}
	out, err := asm.Disassemble(dat)
	if err != nil {
		log.Fatalln(err)
	}
	spew.Dump(out)
	// log.Fatalln(asm.PrintDisassembled(out[0]))
}
