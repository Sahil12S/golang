/**
 * GoRoutines
 * Date: May 16, 2018
 */

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

/*	for  {
		go checkLink(<-c, c)
	}*/
	// fmt.Println(<- c)

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, " is up!")
	c <- link
}






/*
*** GOROUTINES ***

* GoRoutines: Kind of threads but lightweight.
* By default "main goroutine" is created.
* what happens when some function is taking time with just one goroutine is it gets into halt or sleep state until the execution is complete.
* To create additional goroutine we use "go" before function name.
* With goroutines, with every function call, one goroutine will be created such that all can work concurrently.
* By default Go Scheduler uses one core of CPU, so at a time it runs one goroutine.
* But scheduler keeps switching between goRoutines very fast that they all are running concurrently.
* Parallel will be when goRoutines are running on different CPUs at same time.

* So what happens when we do
		go checkLink(link) ?
* Main routine is major routine or routine that will decide when the program is going to end.
* while executing the code, it creates all child routines while doing for loop. And it keep on going through the execution of code.
* if the code ends, main routine will exit without waiting for child routines to come back from halt.

* Solution?
	block Main routine from exiting until all routines are finished. CHANNEL helps in doing this.

*** CHANNELS ***

* It will communicate between all go routines (child and main).
* Channels are typed. that is data they send through must be of same data-type.
* Main routine will wait to receive message from channel, even if its just one message from one checkLink function.

*/