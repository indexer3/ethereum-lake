package utils

type Xor struct {
	result []byte
}

func NewXor(result []byte) Xor {
	return Xor{result: result}
}

func (a Xor) Xor(b Xor) []byte {
	if len(a.result) != len(b.result) {
		return nil
	}

	result := make([]byte, len(a.result))
	for i := 0; i < len(a.result); i++ {
		result[i] = a.result[i] ^ b.result[i]
	}

	return result
}

func XorResultMethods(methodBytes [][]byte) []byte {
	if len(methodBytes) == 0 {
		return nil
	}

	result := make([]byte, len(methodBytes[0]))
	copy(result, methodBytes[0])

	for i := 1; i < len(methodBytes); i++ {
		result = NewXor(result).Xor(NewXor(methodBytes[i]))
	}

	return result
}
