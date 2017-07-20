package aa_channel

import (
	"fmt"
)

func SyncTwoChannelsFeedback() {
	syncChan1 := make(chan bool)
	syncChan2 := make(chan bool)

	feedbackChan := make(chan bool)

	c := fanIn(boring_sync("foobar", syncChan1), boring_sync("baabaz", syncChan2))

	go func() {
		for {
			fmt.Println(<-c)
			feedbackChan <- true
		}
	}()

	for i := 0; i < 10; i++ {
		syncChan1 <- true
		<-feedbackChan
		syncChan2 <- true
		<-feedbackChan
	}

	fmt.Println("I'm leaving now!")
}
