package main

import (
	"crypto/sha1"
	"fmt"
)
//this example very similar to crc32 example,
// it's because they both implement hash.Hash interface. Main difference, while crc32 compute
//32-bit hash,sha1 copmute 160-bit hash. Go hasn't 160 type of data, so we use 20 byte slice
func main() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Print(bs)
}
