package stringalg

func StringContainsHash(s string, substr string) bool {
	//hash
	hash := 0
	for _, c := range s {
		hash |= (1 << (c - 'A'))
	}

	for _, b := range substr {
		if hash&(1<<(b-'A')) == 0 {
			return false
		}
	}
	return true
}
