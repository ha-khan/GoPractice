package main

func main() {
	defer func() {
		if val := recover(); val.(error) != nil {

		}
	}()
}

func Compute() {
	panic("error")
}
