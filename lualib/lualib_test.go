package lualib_test

func reduceMargin(s string) string {
	return s[1 : len(s)-1]
}

func reduceLMargin(s string) string {
	return s[1:]
}

func reduceRMargin(s string) string {
	return s[:len(s)-1]
}
