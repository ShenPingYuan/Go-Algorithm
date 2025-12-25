package main

func main() {

}

func fib(n int) int {
	catch := make([]int, n+1)
	if n < 2 {
		return n
	}
	var innerFib func(n int) int
	innerFib = func(n int) int {
		if n < 2 {
			return n
		}
		if catch[n] != 0 {
			return catch[n]
		}
		var res int = innerFib(n-1) + innerFib(n-2)
		catch[n] = res
		return res
	}
	return innerFib(n)
}

func fib2(n int) int {
	n_1 := 0
	n_2 := 1
	n_i := 1
	for i := 2; i <= n; i++ {
		n_i = n_1 + n_2
		n_1, n_2 = n_2, n_i
	}
	if n < 2 {
		return n
	}
	return n_i
}
