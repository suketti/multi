package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	wg   sync.WaitGroup = sync.WaitGroup{}
	jobs chan string    = make(chan string, 20)
)

func Worker(id int) {
	for job := range jobs {
		fmt.Println(id, " worker got", job)
		time.Sleep(time.Duration(rand.Intn(20)) * 100 * time.Millisecond)
		fmt.Println(id, " worker finished", job)
		wg.Done()
	}

}

func main() {

	jobsToDo := make([]string, 0)
	for i := 0; i < 10; i++ {
		jobsToDo = append(jobsToDo, "Job "+strconv.Itoa(i))
	}

	/* for i := 1; i <= 2; i++ {
		go Worker(i)
	}

	for _, job := range jobsToDo {
		wg.Add(1)
		jobs <- job
		fmt.Println(" ** Job added ** ")
	}
	*/

	/*
			var index int64
			for i := 1; i <= 10; i++ {
				go func() {
					wg.Add(1)
					for i := 1; i <= 10000; i++ {
						atomic.AddInt64(&index, 1)
					}
					wg.Done()
				}()
			}
		wg.Wait()
		fmt.Println("index ", index)
	*/

	ch1 := make(chan int)
	ch2 := make(chan int)
	rand.Seed(time.Now().Unix())
	go func() {
		time.Sleep(time.Second)
		if rand.Intn(2) == 0 {
			ch1 <- 1
		} else {
			ch2 <- 1
		}
	}()

	fmt.Println("waiting")
loop:
	for {
		select {
		case <-ch1:
			fmt.Println("ch1")
			break loop
		case <-ch2:
			fmt.Println("ch2")
			break loop
		default:
			fmt.Println("nothing came yet")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
