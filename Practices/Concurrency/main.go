package main

//GoRoutines
import (
	"fmt"
	"time"
)

func out(from, to int) {
	for i := from; i <= to; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(i)
	}
}

func main() {

	// //sequential program first call executes and second one waits till first one execute is complete
	// out(6, 10)
	// out(0, 5)

	//GoROutines-->this gives no output bec go rountines aloows to run all func concurrently so
	//main func exits before our two func call
	go out(0, 5)
	go out(6, 10)

	//mechanisim to wait till execute

	time.Sleep(500 * time.Millisecond)

}
