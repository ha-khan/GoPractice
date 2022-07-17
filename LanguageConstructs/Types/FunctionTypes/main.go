package main

func main() {

	var t func(...int)
	t = func(i ...int) {}
	t(1, 2, 3, 4, 5, 6, 7, 8)
}
