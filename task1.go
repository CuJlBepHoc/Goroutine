package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	in := make(chan int)
	out := make(chan int)
	wg := &sync.WaitGroup{}

	// Горутина для чтения чисел из стандартного ввода
	go func() {
		defer close(in)
		for {
			var input string
			fmt.Scanln(&input)
			if input == "stop" {
				break
			}
			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Введите корректное число или 'stop' для завершения.")
				continue
			}
			fmt.Println("Ввод:", num)
			in <- num
		}
	}()

	// Квадрат числа
	square := func() {
		defer wg.Done()
		for num := range in {
			out <- num * num
		}
	}

	// Произведение на 2
	product := func() {
		defer wg.Done()
		for num := range out {
			fmt.Println("Квадрат:", num)
			fmt.Println("Произведение:", num*2)
		}
	}

	wg.Add(2)
	go square()
	go product()

	wg.Wait()
}
