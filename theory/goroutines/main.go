package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	//go infront of function makes go routine - concurrent task
	//main function doesn't wait go routine
	people := [4]string{"chang", "amy", "nicolai", "sarah"}
	for _, person := range people {
		//communicate with channel
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println("Received this message: " + <-c)
	}
	//result is received from channel
	// result := <-c
	// fmt.Println(result)
	// fmt.Println(<-c)
}

// func sexyCount(person string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(person, "is sexy", i)
// 		time.Sleep(time.Second)
// 	}
// }

//channel communicates from go routine to go routine or goroutine with main
func isSexy(person string, c chan string) {
	c <- person + " is sexy"
}
