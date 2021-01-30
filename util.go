package fluentffmpeg

func containsString(arr []string, str string) bool {
	for _, val := range arr {
		if val == str {
			return true
		}
	}

	return false
}
