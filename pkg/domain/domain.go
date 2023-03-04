package domain

type OpCode int16

const (
	OP_ADD OpCode = iota
	OP_LOAD
	OP_AND
	OP_HALT
)

type Register int16

const (
	R1 Register = iota
	R2
	R3
	R4
	R5
	R6
	R7
	PC
	COUNT
)
