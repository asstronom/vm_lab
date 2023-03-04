package cpu

import (
	"fmt"
	"testing"

	"example.com/pkg/domain"
	"example.com/pkg/instruction"
)

func TestAdd(t *testing.T) {
	cpu := NewCPU(nil)
	cpu.reg[int16(domain.R2)] = 90
	cpu.reg[int16(domain.R4)] = 153
	fmt.Println(cpu.reg)
	cpu.add(instruction.ADD{
		DR: domain.R1,
		S1: domain.R2,
		S2: domain.R4,
	}.Compile())
	fmt.Println(cpu.reg)
	if cpu.reg[int16(domain.R1)] != 243 {
		t.Errorf("error adding: %d != %d", cpu.reg[int16(domain.R1)], 243)
	}
}

func TestLoad(t *testing.T) {
	mem := []int16{10, 15, 20, 8}
	cpu := NewCPU(mem)
	cpu.load(instruction.LOAD{
		DR:       domain.R1,
		PCOffset: 0,
	}.Compile())

	cpu.load(instruction.LOAD{
		DR:       domain.R2,
		PCOffset: 1,
	}.Compile())

	cpu.load(instruction.LOAD{
		DR:       domain.R3,
		PCOffset: 2,
	}.Compile())

	cpu.load(instruction.LOAD{
		DR:       domain.R4,
		PCOffset: 3,
	}.Compile())

	fmt.Println(mem)
	fmt.Println(cpu.reg)
	fmt.Println(cpu.memory[:7])

	for i, v := range mem {
		if cpu.reg[i] != v {
			t.Errorf("error wrong R%d, %d != %d", i, cpu.reg[i], v)
		}
	}
}

func TestAnd(t *testing.T) {
	cpu := NewCPU(nil)
	cpu.reg[int16(domain.R2)] = 90
	cpu.reg[int16(domain.R3)] = 153
	cpu.and(instruction.AND{
		DR: domain.R1,
		S1: domain.R2,
		S2: domain.R3,
	}.Compile())
	fmt.Println(cpu.reg)
	if cpu.reg[int16(domain.R1)] != 24 {
		t.Errorf("error adding: %d != %d", cpu.reg[int16(domain.R1)], 24)
	}
}
