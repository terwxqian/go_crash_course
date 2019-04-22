package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func multiReturn() (int, int) {
	return 5, 6
}

func MultiParam(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	//fmt.Println(getSum(3, 4))
	//x, y := multiReturn()
	//fmt.Printf("x=%d,y=%d", x, y)

	//fmt.Println(MultiParam(10, 2, 3))
	xs := []int{1, 20, 33}
	fmt.Println(MultiParam(xs...))

}
