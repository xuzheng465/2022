package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func waitForResult() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		ch <- "data"
		fmt.Println("child: sent signal")
	}()
	d := <-ch
	fmt.Println("parent: received signal :", d)
	for i := 0; i < 20; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Print("-")
	}
}

func fanOutSem() {
	children := 20
	ch := make(chan string, children)

	g := runtime.GOMAXPROCS(0)
	fmt.Println("GOMAXPROCS:", g)
	sem := make(chan bool, 5)

	for c := 0; c < children; c++ {
		go func(child int) {
			sem <- true
			{
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				ch <- "data"
				ch <- fmt.Sprintf("Child : sent signal %d", child)
			}
			<-sem
		}(c)
	}

}

func fanOut() {
	emps := 20
	ch := make(chan string, emps)

	for e := 0; e < emps; e++ {
		go func(id int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "paper"
			fmt.Println("employee : sent signal :", id)
		}(e)
	}
	for emps > 0 {
		p := <-ch
		fmt.Println(p)
		fmt.Println("manager : received signal :", emps)
		emps--
	}

	time.Sleep(time.Second * 2)
	fmt.Println("----------------------------------------------------")
}

func pooling() {
	ch := make(chan string)
	const emps = 2
	for e := 0; e < emps; e++ {
		go func(emp int) {
			for p := range ch {
				fmt.Printf("Employee %d: %s\n", emp, p)
			}
			fmt.Printf("employee %d : recv'd shutdown signal\n", emp)
		}(e)
	}

	const work = 10
	for w := 0; w < work; w++ {
		ch <- fmt.Sprintf("Work %d", w)
		fmt.Println("manager : sent signal :", w)
	}

	close(ch)
	fmt.Println("manager : sent shutdown signal")
	time.Sleep(2 * time.Second)
	fmt.Println("-----------------------------------------------------")
}

func main() {
	// pooling()
	//fanOut()
	waitForResult()
}
