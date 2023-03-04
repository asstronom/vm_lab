package instruction

import (
	"example.com/pkg/domain"
)

func makeTemplateWithOpCode(code domain.OpCode) int16 {
	var instruction int16
	instruction = (instruction | int16(code)) << 12
	return instruction
}

type InstructionBuilder interface {
	Compile() int16
}

type ADD struct {
	DR domain.Register
	S1 domain.Register
	S2 domain.Register
}

func (builder ADD) Compile() int16 {
	instruction := makeTemplateWithOpCode(domain.OP_ADD)
	instruction = instruction | (int16(builder.DR) << 9)
	instruction = instruction | (int16(builder.S1) << 6)
	instruction = instruction | (int16(builder.S2))
	return instruction
}

type LOAD struct {
	DR       domain.Register
	PCOffset int16
}

func (builder LOAD) Compile() int16 {
	instruction := makeTemplateWithOpCode(domain.OP_LOAD)
	instruction = instruction | (int16(builder.DR) << 9)
	instruction = instruction | (int16(builder.PCOffset))
	return instruction
}

type AND struct {
	DR domain.Register
	S1 domain.Register
	S2 domain.Register
}

func (builder AND) Compile() int16 {
	instruction := makeTemplateWithOpCode(domain.OP_AND)
	instruction = instruction | (int16(builder.DR) << 9)
	instruction = instruction | (int16(builder.S1) << 6)
	instruction = instruction | (int16(builder.S2))
	return instruction
}

type HALT struct {
}

func (builder HALT) Compile() int16 {
	return makeTemplateWithOpCode(domain.OP_HALT)
}
