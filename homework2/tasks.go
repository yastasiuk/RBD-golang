package main

import (
	"fmt"
	"strconv"
)

func FibonacciIterative(n int) int {
	// Функція вираховує і повертає n-не число фібоначчі
	// Імплементація без використання рекурсії
	arr := [2]int{0, 1}

	for i := 2; i <= n; i++ {
		nextValue := arr[0] + arr[1]
		arr[0], arr[1] = arr[1], nextValue
		fmt.Println(i, i%2, arr)
	}

	return arr[min(n, 1)]
}

var cache = make(map[int]int)

func FibonacciRecursive(n int) int {
	// Функція вираховує і повертає n-не число фібоначчі
	// Імплементація з використанням рекурсії
	if n <= 1 {
		return n
	}

	result := cache[n]
	if result == 0 {
		result = FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
		cache[n] = result
	}

	return result
}

func IsPrime(n int) bool {
	// Функція повертає `true` якщо число `n` - просте.
	// Інакше функція повертає `false`
	if n < 2 {
		return false
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func IsBinaryPalindrome(n int) bool {
	// Функція повертає `true` якщо число `n` у бінарному вигляді є паліндромом
	// Інакше функція повертає `false`
	//
	// Приклади:
	// Число 7 (111) - паліндром, повертаємо `true`
	// Число 5 (101) - паліндром, повертаємо `true`
	// Число 6 (110) - не є паліндромом, повертаємо `false`
	nString := strconv.FormatInt(int64(n), 2)

	for i := 0; i < len(nString)/2; i++ {
		if nString[i] != nString[len(nString)-1-i] {
			return false
		}
	}

	return true
}
func ValidParentheses(s string) bool {
	// Функція повертає `true` якщо у вхідній стрічці дотримані усі правила високристання дужок
	// Правила:
	// 1. Допустимі дужки `(`, `[`, `{`, `)`, `]`, `}`
	// 2. У кожної відкритої дужки є відповідна закриваюча дужка того ж типу
	// 3. Закриваючі дужки стоять у правильному порядку
	//    "[{}]" - правильно
	//    "[{]}" - не правильно
	// 4. Кожна закриваюча дужка має відповідну відкриваючу дужку
	stack := []rune{}
	popLastValue := func() rune {
		if len(stack) == 0 {
			return 0
		}

		lastItem := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return lastItem
	}

	runes := []rune(s)

	for _, char := range runes {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')':
			if popLastValue() != '(' {
				return false
			}
		case '}':
			if popLastValue() != '{' {
				return false
			}
		case ']':
			if popLastValue() != '[' {
				return false
			}
		}
	}

	return len(stack) == 0
}

func Increment(num string) int {
	// Функція на вхід отримує стрічку яка складається лише з символів `0` та `1`
	// Тобто стрічка містить певне число у бінарному вигляді
	// Потрібно повернути число на один більше
	value := 0
	multiplier := 1
	for i := len(num) - 1; i >= 0; i-- {
		value += int(num[i]-'0') * multiplier
		multiplier *= 2
	}
	return value
}
