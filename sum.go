package base58

func Sum(v []int64, h int64) int64 {
	val := int64(0)
	l := int64(1)
	for _, item := range v {
		val += item * l
		l *= h
	}

	return val
}
