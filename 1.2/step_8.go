// https://stepik.org/lesson/526867/step/8?unit=519586

package main

import (
	"fmt"
)

func main() {
	var source string
	var times int
	// гарантируется, что значения корректные
	fmt.Scan(&source, &times)

	// возьмите строку `source` и повторите ее `times` раз
	// запишите результат в `result`
	var result string
	for i := 1; i <= times; i++ {
		result += source
	}

	fmt.Println(result)
}

// ┌─────────────────────────────┐
// │ тесты для запуска в консоли │
// └─────────────────────────────┘

// echo "x 3" | go run step_8.go
// echo "omm 2" | go run step_8.go
// echo "a 5" | go run step_8.go
