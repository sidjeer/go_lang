//In this example we’ll look at how to implement a worker pool using goroutines and channels.
package main

import "fmt"
import "time"

//Here’s the worker, of which we’ll run several concurrent instances. These workers will receive work on the jobs channel and send the corresponding results on results. We’ll sleep a second per job to simulate an expensive task.
func worker(id string, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
func main() {
	//In order to use our pool of workers we need to send them work and collect their results. We make 2 channels for this.
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	a1 := [5]string{"English", "Japanese", "Spanish", "French", "Hindi"}
	//This starts up 3 workers, initially blocked because there are no jobs yet.
	for w := 0; w <= 4; w++ {
		go worker(a1[w], jobs, results)
	}
	//Here we send 5 jobs and then close that channel to indicate that’s all the work we have.
	for j := 1; j <= 20; j++ {
		jobs <- j
	}
	close(jobs)
	//Finally we collect all the results of the work.
	for a := 1; a <= 5; a++ {
		<-results
	}

}