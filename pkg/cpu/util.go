package cpu

func signExtend(value int16, bitCount int) int16 {
	// Якщо молодший біт вказує на від'ємне число,
	// то встановлюємо біти старшого слова в 1
	if value&(1<<(bitCount-1)) != 0 {
		return int16(int32(value) | (^0 << bitCount))
	}
	return value
}
