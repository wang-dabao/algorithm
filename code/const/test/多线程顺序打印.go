package main

import (
	"fmt"
	"sync"
)

func main() {
	cha := make(chan bool)
	chb := make(chan bool)
	chc := make(chan bool)
	wg := sync.WaitGroup{}

	go func() {
		defer wg.Done()
		select {
		case <-chc:
			fmt.Println("C")
			chb <- true
		}
	}()
	go func() {
		defer wg.Done()
		select {
		case <-chb:
			fmt.Println("B")
			cha <- true
		}
	}()
	go func() {
		defer wg.Done()
		select {
		case <-cha:
			fmt.Println("A")
		}
	}()
	wg.Add(3)
	chc<-true
	wg.Wait()
}
