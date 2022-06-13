package section14function

func Greeting(prefix string, who ...string) string {
	s := prefix
	for i, v := range who {
		s += v
		if i != len(who)-1 {
			s += ", "
		}
	}
	return s
}