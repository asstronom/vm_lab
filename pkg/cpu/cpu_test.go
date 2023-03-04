package cpu

import (
	"testing"

	"example.com/pkg/domain"
	"example.com/pkg/instruction"
)

func TestWork(t *testing.T) {
	mem := make([]int16, 4*2)
	mem[0] = instruction.LOAD{
		DR:       domain.R1,
		PCOffset: 5,
	}.Compile()
	mem[5] = 90
	mem[1] = instruction.LOAD{
		DR:       domain.R2,
		PCOffset: 5,
	}.Compile()
	mem[6] = 153
	mem[2] = instruction.ADD{
		DR: domain.R3,
		S1: domain.R1,
		S2: domain.R2,
	}.Compile()
	mem[3] = instruction.AND{
		DR: domain.R4,
		S1: domain.R1,
		S2: domain.R2,
	}.Compile()
	mem[4] = instruction.HALT{}.Compile()
	mycpu := NewCPU(mem)
	mycpu.Work()
	mycpu.Debug()
}
