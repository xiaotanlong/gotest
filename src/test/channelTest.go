package main

import "fmt"

/*func main() {
	fmt.Println("Begin doing something!")
	c := make(chan bool)
	go func() {
		fmt.Println("Doing something…")
		c <- true
		//close(c)
	}()
	x := <-c
	fmt.Println("Done!",x)
}*/
func main() {
	start := make(chan bool)
	for i := 1; i <= 100; i++ {
		go worker(start, i)
	}
	close(start)
	select {} //deadlock we expected
}
func worker(start chan bool, index int) {
	<-start
	fmt.Println("This is Worker:", index)
}
