package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

// Шаг наращивания счётчика
const step int64 = 1

// Конечное значение счетчика
const endCounterValue int64 = 1000

// Количество горутин
const goroutinesCount int = 10

func main() {

    var counter int64 = 0
    var wg sync.WaitGroup
    
    // Функция инкремента, которая будет выполняться каждой горутиной
    // Каждая горутина будет увеличивать счётчик на 100 (1000 / 10)
    increment := func() {
        defer wg.Done()
        // Вычисляем, сколько раз нужно увеличить счётчик для достижения результата 1000
        // при использовании 10 горутин: 1000 / 10 = 100
        incrementsPerGoroutine := endCounterValue / int64(goroutinesCount)
        
        for i := int64(0); i < incrementsPerGoroutine; i++ {
            atomic.AddInt64(&counter, step)
        }
    }
    
    // Запускаем 10 горутин
    for i := 1; i <= goroutinesCount; i++ {
        wg.Add(1)
        go increment()
    }
    
    // Ожидаем завершения всех горутин
    wg.Wait()
    
    // Печатаем результат
    fmt.Printf("Результат: %d\n", counter)
    
    // Проверяем, что результат равен 1000
    if counter == endCounterValue {
        fmt.Println("Успешно! Счётчик равен 1000")
    } else {
        fmt.Printf("Ошибка! Ожидалось %d, получено %d\n", endCounterValue, counter)
    }
}