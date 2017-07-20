package aa_channel

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(tosay string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("I'm saying \"%s\" for the %d time!", tosay, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Int31n(10000)))
		}
	}()
	return c
}

func Simple() {
	c := boring("foobar")
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("I'm leaving now!")
}
