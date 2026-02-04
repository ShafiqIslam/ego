package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func communicateViaChannel() {
	fmt.Println("send and receive values")

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func bufferedChannel() {
	fmt.Println("Channels can be buffered")

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func closingChannels() {
	fmt.Println("range receives values from the channel until it is closed by sender")

	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
}

func fibonacciV2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func selectStatement() {
	fmt.Println("select statement lets a goroutine wait on multiple communication operations")

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciV2(c, quit)
}

func main() {
	communicateViaChannel()
	fmt.Println()
	bufferedChannel()
	fmt.Println()
	closingChannels()
	fmt.Println()
	selectStatement()
}
