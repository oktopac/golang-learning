package aa_channel

import (
	"fmt"
	"math/rand"
	"time"
)

func boring_sync(tosay string, sync <-chan bool) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			// time.Sleep(time.Millisecond * time.Duration(rand.Int31n(1000)))
			c <- fmt.Sprintf("I'm saying \"%s\" for the %d time!", tosay, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Int31n(1000)))
			<-sync
		}
	}()
	return c
}

func SyncTwoChannels() {
	syncChan1 := make(chan bool)
	syncChan2 := make(chan bool)

	c := fanIn(boring_sync("foobar", syncChan1), boring_sync("baabaz", syncChan2))

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	for i := 0; i < 10; i++ {
		syncChan1 <- true
		syncChan2 <- true
	}
	fmt.Println("I'm leaving now!")
}
