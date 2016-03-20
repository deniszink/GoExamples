package main

import "fmt"

func main() {

	fmt.Println("\nSimple functions\n ------>")
	fmt.Println(average([]float64{45.3, 45.1, 568, 23.45, 123.43}))
	fmt.Println(oneReturn())

	text, num := twoReturn() //we receive two returned values
	fmt.Println(text, num)

	fmt.Println(returnVariable())

	fmt.Println("\nsend array as input value")
	xs := []int{12, 23, 1, 6, 23}

	fmt.Println(add(2, 4, 5, 1, 4))
	fmt.Println(add(xs...)) // we put slice of ints

	fmt.Println("\nClosures ------> \n")
	Outer()
	increment()
	CallEvenGenerator()

	var number uint = 5
	fmt.Println("Factorial of",number,"is",Factorial(number))

	fmt.Println("\nUse Defer ------> \n")
	useDefer()

	fmt.Println("\nUse panic with recover -------->\n")
	callRecover()

}

func average(xs []float64) float64 {
	total := 0.0
	for _, value := range xs {
		total += value
	}
	return total / float64(len(xs))
}

func oneReturn() string {
	return "One return value func"
}

func twoReturn() (string, int) {
	return "First v, second =", 200
}

func returnVariable() (r int) {
	r = 100
	return
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func Outer() {
	add := func(x, y int) int {
		fmt.Printf("%d + %d = ", x, y)
		return x + y
	}
	fmt.Println(add(1, 1))
}

func increment() {
	x := 0;
	doIncrement := func() int {
		x++
		return x
	}
	fmt.Println(doIncrement())
	fmt.Println(doIncrement())
}

//region closure
func EvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func CallEvenGenerator() {
	fmt.Println("Generator called")
	nextEven := EvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}
//endregion

func Factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * Factorial(x-1) //example of recursive call
}

//region defer
func printOne(){
	fmt.Println("First")
}

func printTwo(){
	fmt.Println("Second")
}

func useDefer(){
	defer printTwo() //defer move call of printTwo in the end of func useDefer
	printOne()
}
//endregion

//region panic and recover
func callRecover(){
	//this is something similar to try/catch block, but recover will retunr input value which you put in panic,
	// recover should be called in defer function, if it not recover doesn't be called
	defer func(){
		str := recover()
		fmt.Println(str)
	}()

	panic("PANIC")
}