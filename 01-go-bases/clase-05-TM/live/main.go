package main

func horse1(channel chan string) {
	channel <- "Horse 1 finished"
}

func horse2(channel chan string) {
	channel <- "Horse 2 finished"
}

func horse3(channel chan string) {
	channel <- "Horse 3 finished"
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)

	go horse1(channel1)
	go horse2(channel2)
	go horse3(channel3)

	select {
	case winner := <-channel1:
		println(winner)
	case winner := <-channel2:
		println(winner)
	case winner := <-channel3:
		println(winner)
	}
}

// func Process(id int, channel chan<- int) {
// 	fmt.Printf("Process %d started\n", id)
// 	time.Sleep(1 * time.Second)
// 	fmt.Printf("Process %d finished\n", id)

// 	channel <- id
// }

// func main() {
// 	startedOn := time.Now()
// 	channel := make(chan int)

// 	for i := 0; i < 10; i++ {
// 		go Process(i, channel)
// 	}

// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("Process %d finished", <-channel)
// 	}

// 	fmt.Printf("All processes finished in %s", time.Since(startedOn))
// 	close(channel)
// }
