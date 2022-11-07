# Параллельное исполнение

## Цель
Написать функцию для параллельного выполнения N заданий (т.е. в N параллельных горутинах).

Функция должна останавливать свою работу если произошло M ошибок

Сигнатура функции: Run(task []func()error, N int, M int) error

Учесть что задания могут выполняться разное время

Длинна списка задач L = len(tasks) может быть больше или меньше N.

## Описание/Пошаговая инструкция выполнения домашнего задания
Завести в репозитории отдельный пакет (модуль) для этого ДЗ

Реализовать функцию вида Run(task []func()error, N int, M int) error

При необходимости выделить вспомогательные функции

Написать unit-тесты на функцию, проверяющие, что

- если задачи работаю без ошибок, то выполняются все N
- если в первых M задачах происходят ошибки, то всего выполнится не более N+M задач