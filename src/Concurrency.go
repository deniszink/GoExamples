package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	//for more readable console output uncomment only one operation example and see what will happen
	fmt.Println("Goroutines ----->\n")
	//callDoWorkParallel()

	fmt.Println("Channels ----->\n")

	/*var c chan string = make(chan string) //create channel
	go pinger(c)
	go ponger(c)
	go printer(c)*/

	fmt.Println("Operator select(switch for channels) ---->\n")
	//spam()

	selectForTimer()

	var input string
	fmt.Scanln(&input)

}

// region Goroutines

func doWord(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Hello ", i)

		amt := time.Duration(rand.Intn(250)) //waits 250 ms after each Println output
		time.Sleep(time.Millisecond * amt)

	}

}

func callDoWorkParallel() {
	//this is similar to new Thread(Runnable r) in Java, we call new async operation 10 times
	for i := 0; i < 10; i++ {
		go doWord(1)
	}


	/*var input string
	fmt.Scanln(&input)*/
}
//endregion

//region Channels

func pinger(c chan <- string) {
	// channels can be readOnly or writeOnly, in this case pongerChannel sendOnly
	// pinger func will send "ping" word to channel
	for i := 0;; i++ {
		c <- "ping"
	}
}
func ponger(c chan <- string) {
	// channels can be readOnly or writeOnly, in this case pongerChannel sendOnly
	// ponger func will send "ping" word to channel
	for i := 0;; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) {
	// writeOnly channel
	//printer fun will read channel and print message from it
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func spam() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()
}

func selectForTimer() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("Message 1", msg1)
	case msg2 := <-c2:
		fmt.Println("Message 1", msg2)
	//case <-time.After(time.Second):
		fmt.Println("timeout")
		default: fmt.Println("default")
	}



}

//endregion
