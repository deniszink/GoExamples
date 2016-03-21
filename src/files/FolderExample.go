package main

import (
	"os"
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("Read Folder ----->")
	readFolder()
	fmt.Println("\nUse Wall --------->")
	useWalk()
}

func readFolder(){
	dir, err := os.Open(".") //open dir
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1) // n <= 0 - scan all files in folder,
	// otherwise it will scan only first file
	if (err != nil) {
		err.Error()
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func useWalk(){ //this func will walk for all files which allocated in root
	filepath.Walk(".",func(path string, info os.FileInfo, err error) error{
		fmt.Println(path)
		return nil
	})
}
