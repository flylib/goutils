package cache

import (
	"fmt"
	"testing"
	"time"
)

func Benchmark_Cache(b *testing.B) {
	//b.ReportAllocs()
	//b.SetBytes(2)
	c := NewCache[int](time.Second * 30)

	for i := 0; i < b.N; i++ {
		go func(i int) {
			s := i/10 + 1
			c.Story(fmt.Sprintf("%d", i), i, time.Second*time.Duration(s))
		}(i)
	}
	for i := 0; i < 30; i++ {
		b.StopTimer()
		time.Sleep(time.Second)
		b.StartTimer()
		fmt.Println("-------------second:", i)
		for j := 0; j < b.N; j++ {
			v, ok := c.Get(fmt.Sprintf("%d", j))
			if ok {
				fmt.Println(fmt.Sprintf("%d = %d", j, v))
			} else {
				fmt.Println(j, " not foun")
			}
		}
	}

	b.Log(b.N)
}
