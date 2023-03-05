package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandInt(min, max int) <-chan int {
	out := make(chan int, 3)

	go func() {
		for {
			out <- rand.Intn(max-min+1) + min
		}
	}()

	return out
}

func generateRandIntn(count, min, max int) <-chan int {
	out := make(chan int, count)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max-min+1) + min
		}
		close(out)
	}()

	return out
}

func main() {

	rand.Seed(time.Now().UnixNano())

	randInt := generateRandInt(1, 1000)

	fmt.Println("Generator started - infinite")

	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

	fmt.Println("Another Generator started - finite [2 numbers]")
	randIntn := generateRandIntn(2, 1, 1000)

	for i := range randIntn {
		fmt.Println(i)
	}

	fmt.Println("Another Generator started - finite [3 numbers]")
	randIntn2 := generateRandIntn(3, 1, 1000)
	for {
		n, open := <-randIntn2
		if !open {
			break
		}
		fmt.Println(n)
	}

}
