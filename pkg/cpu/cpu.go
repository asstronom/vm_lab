package cpu

import (
	"fmt"
	"log"

	"example.com/pkg/domain"
)

type CPU struct {
	memory []int16
	reg    []int16
}

func NewCPU(memory []int16) *CPU {
	cpu := CPU{
		memory: make([]int16, 1<<16),
	}
	cpu.reg = make([]int16, domain.COUNT)
	copy(cpu.memory, memory)
	return &cpu
}

func (cpu *CPU) memRead(i int16) int16 {
	return cpu.memory[i]
}

func (cpu *CPU) Debug() {
	fmt.Println(cpu.reg)
}

func (cpu *CPU) Work() {
workLoop:
	for {
		instr := cpu.memRead(cpu.reg[int16(domain.PC)])
		op := domain.OpCode(instr >> 12)
		switch op {
		case domain.OP_ADD:
			cpu.add(instr)
		case domain.OP_LOAD:
			cpu.load(instr)
		case domain.OP_AND:
			cpu.and(instr)
		case domain.OP_HALT:
			break workLoop
		default:
			log.Fatalln("unknown instruction")
		}
		cpu.reg[domain.PC]++
	}
}
