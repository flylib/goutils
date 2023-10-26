package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache[int](time.Second * 5)
	//cache.Story("A", 1, time.Second)
	//cache.Story("B", 1, time.Second*6)
	//v, ok := cache.Get("A")
	//if ok {
	//	fmt.Println(v)
	//} else {
	//	fmt.Println("not found the item of A")
	//}
	//time.Sleep(time.Second)
	//fmt.Println("after 1 second")
	//v, ok = cache.Get("A")
	//if ok {
	//	fmt.Println(v)
	//} else {
	//	fmt.Println("not found the item of A")
	//}
	N := 1000000
	for i := 0; i < N; i++ {
		go func(j int) {
			num := time.Duration(2)
			cache.Story(fmt.Sprintf("%d", j), 0, time.Second*num)
		}(i)
	}
	//for i := 0; i < N; i++ {
	//	cache.Story(fmt.Sprintf("%d", i), i, time.Second*10)
	//}

	time.Sleep(time.Second)
	var liveTotal int
	for i := 0; i < N; i++ {
		_, ok := cache.Get(fmt.Sprintf("%d", i))
		if ok {
			liveTotal++
		}
	}
	t.Log("liveTotal=", liveTotal)
}
