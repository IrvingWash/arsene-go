package utils

func RemoveTrailingSlash(value string) string {
	lastChar := value[len(value)-1:]

	if lastChar != "/" {
		return value
	}

	return value[:len(value)-1]
}
