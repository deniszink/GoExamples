package main

import "fmt"

func main() {

	fmt.Println("Variables ------>")
	fmt.Println()
	SimpleVal()
	fmt.Println()

	fmt.Println("Arrays ------> \n")
	Massive()
	fmt.Println()

	fmt.Println("Slices------> \n")
	Slices()
	fmt.Println()

	fmt.Println("Slices Operations------> \n")
	SlicesOperations()
	fmt.Println()

	fmt.Println("Maps------> \n")
	Maps()
	fmt.Println()

	fmt.Println("Map of Maps------> \n")
	MapOfMap()
	fmt.Println()


}

func SimpleVal() {
	var x int = 100
	x1 := 300
	fmt.Println("Long =", x, "short =", x1)
}

func Massive() {
	var x[3] float64 //long form
	x[0] = 95
	x[1] = 23
	x[2] = 34

	x2 := [5]float64{93, 23, 65, 12} //short form

	var total float64
	var total2 float64

	for i := 0; i < len(x); i ++ {
		total += x[i]
	}

	for _, value := range x2 {
		//we use _ to ignore variable
		total2 += value
	}

	fmt.Println(total / float64(len(x)))
	fmt.Println(total2 / float64(len(x)))

}
func Slices() {
	var x[] float64 //slice without length
	var x2[] float64 //slice without length

	x = make([]float64, 2) //slice with type of float64 and length of 5
	x2 = make([]float64, 5, 10) //slice with length of 5 and array length of 10 that slice points to

	fmt.Println(x)
	fmt.Println(x2)

	//another way to create a slice

	arr := [5]float64{1, 2, 3, 4, 5}
	slice := arr[1:3] // [low : high] low - it's a start position of array, high - it's the end position of array

	fmt.Println("Full array", arr)
	fmt.Println("Slice [1:3]", slice)
	fmt.Println("Slice [:3]", arr[:3]) //the same as [0:3], so [:3] == [:3]
	fmt.Println("Slice [:]", arr[:]) //the same as [0:len(arr)]
}

func SlicesOperations() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)

	slice1[2] = 12 // slice2 doesn't change, so append create new slice depends on slice1
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3} //slice with 3 elements
	slice4 := make([]int, 2) //slice with 2 elements than depend on previous slice

	copy(slice4, slice3) //copy values from slice3 to slice4, only first 2 values
	fmt.Println(slice3, slice4)
}

func Maps() {
	var x map[string]int // [key type]value type, smth like HashMap<K,T> in Java or Dictionary in Obj-C, we declare a variable
	x = make(map[string]int) // we initialize it
	x["key"] = 10 //we assign to it

	//or

	x2 := make(map[string]int)
	x2["key2"] = 12
	x2["key3"] = 45

	delete(x2, "key3") //we delete one item of map
	name, ok := x2["key3"]
	fmt.Println("Value =", name, ",It was successfull request -", ok) //name - it's result of request, ok - tell us is request was successful

	x2["key4"] = 200
	if name, ok := x2["key4"]; ok {//we trying to get value associated with key4, ok - it's a boolean flag, if it true, code inside block will run
		fmt.Println(name, ok)
	}

	fmt.Println(x)
	fmt.Println(x2)

}

func MapOfMap(){
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":"Hydrogen",
			"state":"gas",
		},
		"He": map[string]string{
			"name":"Helium",
			"state":"gas",
		},
		"Li": map[string]string{
			"name":"Lithium",
			"state":"solid",
		},
		"Be": map[string]string{
			"name":"Beryllium",
			"state":"solid",
		},
		"B":  map[string]string{
			"name":"Boron",
			"state":"solid",
		},
		"C":  map[string]string{
			"name":"Carbon",
			"state":"solid",
		},
		"N":  map[string]string{
			"name":"Nitrogen",
			"state":"gas",
		},
		"O":  map[string]string{
			"name":"Oxygen",
			"state":"gas",
		},
		"F":  map[string]string{
			"name":"Fluorine",
			"state":"gas",
		},
		"Ne":  map[string]string{
			"name":"Neon",
			"state":"gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}