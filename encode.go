package base58

import "math/big"

const alphaBet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func Encode(b []byte) string {
	x := &big.Int{}
	x.SetBytes(b)

	answer := make([]byte, 0, len(b)*136/100)

	bigZero := big.NewInt(0)
	bigRadix := big.NewInt(58)

	for x.Cmp(bigZero) > 0 {
		mod := &big.Int{}
		x.DivMod(x, bigRadix, mod)
		answer = append(answer, alphaBet[mod.Int64()])
	}

	for _, i := range b {
		if i != 0 {
			break
		}
		answer = append(answer, []byte("1")[0])
	}

	alen := len(answer)
	for i := 0; i < alen/2; i++ {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}

	return string(answer)
}
