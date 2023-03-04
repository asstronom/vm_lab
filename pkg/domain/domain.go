package domain

type OpCode int16

// коди, що дозволяють визначати процесору яку операцію виконувати
const (
	OP_ADD OpCode = iota
	OP_LOAD
	OP_AND
	OP_HALT
)

// енум клас що вказує на регістри процесору
type Register int16

const (
	//R0 - R7 - загальні регістри
	R0 Register = iota
	R1
	R2
	R3
	R4
	R5
	R6
	R7
	//program counter - вказує на наступну операцію
	PC
	//кількість регістрів
	COUNT
)
