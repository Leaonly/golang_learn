package cal

func addUpper(n int) int {
	res := 0
	for i := 1;i<=n;i++{
		res+=i
	}
	return res
}

func subber(n int, j int) int {
	return n - j
}