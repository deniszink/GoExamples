package main

import (
	"hash/crc32"
	"fmt"
	"io/ioutil"
)

//hashes in Go can be cryptographic and non-cryptographic
//non-crypto can find in hash package that include algorithms like adler32,
//crc32, crc64 and fnv. Here is example of crc32
func main() {
	//crc32Examples()

	h1, err := getHash("text") // we get hash of text file
	if err != nil{
		fmt.Println("h1 error")
		return
	}
	h2, err := getHash("emazing") // we get has of emazing file
	if err != nil{
		fmt.Println("h2 error")
		return
	}
	fmt.Println(h1,h2, h1 == h2) // we compare two hashes

}

func crc32Examples() {
	//we can use this algorithm to compare two files,
	// if returned value of Sum32() will be equals, we can say that content of 2 files if identical
	//but no 100% guarantee, if it isn't equal, so content is 100% different
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)
}

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(),nil
}
