package aa_channel

import (
	"fmt"
)

func fanIn(chan1 <-chan string, chan2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-chan1
		}
	}()
	go func() {
		for {
			c <- <-chan2
		}
	}()
	return c
}

func CombineTwoChannels() {
	c := fanIn(boring("foobar"), boring("baabaz"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("I'm leaving now!")
}
