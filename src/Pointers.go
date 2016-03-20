package main

import "fmt"

func main()  {
	x := 5
	receive(x)
	fmt.Println(x)

	y := 100
	receivePtr(&y)
	fmt.Println(y)

	fmt.Println("\nAssing one pointer to another -------> \n")
	obj1 := 44
	obj2 := 12
	toMatch(&obj1,&obj2)
	fmt.Println(obj1)
	fmt.Println(obj2)

	getPtr()
}

func receive(x int){ //value will copy to function
	x = 0;
}

func receivePtr(xPtr *int){
	*xPtr = 50
}

func toMatch(obj *int, obj2 *int){
	*obj = *obj2
}

func getPtr(){
	xPtr := new(int)
	*xPtr = 500
	fmt.Println(*xPtr)
}


