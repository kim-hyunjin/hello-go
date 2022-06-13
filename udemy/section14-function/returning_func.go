package section14function

func ReturnFunc() func() int {
	f := func() int {
		return 451
	}
	return f
}