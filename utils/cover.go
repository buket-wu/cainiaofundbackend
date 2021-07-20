package utils

func Bool2Uint32(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}
