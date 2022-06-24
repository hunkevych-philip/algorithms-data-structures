package dynamic_programming

import "fmt"

var FiboClient = FiboImpl{}

type FiboImpl struct {
	cached map[int]int
}

// Fibonacci returns sequential number of a fibonacci number
func (F *FiboImpl) Fibonacci(num int) int {
	if F.cached == nil {
		F.cached = make(map[int]int)
	}

	if v, ok := F.cached[num]; ok {
		fmt.Println("Returning cached value")
		return v
	}

	res := fibonacci(num)
	if len(F.cached) >= 5 {
		fmt.Println("Will remove some variable from a cache")
		for i, v := range F.cached {
			delete(F.cached, i)
			fmt.Println("Removed:", v)
			break
		}
	}

	fmt.Println("Storing", res, "in a cache")
	F.cached[num] = res

	return res
}

func fibonacci(num int) int {
	if num < 2 {
		return num
	}

	return fibonacci(num-1) + fibonacci(num-2)
}
