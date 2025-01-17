# Урок 3. Многопоточность, работа с каналами
## Задание 1. Конвейер
Цели задания:
- Научиться работать с каналами и горутинами.

- Понять, как должно происходить общение между потоками.

Что нужно сделать:

- Реализуйте паттерн-конвейер:

Программа принимает числа из стандартного ввода в бесконечном цикле и передаёт число в горутину.

- Квадрат: горутина высчитывает квадрат этого числа и передаёт в следующую горутину.

- Произведение: следующая горутина умножает квадрат числа на 2.

- При вводе «стоп» выполнение программы останавливается.

Советы и рекомендации:
Воспользуйтесь небуферизированными каналами и waitgroup.

Что оценивается:

Ввод : 3
Квадрат : 9
Произведение : 18
## Задание 2. Graceful shutdown
Цель задания:

Научиться правильно останавливать приложения.
Что нужно сделать:

- В работе часто возникает потребность правильно останавливать приложения. Например, когда наш сервер обслуживает соединения, а нам хочется, чтобы все текущие соединения были обработаны и лишь потом произошло выключение сервиса. Для этого существует паттерн graceful shutdown.
- Напишите приложение, которое выводит квадраты натуральных чисел на экран, а после получения сигнала ^С обрабатывает этот сигнал, пишет «выхожу из программы» и выходит.

Советы и рекомендации:
Для реализации данного паттерна воспользуйтесь каналами и оператором select с default-кейсом.

Что оценивается:
Код выводит квадраты натуральных чисел на экран, после получения ^С происходит обработка сигнала и выход из программы.