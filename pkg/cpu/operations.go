package cpu

import "example.com/pkg/domain"

func (cpu *CPU) add(instr int16) {
	DR := (instr >> 9) & 0x7
	SR1 := (instr >> 6) & 0x7
	SR2 := instr & 0x7
	cpu.reg[DR] = cpu.reg[SR1] + cpu.reg[SR2]
}

func (cpu *CPU) load(instr int16) {
	DR := (instr >> 9) & 0x7
	pcOffset := signExtend(instr&0x1FF, 9)
	cpu.reg[DR] = cpu.memRead(cpu.reg[domain.PC] + pcOffset)
}

func (cpu *CPU) and(instr int16) {
	DR := (instr >> 9) & 0x7
	SR1 := (instr >> 6) & 0x7
	SR2 := instr & 0x7
	cpu.reg[DR] = cpu.reg[SR1] & cpu.reg[SR2]
}
