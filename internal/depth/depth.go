package depth

func AllowEndpoints(depth int) bool {
	return depth >= 2
}

func AllowSecrets(depth int) bool {
	return depth >= 3
}

func AllowAdvanced(depth int) bool {
	return depth >= 4
}

func AllowFullIntel(depth int) bool {
	return depth >= 5
}
