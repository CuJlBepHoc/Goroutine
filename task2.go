package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Создаем канал для получения сигналов.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Создаем канал для завершения работы программы.
	done := make(chan bool, 1)

	// Горутина для вывода квадратов натуральных чисел.
	go func() {
		i := 1
		for {
			select {
			case <-done:
				return
			default:
				fmt.Println(i * i)
				i++
			}
		}
	}()

	// Ожидаем получения сигнала.
	<-sigs

	// Обрабатываем сигнал и завершаем программу.
	fmt.Println("Выхожу из программы")
	close(done)
}
