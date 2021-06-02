// https://stepik.org/lesson/526868/step/4?unit=519587

package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	// Возьмите первые `width` байт строки `text`,
	// допишите в конце `...` и сохраните результат
	// в переменную `res`
	// ...
	bytes := []byte(text)
	if len(bytes) > width {
		bytes = bytes[:width]
	}

	res := string(bytes)
	if res != text {
		res += "..."
	}

	fmt.Println(res)
}

// ┌─────────────────────────────┐
// │ тесты для запуска в консоли │
// └─────────────────────────────┘

// echo "Eyjafjallajokull 6" | go run step_4.go
// echo "hello 6" | go run step_4.go
