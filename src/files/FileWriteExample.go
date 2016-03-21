package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func main() {
	file, err := os.Open("emazing")
	if err != nil{
		writeToFile()
	}else{
		ReadFile()
	}
	defer file.Close()

}

func writeToFile(){ //we create file and write to it
	file, err := os.Create("emazing")
	if err != nil{
		return
	}
	defer file.Close()
	file.WriteString("Congratulations!")
}
func ReadFile(){
	bs, err := ioutil.ReadFile("emazing")
	if err != nil{
		fmt.Print("error ocured")
		return
	}
	fmt.Print(string(bs))
}

