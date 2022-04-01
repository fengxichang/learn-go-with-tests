package channel

import (
	"fmt"
	"time"
)

func NoBufferChannel() {
	ch := make(chan int)

	go func() {
		fmt.Println("time sleep 5 second...")
		time.Sleep(time.Second * 5)
		<-ch
	}()

	fmt.Println("即将阻塞...")
	ch <- 1
	fmt.Println("ch数据被消费，主协程退出")
}

func BufferChannel() {
	ch := make(chan int, 1)

	go func() {
		fmt.Println("time sleep 5 second...")
		time.Sleep(5 * time.Second)
		i := <-ch
		i = <-ch
		fmt.Printf("i: %d", i)
		fmt.Println("")
	}()

	ch <- 1
	//ch <- 2
	fmt.Println("第三次发送数据到channel,即将阻塞")
	ch <- 3
	fmt.Println("ch数据被消费，主协程退出")
}
