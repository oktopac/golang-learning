package aa_channel

import (
	"fmt"
)

func TwoChannels() {
	c1 := boring("foobar")
	c2 := boring("baabaz")
	for i := 0; i < 5; i++ {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
	fmt.Println("I'm leaving now!")
}
