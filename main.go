package main

import (
	"fmt"
	"sync"
)

const step int64 = 1
const iterationAmount int64 = 1000 // используем int64 для согласованности

func main() {
	var counter int64 = 0
	var completedIterations int64 = 0 // отслеживает количество выполненных итераций
	var c = sync.NewCond(&sync.Mutex{})

	increment := func() {
		c.L.Lock() // захватываем мьютекс перед изменением разделяемых данных
		counter += step
		completedIterations++
		
		// проверяем, достигли ли мы целевого количества итераций
		if completedIterations == iterationAmount {
			c.Broadcast() // уведомляем все ожидающие горутины
		}
		c.L.Unlock() // освобождаем мьютекс
	}

	// запускаем горутины для инкремента
	for i := int64(1); i <= iterationAmount; i++ {
		go increment()
	}

	// ждем, пока счетчик не достигнет целевого значения
	c.L.Lock()
	for completedIterations < iterationAmount { // динамическая проверка условия
		c.Wait() // ожидаем сигнала от одной из горутин
	}
	c.L.Unlock()

	fmt.Printf("Счетчик достиг значения: %d\n", counter)
}