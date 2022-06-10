package calc

func Add(a ... int) (s int) {
	for _, v := a {
		s+=a
	}
	return s
}
