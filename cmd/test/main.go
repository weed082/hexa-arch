package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"
)

var chans = []chan string{}
var cases = []reflect.SelectCase{}
var ctx, _ = context.WithCancel(context.Background())

func main() {
	// for i := 0; i < 5; i++ {
	// 	ch := make(chan string)
	// 	chans = append(chans, ch)
	// 	cases = append(cases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)})
	// }
	go test()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			log.Println("yex")
			chosen, value, _ := reflect.Select(cases)
			log.Println(chosen, value)
		}
	}
}

func test() {
	for {
		time.Sleep(5 * time.Second)
		ch := make(chan string)
		chans = append(chans, ch)
		cases = append(cases, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)})
		chans[len(chans)-1] <- fmt.Sprintf("count : %d", len(chans))
	}
}
