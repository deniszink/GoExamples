package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	readFileFullWay()

	fmt.Println()
	readFile()
}
func readFileFullWay() {
	file, err := os.Open("text")
	if err != nil {
		fmt.Println("can't open file")
		return
	}
	defer file.Close()

	//get the file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Print("can't get the file size")
		return
	}

	//read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)

	if err != nil {
		fmt.Println("can't read the file")
		return
	}

	str := string(bs)
	fmt.Print(str)
}
func readFile() {
	bs, err := ioutil.ReadFile("text") //read file using ioutil
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}
