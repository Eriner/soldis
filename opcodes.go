package main

type op struct {
	Mnemonic string
	Value    int
	Pops     int
	Pushes   int
}

var opcodes map[int]op = map[int]op{
	0x00: op{
		Mnemonic: "STOP",
		Value:    0x00,
		Pops:     0,
		Pushes:   0,
	},
	0x01: op{
		Mnemonic: "ADD",
		Value:    0x01,
		Pops:     2,
		Pushes:   1,
	},
	0x02: op{
		Mnemonic: "MUL",
		Value:    0x02,
		Pops:     2,
		Pushes:   1,
	},
}
