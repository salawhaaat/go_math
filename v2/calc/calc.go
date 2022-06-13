package calc

func Add(a ...int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}
