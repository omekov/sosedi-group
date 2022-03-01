package util

func Increment(num uint64) uint64 {
	return num + 1
}

func Decrement(num uint64) uint64 {
	if num == 0 {
		return 0
	}

	return num - 1
}
