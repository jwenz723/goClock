//go 1.6

package main

import "fmt"
import "sync"
import "time"

var wg sync.WaitGroup

func main() {
	alarmTime := time.Now().Local().Add(time.Minute * time.Duration(1))
	fmt.Println("alarm time:", alarmTime.Format("15:04:05"))
	//wg.Add(2)
	ch := make(chan int)
	alarmCh := make(chan bool)
	//go generate(ch)
	//go test(ch)
	//wg.Wait()
	
	go func(a chan bool) {
		for {
			t := time.Now().Local()
			s := t.Format("15:04:05")
			fmt.Println(s)

			if s == alarmTime.Format("15:04:05") {
				a <- true
			}
			time.Sleep(1 * time.Second)
		}
	}(alarmCh)

	go func(a chan bool) {
		for b := range a {
			fmt.Println("alarm time!", b)
		}
	}(alarmCh)

	<-ch
}

func generate(out chan int) {
	for i := 0; i<10; i++ {
		out <- i
	}
	
	close(out)
	wg.Done()
}

func test(in chan int) {
	for i := range in {
		fmt.Printf("received %d\n", i)
	}
	wg.Done()
}
