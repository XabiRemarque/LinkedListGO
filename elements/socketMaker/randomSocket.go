package socketmaker

func RandomSocket() (string, int) {
	return randomIPv4(), randomPort()
}
