package main

import "time"

/*
*
复现map并发读写问题, map并发读不会有问题
并不会每次都出现
*/
func main() {
	m := make(map[int]struct{})

	loop := 1000
	go func() {
		for i := 0; i < loop; i++ {
			m[i] = struct{}{}
		}
	}()

	go func() {
		for i := 0; i < loop; i++ {
			_ = m[i]
		}
	}()

	time.Sleep(10 * time.Second)
}
