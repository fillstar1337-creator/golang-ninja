package main

import (
	"fmt"

	"sync"
)

var (
	counter int
	mu sync.Mutex
)

func main(){
	var wg sync.WaitGroup

	masters := make(map[string]int)
	masters["odin"] = 1
	masters["dva"] = 2
	masters["tri"] = 3

	mu.Lock()
	fmt.Println("Безопасная работа началась")

	defer fmt.Println("Конец безопасной работы")
	defer mu.Unlock()

	delete(masters, "odin")
	
	value, ok := masters["dva"]

	if ok {
		fmt.Println("Мастер на работе", "dva", value)
	} else {
		fmt.Println("Мастер на выходном", "dva")
	}
	fmt.Println("\nПеребор элементов map:")
	for key, value := range masters {
		fmt.Printf("Мастер: %s, Значение: %d\n", key, value)
	}
	wg.Wait()
}